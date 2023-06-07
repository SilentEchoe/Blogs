---
title: Kubernetes 存储
date: 2022-5-10 11:16:00
tags: [Kubernetes,学习笔记,存储]
category: Kubernetes
---



## Kubernetes 存储卷

Container中的文件在磁盘中是临时存放的，当容器崩溃时文件会丢失，Kubernetes 会重新启动容器，但容器会以一个全新的状态重启。

Docker 当中也有卷(Volume)的概念，但是Docker卷是磁盘或另外一个容器内的某个目录，其提供的功能非常有限。

相比之下，Kubernetes支持很多类型的卷，Pod 可以使用任意数目的卷类型。还分为了临时卷和持久卷，它们拥有不同的生命周期或不同的挂在方式，Kubernetes会销毁临时卷，但是不会销毁持久卷。对于持久卷，在容器重启期间数据都不会丢失。



### 持久卷(Persistent Volume)

持久卷是集群中的一块存储，由管理员事先制备，或者使用存储类(Storage Class)来动态制备。持久卷是集群资源，PV持久卷和普通的 Volume 一样，也是使用卷插件来实现的，它们拥有独立于任何使用PV的Pod的声明周期。

**持久卷申领（PersistentVolumeClaim，PVC）** 表达的是用户对存储的请求。概念上与 Pod 类似。 Pod 会耗用节点资源，而 PVC 申领会耗用 PV 资源。Pod 可以请求特定数量的资源（CPU 和内存）；同样 PVC 申领也可以请求特定的大小和访问模式。







ConfigMap用来保存Key-Value的配置数据，这个数据可以在Pod里使用，ConfigMap跟Secrets类似，但是ConfigMap一般用来管理配置，并且不包敏感信息的字符串.

ConfigMap中的每个data项都会成为一个新文件，一般用来：

1.设置环境变量的值

2.在容器里面设置命令行参数

3.在数据卷里面创建config文件

以官方文档为例：

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: special-config
  namespace: default
data:
  special.how: very
  special.type: charm
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: env-config
  namespace: default
data:
  log_level: INFO
```

Pod中使用ConfigMap

```
apiVersion: v1
kind: Pod
metadata:
  name: dapi-test-pod
spec:
  containers:
    - name: test-container
      image: gcr.io/google_containers/busybox
      command: [ "/bin/sh", "-c", "echo $(SPECIAL_LEVEL_KEY) $(SPECIAL_TYPE_KEY)" ]
      env:
        - name: SPECIAL_LEVEL_KEY
          valueFrom:
            configMapKeyRef:
              name: special-config
              key: special.how
        - name: SPECIAL_TYPE_KEY
          valueFrom:
            configMapKeyRef:
              name: special-config
              key: special.type
      envFrom:
        - configMapRef:
            name: env-config
  restartPolicy: Never
```

执行 `kubectl apply -f cm.yaml`会输出

![image-20230510162625449](https://raw.githubusercontent.com/AnAnonymousFriend/images/main/image-20230510162625449.png)
