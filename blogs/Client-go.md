---
title: Client-go 源码分析
date: 2023-2-12 18:56:00
tags: [Kubernetes,学习笔记,Operator开发]
category: Kubernetes
---

Client-go是与kube-apiserver通信的clients的具体实现。

<div align="center">
    	<img src="https://s1.ax1x.com/2023/02/15/pS7DEOx.jpg">  
</div>



Reflector: 从apiserver 监听（watch）特定类型的资源,拿到变更通知后,将其放入 DeltaFIFO 队列中

Informer: 从DeltaFIFO 中弹出(pop)相应对象,然后通过 Indexer将对象和索引丢到本地 cache 中,再触发相应的事件处理函数(Resource Event Handlers)

Indexer: 提供一个对象根据一定条件检索能力,典型的实现是通过 namespace/name 来构造key,通过 Thread Safe Store 来存储对象

WorkQueque: 使用延迟队列实现,在Resource Event Handlers中会完成将对象的key放入WorkQueue的过程,然后在自己的逻辑代码里消费这些key

ClientSet: 提供资源的CURD能力,能与apiserver交互

Resource Event Handlers: 在Resource Event Handlers中添加一些简单的过滤功能，能判断哪些对象需要加入到WorkQueque中处理,对于需要加到WorkQueque中的对象,就提取其key然后入队

Worker: 指业务代码处理过程,可以直接收到WorkQueque中的任务,可以通过Indexer从本地缓存检索对象,通过ClientSet实现对象的增删改查逻辑



### WorkQueue 源码分析

使用`WorkQueue`来处理`Event`,而不是直接在`Event`中编写业务逻辑是因为：Event创建的速度比处理它的速度要快，为了解决速度不一致的问题，所以引入WorkQueue机制。

WorkQueue 一般使用延时队列实现,在`Resource Event Handlers`中完成将对象的key放入WorkQueue的过程，然后在自己的逻辑代码里从WorkQueue中消费这些key。

client-go主要有三个队列,分别为普通队列,延迟队列和限速队列,后一个队列以前一个队列的实现为基础,层层添加新功能。通常我们直接使用限速队列。

#### 普通队列

```go
type Interface interface {
	Add(item interface{}) //添加一个元素
	Len() int						  //获取元素个数
	Get() (item interface{}, shutdown bool) //获取一个元素,shutdown 队列是否关闭
	Done(item interface{}) //标记元素已经处理完毕
	ShutDown()	 //关闭队列
	ShutDownWithDrain() //关闭队列,但是等待队列中的元素处理完毕
	ShuttingDown() bool //标记当前 channel 是否关闭
}


// Type is a work queue (see the package comment).
type Type struct {
	queue []t  //定义元素的处理顺序,里面所有的元素在 dirty集合中应该都有,但是不能出现在processing集合中
	dirty set //标记所有需要被处理的元素
	processing set //当前正在被处理的元素,当处理完毕后,需要检查该元素是否在 dirty 集合中,如果在则添加到 queue 队列中
	cond *sync.Cond
	shuttingDown bool
	drain        bool
	metrics queueMetrics
	unfinishedWorkUpdatePeriod time.Duration
	clock                      clock.WithTicker
}

// set是一个map,使用map key的唯一性当作set使用
type set map[t]empty

func (s set) has(item t) bool {
	_, exists := s[item]
	return exists
}

func (s set) insert(item t) {
	s[item] = empty{}
}

func (s set) delete(item t) {
	delete(s, item)
}

func (s set) len() int {
	return len(s)
}
```

普通队列中包含几个比较重要的函数：

`Add`函数

1.判断队列是否关闭,如果关闭直接返回

2.判断是否已经在dirty集合中,如果存在则直接返回

3.metrics 队列中添加该元素,同时dirty集合中添加该元素

4.如果processing（元素已经存在正在被处理的元素集合中）则返回

5.将该元素添加到队列中，并通知等待该Cond的goroutine

