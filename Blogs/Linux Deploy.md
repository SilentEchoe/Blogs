---
title: Linux 常用软件安装
date: 2020-09-05 15:00:00
tags: [Linux,学习笔记,软件安装]
category: Linux

---



[TOC]

# Linux 常用软件安装

## Ubuntu

### 1.1 软件名单

| 软件     | 版本  | 描述 |
| -------- | ----- | ---- |
| Node.js  | 12.3+ |      |
| MongoDB  | 3.6 + |      |
| Genieacs | 最新  |      |

#### 1.1.1 Node

##### 1.1.1.1 安装

```bash
$ sudo apt update
$ sudo apt install nodejs
$ sudo apt install npm

# 升级node版本
$ sudo npm cache clean -f
$ sudo npm install -g n
$ sudo n stable

```



#### 1.1.2 MongoDB

##### 1.1.2.1 部署

```
# 更新软件包列表以获取最新版本的存储库列表
$ sudo apt update

$ sudo apt install -y mongodb
```

##### 1.1.2.2 服务管理

- 启动：sudo systemctl start mongodb
- 停止：sudo systemctl stop mongodb
- 重启：sudo systemctl restart mongodb



##### 1.1.2.3 开启远程连接

注：内网环境无需开启，部署云服务器时设置,需要设置密码访问



$ sudo vim /etc/mongodb.conf

将bind_ip = 127.0.0.1   #注意开启远程连接应该改为 bind_ip = 0.0.0.0 然后重启mongodb

```
# mongodb.conf
  
# Where to store the data.
dbpath=/var/lib/mongodb

#where to log
logpath=/var/log/mongodb/mongodb.log

logappend=true


bind_ip = 127.0.0.1   #注意开启远程连接应该改为 bind_ip = 0.0.0.0
#port = 27017

# Enable journaling, http://www.mongodb.org/display/DOCS/Journaling
journal=true

# Enables periodic logging of CPU utilization and I/O wait
#cpu = true

# Turn on/off security.  Off is currently the default
#noauth = true
#auth = true

# Verbose logging output.
#verbose = true

# Inspect all client data for validity on receipt (useful for
# developing drivers)
#objcheck = true

```



##### 1.1.2.4 开放防火墙端口

```
$ iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 27017 -j ACCEPT
```



##### 1.1.2.5  设置用户名和密码



```
use admin

db.createUser(
    {
	user: "admin",
	pwd: “huoshen.info2020”,
	roles: [{role: ”userAdminAnyDatabase”,db:”genieacs”}]
    }
  )


```



### 1.2 应用部署

#### 1.2.1 Genieacs 部署

注：Genieacs 原版无法适配锐捷设备，需更改部分源码



1. 更改soap.ts 138 行代码 改为：

```
let valueType = getValueType(valueElement.attrs);

        if (valueType) {
          valueType = valueType["value"].trim();
        } else {
          valueType = 'xsd:string';
        }
```

2. 更改soap.ts  273 行代码 改为：

```
  //return `<ParameterValueStruct><Name>${p[0]}</Name><Value xsi:type="${
    //    p[2]}">${encodeEntities('' + val)}</Value></ParameterValueStruct>`;
    
      return `<ParameterValueStruct><Name>${p[0]}</Name><Value>${encodeEntities("" + val)}</Value></ParameterValueStruct>`;
```



#### 1.2.2 启动Genieacs

将Genieacs源码上传至服务器 /home/ 目录

```
$ cd /genieacs/

# 编译
$ npm run build

# 进入到编译之后的文件
$ cd /dist/bin/

# 启动3000 端口页面
$ genieacs-ui --ui-jwt-secret secret

#启动cwmp 包检测
$ genieacs-cwmp

# 启动API
$ genieacs-nbi
```



#### 1.2.3 编写批处理脚本

脚本为 start.bat 文件，linux 启动时需给该文件权限

```
cd /home/genieacs/dist/bin

nohup  genieacs-ui --ui-jwt-secret secret &

nohup  genieacs-cwmp &

nohup  genieacs-nbi &
```



停止服务脚本

```
fuser -k -n tcp 3000
fuser -k -n tcp 7547
fuser -k -n tcp 7557
```



## Docker 安装软件

### 1.1 软件名单

| 软件    | 版本 | 描述 |
| ------- | ---- | ---- |
| MongoDB | 最新 |      |
|         |      |      |
|         |      |      |

#### Centos 安装Docker 

## 注意：Centos 需升级为最新版本

```
$ yum install docker
```

##### 服务管理

- 启动：sudo systemctl start docker
- 停止：sudo systemctl stop docker
- 重启：sudo systemctl restart docker

#### 1.1.1 MongoDB

```
# 拉取mongo
$ docker pull mongo

# 启动mongo
$ docker run --name <YOUR-NAME> -p 27017:27017 -v /data/db:/data/db -d mongo:3.4 --auth
```



#### 1.1.2 .Net Core

#### 1.1.3 持续集成部署

```
# 安装git
$ yum install -y git
```

