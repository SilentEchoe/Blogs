---
title: Kubernetes 术语解释
date: 2021-6-29 15:22:00
tags: [Kubernetes,学习笔记]
category: Kubernetes 	
---



## **什么是 Kubernetes**

Kubernetes是一个自动化的容器编排平台，它负责应用的部署，应用的弹性以及应用的管理，这些都是基于容器的。



# **Kubernetes 有如下几个核心的功能：**

服务的发现与负载均衡

容器的“调度”，把一个容器放到一个集群的某一台机器上，K8s帮助我们去做存储的编排，让存储的声明周期与容器的生命周期能有一个连接。

k8s 会对容器做自动化的恢复，如果在一个集群中，经常出现宿主机的问题，导致容器不可用，K8s会自动对不可用的容器进行恢复。

k8S 会帮助我们去做应用的自动发布和回滚，以及与应用相关的配置密文的管理。

批量执行Job类型任务

支持水平的伸缩



# **核心讲解**

**调度**

K8s 可以将用户提交的容器放到管理的集群中的某一台节点上去。“调度”是执行这项功能的组件，同时它会观察正在被调度的容器的大小，规格。



**自动修复**

K8s有一个节点健康检查的功能，它会检测集群中的所有宿主机，当宿主机本身出现故障时，或者某个软件出现故障时，这个节点健康检查会自动对它进行发现



**水平伸缩**

K8s 有业务负载检查的能力，它会监测业务上所承担的负载，如果这个业务本身的CPU利用率过高，或者响应时间过长，它可以对这个业务进行一次扩容。



# Kubernetes概念和术语

Kubernetes 使用共享网络将多个物理机或虚拟机汇集到一个集群中，在各服务器之间进行通信，该集群是配置 Kubernetes 的所有组件，功能和工作负载的物理平台。

集群中一台服务器作为 Master 节点，负责管理整个集群。



### Master

Master 主要职责是调度，决定将应用放在那里运行。Master 运行Linux 操作系统，可以是物理机或者虚拟机。为了实现高可用，可以运行多个 Master。

> Master 是集群的网关和中枢，例如追踪其他服务器健康状态，以最优方式调度工作负载，以及编排等任务

余下的其他机器用作 Worker Node ，它们是使用本地和外部资源接收和运行工作负载的服务器。

集群中的主机可以是物理服务器也可以是虚拟机。



### Node

Node 的职责是运行容器应用。Node 由 Master 管理，Node 负责监控并汇报容器的状态，同时根据 Master 的要求管理容器的生命。Node 运行在 Linux 操作系统上，可以是物理机或是虚拟机。



### Pod

**Pod 是 Kubernetes 调度的最小单位，容器的集合，一组紧密相关的容器放在一个Pod中，同一个Pod 中的所有容器共享IP地址和 Port空间，他们在一个 network namespace 中。同一Pod中的容器始终被一起调度。**



### Controller

Kubernetes 提供多种 Controller 来管理 Pod ，包括 Deployment，ReplicaSet，DaemonSet，StatefuleSet，Job等。

**Deployment：可以理解为应用。**



### Service

Kubernetes Service 定义了外界访问一组特定的 Pod 的方式。Service 有自己的 IP 和端口，Service 为 Pod 提供了负载均衡。

**Kubernetes 运行容器（Pod）与访问容器 （Pod）这两项任务分别由 Controller 和 Service 执行。**



### Cluster

Cluster 是计算，存储和网络资源的集合，Kubernetes 利用这些资源运行各种基于容器的应用。



### Namespace

Namespace 可以将一个物理的 Cluster 逻辑上划分成多个虚拟 Cluster，每个 Cluster 就是一个 Namespace。不同 Namespace 里的资源是完全隔离的。



### Ingress

Kubernetes 将Pod对象和外部网络环境进行了隔离，Pod 和 Service 等对象间的通信都是用其内部专用地址进行。除了 Service 之外，Ingress 也是实现一个通往集群内部通道，供给外部使用的实现方式。