---
title: Rust life cycle
date: 2024-8-16 10:58:00
tags: [Rust,学习笔记]
category: Rust
---

### 生命周期

生命周期是Rust独有的概念，因为没有其它编程语言的经验作为借鉴，这可能是Rust中最困难的部分。生命周期的主要作用在于避免悬垂引用，包含指针的编程语言中很容易遇到这种错误：指针指向的内存可能已经被分配给其它的持有者。



即使是Go语言这种带GC并包含指针的编程语言也会遇到[野指针]错误：

![image-20240816111531838](https://raw.githubusercontent.com/SilentEchoe/images/main/image-20240816111531838.png)









