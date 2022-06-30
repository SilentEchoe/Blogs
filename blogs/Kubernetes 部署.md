---
title: Kubernetes 安装
date: 2021-7-3 19:03:00
tags: [Kubernetes,学习笔记]
category: Kubernetes 	
---



## Ubuntu 20.04 部署 Kubernetes

### 安装组件

```shell
apt-get update && apt-get install -y \
apt-transport-https ca-certificates curl software-properties-common gnupg2

# 先安装ssh
sudo  apt-get install   openssh-server
sudo /etc/init.d/ssh start
sudo ufw allow ssh

# Add Docker’s official GPG key:
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

#禁止交换内存
sudo swapoff  -a
sed -ri 's/.*swap.*/#&/' /etc/fstab

#安装 docker
sudo apt install docker.io
sudo systemctl enable docker
sudo systemctl status docker
sudo systemctl start docker

# Set up the Docker daemon
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

mkdir -p /etc/systemd/system/docker.service.d
systemctl daemon-reload
systemctl restart docker


# 安装 Kubernetes 必要组件
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -

sudo tee /etc/apt/sources.list.d/kubernetes.list <<EOF 
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
sudo systemctl enable kubelet && sudo systemctl start kubelet

#关闭防火墙
systemctl stop firewalld
systemctl disable firewalld

# master 节点安装
kubeadm init   --image-repository=registry.aliyuncs.com/google_containers  --apiserver-advertise-address=master_ipadree   --pod-network-cidr=10.244.0.0/16  --service-cidr=10.244.0.0/12

#安装好以后,看一下状态
kubectl get ns
kubectl get nodes
kubectl get pods --all-namespaces
```



### 可能出现问题

```shell
#问题一：8080端口未开
#解决方案：
echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile
source ~/.bash_profile

#问题二：
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
#解决方案：
vim /usr/lib/systemd/system/docker.service
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock --exec-opt native.cgroupdriver=systemd

sudo systemctl daemon-reload
sudo systemctl restart docker

#问题三：
node  Failed to get imageFs info: non-existent label "docker-images"
解决方案：
重启一下 k8s

#问题五
The connection to the server 10.0.2.15:6443 was refused - did you specify the right host or port?
解决方案：
sudo -i
swapoff -a
exit
strace -eopenat kubectl version

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