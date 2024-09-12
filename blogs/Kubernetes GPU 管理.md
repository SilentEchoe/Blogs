---
title: Kubernetes GPU 管理
date: 2024-9-06 13:49:00
tags: [Kubernetes]
category: Kubernetes
---

### 前言

2022年5月13日，英伟达发布了Linux开源[GPU内核模块](https://github.com/NVIDIA/open-gpu-kernel-modules)，支持数据中心所用的GPU和消费级显卡，这意味着开发者可以通过代码而观察到内核驱动是如何工作的，同时还可以将 [NVIDIA](https://www.nvidia.com/en-us/) 驱动直接安装在企业内部的服务器上。

处于AI浪潮的大背景下，背后的推手不言而喻，长期以来 NVIDIA 一直以高性能 GPU 和闭源驱动程序而闻名。闭源一直是 Linux 社区和其他开源社区所厌恶的，这场由 Linus 与 NVIDIA 长达十年的冲突在开源内核模块后似乎画上了句号。

看似开源社区取得了胜利，但 CUDA 作为 GPU 的核心依然是闭源的。这个由 NVIDIA 推出的并行计算架构和编程模型，它允许开发者使用 C/C++ 直接在 NVIDIA 的 GPU 上进行通用计算，通过 CUDA 开发者可以更加高效利用 GPU 的并行计算能力，加速各种计算密集型任务……CUDA 的存在让 GPU 资源变成一种通用的计算资源。这是 NVIDIA 在软件领域上的核心，从目前的趋势和 CEO 的态度来看，几乎不存在开源 CUDA 的可能。

但无论如何，至少开源社区拿到了 GPU 内核模块的源码。



### Kubernetes GPU 管理

从2016年起，Kubernetes 社区中希望能在集群中对 GPU 硬件加速设备进行管理的声音越来越大，对于云用户来说，他们想在 Pod 中能访问到 GPU 设备和驱动目录，以便进行大规模计算。

阿里的魔搭社区依靠着扎实的虚拟化技术和GPU资源，阔绰地赠送上百小时的计算时长，但依靠的是云主机，在面对大规模计算时可能因为无法抢到有限的计算资源而陷入停摆。

早在2021年，OpenAI 分享了一篇文章《Scaling Kubernetes to 7,500 nodes》其中记录了他们 Kubernetes 集群规模的成长与经验的分享，这说明在21年前 Kubernetes 社区就已经解决了 GPU 硬件的挂载和管理。

那么，如果想要在集群内实现 GPU 设备的管理，需要使用哪些技术？

Linux 中 Cgroups 暴露出来的操作接口是文件系统，它以文件和目录的方式出现在 `/sys/fs/cgroup` 路径下，可以通过挂载的方式自行挂载 Cgroups，在这个文件夹下会包含 cpuset cpu memory 这样的子目录，这些子目录代表着可以被 Cgroups 所限制的资源种类。

根据上述推理，如果想要在容器内使用 NVIDIA 的 GPU 设备，那么这个容器必须挂载 GPU 的设备和驱动目录。

![image-20240912143357471](https://raw.githubusercontent.com/SilentEchoe/images/main/image-20240912143357471.png)

Kubernetes 在实现 GPU 设备时，直接设置容器的 CRI (Container Runtime Interface) 参数就可以通过 Volume 将 GPU 驱动信息挂载到容器内，这也是 Linux 系统的特点，一切皆文件，只要设置好参数挂载驱动目录后就能直接使用该设备。

NVIDIA 开源了一个名为[nvidia-container-toolkit](https://github.com/NVIDIA/nvidia-container-toolkit)的项目，包含了容器运行时库，用于自动配置容器使用 NVIDIA GPU ，这样可以不需要在启动容器时设置额外的参数。

![image-20240912150116210](https://raw.githubusercontent.com/SilentEchoe/images/main/image-20240912150116210.png)

容器化只是第一步，如果要将 Kubernetes 与 Docker 一起使用需要将 Docker 配置为 NVIDIA Container Runtime 的引用，并设置为默认运行时。docker 的 daemon.json 文件内容如下：

```json
{ 
  "default-runtime": "nvidia",
    "exec-opts": [
        "native.cgroupdriver=systemd"
    ],
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "100m"
    },
    "runtimes": {
        "nvidia": {
            "args": [],
            "path": "nvidia-container-runtime"
        }
    },
    "storage-driver": "overlay2"
}
```

在 Kubernetes 支持的 GPU 方案里所有对硬件加速设备进行管理的功能，都是通过 Device Plugin 插件来实现的。NVIDIA 会实现一个叫做 **[k8s-device-plugin](https://github.com/NVIDIA/k8s-device-plugin)** 的插件，以官方的调度示例为例：

```yaml
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: gpu-pod
spec:
  restartPolicy: Never
  containers:
    - name: cuda-container
      image: nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda10.2
      resources:
        limits:
          nvidia.com/gpu: 1 # requesting 1 GPU
  tolerations:
  - key: nvidia.com/gpu
    operator: Exists
    effect: NoSchedule
EOF
```

Device Plugin 会通过 Kubernetes 的 ListAndWatch API 定期向 kubelet 上报该 Node 上 GPU ID 信息，不会包含 GPU 设备的信息，这一点将插件和实际的显卡信息做了解耦，kubelet 通过双层缓存来维护这些 GPU 的 ID 列表，并通过 ListAndWatch API 定时更新。

当一个 Pod 想使用一个 GPU 时，开发者只需要在 resources 中添加 nvidia.com/gpu: 1 这样调度器会从缓存中查询符合条件的 Node 再对双重缓存里的 Gpu 数量减去相应的数量，至此完成 Pod 和 Node 的绑定。

