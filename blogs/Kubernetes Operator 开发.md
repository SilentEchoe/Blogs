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



### Controller-runtime

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

[![pSh9dyt.png](https://s1.ax1x.com/2023/02/10/pSh9dyt.png)](https://imgse.com/i/pSh9dyt)

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



#### **Manager**

Manager提供了Controller的依赖并控制Controller的运行,Manager是控制"可运行程序",只要实现`Runnable`接口都可以向Manager中注册。

可以注册一个 HttpServer 到Manager中,然后使用Manager来启动 HttpServer,只需要HttpServer实现对应的`Start`接口。

```go
type Runnable interface {
	Start(context.Context) error
}
```

上述代码中`Builder.Complete`调用`Manager.Add`实现注册服务。当服务被注册后,Manager 会根据`Runnable`是否受“选举机制”的影响,将其分类到 `leaderElectionRunnables`或`nonLeaderElectionRunnables`两个数组中,依据`Runnable`可能实现的`func NeedLeaderElection() bool`方法的返回值进行划分，未实现此方法的会被归类到`leaderElectionRunnables`中。

Manager会通过goroutine运行所有注册的Runnable,`nonLeaderElectionRunnables`会直接运行。

`leaderElectionRunnables`会根据选举结果运行，Manager的选举机制使用`k8s.io/client-go/tools/leaderelection`实现,可以通过创建manager时传入的`manager.Options`参数设置，其他更详细的实现可以查看`pkg/manager/internal.go`。

```
type Manager interface {
	// Cluster holds objects to connect to a cluster
	cluser.Cluster

	// Add will set requested dependencies on the component, and cause the component to be
	// started when Start is called.  Add will inject any dependencies for which the argument
	// implements the inject interface - e.g. inject.Client.
	// Depending on if a Runnable implements LeaderElectionRunnable interface, a Runnable can be run in either
	// non-leaderelection mode (always running) or leader election mode (managed by leader election if enabled).
	Add(Runnable) error

	// Elected is closed when this manager is elected leader of a group of
	// managers, either because it won a leader election or because no leader
	// election was configured.
	Elected() <-chan struct{}

	// SetFields will set any dependencies on an object for which the object has implemented the inject
	// interface - e.g. inject.Client.
	SetFields(interface{}) error

	// AddMetricsExtraHandler adds an extra handler served on path to the http server that serves metrics.
	// Might be useful to register some diagnostic endpoints e.g. pprof. Note that these endpoints meant to be
	// sensitive and shouldn't be exposed publicly.
	// If the simple path -> handler mapping offered here is not enough, a new http server/listener should be added as
	// Runnable to the manager via Add method.
	AddMetricsExtraHandler(path string, handler http.Handler) error

	// AddHealthzCheck allows you to add Healthz checker
	AddHealthzCheck(name string, check healthz.Checker) error

	// AddReadyzCheck allows you to add Readyz checker
	AddReadyzCheck(name string, check healthz.Checker) error

	// Start starts all registered Controllers and blocks until the Stop channel is closed.
	// Returns an error if there is an error starting any controller.
	// If LeaderElection is used, the binary must be exited immediately after this returns,
	// otherwise components that need leader election might continue to run after the leader
	// lock was lost.
	Start(<-chan struct{}) error

	// GetWebhookServer returns a webhook.Server
	GetWebhookServer() *webhook.Server
}
```



#### Cluster

`Cluster`提供各种封装的调用集群资源的接口,包括：

- 通过`Cluster.GetEventRecorderFor()`获取用于记录K8s Event的Recorder。

- 通过`Cluster.GetClient()`获取K8s的Client，用于读写。

- 通过`Cluster.GetCache()`获取后端的Cache。

  

`Cluster`同时也会为Controller的创建提供共同的依赖：

- 在`Controller.Watch()`中，会通过`Cluster.SetFields()`注入`cache`、`RESTMapper`等。
- 在`Builder.Complate()`中，会通过`Cluster.Scheme()`获取`Scheme`，从而获取`Source`对应的`GroupVersionKind`。

Cluster 主要使用`Client`和`Cache`来实现



##### Cache

一般情况下,开发者不会直接对Cache进行操作,Controller-runtime中,使用`InformerCache`实现`Cache`,也可以通过设置`Manager.Options.NewCache`参数,传入 Cache 的创建函数。

InormerCache 包含三组`specificInformersMap`，分别用于支持`structured`、`unstructured`、`metadata`三种资源类型，实现三类资源的List-watch。

`specificInformersMap`中包含了一个key为`GroupVersionKind`、value为`MapEntry`的Map类型，是为了支持多种资源的监听。

##### Client

Controller-runtime实现了多种Client：

`Manager.GetClient`返回的Client可以用于Get、Update、Patch、Create等多种操作，但在Get、List时，优先从`cache`中的读取；

`Manager.GetAPIReader`返回的Reader对象用于读操作，但会直接通过请求Kube-apiserver来获取结果。

Client 也支持三种资源类型：

```
type client struct {
	typedClient        typedClient
	unstructuredClient unstructuredClient
	metadataClient     metadataClient
	...
}
```



### **Kubebuilder 开发示例**

```shell
# Mac 安装 Kubebuilder
brew install Kubebuilder

mkdir $GOPATH/src/projectName
cd $GOPATH/src/projectName
#使用 demo.kubebuilder.io 域，所有的 API 组将是<group>.demo.kubebuilder.io.

#创建项目
kubebuilder init --domain app.kubebuilder.io --repo=github.com/AnAnonymousFriend/KubeMin-Cli --owner kai

#创建API(APP)
#gourp（资源组）
Kubebuilder create api apps --version v1 --kind Application

#更改Crd后需要输入命令生成
make manifests generate

#安装CRD
make install

#启动服务
cd $GOPATH/src/projectName
make run
```





### 学习资料

[controller-runtime](https://maao.cloud/2021/02/26/Kubernetes-Controller%E5%BC%80%E5%8F%91%E5%88%A9%E5%99%A8-controller-runtime)

[controller-runtime 源码分析](https://qiankunli.github.io/2020/08/10/controller_runtime.html)



