---
title: Go并发编程——学习笔记
date: 2020-12-1 10:99:00
tags: [Go,学习笔记,并发编程]
category: Go
---

## Go 标准库Cond

Cond 的目的是为 等待/通知场景下的并发提供支持。Cond 是和某个条件相关，这个条件满足后会触发一组 goroutine 协作完成，在条件还没有满足的时候，所有等待这个条件的 goroutine 都会被阻塞。

标准库中的 Cond 初始化时，需要关联一个 Locker 接口的实例，一般我们使用Mutex 或者 RWMutex。

该标准库包含三个方法:

**Signal 方法，**允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，则无需通知 waiter。如果 Cond 等待队列中有一个或者多个等待的 goroutine, 则需要从等待队列中移除第一个 goroutine 并把它唤醒。



**Broadcast 方法,** 允许调用者  Caller 唤醒所有等待此 Cond 的 goroutine 。 如果此时没有等待的 goroutine，显然无需通知 waiter。如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。



**Wait 方法，**会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。

调用 Wait 方法时必须要持有 c.L 的锁。其他方法者无需。

>调用 cond.Wait 方法之前一定要加锁。

Cond 和 一个Locker 关联, 可以利用这个 Locker 对相关依赖条件更改提供保护。

Cond 可以同时支持 Signal 和 Broadcast 方法， 而 Channel 只能支持其中一种。

Cond 的 Broadcast 方法可以被重复调用。等待条件再次变成不满足的状态后，我们又可以调用 Broadcast 再次唤醒等待的 goroutine。这也是 Channel 不能支持的， Channel 被 close 掉了之后不支持再 open。



## Once

Once 可以用来执行且仅仅执行一次动作，常常用于单例对象的初始化场景。作用和 init 函数类似，但是存在区别。

init 函数是当所在的 package 首次被加载时执行，若迟迟未被使用，则浪费了内存，又延长了程序的加载时间。

sync.Once 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下线程是安全的。

sync.Once 只暴露了一个方法 Do, 你可以多次调用 Do 方法，但是只有第一次调用 Do 方法时 f 参数才会执行，这里的 f 是一个无参数无返回值的函数。

Once 常用于初始化单例资源，或者并发访问只需初始化一次的共享资源，或者在测试的时候初始化一次测试资源。 

一个正确的 Once 实现要使用一个互斥锁，这样初始化的时候如果有并发的 goroutine，就会进入doSlow 方法。互斥锁的机制保证只有一个 goroutine 进行初始化，同时利用双检查机制。

**done 为什么是第一个字段**

```
type Once struct {
    // done indicates whether the action has been performed.
    // It is first in the struct because it is used in the hot path.
    // The hot path is inlined at every call site.
    // Placing done first allows more compact instructions on some architectures (amd64/x86),
    // and fewer instructions (to calculate offset) on other architectures.
    done uint32
    m    Mutex
}
```



**done 在热路径中，done 放在第一个字段，能够减少CPU的指令，这样做能够提升性能。**

1.热路径是程序中非常频繁执行的一系列指令，sync.Once 绝大部分场景都会访问 o.done，在热路径上是比较好理解的，如果 hot path 编译后机械码指令更少，更直接，必然是能够提升性能的。

2.因为结构体第一个字段的地址和结构体指针是相同的，如果是第一个字段，直接对结构体的指针解引用即可。如果是其他的字段，除了结构体指针外，还需要己算与第一个值的偏移。在机械码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址。因为访问第一个字段的机器代码更紧凑，速度更快。



## 线程安全的 Map 

在 Go 中，map[key] 函数返回结果可以是一个值，也可以是两个值。这是因为，如果获取一个不存在的key 对应的值时，会返回零值。

Go 内建的 map 对象不是线程安全的，并发读写的时候运行时会有检查，遇到并发问题就会导致 panic 。

**避免 map 并发读写 panic 的方式之一就是加锁，考虑到读写性能，可以使用读写锁提供性能。**

**分片加锁：更高效的并发 map** 在大量并发读写的情况下，锁的竞争会非常激烈。在并发编程中，我们的一条原则就是尽量减少锁的使用。一些单线程的应用（Redis 等 ），基本上不需要使用锁去解决并发线程访问的问题，所以可以实现很高的性能。但是对于Go 来说，并发是常用的一个特性，在这种情况下我们应该，**尽量减少锁的粒度和锁持有的时间。**

**减少锁的粒度常用方法就是分片（Shard）**，将一把锁分成几把锁，每个锁控制一个分片。

它默认采用 32个分片，**GetShard 是一个关键的方法，能根据 key 计算出分片索引。**

解决 map 并发 panic 的两个方法：加锁和分片。

如果追求更高的性能，分片加锁。因为它可以降低锁的颗粒度，进而提高访问此 map 对象的吞吐。如果并发性能要求不是那么高的场景，简单加锁方式更简单。



**Sync.Map**

Go 内建的 map 类型不是线程安全的，sync.Map 并不是用来替换内建的 map 类型的，它只能被应用在一些特殊场景内：

