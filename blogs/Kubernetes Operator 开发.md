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

Kubernetes 中，有一组内置的 Controller 在主节点中的 controller-manager 内部运行。比如：Deployment  ReplicaSet  DaemonSet  Service  Job  CronJob Endpoint  StatefulSet 等



### **Controller-runtime **

**Controller-runtime** 是一个用于开发 Kubernetes Controller 的库，包含了各种Controller 常用的模块。而 Kubebuilder 渲染出的框架使用的就是 Controller-runtime, 在了解怎么使用 Kubebuilder 进行开发前，先对 Controller-runtime 进行一些了解，这会更好帮助我们开发 Controller

<div align="center">
    	<img src="https://s1.ax1x.com/2023/02/09/pSWQEcT.png">  
</div>
Controller-runtime 的整体架构图

<div align="center">
    	<img src="https://s1.ax1x.com/2023/02/10/pSfHkVI.png">  
</div>





**Builder 阶段**

`Builder` 用于生成 Controller 或 Webhook，通过链式调用可以组装出所需的 Controller ,一般情况下会先创建出一个 Manager 然后再创建 Controller :

![image-20230209152617781](/Users/kai/Library/Application Support/typora-user-images/image-20230209152617781.png)

在Kubebuilder 中，组装 Controller 的函数是 `SetupWithManager`

```go
// SetupWithManager sets up the controller with the Manager.
func (r *DemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Demo{}).
		Complete(r)
}
```

在上述代码片段中创建一个新的 Controller 并且使用 `For` （主要监听资源）去监听&webappv1.Demo（自定义资源）在`ControllerManagedBy`和`Complete` 之前是一系列对 Controller 的配置，比如可以使用 `Owns` 监听其他资源 (Pod)

```go
// SetupWithManager sets up the controller with the Manager.
func (r *DemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Demo{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
```

最后通过`Complete` 创建新的 Controller ，并将创建的 Controller 注册到 Manager 中。

当然除了上述的 `For` 和 `Owns` 函数以外,还能对 Controller 进行一些其他的配置, `WithEventFilter()`对Controller的事件进行过滤; 用`Named()`配置Controller的名称等;用`Watches()`配置其他需要监听的资源。



#### **Watches 函数**

Watches 函数包含三个参数：`Source`,`EventHandler`,`WatchesOption`

[![pSfOong.png](https://s1.ax1x.com/2023/02/10/pSfOong.png)](https://imgse.com/i/pSfOong)



1.`Source` 负责 watch 相应的资源，能将资源的 event（事件）发送到队列中。用于初始化 watch 操作所需的结构，比如 `eventHandler`, `queue`,实际开发中用于监听 Kubernetes 的`Source`实现：`Kind`。

Kind 不实现真正的 Watch 操作，而是通过 `cache` 来 Watch 指定的资源，Kind 的作用是将`eventHandler`,`queue` 等注册到 Cache 中。`Complete`执行时会调用 Controller.Watch 会自动调用 Manager 的`SetFields`方法注入`cache`。



2.`EventHandler`是一个处理各种 Event 的接口，可以实现 "对指定事件入队列的逻辑",它需要实现四个函数：

```go
type EventHandler interface {
	// Create is called in response to an create event - e.g. Pod Creation.
	Create(event.CreateEvent, workqueue.RateLimitingInterface)

	// Update is called in response to an update event -  e.g. Pod Updated.
	Update(event.UpdateEvent, workqueue.RateLimitingInterface)

	// Delete is called in response to a delete event - e.g. Pod Deleted.
	Delete(event.DeleteEvent, workqueue.RateLimitingInterface)

	// Generic is called in response to an event of an unknown type or a synthetic event triggered as a cron or
	// external trigger request - e.g. reconcile Autoscaling, or a Webhook.
	Generic(event.GenericEvent, workqueue.RateLimitingInterface)
}
```

 Controller-runtime 提供了四类 handler:

- `EnqueueRequestForObject`：一个简单的实现，直接将Object的metadata入队列。`For()`就使用的它。
- `Enqueue_mapped`：用的较多。Object在入队前使用用户实现的映射方法`MapFunc()`做映射，将映射后的结果入队列。例如可以将`Endpoint Event`映射为对应的Service，从而入队Service。
- `Enqueue_owner`：将Object的Owner资源入队列，`Owns()`使用的它。
- `Funcs`：一个空的父类，需要你实现接口的四个方法。



3.`WatchesOption`用于修改Watch配置,提供两种配置

- `Predicates`：用于过滤事件,或者使用预设参数。

  [![pSfXDCq.png](https://s1.ax1x.com/2023/02/10/pSfXDCq.png)](https://imgse.com/i/pSfXDCq)

- `OnlyMetadate`用于告诉Controller，只缓存Watch对象的Metadata数据，提升性能。





### **Kubebuilder 开发**

```shell
# Mac 安装 Kubebuilder
brew install Kubebuilder

mkdir $GOPATH/src/projectName
cd $GOPATH/src/projectName
#使用 demo.kubebuilder.io 域，所有的 API 组将是<group>.demo.kubebuilder.io.

#创建项目
kubebuilder init --domain demo.kubebuilder.io --owner kai

#创建API(APP)
#gourp（资源组）
kubebuilder create api --group webapp --version v1 --kind Demo

#更改Crd后需要输入命令生成
make manifests generate

#安装CRD
make install
```



```shell
#启动服务
cd $GOPATH/src/projectName
make run
```

![image-20230206150437195](/Users/kai/Library/Application Support/typora-user-images/image-20230206150437195.png)







### 学习资料



[controller-runtime](https://maao.cloud/2021/02/26/Kubernetes-Controller%E5%BC%80%E5%8F%91%E5%88%A9%E5%99%A8-controller-runtime)

[controller-runtime 源码分析](https://qiankunli.github.io/2020/08/10/controller_runtime.html)



