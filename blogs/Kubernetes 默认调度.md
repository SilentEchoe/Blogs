---
title: Kubernetes 默认调度
date: 2022-5-26 16:26:00
tags: [Kubernetes,学习笔记,调度]
category: Kubernetes
---

# Kubernetes 默认调度器

Kubernetes 中，默认调度器的主要职责，就是为一个新创建的 Pod ，寻找一个最合适的节点（Node）

那怎么定义“最合适”？

1.从集群所有的节点中，根据**调度算法**挑选出可以运行该 Pod 的节点

2.从挑选出的节点中，再根据**调度算法**选出一个最符合条件的节点作为最终节点

在具体的调度流程中，默认调度器会首先调用一组叫作 **Predicate** 的调度算法，来检查每个 Node。然后，再调用一组叫作 **Priority** 的调度算法，来给上一步得到的结果里的每个 Node 打分。最终的调度结果，就是得分最高的那个 Node。



**Informer Path**

启动一系列 Informer ，用来监听 Etcd 中 Pod, Node, Service 等 APi 对象。比如一个待调度的 Pod 被创建出来后，调度器就会通过 Pod Informer 的 Handler，将这个待调度 Pod 添加进调度队列。

默认情况下，Kubernetes 的调度队列是一个 PriorityQueue（优先级队列），并且当某些集群信息发生变化的时候，调度器还会对调度队列里的内容进行一些特殊操作。这里的设计，主要是出于调度优先级和抢占的考虑。

此外，Kubernetes 的默认调度器还要负责对调度器缓存（即：scheduler cache）进行更新。事实上，Kubernetes 调度部分进行性能优化的一个最根本原则，就是尽最大可能将集群信息 Cache 化，以便从根本上提高 Predicate 和 Priority 调度算法的执行效率。



**Scheduling Path**

Scheduling Path 的主要逻辑，就是不断地从调度队列里出队一个 Pod。然后，调用 Predicates 算法进行“过滤”。这一步“过滤”得到的一组 Node，就是所有可以运行这个 Pod 的宿主机列表。当然，Predicates 算法需要的 Node 信息，都是从 Scheduler Cache 里直接拿到的，这是调度器保证算法执行效率的主要手段之一。

调度器就会再调用 Priorities 算法为上述列表里的 Node 打分，分数从 0 到 10。得分最高的 Node，就会作为这次调度的结果。

调度算法执行完成后，调度器就需要将 Pod 对象的 nodeName 字段的值，修改为上述 Node 的名字。这个步骤在 Kubernetes 里面被称作 Bind。

但是，为了不在关键调度路径里远程访问 APIServer，Kubernetes 的默认调度器在 Bind 阶段，只会更新 Scheduler Cache 里的 Pod 和 Node 的信息。这种基于“乐观”假设的 API 对象更新方式，在 Kubernetes 里被称作 Assume。

Assume 之后，调度器才会创建一个 Goroutine 来异步地向 APIServer 发起更新 Pod 的请求，来真正完成 Bind 操作。如果这次异步的 Bind 过程失败了，也不会有太大的影响，等 Scheduler Cache 同步之后一切就会恢复正常。

因为当一个新的 Pod 完成调度，并需要再某个节点上运行起来时，kubelet 还会做一个 Admit 的验证操作来确保这个 Pod 是否能运行再这个节点上。



**无锁化**

在 Scheduling Path 上，调度器会启动多个 Goroutine 以节点为粒度并发执行 Predicates 算法，从而提高这一阶段的执行效率。而与之类似的，Priorities 算法也会以 MapReduce 的方式并行计算然后再进行汇总。而在这些所有需要并发的路径上，调度器会避免设置任何全局的竞争资源，从而免去了使用锁进行同步带来的巨大的性能损耗。

