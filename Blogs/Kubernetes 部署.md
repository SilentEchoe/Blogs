---
title: Kubernetes 安装
date: 2021-7-3 19:03:00
tags: [Kubernetes,学习笔记]
category: Kubernetes 	
---



## Docker 安装

```bash
sudo apt-get update
sudo apt-get install docker.io
```



## K8s 安装

```bash
# 使得 apt 支持 ssl 传输
apt-get update && apt-get install -y apt-transport-https
# 下载 gpg 密钥
curl <https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg> | apt-key add - 
# 添加 k8s 镜像源
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb <https://mirrors.aliyun.com/kubernetes/apt/> kubernetes-xenial main
EOF
# 更新源列表
apt-get update
# 下载 kubectl，kubeadm以及 kubelet
apt-get install -y kubelet kubeadm kubectl
```



## 关闭Swap

```bash
$ vim /etc/fstab
# UUID=9224d95f-cd87-4b56-b249-3dc7de4491d3 none            swap    sw              0       0
```



## Kubeadm 使用阿里云拉取需要的镜像

```bash
kubeadm config images pull --image-repository=registry.aliyuncs.com/goole_containers
```

PS:这里拉取 coredns 会有问题，所以直接拉取官方的镜像。

```bash
docker pull coredns/coredns:1.8.0
```

然后修改镜像的tag

```bash
sudo docker tag 296a6d5035e2 registry.aliyuncs.com/google_containers/coredns/coredns:V1.8.0
```

先查看 kubeadm 需要的镜像

```bash
kubeadm config images list
```

**PS：可以先从阿里云等国内源拉取需要的镜像，然后更改Tag。**



## Kubeadm 初始化

```bash
kubeadm init --apiserver-advertise-address 192.168.56.105 --pod-network-cidr=10.244.0.0/16
```

配置 Kubectl

```bash
su - username
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

echo "source <(kubectl completion bash)" >> ~/.bashrc
```



### The Connection to the server [localhost:8080](http://localhost:8080) was refused did you specify the right host or port?

解决方案：

Kubernetes master 没有与本机绑定，集群初始化的时候没有绑定，此时设置在本机的环境变量可解决

```bash
sudo su

echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /etc/profile

soure /etc/profile
```



## 安装Pod 网络

```bash
kubectl apply -f <https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml>

# 无外网可使用以下命令
kubectl apply -f [<https://github.com/flannel-io/flannel/blob/master/Documentation/kube-flannel.yml>](<https://github.com/flannel-io/flannel/blob/master/Documentation/kube-flannel.yml>)
```



## 添加子节点

```bash
#在master 节点上查询Token
kubeadm token list 

#Token 默认存在24小时，如果不存在可以自己创建一个新token
#kubeadm token create

#添加节点
kubeadm join --token sq10bo.yklylpvoavkbg9om Ip:6443
```

可能遇见问题：

discovery.bootstrapToken: Invalid value: "": using token-based discovery without caCertHashes can be unsafe. Set unsafeSkipCAVerification to continue

解决方案：

如果出现以下错误，说明需要进行ca校验可以使用--discovery-token-unsafe-skip-ca-verification参数忽略校验

```bash
kubeadm join --token tokenId Ip:6443 --discovery-token-unsafe-skip-ca-verification
```

遇见问题

```bash
[ERROR Port-10250]: Port 10250 is in use
        [ERROR DirAvailable--etc-kubernetes-manifests]: /etc/kubernetes/manifests is not empty
        [ERROR FileAvailable--etc-kubernetes-pki-ca.crt]: /etc/kubernetes/pki/ca.crt already exists
        [ERROR FileAvailable--etc-kubernetes-kubelet.conf]: /etc/kubernetes/kubelet.conf already exists
```

解决办法：

使用虚拟机时，子节点加入集群时，可能会因为本身克隆的Master节点的镜像，照成HostName 一致，所以需要更改HostName

```bash
#修改以下两个文件，然后重启虚拟机
vi /etc/hostname
vi /etc/hosts
kubeadm reset
```





## 创建 Nginx 应用

```bash
kubeadm run nginx-deploy --image=nginx:1.14-alpine --port=80 --replicas=2

PS:在 K8S V1.18.0 版本后，不支持 --replicas =2 这种语法格式，已弃用。推荐使用 deployment 创建 Pods
```

使用 Yaml 文件来创建：

```bash
# API 版本号
apiVersion: apps/v1
# 类型，如：Pod/ReplicationController/Deployment/Service/Ingress
kind: Deployment
metadata:
  # Kind 的名称
  name: nginx-app
spec:
  selector:
    matchLabels:
      # 容器标签的名字，发布 Service 时，selector 需要和这里对应
      app: nginx
  # 部署的实例数量
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx
    spec:
      # 配置容器，数组类型，说明可以配置多个容器
      containers:
      # 容器名称
      - name: nginx
        # 容器镜像
        image: nginx:1.17
        # 只有镜像不存在时，才会进行镜像拉取
        imagePullPolicy: IfNotPresent
        ports:
        # Pod 端口
        - containerPort: 80
kubectl apply -f nginx.yaml

#暴露服务地址
kubectl expose deployment nginx-app --port=80 --type=LoadBalancer
```