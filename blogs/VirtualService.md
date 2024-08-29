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

架构图来自istio官网：

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
kubectl label namespace test istio-injection=enabled

#准备一个Nginx的Pod及Service
---
apiVersion: v1
kind: Pod
metadata:
  name: demo-nginx
  labels:
    app.kubernetes.io/name: demo-nginx
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
      - containerPort: 80
        name: http-web-svc

---
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  selector:
    app.kubernetes.io/name: demo-nginx
  ports:
  - name: name-of-service-port
    protocol: TCP
    port: 80
    targetPort: http-web-svc
    
    
---
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: nginx-gateway
spec:
  selector:
    istio: ingressgateway # use istio default controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - '*'
    
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: nginx
spec:
  hosts:
  - '*'
  gateways:
  - nginx-gateway
  http:
  - match:
    - uri:
        exact: /
    route:
    - destination:
        host: nginx-service
        port:
          number: 80
          
```

为了方便演示，可以直接将 istio 的 ingressgateway 更改为 NodePort 模式

```shell
kubectl edit svc istio-ingressgateway -n istio-system
```

当通过curl 命令请求 istio-gateway 地址就变成了 `curl IP + {ingressgatewayProt}` 

```html
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html
```

上述虚拟服务中包含了[路由规则]，这是用来描述匹配条件和路由行为，当匹配到 `/` 时会 destination 按照虚拟服务的路由规则像指定的 `Service` 转发流量。这里要注意的是：destination 的 host 必须是存在于 istio 服务注册中心的实际目标地址。

路由的规则是从上到下的顺序进行优先级排序，这点很像代码，一般情况下在构建路由规则时都会提供一个默认的规则，以确保流量在经过虚拟服务时至少能匹配到一条路由规则。

更多的路由规则可以在[[官方文档](https://istio.io/latest/zh/docs/reference/config/networking/virtual-service/#HTTPMatchRequest)]中找到。

在代理前端服务时，比较常见的是需要虚构出一个动态路由，在 istio 中可以通过配置反向代理对进入服务的 URL 进行重写：

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: nginx
spec:
  gateways:
  - nginx-gateway
  hosts:
  - '*'
  http:
  - match:
    - uri:
        prefix: /v1/
    - uri:
        prefix: /v2/
    rewrite:
      uri: /
    route:
    - destination:
        host: nginx-service
        port:
          number: 80
```

使用 rewirte 中的 uri 替换 match 中的 uri 这步操作是强绑定的，只要匹配到符合的路由规则都会重定向到 / ，在 nginx-ingress 实现相同的目标要更为复杂，需要在 Ingress 配置中加上rewrite 的配置，在Path中还要编写正则进行路由的匹配。

```yaml
 nginx.ingress.kubernetes.io/configuration-snippet: |
      rewrite ^(/v1)$ $1/ redirect;
      nginx.ingress.kubernetes.io/rewrite-target: /$2
      nginx.ingress.kubernetes.io/use-regex: "true"
     
     
...

path: /v1(/|$)(.*)
pathType: Prefix
```

在验证上述配置时，往往会遇到 `404 503` 等错误，在调试时会相对复杂，当然也不是没有对应的[解决办法](https://www.danielhu.cn/2404-k8s-ingress-debug)。

Istio 中包含丰富的网络相关的策略，比如随机，按照百分比权重转发，最少被访问的实例。利用 istio 提供的强大能力和灵活性，可以覆盖相当多的实际场景。

网关既可以用于管理进入的流量，在流入服务网格后也可以配置出口网关，限制哪些服务可以访问外部网络，哪些不能。istio 还能配置各种安全方面的配置，比如Jwt，TLS 加密，认证，授权。

在测试方面通过故障注入，延迟，超时，熔断对应用等容忍度进行验证，提前预知局部故障的影响范围。



### Ambient 模式

不同于 Sidecar 模式，Ambient 模式下不需要在 Pod 中运行代理才能参与到服务网格的治理中，Ambient 模式也被称为 “无 Sidecar网格”，虽然这个称呼并不官方，但也能透露出不同之处。

Ambient 模式下不再使用 Envoy 进行代理，而是使用 ztunnel (Zero Trust tunnel，零信任隧道)， 这个组件不再挂在到每一个 Pod 作为“边车”服务，它在 Kubernetes 集群中的每一个节点中进行代理。

ztunnel 代理是用 Rust 编写的， 旨在处理 L3 和 L4 功能，例如 mTLS、身份验证、L4 鉴权和遥测。 

在 istio 最新的版本中可以在没有 sidecar 的情况下运行，sidecar 提供了服务网格的最初模式，在过去十几年的云原生发展中取得巨大成功，但不可避免的是：sidecar 在集群环境中被滥用，并且带来开销，当集群规模较大时，这些 sidecar 将不可忽视地算在成本的一部分。



### 后记

云原生中 istio Kubernetes Prometheus 这些技术已经成为某一领域的事实基础，服务网格是很好的技术，它解决了微服务架构下网络治理的问题，并且它还在不断进步，潜在提高Kubernetes 集群的性能。

从 istio 模式变迁中也能看出，在基础设施和系统级编程中，Rust 这门号称最安全语言的身影出现的越来越频繁，支持的模式也越来越多元化。

从TCP/IP 到服务网格，我们从中能看到技术的变迁，最后我们用《深入高可用架构原理与实践》中的一句话来结尾：

> 云原生中大部分技术栈，无论是 ServiceMesh 还是 Serverless（无服务器计算）等等，虽然它们各个维度、领域不同，但核心都是将非功能逻辑从应用中剥离，让业务开发更简单。这也是本书反反复复提及的以“以应用为中心”的软件设计理念。



### 学习资料

https://www.thebyte.com.cn/ServiceMesh/What-is-ServiceMesh.html

https://istio.io/latest/zh/docs/ops/deployment/architecture

https://blog.fleeto.us/post/istio-route-rules

https://www.danielhu.cn/2404-k8s-ingress-debug

https://mp.weixin.qq.com/s/nM1qPh9ZxpLV92vbl-pKrA
