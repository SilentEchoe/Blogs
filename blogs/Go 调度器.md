---
title: Go 调度器
date: 2020-12-31 11:29:00
tags: [Go,学习笔记,调度器]
category: Go
---

### 为什么会需要调度器？

Go 语言中使用 Goroutine 来代替线程，在编程语言层面中实现，它的特点就是占用的内容空间比线程更小，而且上下文切换不需要经过内核态，同时这意味着上下文切换的开销也更低。(如果引入使用线程的操作系统中，线程是独立调度的基本单位，进程是资源拥有的基本单位。而在同一进程中，线程的切换不会引起进程切换。在不同进程中进行线程切换，意味着额外资源需要消耗会更多)

> 一个线程分为 “内核态”线程和“用户态”线程。 一个用户态线程必须要绑定一个内核态线程，但是 CPU并不知道有 **用户态线程** 的存在，它只知道它运行的是一个 **内核态线程。
> **
>
> **一个协程可以绑定一个线程，那么也可以通过调度器把多个协程与一个或者多个线程进行绑定。**



不同于其他语言直接使用操作系统来调度线程，Go 有一个专门的调度器来调度 Goroutine 。



### 如果 Goroutine 直接由操作系统调度，那会出现什么问题？

1.操作系统去操作 Goroutine 可能不会做出好的调度决策。比如 Go GC 在执行回收任务时，会需要所有的 Goroutine 停止工作，然后去标记需要清理的内存，操作系统去对 Goroutine 做操作可能没有 Go 调度器直接对 Goroutine 操控更方便。

2.当 Goroutine 过多时，比如 GC 时，必须等待它们达到一致性状态。而 Go 调度器可以更容器确认内存是否是一致。

> Go 会有一个调度器，一方面是因为 Goroutine 的轻量可以轻而易举地创建多个，直接使用操作系统去调用 Goroutine 会导致几个不便，例如上文提到的 GC 执行回收时确定内存一致性。另一方面，由 Go 语言调度器去调度 Goroutine 可以减少很多额外的开销，以提供高并发能力。 Goroutine 和调度器是 Go 语言能够高效地处理任务并且最大化利用资源的基础。

**Go 语言的调度器通过使用与 CPU 数量相等的线程减少线程频繁切换的内存开销，同时在每一个线程上执行额外开销更低的 Goroutine 来降低操作系统和硬件的负载。**



### 抢占式调度器

Go语言从1.14版本至今，实现的是**基于信号的真抢占式调度。**在调度器的演进过程中，1.1版本的**任务窃取调度器**引入了处理器P，构成了目前的G-M-P模型，该模型沿用至今，在讨论抢占式调度前，我们先理解什么是G-M-P模型。

#### **G-M-P 模型**

G

Goroutine 在调度器中的地位与线程在操作系统中差不多，我们通常称它为'协程'。但是实际上 Goroutine不是传统意义上的协程，主流的线程模型分为三种：内核级线程模型，用户级线程模型和两级线程模型(混合型线程模型)，传统的协程属于用户级线程模型，但是 Goroutine 和调度器在底层实现上其实是属于两级线程模型。

Goroutine 只存在与Go语言的运行时中，它是Go 语言在用户态提供的线程，是一种颗粒度更细的资源调度单元，一般情况下Goroutine 中是一段用户的业务代码。



M

Go 语言并发模型中的 M 是操作系统线程。调度器最多可以创建一万个线程，但是其中的大多数线程都不会执行用户代码(它们可能陷入系统调用)，最多只有**当前机器的核数**个活跃线程能够正常运行。

默认情况下八核的机器会创建八个活跃的操作系统线程，每一个线程都对应一个运行时中的M。我们可以通过`runtime.NumCPU()`函数来获取线程数量，线程数量等用于CPU数，默认的设置不会频繁触发操作系统的线程调度和上下文切换，所有的调度都会发生在用户态，由 Go 语言调度器触发，能够减少很多额外开销。

需要注意的是：M中会包含一个 `g0`的Goroutine，这是一个比较特殊的Goroutine，它会深度参与运行时的调度过程，包括Goroutine的创建，大内存分配和CGO函数的执行。



P

调度器中的处理器 P 是线程和 Goroutine 的中间层，它能提供线程需要的上下文环境，也会负责线程上的等待队列，通过处理器 P 的调度，每一个内核线程都能执行多个 Goroutine，它能在 Goroutine 进行 I/O 操作时及时让出计算资源，提高线程的利用率。

P会持有一个运行队列，其中存储等待执行的Goroutine列表。



介绍完了GMP模型后，我们可以尝试描述整个调度的过程了：

调度器首先会持有一个全局的队列，里面存放着等待运行的Goroutine，同时 P 也会持有一个队列，里面也存放着 Goroutine ，但是P所持有的这个队列数量不会超过256 个。

当一个Goroutine被新创建出来，会优先加入到 P 所持有的队列，如果 P 所持有的队列满了，则会把本地队列中一半的 G 移动到全局队列。

M(线程)从 P 所持有的队列中获取Goroutine 并执行，如果 P 的队列为空，调度器也会尝试从全局队列拿一批 G 放入到 P 的队列中，或者从其他的 P 队列中拿一部分放入到较少队列的 P中。





### 抢占式调度器的作用

从调度器的演变可以看出，调度器的作用就是为了更快，更均匀地调度Goroutine。最开始设计**抢占式**调度器是为了避免：

1.某些Goroutine长时间占用线程，导致其他Goroutine饥饿

2.垃圾回收时需要暂停整个程序，最长可能需要几分钟，导致整个程序无法工作

但是基于写协作的抢占式调度方案也会存在一个问题：边缘情况下一些Goroutine无法被抢占。同样的为了解决这个问题，基于信号的抢占式调度方案被实现出来。

该方案可以在垃圾回收在扫描栈时会触发抢占调度，从而减少边缘情况的产生。













## 学习资料

《Go语言设计与实现》

https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/

https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit

https://zhuanlan.zhihu.com/p/77620605

https://mp.weixin.qq.com/s/GhC2WDw3VHP91DrrFVCnag