1.只会增长的缓存系统中，一个key 只写入一次而被读很多次。

2.多个 goroutine 为不相交的键集读，写和重写键值对。



Snyc.Map 的优点：

空间换时间。通过冗余的两个数据结构（只读的 read 字段，可写的 dirty ）,来减少加锁对性能的影响。对只读字段（read）的操作不需要加锁。

优先从 read 字段读取，更新，删除，因为对read字段的读取不需要锁。

动态调整。miss 次数过多，将 dirty 数据提升为read，避免总是从 dirty 中加锁读取。

double-checking。加锁之后还要再检查 read 字段，确定真的不存在才操作 dirty 字段。

延迟删除。删除一个键值只是打标记，只有在提升dirty 字段为 read 字段的时候才清理删除的数据。



## sync.Pool

Go 标准库中提供了一个通用的 Pool 数据结构, 也就是 sync.pool，它可以创建池化对象。

> 但是它的池化的对象可能会被垃圾回收掉，所以对于数据库长连接等场景是不合适的。

sync.Pool 用来保存一组可独立访问的临时对象。它池化的对象会在未来的某个时候被移除掉。如果没有其他对象引用这个移除对象，那么这个被移除的对象就会被垃圾回收掉。

sync.Pool 本身就是线程安全的，多个 goroutine 可以并发地调用它的方法存取对象。

sync.Pool 不可在使用之后再复制使用。

sync.Pool 提供三个对外方法：New，Get 和 Put。

**New** 是一个函数 func() interface{}。当调用 Pool 的 Get 方法从池中获取元素，没有更多的空闲元素可以返回时，会调用 New 方法来创建新的元素。如果没有设置 New 字段，将返回 nil。



**Get** 调用时会从Pool 取走一个元素，这代表着Pool 会移除这个元素。同时Pool 可能没有元素返回，这时Get 方法会去除一个 nil 。



**Put** 调用会将一个元素返回给 Pool ，Pool 将这个元素保存到池中，并且可以复用。但是如果 Put 一个 nil 值，Pool 则会忽略这个值。



**内存泄漏**

取出来的 bytes.Buffer 在使用的时候，我们可以在这个元素中增加大量的 byte 数据，这会导致底层的 byte slice 的容量可能会变得很大。即使 Reset 再放回到池子中，这些 byte slice 的容量不会改变，占据空间依然很大。又因为 Pool 回收机制，这些大的 Buffer 可能不被回收，进而一直占用很大的空间，这属于内存泄漏的问题。



**内存浪费**

池中的 buffer 都比较大，但是在实际使用的时候，很多时候只需要一个比较小的 buffer,这属于内存浪费。

可以将 buffer 池分成几层。列如：小于512 byte的元素占一个池，小于1k byte 的元素占一个池。这样分成多个池以后，可以根据需要到所需大小的池中获取 buffer。



## Context 

Context 接口包含四个方法：Deadline, Done ， Err 和 Value 。



**Deadline** 方法会返回这个 Context 被取消的截止日期。如果没有截止日期，ok的值是 false。后续每次调用这个对象的 Deadline 方法时，都会返回第一次调用相同的结果。



**Done** 方法返回一个 Channel 对象。在 Context 被取消时，此 Channel 会被 close，如果没有被取消，可能会返回nil 。后续的 Done 调用总是返回相同的结果。当 Done 被 close 的时候，你可以通过 ctx.Err 获取错误信息。

如果 Done 没有被 close，Err 方法返回 nil。 如果Done 被 close ，Err方法会返回 Done 被 Close的原因。



**Value** 返回此 ctx 中和指定的 key 相关联的 value 。

Context 中实现了两个生成顶层 Context 的方法： context.Background()  和 context.TODO() 。 这两个方法底层实现是一模一样的，用任意一个都可以。



1.一般函数在使用Context 的时候，会把这个参数放在第一个参数的位置。

2.从来不把 nil 当作 Context 类型的参数值，可以使用 context.Backgroud() 创建一个空的上下文对象，也不要使用nil。

3.Context 只用来临时做函数之间的上下文透传，不能持久化 Context 或者把 Context 长期保存。把 Context 持久化到数据库，本地文件，全局变量，缓存等方式都是错误的用法。

4.Key 的类型不应该是字符串类型或者其他内建类型，否则容易在包之间使用 Context 时产生冲突。使用 WithValue 时，Key 的类型应该是自己定义的类型。

5.常常使用 struct{} 作为底层类型定义 key的类型。对于 exported key 的静态类型，常常是接口或者指针。这样可以金量减少内存分配。



Context 常用来取消一个 goroutine 的运行。



## 原子操作基础

Package sync/atomic 实现了同步算法底层的原子的内存操作语句。

在一个原子在执行的时候，其他线程不会看到执行一半的操作结果。原子操作要么执行完毕，元哦么还没执行。不同架构的系统原子操作是不一样的。

对于单处理器单核系统来说，如果一个操作是由一个 CPU 指令来实现的，那么它就是原子操作。如果操作是基于多条指令来实现的，那么执行的过程中可能会被中断，并执行上下文茄换，这样原子性的保证就会被打破，因为这个操作可能只执行了一半。