```go
func (q *Type) Add(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if q.shuttingDown {
		return
	}
	if q.dirty.has(item) {
		return
	}

	q.metrics.add(item)

	q.dirty.insert(item)
	if q.processing.has(item) {
		return
	}

	q.queue = append(q.queue, item)
	q.cond.Signal()
}
```



`Get`函数

1.如果队列是空的,并且队列为开启状态,则等待cond，并且将该元素添加到 "当前正在被处理的元素集合"中,

2.如果该队列开启了,但是队列为空，则直接返回

```go
func (q *Type) Get() (item interface{}, shutdown bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 && !q.shuttingDown {
		q.cond.Wait()
	}
	if len(q.queue) == 0 {
		// We must be shutting down.
		return nil, true
	}

	item = q.queue[0]
	// The underlying array still exists and reference this object, so the object will not be garbage collected.
	q.queue[0] = nil  //设置为nil,让该元素可以被垃圾回收掉
	q.queue = q.queue[1:]

	q.metrics.get(item)

	q.processing.insert(item)
	q.dirty.delete(item)

	return item, false
}
```



#### 延迟队列

延迟队列在普通队列的基础上,新增了`AddAfter`函数

```go
type delayingType struct {
	Interface					
	clock clock.Clock  //计时器
	stopCh chan struct{}
	stopOnce sync.Once

	// heartbeat ensures we wait no more than maxWait before firing
	heartbeat clock.Ticker

	// waitingForAddCh is a buffered channel that feeds waitingForAdd
	waitingForAddCh chan *waitFor  //传递 waitfor的channel,默认大小为1000

	// metrics counts the number of retries
	metrics retryMetrics
}

type DelayingInterface interface {
	Interface
	// AddAfter adds an item to the workqueue after the indicated duration has passed
	AddAfter(item interface{}, duration time.Duration)
}
```

延迟队列,主要使用`waitingLoop`实现延迟功能：

循环分为两部分：

1.从堆（优先队列,waitingForQueue）中拿一个数据,添加到通用队列中的逻辑

2.通过`waitingForAddCh`拿到新的元素,然后通过判断时间,再将它放入到堆中。

```
func (q *delayingType) waitingLoop() {
	defer utilruntime.HandleCrash()

	// Make a placeholder channel to use when there are no items in our list
	never := make(<-chan time.Time)

	// 创建一个计时器
	var nextReadyAtTimer clock.Timer

  // 初始化一个优先级队列
	waitingForQueue := &waitForPriorityQueue{}
	heap.Init(waitingForQueue)

	waitingEntryByData := map[t]*waitFor{}

	for {
		if q.Interface.ShuttingDown() {
			return
		}

		now := q.clock.Now()

		// Add ready entries
		// 从堆中取出数据,判断时间是否已到达预定时间
		// 如果没有到,就进入下次循环
		for waitingForQueue.Len() > 0 {
			entry := waitingForQueue.Peek().(*waitFor)
			if entry.readyAt.After(now) {
				break
			}

			entry = heap.Pop(waitingForQueue).(*waitFor)
			q.Add(entry.data)
			delete(waitingEntryByData, entry.data)
		}

		// Set up a wait for the first item's readyAt (if one exists)
		nextReadyAt := never
		if waitingForQueue.Len() > 0 {
			if nextReadyAtTimer != nil {
				nextReadyAtTimer.Stop()
			}
			entry := waitingForQueue.Peek().(*waitFor)
			nextReadyAtTimer = q.clock.NewTimer(entry.readyAt.Sub(now))
			nextReadyAt = nextReadyAtTimer.C()
		}

		select {
		case <-q.stopCh:
			return

		case <-q.heartbeat.C():
			// continue the loop, which will add ready items

		case <-nextReadyAt:
			// continue the loop, which will add ready items

		case waitEntry := <-q.waitingForAddCh:
			if waitEntry.readyAt.After(q.clock.Now()) {
				insert(waitingForQueue, waitingEntryByData, waitEntry)
			} else {
				q.Add(waitEntry.data)
			}

			drained := false
			for !drained {
				select {
				case waitEntry := <-q.waitingForAddCh:
					if waitEntry.readyAt.After(q.clock.Now()) {
						insert(waitingForQueue, waitingEntryByData, waitEntry)
					} else {
						q.Add(waitEntry.data)
					}
				default:
					drained = true
				}
			}
		}
	}
}
```

