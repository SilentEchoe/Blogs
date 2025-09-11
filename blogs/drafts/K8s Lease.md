K8s Lease

分布式系统中通常需要**租约**(Lease);租约提供了一种机制来锁定共享资源并协调集合成员之间的活动。

在Kubernetes中，租约概念表示为 `coordination.k8s.io` [API 组](https://kubernetes.io/zh-cn/docs/concepts/overview/kubernetes-api/#api-groups-and-versioning)中的 [Lease](https://kubernetes.io/zh-cn/docs/reference/kubernetes-api/cluster-resources/lease-v1/) 对象，常用于类似节点心跳和组件级领导者选举等系统核心能力。



K8s使用 Lease API 将Kubelet节点心跳传递到 Kubernetes API 服务器。对于每个Node，在Kube-node-lease名字空间中都有一个具备匹配名称到 Lease 对象
