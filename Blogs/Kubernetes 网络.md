---
title: Kubernetes 网络
date: 2022-1-17 20:35:00
tags: [Kubernetes,学习笔记,网络]
category: Kubernetes
---



## Kubernetes 网络

“网络栈” 包括了：网卡（Network Interface），回环设备（Loopback Device）,路由表（Routing Table）和 iptables 规则。

容器可以声明直接使用宿主机的网络栈，比如:

```go
$ docker run -d -net=host --name nginx-host nginx
```

容器会直接监听宿主机的 80 端口。

> 直接使用宿主机网络栈的方式，虽然可以为容器提供良好的网络性能，但可能会引起其他的问题，比如端口冲突。 **在大多数情况下，我们希望容器进程使用自己 Network Namespace 里的网络栈，拥有属于自己的 IP 地址和端口。**

在Linux 中，网桥会起到虚拟交换机作用的网络设备。它是一个工作数据链路层的设备，主要功能是根据 MAC 地址学习来将数据包转发到网桥的不同端口上。

Docker 会默认在宿主机上创建一个叫 docker0 的网桥，凡是连接在 docker0 网桥上的容器，可以通过它来进行通信。



** Veth Pair 虚拟设备**

它被创建出来后，总是以两张虚拟网卡（Veth Peer）的形式成对出现的。并且，从其中一个“网卡”发出的数据包，可以直接出现在与它对应的另一张“网卡”上，哪怕这两个“网卡”在不同的 Network Namespace 里。

它常常被用作连接不同 Network Namespace 的 “网线”。

```go
// 可以看到网桥信息
$ brctl show  
```



## 学习资料

《深入剖析Kubernetes》