`AddAfter`函数的作用是在指定的延迟时长到达之后,在 work queue中新增一个元素

```go
// AddAfter adds the given item to the work queue after the given delay
func (q *delayingType) AddAfter(item interface{}, duration time.Duration) {
	// don't add if we're already shutting down
	if q.ShuttingDown() {
		return
	}

	q.metrics.retry()

	// immediately add things with no delay
	if duration <= 0 {
		q.Add(item)
		return
	}

	select {
	case <-q.stopCh:
		// unblock if ShutDown() is called
	case q.waitingForAddCh <- &waitFor{data: item, readyAt: q.clock.Now().Add(duration)}:
	}
}
```



#### 限速队列

限速器的目的：根据相应的算法获取元素的延迟时间,然后利用延迟队列来控制队列的速度

```go
// RateLimitingInterface is an interface that rate limits items being added to the queue.
type RateLimitingInterface interface {
	DelayingInterface //内嵌延迟队列

	// AddRateLimited adds an item to the workqueue after the rate limiter says it's ok
	AddRateLimited(item interface{})  //使用限速方式往队列中加入一个元素

	// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for perm failing
	// or for success, we'll stop the rate limiter from tracking it.  This only clears the `rateLimiter`, you
	// still have to call `Done` on the queue.
	Forget(item interface{})  //标识一个元素结束重试

	// NumRequeues returns back how many times the item was requeued
	NumRequeues(item interface{}) int //标识这个元素被处理了多少次
}

// 限速器的定义
type RateLimiter interface {
	// When gets an item and gets to decide how long that item should wait
	When(item interface{}) time.Duration
	// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for failing
	// or for success, we'll stop tracking it
	Forget(item interface{})
	// NumRequeues returns back how many failures the item has had
	NumRequeues(item interface{}) int
}
```

`RateLimiter`接口存在五个实现:

BucketRateLimiter: 使用Go语言标准库 "golfing.org/x/time/rate.Limiter"包实现,BucketRateLimiter实例化的时候,会设置令牌桶相关参数,比如设置令牌桶里面最多有100个令桶,每秒发放10个令牌

ItemExponentialFailureRateLimiter: 失败次数越多,限速越长,而且是呈指数级增长的一种限速器

ItemFastSlowRateLimiter: 快慢限速器,快慢指的是定义一个阈值,达到阈值之前快速重试,超过了就满满重试

MaxOfRateLimiter: 通过维护多个限速器列表,返回其中最严格的一个延迟

WithMaxWaitRateLimiter: 在其他限速器上包装一个最大延迟的属性,如果到了最大延时,则直接返回

Resource Event Handlers 会完成将对象的 key 放入到 WorkQueue的过程,我们可以在自己的逻辑代码里从 WorkQueue 中消费这些 Key。延迟队列实现了 item的延迟入队效果,内部是一个"优先级队列",用了"最小堆"（有序完全二叉树）,所以"在requeueAfter中指定一个凋谐过程1分钟后重试"的实现原理也就清晰了。



### DeltaFIFO 源码分析

DeltaFIFO 是一个生产者-消费者的队列,生产者是Reflector,消费者是Pop函数。

DeltaFIFO 的数据来源为 Reflector，通过 Pop 操作消费数据，消费的数据一方面存储到 Indexer 中，另一方面可以通过 Informer 的 handler 进行处理，Informer 的 handler 处理的数据需要与存储在 Indexer 中的数据匹配。需要注意的是，Pop 的单位是一个 Deltas，而不是 Delta。

