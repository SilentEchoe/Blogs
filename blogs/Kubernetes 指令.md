---
title: Kubernetes 常用指令
date: 2022-2-17 15:21:00
tags: [Kubernetes,学习笔记,指令]
category: Kubernetes
---



# Kubernetes 常用指令



## 查看集群信息

```bash
kubectl cluster-info
```



## 部署应用

```bash
kubectl run kubernetes-bootcamp --image=docker.io/jocatalin/kubernetes-bootcamp:v1 --port=8080

kubectl create deployment kubernetes-bootcamp --image=docker.io/jocatalin/kubernetes-bootcamp:v1 --port=8080

#创建一个node 应用 在 deployment下
kubectl create deployment node-hello --image=gcr.io/google-samples/node-hello:1.0 --port=8080
```

通过 kubectl 部署了一个应用，命名它为 kubernetes-bootcam。



## 查看当前的Pod

```bash
kubectl get pods
```



## 访问应用/更改端口映射

默认情况下，所有Pod 只能在集群内部访问，为了能从外部访问应用，需要将容器的8080 端口映射到节点的端口。

```bash
#pods 为命名空间下 不同的应用创建
kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
```



## Scale 应用

执行命令将副本数量增加到3个

```bash
kubectl scale deployments/node-hello --replicas=3
```



# 排错命令

```bash
kubectl describe pods ${POD_NAME}
```



## 获取命名空间下的所有 Pods

```bash
kubectl get pods --namespace name

示例：
kubectl get pods --namespace zadig
```



## 查询状态异常的 Pod

```go
kubectl get pod -n <namespace>

# 异常状态 pod 的详细信息
kubectl desribe pod <podName> -n <namespace>

# 查看日志
kubectl logs -f <pod/podName> -n <namespace>
```

