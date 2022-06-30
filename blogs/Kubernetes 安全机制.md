---
title: Kubernetes 安全机制
date: 2022-3-30 11:25:00
tags: [Kubernetes,学习笔记,安全]
category: Kubernetes
---

# Kubernetes 安全机制

Kubernetes 通过一系列安全机制来保证集群的安全控制



## Authentication 认证

Kubernet 对 API 调用使用 CA （Client Authentication）, Token 和 HTTP Base 方式实现用户认证。

CA 被称为可信第三方（Trusted Third Party，TTP）CA 认证涉及诸多概念，比如根证书，自签名证书，密钥，私钥，加密算法及HTTPS 等。

Kubernetes 的 CA 认证方式通过添加 API Server 的启动参数 “—client_ca_file=SOMEFILE” 实现，其中 “SOMEFILE”  为认证授权文件，该文件包含一个或多个证书办法机构。

Token 认证方式通过添加 API Server 的启动参数 “—token_auth_file=SOMEFILE” 实现，其中 “SOMEFILE” 指的是存储 Token 的 Token 文件。目前，Token 认证中 Token 是永久有效的，而且 Token 列表不能被修改，除非重启 API Server.



## Authorization 授权

在 Kubernetes 中，授权是认证后的一个独立步骤，作用于 API Server 主要端口的所有 HTTP 访问。

访问策略有三种

AlwaysDeny 表示拒绝所有的请求，一般用于测试；

AlwaysAllow 表示接收所有的请求，如果集群不需要授权流程，则可以采用该策略；

ABAC 表示使用用户配置的授权策略去管理访问 API Server 的请求。

在 Kubernetes 中，一个 HTTP 请求包含四个能被授权进程识别的属性：

用户名（代表一个已经被认证的用户的字符型用户名）；

是否是只读请求（Get 操作是只读的）；

被访问的是哪一类资源

被访问对象所属的 Namespace ，如果被访问的资源不支持 Namespace ,则是空字符串。

如果选用ABAC 模式，需要通过设置 API Server 的 “—authorization_policy_file=SOME_FILENAME” 参数来指定授权策略文件，其中 “SOME_FILENAME” 为授权策略文件。

授权策略文件中的策略对象的一个未设置属性，表示匹配 HTTP 请求中该属性的任何值。对请求的四个属性值和授权文件中的所有策略对象逐个匹配，如果至少有一个策略对象被匹配，则该请求将被授权通过。

示例：

```yaml
# 用户 A 指定读取资源 Pods
{
	"user":"A",
	"resource":"pods",
	"readonly":true
}

# 用户 B 只能读取 Namespace “test” 中的资源 Pods
{
	"user":"B",
	"resource":"pods",
	"readonly":true，
	"ns":"test"
}
```



## Admission Control 准入控制

用于拦截所有经过认证和鉴权后的访问 API Server 请求的可插入代码（或插件）。这些可插入代码运行于 API Server 进程中，在被调用前必须被编译成二进制文件。在请求被 API Server 接收前，每个 Admission Control 插件按配置顺序执行。如果其中的任意一个插件拒绝该请求，就意味着这个请求被 API Server 拒绝，同时 API Server 反馈一个错误信息给请求发起方。

Admission Control 插件回使用系统配置的默认值去改变进入集群对象的内容，而且可能回改变请求处理所使用的资源的配额，比如增加请求处理的资源配额。



## Secret 私密凭据

作用于保管私密数据，比如密码，OAuth Tokens , SSH Keys 等信息。这些私密信息放在 Secret 对象中比直接放在 Pod 或 Docker Image 中更安全，也便于使用。

Kubernetes 在 Pod 创建时，可以指定 Service Account 用于访问 API Server 和下载 Image .

```yaml
$ kubectl namespace myspace
$ cat <<EOF> secrets.yaml

apiVersion:v1
kind: Secret
metadata:
 name: mysecret
type: Opaque
data:
 password: base64 编码值
 username: base64 编码值

$ kubectl create -f secrets.yaml
```

注意：data 域的各子域的值必须为 base64 编码值

一旦 Secret 被创建，可以通过三种方式使用它：

1.创建 Pod 时，通过 Pod 指定 Service Account 来自动使用它

2.通过挂载该 Secret 到 Pod 来使用它

3.在创建 Pod 时，指定 Pod 的 spc.ImagePullSecrets 来引用它

Pod 创建时候会验证所挂载的 Secret 是否真的指向一个 Secret 对象，因此 Secret 必须在任何引用它的 Pod 之前被创建。并且 Secret 对象属于 Namespace ，它们只能被同一个 Ns 中的 Pod 所引用。

Secret 大小不能超过 1M

可以通过 Secret 保管其他系统的敏感信息，并以 Mount 的方式将 Secret 挂载到 Container 中，然后通过访问目录中的文件的方式获取敏感信息。



## Service Account

Service Account 是多个 Secret 的集合，一种是普通的 Secret，用于访问 API Server；另外一种用于下载容器镜像。

```bash
# 查询 Service Account 列表
kubectl get serviceAccounts 
```

如果创建 Pod 时没有为 Pod 指定 Service Account ，则系统会自动为其指定一个在同一个命名空间下的 “default”的 Service Account 。

如果后续想更改，可以更改 yaml 文件：

```
apiVersion: v1
kind: Pod
metadata:
 name: mypod
spec:
 containers:
   - name: mycontainter
     image: nginx:v1
 serviceAccountName: myserviceaccount
```



# 学习资料

《Kubernetes 权威指南》