```go
type Queue interface {
	Store
	Pop(PopProcessFunc) (interface{}, error)
	AddIfNotPresent(interface{}) error
	HasSynced() bool
	Close()
}

type Store interface {
	Add(obj interface{}) error
	Update(obj interface{}) error
	Delete(obj interface{}) error
	List() []interface{}
	ListKeys() []string
	Get(obj interface{}) (item interface{}, exists bool, err error)
	GetByKey(key string) (item interface{}, exists bool, err error)
	Replace([]interface{}, string) error
	Resync() error
}

type Delta struct {
	Type   DeltaType
	Object interface{}
}

type Deltas []Delta


type DeltaFIFO struct {
	lock sync.RWMutex
	cond sync.Cond
	items map[string]Deltas
	queue []string
	populated bool
	initialPopulationCount int
	keyFunc KeyFunc
	knownObjects KeyListerGetter
	closed bool
	emitDeltaTypeReplaced bool
}

// DeltaType 是一个字符串类型,对应的是Added描述一个Delta类型
type DeltaType string

// Change type definition
const (
	Added   DeltaType = "Added"
	Updated DeltaType = "Updated"
	Deleted DeltaType = "Deleted"
	Replaced DeltaType = "Replaced"
	Sync DeltaType = "Sync"
)
```

DetlaFIFO 同时实现了 Queue 和 Store 接口，使用 Deltas 保存了对象状态的变更信息(如Pod的删除或添加)，Deltas 缓存了针对相同对象的多个状态变更信息,如 Pod 的 Deltas[0]可能更新了标签，Deltas[1]可能删除了该 Pod。最老的状态变更信息为 Oldest()，最新的状态变更信息为 Newest()，使用中，获取 DeltaFIFO 中对象的 key 以及获取 DeltaFIFO 都以最新状态为准。

[^DeltaFIFO结构图如下]: 图来源于《Kubernetes Operator 开发进阶》



<div align="center">
    	<img src="https://s1.ax1x.com/2023/02/19/pSO9kKf.png">  
</div>

#### 核心函数

store接口中的`Add() Update()`等函数都会调用`queueActionLocked`函数

`queueActionLocked`函数的作用主要是构建一个Delta添加到[]Deltas中,其中包含一个去重判断,如果已经存在,则只更新items map中对应这个key的[]Deltas

```go
func (f *DeltaFIFO) queueActionLocked(actionType DeltaType, obj interface{}) error {
	id, err := f.KeyOf(obj)
	if err != nil {
		return KeyError{obj, err}
	}
	oldDeltas := f.items[id]
	newDeltas := append(oldDeltas, Delta{actionType, obj})
	newDeltas = dedupDeltas(newDeltas)

	if len(newDeltas) > 0 {
		if _, exists := f.items[id]; !exists {
			f.queue = append(f.queue, id)
		}
		f.items[id] = newDeltas
		f.cond.Broadcast()
	} else {
		// This never happens, because dedupDeltas never returns an empty list
		// when given a non-empty list (as it is here).
		// If somehow it happens anyway, deal with it but complain.
		if oldDeltas == nil {
			klog.Errorf("Impossible dedupDeltas for id=%q: oldDeltas=%#+v, obj=%#+v; ignoring", id, oldDeltas, obj)
			return nil
		}
		klog.Errorf("Impossible dedupDeltas for id=%q: oldDeltas=%#+v, obj=%#+v; breaking invariant by storing empty Deltas", id, oldDeltas, obj)
		f.items[id] = newDeltas
		return fmt.Errorf("Impossible dedupDeltas for id=%q: oldDeltas=%#+v, obj=%#+v; broke DeltaFIFO invariant by storing empty Deltas", id, oldDeltas, obj)
	}
	return nil
}
```



