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





### **Kubebuilder**

**Controller-runtime** 是一个用于开发 Kubernetes Controller 的库，包含了各种Controller 常用的模块。而 Kubebuilder 渲染出的框架使用的就是 Controller-runtime, 在了解怎么使用 Kubebuilder 进行开发前，先对 Controller-runtime 进行一些了解，这会更好帮助我们开发 Controller

[![pSWQEcT.png](https://s1.ax1x.com/2023/02/09/pSWQEcT.png)](https://imgse.com/i/pSWQEcT)

Controller-runtime 的整体架构图

![Untitled](/Users/kai/Downloads/Untitled.png)



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



**Watches 函数**







**Kubebuilder 开发**

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
