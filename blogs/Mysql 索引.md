---
title: Mysql 索引
date: 2024-11-10 19:51:00
tags: [mysql]
category: 存储
---

### 

Mysql 只能高效地使用索引的最左前缀列，创建一个包含两个列的索引，和创建两个只包含一列的索引是不同的。

不同的存储引擎的索引工作方式是不一样的，这一章我们只讨论B+Tree索引，因为Mysql中默认使用InnoDB作为存储引擎，它使用就是B+tree索引。