`Pop` 函数会按照元素的添加或更新顺序有序返回一个元素(Deltas),在队列为空时会阻塞。Pop过程中会先从队列中删除一个元素后返回,如果处理失败了,需要通过 AddIfNotPresent函数将这个元素重新加回到队列汇总。

Pop的参数是 Type PopProcessFunc func(interface{}) error 类型的process,在Pop函数中,直接将队列中第一个元素出队,然后丢给process处理,如果处理失败会重新入队,但是这个 Deltas 和对应的错误信息会被返回

```
func (f *DeltaFIFO) Pop(process PopProcessFunc) (interface{}, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	for {
		for len(f.queue) == 0 {
			// When the queue is empty, invocation of Pop() is blocked until new item is enqueued.
			// When Close() is called, the f.closed is set and the condition is broadcasted.
			// Which causes this loop to continue and return from the Pop().
			if f.closed {
				return nil, ErrFIFOClosed
			}

			f.cond.Wait()
		}
		isInInitialList := !f.hasSynced_locked()
		id := f.queue[0]
		f.queue = f.queue[1:]
		depth := len(f.queue)
		if f.initialPopulationCount > 0 {
			f.initialPopulationCount--
		}
		item, ok := f.items[id]
		if !ok {
			// This should never happen
			klog.Errorf("Inconceivable! %q was in f.queue but not f.items; ignoring.", id)
			continue
		}
		delete(f.items, id)
		// Only log traces if the queue depth is greater than 10 and it takes more than
		// 100 milliseconds to process one item from the queue.
		// Queue depth never goes high because processing an item is locking the queue,
		// and new items can't be added until processing finish.
		// https://github.com/kubernetes/kubernetes/issues/103789
		if depth > 10 {
			trace := utiltrace.New("DeltaFIFO Pop Process",
				utiltrace.Field{Key: "ID", Value: id},
				utiltrace.Field{Key: "Depth", Value: depth},
				utiltrace.Field{Key: "Reason", Value: "slow event handlers blocking the queue"})
			defer trace.LogIfLong(100 * time.Millisecond)
		}
		err := process(item, isInInitialList)
		if e, ok := err.(ErrRequeue); ok {
			f.addIfNotPresent(id, item)
			err = e.Err
		}
		// Don't need to copyDeltas here, because we're transferring
		// ownership to the caller.
		return item, err
	}
}

```



`Replace` 主要做了两件事:

1.给传入的对象列表添加一个 Sync/Replace DeltaType的Delta

2.执行一些与删除相关的程序逻辑

```go
func (f *DeltaFIFO) Replace(list []interface{}, _ string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	keys := make(sets.String, len(list))

	// keep backwards compat for old clients
	action := Sync
	if f.emitDeltaTypeReplaced {
		action = Replaced
	}

	// Add Sync/Replaced action for each new item.
	for _, item := range list {
		key, err := f.KeyOf(item)
		if err != nil {
			return KeyError{item, err}
		}
		keys.Insert(key)
		if err := f.queueActionLocked(action, item); err != nil {
			return fmt.Errorf("couldn't enqueue object: %v", err)
		}
	}

	if f.knownObjects == nil {
		// Do deletion detection against our own list.
		queuedDeletions := 0
		for k, oldItem := range f.items {
			if keys.Has(k) {
				continue
			}
			// Delete pre-existing items not in the new list.
			// This could happen if watch deletion event was missed while
			// disconnected from apiserver.
			var deletedObj interface{}
			if n := oldItem.Newest(); n != nil {
				deletedObj = n.Object
			}
			queuedDeletions++
			if err := f.queueActionLocked(Deleted, DeletedFinalStateUnknown{k, deletedObj}); err != nil {
				return err
			}
		}

		if !f.populated {
			f.populated = true
			// While there shouldn't be any queued deletions in the initial
			// population of the queue, it's better to be on the safe side.
			f.initialPopulationCount = keys.Len() + queuedDeletions
		}

		return nil
	}

	// Detect deletions not already in the queue.
	knownKeys := f.knownObjects.ListKeys()
	queuedDeletions := 0
	for _, k := range knownKeys {
		if keys.Has(k) {
			continue
		}

		deletedObj, exists, err := f.knownObjects.GetByKey(k)
		if err != nil {
			deletedObj = nil
			klog.Errorf("Unexpected error %v during lookup of key %v, placing DeleteFinalStateUnknown marker without object", err, k)
		} else if !exists {
			deletedObj = nil
			klog.Infof("Key %v does not exist in known objects store, placing DeleteFinalStateUnknown marker without object", k)
		}
		queuedDeletions++
		if err := f.queueActionLocked(Deleted, DeletedFinalStateUnknown{k, deletedObj}); err != nil {
			return err
		}
	}

	if !f.populated {
		f.populated = true
		f.initialPopulationCount = keys.Len() + queuedDeletions
	}

	return nil
}
```



