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



### 持久卷

ConfigMap用来保存Key-Value的配置数据，这个数据可以在Pod里使用，ConfigMap跟Secrets类似，但是ConfigMap一般用来管理配置，并且不包敏感信息的字符串.

ConfigMap中的每个data项都会成为一个新文件，一般用来：

1.设置环境变量的值

2.在容器里面设置命令行参数

3.在数据卷里面创建config文件

```yaml
---
apiVersion: v1
kind: ConfigMap
metadata:
   name: demo-config
data:
  demodata: "TestDemoData"
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
      command: [ "/bin/sh", "-c", "env" ]
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

