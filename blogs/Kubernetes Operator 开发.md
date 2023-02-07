---
title: Kubernetes Operator 开发
date: 2023-2-4 22:52:00
tags: [Kubernetes,学习笔记,Operator开发]
category: Kubernetes
---



在 Kubernetes 上运行工作负载的人们都喜欢通过自动化来处理重复的任务。 Operator 模式会封装我们编写的（Kubernetes 本身提供功能以外的）任务自动化代码。

> **Operator 通过扩展 Kubernetes 控制平面和 API 进行工作。Operator 将一个 endpoint（称为自定义资源 CR）添加到 Kubernetes API 中，该 endpoint 还包含一个监控和维护新类型资源的控制平面组件。**
>
> **Operator 由一组监听 Kubernetes 资源的 Controller 组成。Controller 可以实现调协（reconciliation loop），另外每个 Controller 都负责监视一个特定资源，当创建、更新或删除受监视的资源时就会触发调协。**

Kubernetes 中，有一组内置的 Controller 在主节点中的 controller-manager 内部运行。比如：

Deployment  ReplicaSet  DaemonSet  Service  Job  CronJob Endpoint  StatefulSet 等





Kubernetes Operator 扩展开源项目有很多，但本文只涉及 Operator-SDK 和 Kubebuilder



**Operator-SDK**

```shell
# Mac 安装 operator-sdk
brew install Kubebuilder

operator-sdk version

$ operator-sdk version: "v1.23.0", commit: "1eaeb5adb56be05fe8cc6dd70517e441696846a4", kubernetes version: "v1.24.2", go version: "go1.19", GOOS: "darwin", GOARCH: "arm64"
```



### **Kubebuilder**

**Controller-runtime** 是一个用于开发 Kubernetes Controller 的库，包含了各种Controller 常用的模块。后来社区推出了 `Kubebuilder` 来渲染 Controller 的整个框架，而 Kubebuilder 渲染出的框架使用的就是 Controller-runtime。

Controller-runtime中为Controller的开发提供了各种功能模块，主要包括：

- `Client`：用于读写Kubernetes资源
- `Cache`：本地缓存，可供Client直接读取资源。
- `Manager`：可以管理协调多个Controller，提供Controller共用的依赖。
- `Controller`："组装"多个模块（例如 Source ,Queue , Reconciler），实现Kubernetes Controller 的通用逻辑
  - 1）监听k8s资源，缓存资源，并根据`EventHandler`入队事件；
  - 2）启动多个goroutine，每个goroutine会从队列中获取event，并调用`Reconciler`方法处理。
- `Reconciler`：状态同步的逻辑所在，是开发者需要实现的主要接口，供Controller调用。Reconciler的重点在于“状态同步”，由于Reconciler传入的参数是资源的`Namespace`和`Name`，而非event，Reconciler并非用于“处理事件”，而是根据指定资源的状态，来同步“预期集群状态”与“当前集群状态”。
- `Webhook`：用于开发webhook server，实现Kubernetes Admission Webhooks机制。
- `Source`：source of event，Controller从中获取event。
- `EventHandler`：顾名思义，event的处理方法，决定了一个event是否需要入队列、如何入队列。
- `Predicate`：相当于event的过滤器。





**Kubebuilder 开发**

```shell
# Mac 安装 Kubebuilder
brew install Kubebuilder

mkdir $GOPATH/src/projectName
cd $GOPATH/src/projectName
#使用 demo.kubebuilder.io 域，所有的 API 组将是<group>.demo.kubebuilder.io.

#创建项目
kubebuilder init --domain demo.kubebuilder.io

#创建API(APP)
#gourp（资源组）
kubebuilder create api --group webapp --version v1 --kind Demo
```



```shell
#启动服务
cd $GOPATH/src/projectName
make run
```

![image-20230206150437195](/Users/kai/Library/Application Support/typora-user-images/image-20230206150437195.png)