所以，在这种思想的指导下，如果你再去查看一下前面的调度器原理图，你就会发现，Kubernetes 调度器只有对调度队列和 Scheduler Cache 进行操作时，才需要加锁。而这两部分操作，都不在 Scheduling Path 的算法执行路径上。



**Predicates 调度策略**

Predicates 可以当作一个"过滤"策略，筛选出一部分符合条件的节点，这些节点都是可以运行待调度 Pod 的宿主机。

默认的调度策略有四种：



GeneralPredicates 规则

过滤规则，负责的是最基础的调度策略。

比如，PodFitsResources 计算的就是宿主机的 CPU 和内存资源等是否够用。

PodFitsHostPorts 检查：Pod 申请的宿主机端口（spec.nodePort）是不是跟已经被使用的端口有冲突。

PodMatchNodeSelector 检查：Pod 的 nodeSelector 或者 nodeAffinity 指定的节点，是否与待考察节点匹配，等等。

PS：Admit 操作做的二次确认，就是执行一遍 GeneralPredicates



Volume 过滤规则

负责的是跟容器持久化 Volume 相关的调度策略。

NoDiskConflict 检查：多个 Pod 声明挂载的持久化 Volume 是否有冲突。比如，AWS EBS 类型的 Volume，是不允许被两个 Pod 同时使用的。所以，当一个名叫 A 的 EBS Volume 已经被挂载在了某个节点上时，另一个同样声明使用这个 A Volume 的 Pod，就不能被调度到这个节点上了。

MaxPDVolumeCountPredicate 检查：判断一个节点上某种类型的持久化 Volume 是不是已经超过了一定数目，如果是的话，那么声明使用该类型持久化 Volume 的 Pod 就不能再调度到这个节点了。

VolumeZonePredicate 检查：持久化 Volume 的 Zone（高可用域）标签，是否与待考察节点的 Zone 标签相匹配。

VolumeBindingPredicate 检查：该 Pod 对应的 PV 的 nodeAffinity 字段，是否跟某个节点的标签相匹配。

**在 Predicates 阶段，Kubernetes 就必须能够根据 Pod 的 Volume 属性来进行调度。**

如果该 Pod 的 PVC 还没有跟具体的 PV 绑定的话，调度器还要负责检查所有待绑定 PV，当有可用的 PV 存在并且该 PV 的 nodeAffinity 与待考察节点一致时，这条规则才会返回“成功”



宿主机过滤规则

负责判断调度 Pod 是否满足 Node 本身的某些条件。

PodToleratesNodeTaints 检查：Node 的“污点”机制。

NodeMemoryPressurePredicate 检查：当前节点的内存是不是已经不够充足，如果是的话，那么待调度 Pod 就不能被调度到该节点上。



Pod 过滤规则

这一组规则，跟 GeneralPredicates 大多数是重合的。

PodAffinityPredicate 检查：待调度 Pod 与 Node 上的已有 Pod 之间的亲密（affinity）和反亲密（anti-affinity）关系。



上面这四种类型的 Predicates，构成了调度器确定一个 Node 可以运行待调度 Pod 的基本策略。

**在具体执行的时候， 当开始调度一个 Pod 时，Kubernetes 调度器会同时启动 16 个 Goroutine，来并发地为集群里的所有 Node 计算 Predicates，最后返回可以运行这个 Pod 的宿主机列表。**



**Priorities** **调度策略**

得到节点列表后，Priorities 阶段的工作就是为这些节点打分。这里打分的范围是 0-10 分，得分最高的节点就是最后被 Pod 绑定的最佳节点。

Priorities 打分规则可以简单总结为：**选择空闲资源（CPU 和 Memory）最多的宿主机。**

选择的，其实是调度完成后，所有节点里各种资源分配最均衡的那个节点，从而避免一个节点上 CPU 被大量分配、而 Memory 大量剩余的情况。