不涉及到对资源复杂的竞争逻辑，只是会并发地读写某个标志时，就适合使用 atomic 的原子操作。

> 可以使用 atomic 实现自己定义的基本并发原语。



atomic 原子操作是实现 lock-free 数据结构的基石。

**atimic操作的对象是一个地址，你需要把可寻址的变量的地址作为参数传递给方法，而不是把变量的值传递给方法。** 



## Channel 

**执行业务处理的 goroutine 不要通过共享内存的方式通信，而是要通过 Channel 通信的方式分享数据。**

Channel 的应用场景分为五种类型。

**1.数据交流**：当作并发的 buffer 或者 queue ,解决生产者-消费者问题。多个 goroutine 可以当作生产者（Producer）和 消费者（Consumer）。

**2.数据传递**：一个 goroutine 将数据交给另一个 goroutine ，相当于把数据的拥有权托付出去。

**3.信号通知**： 一个 goroutine  可以将信号 （closing, closed, data ready 等） 传递给另外一个或者另外一组 goroutine 。

**4.任务编排**：可以让一组 goroutine  按照一定的顺序并发或者串行的执行。

**5. 锁**：利用 Channel 也可以实现互斥锁的机制。



Channel 分为三种类型：**只能接受，只能发送，可接受也可以发送** 。 

```
chan string          // 可以发送接收string
chan<- struct{}      // 只能发送struct{}
<-chan int           // 只能从chan接收int
```



我们把既能接收又能发送的 chan 叫做双向的 chan，把只能发送和只能接收的 chan 叫做单向的 chan。其中，“<-”表示单向的 chan，如果你记不住，我告诉你一个简便的方法：**这个箭头总是射向左边的，元素类型总在最右边。如果箭头指向 chan，就表示可以往 chan 中塞数据；如果箭头远离 chan，就表示 chan 会往外吐数据。**



**Chan 通过 make 来进行初始化，未初始化的 chan 的值为 nil 。Chan 可以设置容量，这样的 chan 叫做 buffered chan 。如果没有设置，它的容量为 0 ，叫做 unbuffered chan。**

如果 chan 种还有数据，那么从这个 chan 接受数据时不会阻塞。如果 chan 没达到容量数，给它发送数据也不会阻塞，否则会阻塞。

**unbuffered chan 只有读写都准备好之后才不会阻塞。**

**nil 是 chan 的零值，是一种特殊的 chan，对值是 nil 的 chan 的发送节后调用者总是会阻塞。**



```
// 发送数据到 chan
ch <- 100
// 接收数据
x := <-ch  //变量接受数据
a(<-ch) // 函数传参接受数据
<-ch	//丢弃接受的数据
```



## **使用反射操作 Channel**



在使用 Chan 时，我们有时无法确定它的数量，这样我们也没办法在编译前写成字面意义上的 select 。

这个时候 我们需要 reflect.Select 函数，它可以将一组运行时的 case clause 传入，当作参数执行。

```
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```





## 合并 SingleFlight

SingleFlight 是 Go 开发组提供的一个扩展并发原语。它的作用是：**在处理多个 goroutine 同时调用一个函数的时候，只让一个 goroutine 去调用这个函数，等这个 goroutine 返回结果时，再把结果返回给这几个同时调用的 goroutine。**

SingleFlight 主要用在合并并发请求的场景中。比如秒杀场景，你可以把这些请求合并为一个请求。

SingleFlight 的数据结构是 Group ，提供了三个方法：

**Do:** 这个方法执行一个函数，并返回函数执行的结果。你需要提供一个Key，对于同一个 key,在同一个时间只有一个在执行，同一个 key并发的请求会等待。第一个执行的请求返回的结果，就是它的返回结果。函数 fn 是一个无参的函数，返回一个结果或者error，而 Do 方法会返回函数执行的结果或者是 error ，shared 会指示 v 是否返回给多个请求。

```
func(g *Group)Do(key string,fu func()(interface{},error)) (v interface{},err error,shared bool)
```



**DoChan：**类似 Do 方法，只不过返回一个chan，等 fn 函数执行完，产生了结果以后，就能充chan 中接受这个结果。

```
func(g *Group)DoChan(key string,fn func()(interface{},error)) <- chan Result
```



**Forget:** 告诉Group 忘记这个key。之后这个 key 请求会执行 f，而不是等待前一个未完成的 fn 函数的结果。

```
func (g *Group) Forget(key string)
```



## 循环栅栏 CyclicBarrier 

CyclicBarrier，它常常应用于重复进行一组 goroutine 同时执行的场景中。它允许一组 goroutine 彼此等待，到达一个共同的执行点。同时它可以被重复使用，所以叫循环栅栏。

CyclicBarrier 和 WaitGroup 的功能有点类似，不过CyclicBarrier  更适合用在 “固定数量的 goroutine 等待同一个执行点”的场景中，而且它可以重复利用。





## 学习资料

《Go 并发编程实战》

《Go语言高性能编程》