### Indexer 和 ThreadSafeStore

Indexer是Client-go用来存储资源对象并自带索引功能的本地存储,Reflector从DeltaFIFO中将消费出来的资源对象存储至Indexer。而且Indexer中的数据与Etcd集群中的数据保持完全一致。

Index主要为对象提供根据一定条件进行检索的能力,比如通过namespace/name来构造key,通过ThreadSafeStore来存储对象。Index主要依赖ThreadSafeStore的实现,是client-go 提供的一种缓存机制,通过检索本地缓存可以有效降低apiserver的压力。

Indexer主要在Store接口的基础上拓展了对象的检索功能,而ThreadSafeStore才是Indexer的核心逻辑

#### 数据结构

```
type ThreadSafeStore interface {
	Add(key string, obj interface{})
	Update(key string, obj interface{})
	Delete(key string)
	Get(key string) (item interface{}, exists bool)
	List() []interface{}
	ListKeys() []string
	Replace(map[string]interface{}, string)
	Index(indexName string, obj interface{}) ([]interface{}, error)
	IndexKeys(indexName, indexedValue string) ([]string, error)
	ListIndexFuncValues(name string) []string
	ByIndex(indexName, indexedValue string) ([]interface{}, error)
	GetIndexers() Indexers

	// AddIndexers adds more indexers to this store.  If you call this after you already have data
	// in the store, the results are undefined.
	AddIndexers(newIndexers Indexers) error
	// Resync is a no-op and is deprecated
	Resync() error
}

// threadSafeMap 对应的实现 ThreadSafeStore
type threadSafeMap struct {
	lock  sync.RWMutex
	items map[string]interface{}

	// index implements the indexing functionality
	index *storeIndex
}
```



storeIndex 主要由以下数据结构组成：

```go
type Index map[string]sets.String

// Indexers maps a name to an IndexFunc
type Indexers map[string]IndexFunc

// Indices maps a name to an Index
type Indices map[string]Index
```



下图为数据结构关系图（来源于《Kubernetes Operator 开发进阶》）：

<div align="center">
    	<img src="https://s1.ax1x.com/2023/02/23/pSxt4ns.png">    
 </div>




Indexers 中保存的是 Index函数map,字符串namesapce作为key,IndexFunc 类型的实现`MetaNamespaceIndexFunc`函数作为value。通过namespace来检索时,借助IndexFunc可以拿到对应的计算Index的函数,然后调用这个函数把对象传入进去,就可以计算出这个对象对应的key,就是具体的namespace值,比如上图中的"default"或"system"。通过key一层一层向下找,最终找到对应的对象信息。

#### 核心方法

接口定义`Add,Update`等方法,比如Add方法就是直接调用Update,Update和Delete函数都会调用`updateIndices`函数

