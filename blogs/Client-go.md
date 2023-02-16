---
title: Client-go 源码分析
date: 2023-2-12 18:56:00
tags: [Kubernetes,学习笔记,Operator开发]
category: Kubernetes
---

Client-go是与kube-apiserver通信的clients的具体实现。

[![pS7DEOx.jpg](https://s1.ax1x.com/2023/02/15/pS7DEOx.jpg)](https://imgse.com/i/pS7DEOx)

Reflector: 从apiserver 监听（watch）特定类型的资源,拿到变更通知后,将其放入 DeltaFIFO 队列中

Informer: 从DeltaFIFO 中弹出(pop)相应对象,然后通过 Indexer将对象和索引丢到本地 cache 中,再触发相应的事件处理函数(Resource Event Handlers)

Indexer: 提供一个对象根据一定条件检索能力,典型的实现是通过 namespace/name 来构造key,通过 Thread Safe Store 来存储对象



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





Resource Event Handlers 会完成将对象的 key 放入到 WorkQueue的过程,我们可以在自己的逻辑代码里从 WorkQueue 中消费这些 Key。延迟队列实现了 item的延迟入队效果,内部是一个"优先级队列",用了"最小堆"（有序完全二叉树）,所以"在requeueAfter中指定一个凋谐过程1分钟后重试"的实现原理也就清晰了。





### 学习资料

《Kubernetes Operator 开发进阶》

