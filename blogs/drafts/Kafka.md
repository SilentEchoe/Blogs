---
title: Kafka
date: 2025-8-25 10:06:00
tags: [中间件,消息队列,学习笔记]
category: 中间件
---

一个典型的Kafka体系架构包括若干Producer(生产者),若干Broker,若干Consumer以及一个ZooKeeper集群。

Zookeeper是Kafka用来负责集群元数据的管理，控制器的选举等操作的。



Kafka中有两个特别重要的概念：主题(Topic)和分区(Partition)这两个概念会贯穿全文。

主题是一个逻辑上的概念，它可以细分为多个分区，一个分区只属于单个主题。同一主题下的不同分区包含的消息是不同的，分区在存储层面可以看着一个可追加的日志(Log)文件，消息在被追加到分区日志文件的时候都会分配一个特定的偏移量(offset)

offset是消息在分区中的唯一标识，Kafka通过它来保证消息在分区内的顺序性，不过offset并不跨越分区，这说明Kafka保证的分区有序而不是主题有序。



Kafka中的分区可以分布在不同的服务器(broker)上，一个主题可以横跨多个broker，以此来提供比单个broker更强大的性能。

每一条消息被发送到broker之前，会根据分区规则选择存储到哪个具体的分区。如果分区规则设定得合理，所有消息都可以均匀地分配到不同的分区中。如果一个主题只对应一个文件，那么这么文件所在的机器I/O将会成为这个主题的性能瓶颈，分区则解决了这个问题，无论是创建主题时，还是在主题创建完后都可以更改分区的数量，通过增加分区的数量可以实现水平扩展。

> Kafka 为分区引入了多副本（Replica）机制，通过增加副本数量可以提升容灾能力。同一分区的不同副本中保存的是相同的消息（在同一时刻，副本之间并非完全一样），副本之间是“一主多从”的关系，其中leader副本负责处理读写请求，follower副本只负责与leader副本的消息同步。副本处于不同的broker中，当leader副本出现故障时，从follower副本中重新选举新的leader副本对外提供服务。Kafka通过多副本机制实现了故障的自动转移，当Kafka集群中某个broker失效时仍然能保证服务可用。



Kafka的消费端也具有一定的容灾能力。Consumer使用拉(Pull)模式从服务端拉取消息，并且保存消费的具体位置，当消费者宕机后恢复上线时可以根据之前保存的消费位置重新拉取需要的消息进行消费。



分区中所有的副本成为AR,所有与leader副本保持一定程度同步的副本(包括leader副本在内)组成ISR(in-Sync Replicas) ISR集合时AR集合中的一个子集。消息会先发送到leader副本，然后follower副本才能从leader副本中拉取消息进行同步，同步期间内follower副本相对于leader副本而言会有一定程度的滞后。

> 这个范围可以通过参数进行配置。leader副本同步滞后过多的副本（不包括leader副本）组成OSR（Out-of-Sync Replicas），由此可见，AR=ISR+OSR。在正常情况下，所有的 follower 副本都应该与 leader 副本保持一定程度的同步，即 AR=ISR，OSR集合为空。



Kafka的复制机制既不是完全的同步复制，也不是单纯的异步复制。事实上同步复制要求所有能工作的follower副本都能复制完，这条消息才会被确认为已成功提交，这种复制方式极大地影响了性能。

异步复制下，follwer副本异步地从leader副本中复制数据，数据只要被leader副本写入就被认为已经成功提交。如果follower副本都还没复制完而落后leader副本，突然leader副本宕机，则会造成数据丢失。



KafkaProducer 是线程安全的，可以在多个线程中共享单个KafkaProducer实例，也可以将KafkaProducer实例进行池化来供其他线程调用。

生产者发送消息主要有三种模式：发后即忘，同步，异步。

发后即忘：向Kafka发送消息而不关心消息是否正确到达，某些情况下(比如发生不可重试异常时)会造成消息丢失。性能最好，可靠性最差。

同步：可靠性高，消息要么发送成功，要么失败。如果失败可捕捉异常并进行处理。同步发送的性能很差，需要阻塞等待一条消息发送完成后才能发送下一条。

异步：在Send函数中指定一个Callback的回调函数，Kafka在返回响应时调用该函数来实现异步的发送确认。Kafka有响应就会回调，要么发送成功，要么返回异常。



消息在通过Send方法发往Broker的过程中，可能需要经过拦截器，序列化器和分区器等一系列“中间件”处理后才能发往broker。

如果消息ProducerRecord中指定了 Partition字段，那么就不需要分区器，因为Partition代表了分区号。如果没有指定Partition字段，那么就需要依赖分区器，根据Key这个字段来计算Partition的值。

默认分区器回对Key进行哈希(MurmurHash2算法)，最终根据得到的哈希值来计算分区号，拥有相同key的消息会被写入同一个分区。如果key为nil，那么消息回以轮询的方式发往主题内的各个可用分区。

> 在不改变主题分区数量的情况下，key与分区之间的映射可以保持不变。不过，一旦主题中增加了分区，那么就难以保证key与分区之间的映射关系了。
> 除了使用 Kafka 提供的默认分区器进行分区分配，还可以使用自定义的分区器，只需同DefaultPartitioner一样实现Partitioner接口即可。默认的分区器在key为null时不会选择非可用的分区，我们可以通过自定义的分区器DemoPartitioner来打破这一限制



### 拦截器

Kafka中有两种拦截器：生产者拦截器和消费者拦截器。

生产者拦截器可以在消息发送钱做一些准备工作，比如规则过滤，修改消息的内容，统计等工作。

```go
package main

import (
	"time"

	"github.com/IBM/sarama"
)

type TraceHdrInterceptor struct{}

func (TraceHdrInterceptor) OnSend(msg *sarama.ProducerMessage) {
	// 这里可以读取上下文的 traceId，这里为演示直接生成/占位
	traceID := time.Now().Format("20060102150405")

	msg.Headers = append(msg.Headers,
		sarama.RecordHeader{Key: []byte("x-trace-id"), Value: []byte(traceID)},
		sarama.RecordHeader{Key: []byte("x-biz-tag"), Value: []byte("order")},
	)
}

func newSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_8_0_0 
	cfg.Producer.Return.Successes = true
	cfg.Producer.Interceptors = []sarama.ProducerInterceptor{
		TraceHdrInterceptor{},
	}
	return sarama.NewSyncProducer(brokers, cfg)
}
```



如果某个拦截器执行需要依赖前一个拦截器的输出，那么可能会产生副作用，如果前一个拦截器由于异常而执行失败，那么这个拦截器也无法执行。在拦截链中，如果某个拦截器执行失败，那么下一个拦截器会接着从上一个执行成功的拦截器继续执行。