```
func (c *threadSafeMap) Add(key string, obj interface{}) {
	c.Update(key, obj)
}

func (c *threadSafeMap) Update(key string, obj interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	oldObject := c.items[key]
	c.items[key] = obj
	c.index.updateIndices(oldObject, obj, key)
}

func (c *threadSafeMap) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if obj, exists := c.items[key]; exists {
		c.index.updateIndices(obj, nil, key)
		delete(c.items, key)
	}
}

func (c *threadSafeMap) Get(key string) (item interface{}, exists bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, exists = c.items[key]
	return item, exists
}
```



`updateIndices`函数对于Create只提供newObj,对于Update需要同时提供oldObj和newObj,对于Delete只需要提供oldObj

```
func (i *storeIndex) updateIndices(oldObj interface{}, newObj interface{}, key string) {
	var oldIndexValues, indexValues []string
	var err error
	for name, indexFunc := range i.indexers {
	// oldObj是否存在,如果不存在就置空,如果存在则取出相应值 
		if oldObj != nil {
			oldIndexValues, err = indexFunc(oldObj) 
		} else {
			oldIndexValues = oldIndexValues[:0]
		}
		if err != nil {
			panic(fmt.Errorf("unable to calculate an index entry for key %q on index %q: %v", key, name, err))
		}
 	// 和上面判断oldObj是否存在的逻辑一样
		if newObj != nil {
			indexValues, err = indexFunc(newObj)
		} else {
			indexValues = indexValues[:0]
		}
		if err != nil {
			panic(fmt.Errorf("unable to calculate an index entry for key %q on index %q: %v", key, name, err))
		}
   
    // 拿到一个index
		index := i.indices[name]
		if index == nil {
			index = Index{}
			i.indices[name] = index
		}

		if len(indexValues) == 1 && len(oldIndexValues) == 1 && indexValues[0] == oldIndexValues[0] {
			// We optimize for the most common case where indexFunc returns a single value which has not been changed
			continue
		}
    // 删除oldIndex
		for _, value := range oldIndexValues {
			i.deleteKeyFromIndex(key, value, index)
		}
		
		// 添加一个新的index
		for _, value := range indexValues {
			i.addKeyToIndex(key, value, index)
		}
	}
}
```



### Reflector

Reflector 用于监控(Watch)制定的Kubernetes资源,当监控的资源发生变化时,触发相应的变更事件,例如Added事件,Updated事件等,将其资源对象存放到本地缓存DeltaFIFO中。

#### ListerWatcher

ListerWatcher是Reflector的主要能力提供者,通过一种叫作 ListAndWatch 的方法，把 APIServer 中的 API 对象缓存在了本地，并负责更新和维护这个缓存。ListAndWatch通过 APIServer 的 LIST API“获取”所有最新版本的 API 对象；然后，再通过 WATCH API 来“监听”所有这些 API 对象的变化；

List-watch主要分为两部分:List调用API展示资源列表,watch监听资源变更事件,基于HTTP长链接实现。

> Watch 通过Chunked transfer enconding(分块传输编码)在Http长链接接受 apiserver发来的资源变更事件。
> HTTP 分块传输编码允许服务器为动态生成的内容维持 HTTP持久链接。使用分块传输编码，数据分解成一系列数据块，并以一个或者多个块发送，这样服务器可以发送数据而不需要预先知道发送内容的总大小。

#### 核心方法

`Reflector.ListAndWatch`方法是Reflector的核心逻辑之一

ListAndWatch方法是先列出特定资源的所有对象,然后获取其资源版,使用这个资源版本开始监听的流程。监听到新版本资源后,将其加入DeltaFIFO的动作是在 watchHandler方法中具体实现。

在此之前list(列选)到到最新元素会通过syncWith方法添加一个Sync类型的DeltaType到DeltaFIFO中,所以list操作本身也会触发后面调谐逻辑。

