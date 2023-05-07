---
title: Dex LDAP SSO 
date: 2023-5-01 20:58:00
tags: [权限认证,学习笔记]
category: Go
---



## 什么是SSO?

单点登录（SSO *Single Sign* On）是一种身份验证解决方案，可以让用户通过一次用户身份验证登录多个应用程序和网站。

使用SSO可以简化用户登录，加强密码安全从而提高生产效率，也有更好的用户体验。



## 什么是身份认证？

身份认证（Authentication）是指通过一定的手段，完成对用户身份的确认。这种确认需要保证其他用户不容易伪装其身份，一般只有经过身份认证得到可靠的用户身份后，才能基于该身份进行后续的权限验证流程。



## 基于Dex实现身份认证系统

### Dex介绍

Dex是一个开源的身份认证系统，它支持广泛的身份提供程序，如LDAP和OAuth2，并实现OpenID Connect。它可以很方便提供一个SSO，开发者可以基于Dex快速完成一个身份认证系统。



### 1.创建Ladp服务

`kubectl create -f lady.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ldapdemo
  namespace: ssodemo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ldapdemo
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: ldapdemo
    spec:
      containers:
        - env:
            - name: LDAP_ADMIN_USERNAME
              value: admin
            - name: LDAP_ADMIN_PASSWORD
              value: adminpassword
            - name: LDAP_ROOT
              value: dc=demo,dc=cn
          image: bitnami/openldap:latest
          imagePullPolicy: Always
          name: openldap
          ports:
            - containerPort: 1389
              name: service
              protocol: TCP
---
apiVersion: v1
kind: Service
metadata:
  name: ldapdemo
  namespace: ssodemo
spec:
  type: NodePort
  ports:
    - port: 1389
      protocol: TCP
      targetPort: 1389
      name: http
  selector:
    app: ldapdemo
```



创建apache-directory-studio连接Ldap,在Safari浏览器中输入: vnc://127.0.0.1:5901

```shell
docker run --rm -ti \
        -v apache-directory-studio:/root/.ApacheDirectoryStudio \
        -p '127.0.0.1:5901:5901' \
        brmzkw/apache-directory-studio
```



