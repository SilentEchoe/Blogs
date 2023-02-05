---
title: Kubernetes Operator 开发
date: 2023-2-4 22:52:00
tags: [Kubernetes,学习笔记,Operator开发]
category: Kubernetes
---



在 Kubernetes 上运行工作负载的人们都喜欢通过自动化来处理重复的任务。 Operator 模式会封装我们编写的（Kubernetes 本身提供功能以外的）任务自动化代码。

> **Operator 通过扩展 Kubernetes 控制平面和 API 进行工作。Operator 将一个 endpoint（称为自定义资源 CR）添加到 Kubernetes API 中，该 endpoint 还包含一个监控和维护新类型资源的控制平面组件。**

Kubernetes Operator 扩展开源项目有很多，但本文只涉及 Operator-SDK 和 Kubebuilder

Kubebuilder 是一个基于 [CRD](https://jimmysong.io/kubernetes-handbook/concepts/crd.html) 来构建 Kubernetes API 的框架，可以使用 [CRD](https://jimmysong.io/kubernetes-handbook/GLOSSARY.html#crd) 来构建 API、Controller 和 Admission Webhook。

```
# Mac 安装 Kubebuilder
brew install Kubebuilder

mkdir $GOPATH/src/projectName
cd $GOPATH/src/projectName
#使用 demo.kubebuilder.io 域，所有的 API 组将是<group>.demo.kubebuilder.io.

#创建项目
kubebuilder init --domain demo.kubebuilder.io

#创建API
kubebuilder create api --group batch --version v1 --kind CronJob

```

