---
title: Nocalhost 插件使用
date: 2022-2-23 16:22:00
tags: [云原生,学习笔记]
category: 插件
---





## Nocalhost 插件使用

1.Goland 安装 Nocalhost

[安装步骤]: 	"https://nocalhost.dev/zh-CN/docs/quick-start"



2.在集群所在服务器上，输入以下命名：

```
kubectl config view --minify --raw --flatten
```



[![qCTob8.png](https://s1.ax1x.com/2022/03/17/qCTob8.png)](https://imgtu.com/i/qCTob8)   



3.服务配置：

```go
name: "user"
  serviceType: "deployment"
  containers: 
    - 
      name: "user"
      dev: 
        gitUrl: ""
        image: "nocalhost-docker.pkg.coding.net/nocalhost/dev-images/golang:latest"
        shell: "bash"
        workDir: ""
        storageClass: ""
        resources: 
          limits: 
            memory: "2048Mi"
            cpu: "2"
          requests: 
            memory: "512Mi"
            cpu: "0.5"
        persistentVolumeDirs: []
        command: 
          run: 
            - "./run.sh"
            - "user"
          debug: 
            - "./debug.sh"
            - "user"
        debug: 
          language: "go"
          remoteDebugPort: 9009
        hotReload: true
        sync: 
          type: "send"
          mode: "pattern"
          filePattern: 
            - "."
          ignoreFilePattern: 
            - ".git"
          deleteProtection: true
        env: []
        portForward: []
        sidecarImage: ""
```

参数说明：

hotReload: 热加载，最开开启热加载，方便调试。

command：run or debug 时执行的脚本，必备

```
debug.sh 脚本

#! /bin/sh
export GOPROXY=https://goproxy.cn
dlv --headless --log --listen :9009 --api-version 2 --accept-multiclient debug ./cmd/"$1"/main.go

run.sh 脚本
#! /bin/sh
export GOPROXY=https://goproxy.cn
go run ./cmd/"$1"/main.go
```



### 注意事项

#### 选择DevMode(Duplicate)

这里要注意，构建完 DevMode 后，可以在 Terminal 信息中看到构建的容器，最好使用 ls 命令查看源码是否包含在内。（如果使用 ls 命令后，容器内什么文件都没有，就存在异常）





## 学习资料

[**Nocalhost**](https://nocalhost.dev/)

