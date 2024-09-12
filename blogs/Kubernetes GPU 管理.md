---
title: Kubernetes GPU 管理
date: 2024-9-06 13:49:00
tags: [Kubernetes]
category: Kubernetes
---

### 前言

2022年5月13日，英伟达发布了Linux开源[GPU内核模块](https://github.com/NVIDIA/open-gpu-kernel-modules)，支持数据中心所用的GPU和消费级显卡，这意味着开发者可以通过代码而观察到内核驱动是如何工作的，同时还可以将 [NVIDIA](https://www.nvidia.com/en-us/) 驱动直接安装在企业内部的服务器上。

处于AI浪潮的大背景下，背后的推手不言而喻，长期以来 NVIDIA 一直以高性能 GPU 和闭源驱动程序而闻名。闭源一直是 Linux 社区和其他开源社区所厌恶的，这场由 Linus 与 NVIDIA 长达十年的冲突在开源内核模块后似乎画上了句号。

看似开源社区取得了胜利，但 CUDA 作为 GPU 的核心依然是闭源的。这个由 NVIDIA 推出的并行计算架构和编程模型，它允许开发者使用 C/C++ 直接在 NVIDIA 的 GPU 上进行通用计算，通过 CUDA 开发者可以更加高效利用 GPU 的并行计算能力，加速各种计算密集型任务……CUDA 的存在让 GPU 资源变成一种通用的计算资源。这几乎是 NVIDIA 在软件领域上的核心，从目前的趋势和 CEO 的态度来看，几乎不存在开源 CUDA 的可能。

但无论如何，至少开源社区拿到了 GPU 内核模块的源码。





### Kubernetes GPU 管理

从2016年起，Kubernetes 社区中希望能在集群中对 GPU 硬件加速设备进行管理的声音越来越大，对于云用户来说，他们需要在 Pod 中能访问到 GPU 设备以及驱动目录。