在默认 Priorities 里，还有一个叫作 ImageLocalityPriority 的策略。它是在 Kubernetes v1.12 里新开启的调度规则，即：如果待调度 Pod 需要使用的镜像很大，并且已经存在于某些 Node 上，那么这些 Node 的得分就会比较高。

为了避免算法引发调度堆叠，调度器在计算得分的时候还会根据镜像的分布进行优化，即：如果大镜像分布的节点数目很少，那么这些节点的权重就会被调低，从而“对冲”掉引起调度堆叠的风险。

对于比较复杂的调度算法来说，比如 PodAffinityPredicate，它们在计算的时候不只关注待调度 Pod 和待考察 Node，还需要关注整个集群的信息，比如，遍历所有节点，读取它们的 Labels。这时候，Kubernetes 调度器会在为每个待调度 Pod 执行该调度算法之前，先将算法需要的集群信息初步计算一遍，然后缓存起来。这样，在真正执行该算法的时候，调度器只需要读取缓存信息进行计算即可，从而避免了为每个 Node 计算 Predicates 的时候反复获取和计算整个集群的信息。



# 优先级与抢占机制

优先级和抢占机制，解决的是 Pod 调度失败后的应对问题。

正常情况下，当一个 Pod 调度失败后，它就会被暂时“搁置”起来，直到 Pod 被更新，或者集群状态发生变化，调度器才会对这个 Pod 进行重新调度。

但是在某种应用场景中，也会希望一个高优先级的 Pod 调度失败后，不会被“搁置”，而是占用一些低优先级 Pod 的资源，让高优先级的 Pod 调度成功。

> 要使用这个机制，需要在 Kubernetes 里提交一个 PriorityClass 的定义

```yaml
# PriorityClass 创建PriorityClass对象
apiVersion: scheduling.k8s.io/v1beta1
kind: PriorityClass
metadata:
  name: high-priority
value: 1000000
globalDefault: false
description: "This priority class should be used for high priority service pods only."


# pod使用PriorityClass
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  priorityClassName: high-priority

```

调度器里维护着一个调度队列。所以，当 Pod 拥有了优先级之后，高优先级的 Pod 就可能会比低优先级的 Pod 提前出队，从而尽早完成调度过程。这个过程，就是“优先级”这个概念在 Kubernetes 里的主要体现。

当一个高优先级的 Pod 调度失败的时候，调度器的抢占能力就会被触发。调度器就会试图从当前集群里寻找一个节点，使得当这个节点上的一个或者多个低优先级 Pod 被删除后，待调度的高优先级 Pod 就可以被调度到这个节点上。这个过程，就是“抢占”这个概念在 Kubernetes 里的主要体现。

当发生抢占时，高优先级的 Pod 不会立即被调度到被抢占到 Node上，调度器只会将将抢占者的 spec.nominatedNodeName 字段，设置为被抢占的 Node 的名字。然后，抢占者会重新进入下一个调度周期，然后在新的调度周期里来决定是不是要运行在被抢占的节点上。这当然也就意味着，即使在下一个调度周期，调度器也不会保证抢占者一定会运行在被抢占的节点上。

这是因为调度器只会通过标准的 DELETE API 来删除被抢占的 Pod，这些 Pod 必然是有一定的“优雅退出”时间（默认30s）而在这段时间里，其他的节点也是有可能变成可调度的，或者直接有新的节点被添加到这个集群中来。所以，鉴于优雅退出期间，集群的可调度性可能会发生的变化，把抢占者交给下一个调度周期再处理，是一个非常合理的选择。

而在抢占者等待被调度的过程中，如果有其他更高优先级的 Pod 也要抢占同一个节点，那么调度器就会清空原抢占者的 spec.nominatedNodeName 字段，从而允许更高优先级的抢占者执行抢占，并且，这也就使得原抢占者本身，也有机会去重新抢占其他节点。这些，都是设置 nominatedNodeName 字段的主要目的。





# 学习资料

《Kubernetes 权威指南》

《深入剖析 Kubernetes》
