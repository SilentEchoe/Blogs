---
title: Client-go 源码分析
date: 2023-2-12 18:56:00
tags: [Kubernetes,学习笔记,Operator开发]
category: Kubernetes
---

Client-go是与kube-apiserver通信的clients的具体实现。

### WorkQueue 源码分析

WorkQueue 一般使用延时队列实现,在`Resource Event Handlers`中完成将对象的key放入WorkQueue的过程，然后在自己的逻辑代码里从WorkQueue中消费这些key。

client-go主要有三个队列,分别为普通队列,延迟队列和限速队列,后一个队列以前一个队列的实现为基础,层层添加新功能。

#### 普通队列







### 学习资料

《Kubernetes Operator 开发进阶》