```go
func (r *Reflector) ListAndWatch(stopCh <-chan struct{}) error {
	klog.V(3).Infof("Listing and watching %v from %s", r.typeDescription, r.name)
  //获取资源列表
	err := r.list(stopCh)
	if err != nil {
		return err
	}

	resyncerrc := make(chan error, 1)
	cancelCh := make(chan struct{})
	defer close(cancelCh)
	go func() {
		resyncCh, cleanup := r.resyncChan()
		defer func() {
      //调用完最后一个，然后清理资源
			cleanup() 
		}()
		for {
			select {
			case <-resyncCh:
			case <-stopCh:
				return
			case <-cancelCh:
				return
			}
			if r.ShouldResync == nil || r.ShouldResync() {
				klog.V(4).Infof("%s: forcing resync", r.name)
				if err := r.store.Resync(); err != nil {
					resyncerrc <- err
					return
				}
			}
			cleanup()
			resyncCh, cleanup = r.resyncChan()
		}
	}()

  //截止时间后重试
	retry := NewRetryWithDeadline(r.MaxInternalErrorRetryDuration, time.Minute, apierrors.IsInternalError, r.clock)
	for {
		// 给stopCh一个停止循环的机会，即使在continue语句进一步错误的情况下
		select {
		case <-stopCh:
			return nil
		default:
		}

		timeoutSeconds := int64(minWatchTimeout.Seconds() * (rand.Float64() + 1.0))
		options := metav1.ListOptions{
			ResourceVersion: r.LastSyncResourceVersion(),
		  //避免死循环，设置一个超时时间
			TimeoutSeconds: &timeoutSeconds,
      //通过启动AllowWatchBookmarks选项,减少负载
			AllowWatchBookmarks: true,
		}

		
		start := r.clock.Now()
		w, err := r.listerWatcher.Watch(options)
		if err != nil {
			//如果错误,可能是apiserver没有相应
			if utilnet.IsConnectionRefused(err) || apierrors.IsTooManyRequests(err) {
				<-r.initConnBackoffManager.Backoff().C()
				continue
			}
			return err
		}

		err = watchHandler(start, w, r.store, r.expectedType, r.expectedGVK, r.name, r.typeDescription, r.setLastSyncResourceVersion, r.clock, resyncerrc, stopCh)
		retry.After(err)
		if err != nil {
			if err != errorStopRequested {
				switch {
				case isExpiredError(err):
					// Don't set LastSyncResourceVersionUnavailable - LIST call with ResourceVersion=RV already
					// has a semantic that it returns data at least as fresh as provided RV.
					// So first try to LIST with setting RV to resource version of last observed object.
					klog.V(4).Infof("%s: watch of %v closed with: %v", r.name, r.typeDescription, err)
				case apierrors.IsTooManyRequests(err):
					klog.V(2).Infof("%s: watch of %v returned 429 - backing off", r.name, r.typeDescription)
					<-r.initConnBackoffManager.Backoff().C()
					continue
				case apierrors.IsInternalError(err) && retry.ShouldRetry():
					klog.V(2).Infof("%s: retrying watch of %v internal error: %v", r.name, r.typeDescription, err)
					continue
				default:
					klog.Warningf("%s: watch of %v ended with: %v", r.name, r.typeDescription, err)
				}
			}
			return nil
		}
	}
}
```

监控资源对象

Watch（监控）操作通过HTTP协议与Kubernetes API Server 建立长链接,接收 Kubernetes API Server 发来的资源变更事件。Watch操作的实现机制使用HTTP协议的分块传输编码(Chunked Transfer Enconding)。当client-go 调用Kubernetes API Server时,Kubernetes API Server 在Response的 HTTP Header 中设置 Transfer-Encoding的值为chunked,表示采用分块传输编码,客户端收到该信息后,便于服务端进行连接，并等待下一个数据块(资源的事件信息)



### 学习资料

《Kubernetes Operator 开发进阶》

《Kubernetes源码剖析》

[DeltaFIFO](https://www.qikqiak.com/k8strain/k8s-code/client-go/deltafifo/)

[理解 K8S 的设计精髓之 List-Watch机制和Informer模块](https://zhuanlan.zhihu.com/p/59660536)

