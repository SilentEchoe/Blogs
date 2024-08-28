---
title: VirtualService
date: 2023-8-27 11:20:00
tags: [VirtualService]
category: VirtualService
---

### 什么是服务网格？

服务网格一词最早出现于 William Morgan 的博文《What’s a service mesh？And why do I need one?》

> 服务网格是一个专门用于处理服务通讯的基础设施层。它可以在云原生应用组成的复杂拓扑结构下进行可靠的请求传送。在实践中，它是一组和应用服务部署在一起的轻量级网络代理，对应用服务透明。

当TCP/IP协议出现以后，机器之间的网络通信不再是问题，网线蛛丝般蔓延开来，互联网的基础设施像自然资源一般存在。在计算机技术不断进步中，为了追寻更好的性能，可用性，稳定性。从单体架构到分布式架构再到微服务，随着系统规模不断扩大，技术门槛也在不断提高，服务注册，服务发现，负载均衡，熔断，限流等技术词汇的出现都是解决某一问题的副产物。

从开发者角度来说，非业务逻辑应该从应用中剥离开来，即使各种微服务框架中提供分布式解决方案，维护和升级依然会在某一时段内影响应用正常运行。随着容器技术应运而生，编排技术极大推动了云原生技术的发展，微服务数量爆发性增长，虽然将服务之间的通信早已下沉到基础设施，但网络拓扑变得更加复杂且难以观测。

服务网格在 Kubernetes 的基础上建立可编程的网络，网络不再是看不见摸不着又重要的基础设施，它与 Kubernetes 和传统工作负载配合使用，带来了通用的流量管理和可观测性，以及安全性。

于是工程师们能使用服务网格更轻松实现金丝雀部署，A/B测试，负载均衡，故障恢复…而且它并不局限于单个集群，就像官方描述的那样：在 Kubernetes 或 VM、多云、混合或本地上运行的服务都可以包含在单个网格中。



### Istio 

Istio 由Google、IBM 和 Lyft 创立，它与Kubernetes 和 Prometheus 等项目并列，是最受欢迎的服务网格解决方案，也是现在最快毕业的 CNCF 项目。

Istio 分为数据平面和控制平面，数据平面由 Envoy 组成，这个由 C++ 开发的高性能七层代理与 Nginx 的技术架构相似，代理服务用于控制微服务之间的网络通信，相当于给每个 Pod 分配一个代理；控制平面则用于管理/配置规则策略，这种方式可以让 Istio 完成细颗粒度的流量控制，故障注入，安全性策略和各种身份认证……

架构图如下：

![image-20240827143854780](https://raw.githubusercontent.com/SilentEchoe/images/main/image-20240827143854780.png)

Istio 提供两种数据平面的模式：

1.**Sidecar 模式**，它会与您在集群中启动的每个 Pod 一起部署一个 Envoy 代理，或者与在虚拟机上运行的服务一同运行。
2.**Ambient 模式**，使用每个节点的四层代理，并且可选地使用每个命名空间的 Envoy 代理来实现七层功能。

自 Istio 发布以来就基于 Sidecar 模式构建，从泛用性来说本篇会着重介绍 Sidecar 模式，因为该模式容易理解且经历过时间的验证。





### 虚拟服务和目标规则

虚拟服务和目标规则是 Istio 流量路由功能的核心模块。虚拟服务这一概念是为了增强流量管理的灵活性和有效性，目的就是为了解耦客户端请求的目标地址与实际响应的载体，这也符合计算机领域的一贯做法：增加抽象层来解耦。

试想一下，如果没有虚拟服务，Envoy 将会沿用 Kubernetes 的默认网络策略，即轮询的负载均衡策略分发请求。这会增加更多的成本，比如在A/B测试中可能要创建更多的 SVC 或 Ingress 以不同的路由进行测试；如果配置不用服务版本的流量百分比路由就更麻烦了，需要按照流量比例部署多个副本以达到其目的。

使用虚拟服务，可以为一个或多个主机名指定流量行为，在虚拟服务中使用路由规则就是告诉 Envoy 如何将流量传递到适合的目标，这点很像 Nginx 的 Conf 配置，但要更为强大。如果想要将流量路由到不同的版本，则可以配置百分比调用到某个目标，再通过逐步增加流量比例完成金丝雀发布，这对于用户来说是无感知的。流量路由完全独立于实例部署，这意味着工程师可以创建任意流量路由来实现更为复杂的操作，相比直接使用 Kubernetes 只支持实例缩放的流量分发，要简单得多。



```yaml
#命名空间标记
kubectl label namespace default istio-injection=enabled

#准备一个Nginx的Pod及Service
---
apiVersion: v1
kind: Service
metadata:
  name: demo-svc
spec:
  selector:
    app.kubernetes.io/name: demo
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: http-web-svc

---
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app.kubernetes.io/name: demo
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
      - containerPort: 80
        name: http-web-svc
---
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: demo-gateway
spec:
  selector:
    istio: ingressgateway # use istio default controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: demoinfo
spec:
  hosts:
  - "*"
  gateways:
  - demo-gateway
  http:
  - match:
    - uri:
        exact: /
    route:
    - destination:
        host: demopage
        port:
          number: 80
```











### 学习资料

https://www.thebyte.com.cn/ServiceMesh/What-is-ServiceMesh.html

https://istio.io/latest/zh/docs/ops/deployment/architecture/

