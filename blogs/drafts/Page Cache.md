---
title: Page Cache
date: 2025-01-09 10:29:00
tags: [Linux,学习笔记]
category: 计算机基础
---

Page Cache 中 Active + Inactive 是 file-backed page 与文件对应的内存页是最需要关注的部分。

Mmap()内存映射方式和 buffered I/O 消耗的内存就来源于这部分。



SwapCached 在打开 Swap 分区后，把 inactive(anon) + Active(anon)这两项里的匿名页交换到磁盘(swap out)，然后再读入到内存(swap in)后分配的内存。**由于读入到内存后原来的SwapFile还在，所以SwapCached 也可以认为是 File-backed page 属于Page Cache**，这样做的目的是为了减少I/O

生产环境需要关闭Swap分区，因为Swap过程产生的I/O 容易引起性能抖动，在Kuberentes集群中，必须要将Swap分区关闭才能正常启动集群服务。

Shmem 指匿名共享映射的方式分配的内存

Page Cache存在的意义在于：减少I/O 提升应用的I/O 速度



### Page Cache 是如何产生和释放的

Page Cache 的产生有两种不同的方式：

Buffered I/O (标准I/O)

标准I/O写(Write(2))用户缓存区(Userpace Page对应的内存)，然后再将用户缓存区里面的数据拷贝到内核缓冲区(Pagecache Page 对应的内存)。

读则需要先从内核缓存区拷贝到用户缓冲区，再从用户缓存区读数据。



Memory-Mapped I/O (存储映射I/O)

相比标准I/O，存储映射I/O会直接将Pagecache Page给映射到用户地址，用户则直接读写Pagecache Page中的内容。



1.向用户缓冲区写数据

2.用户缓冲区拷贝数据到内核缓冲区，发生缺页中断

3.将数据拷贝到内核缓冲区

4.将脏页同步到磁盘



如何判断系统中脏页过多引起Load飙升

如果系统中存储设备既有很快速的设备，比如SSD。又有很慢的设备，比如Ceph RBD，则容易出现被回收的Page写回到慢速设备上，这回引起比较大的性能抖动。

> 控制好系统中积压的脏页数据是解决方式之一：
>
> 可以使用 sar -r 观察系统中的脏页个数