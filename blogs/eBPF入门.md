---
title: eBPF 入门
date: 2024-12-12 11:04:00
tags: [eBPF]
category: eBPF
---

### eBPF是什么？

BPF(Berkeley Packet Filter)最初代表一种网络过滤器，它起源于Linux内核，现如今它代表的不仅仅是包过滤，可以说它现在是一个独立的术语。它可以在特权上下文中(比如操作系统内核)运行沙盒程序，安全有效地扩展内核功能，这意味着无需更改内核源码或加载内核模块就能完成拓展。开发人员可以编写eBPF程序，在运行时向操作系统添加额外的功能。



在操作系统中是实现可观测性，安全性和网络功能的理想场所，内核一直拥有控制和监督整个系统的特权，但又因为内核安全性和稳定性太高，相较于应用层，内核很难快速迭代。

eBPF允许在操作系统内核中运行沙盒程序，开发人员可以编写eBPF程序，在运行时向操作系统添加额外的功能。在JIT编译器和验证引擎的帮助下，操作系统确保它像本地编译的程序一样具备安全性和执行效率。比较常见的用途是网络监控，安全过滤，性能分析等。



### 如何工作？

eBPF主要分为三个步骤：加载，编译和执行。

1.eBPF 需要在内核中运行。一般由用户态的应用程序完成，它会通过系统调用来加载 eBPF程序。在加载过程中，内核会将eBPF程序的代码复制到内核空间。

2.eBPF 需要经过编译，通常情况下是由Clang/LLVM 编译器完成，形成字节码后将用户态的字节码装载进内核，Verifier 会对要注入内核的程序进行一些内核安全机制的检查，这点非常重要，这保证了eBPF程序不会破坏Linux内核的稳定性和安全性。在检查过程中，内核会对 eBPF程序的代码进行分析，比如一些敏感操作是会被禁止的，比如系统调用，内存访问等。

3.当通过内核安全机制检查后，它就可以在内核中正常运行了，通过一个JIT编译步骤将程序的通用字节码转换为机器特定指令集，以优化程序的执行速度。



eBPF程序由事件驱动组成，当内核或应用程序通过某个钩子点时运行。预定义的钩子包括系统调用，函数入口/退出，内核跟踪点，网络事件等。当预定义的钩子不能满足特定需求，可以创建内核探针(kprobe)或用户探针(uprobe)在内核或用户程序中附加eBPF程序。

大部分情况，eBPF不会直接使用，而是通过Cilium bcc bpftrace 项目间接使用，这些项目提供了 eBPF 之上的抽象，不需要直接编写程序，提供基于意图来定义实现的能力，当无法满足特定需求时，才考虑编程实现eBPF程序。



Linux 内核的主要目的是对硬件或虚拟硬件进行抽象，并提供一致的API(系统调用)，允许应用程序运行和共享资源。内核会维护一组广泛的子系统和层来分配这些指责，每个子系统通常允许某种级别的配置，以满足用户不同需求。如何无法配置所需的行为，则需要更改内核。



### 如何开发eBPF程序？

虽然社区提供了很多eBPF相关的工具库，比如 ebpf-go libbpfgo gobpf 等，不同的编程语言也有相应的实现库，但是Go开发库只适用于用户态程序中，它可以完成 eBPF程序编译，加载，事件挂载，以及 BPF 映射交互等用户态的能力。内核态的 eBPF 程序还是需要使用 C 语言开发。



```
# For Ubuntu20.10+
sudo apt-get install -y  make clang llvm libelf-dev libbpf-dev bpfcc-tools libbpfcc-dev linux-tools-$(uname -r) linux-headers-$(uname -r)

# For RHEL8.2+
sudo yum install libbpf-devel make clang llvm elfutils-libelf-devel bpftool bcc-tools bcc-devel
```



```c
int hello_world(void *ctx)
{
    bpf_trace_printk("Hello, World!");
    return 0;
}
```



```python
#!/usr/bin/env python3
# 1) import bcc library
from bcc import BPF

# 2) load BPF program
b = BPF(src_file="hello.c")
# 3) attach kprobe
b.attach_kprobe(event="do_sys_openat2", fn_name="hello_world")
# 4) read and print /sys/kernel/debug/tracing/trace_pipe
b.trace_print()
```

调用BPF()加载BPF源码，也就是C语言编写的内容。然后将BPF程序挂载到内核探针(kprobe) do_sys_openat2()是系统调用 openat() 在内核中的实现。最后读取内核调试文件中 /sys/kernel/debug/tracing/trace_pipe 的内容，并打印到标准输出中。



![QQ_1734400238779](https://raw.githubusercontent.com/SilentEchoe/images/main/QQ_1734400238779.png)



eBPF 是一个运行在内核中的虚拟机，但是系统虚拟化和 eBPF 虚拟机有着本质的不同。系统虚拟化基于 x86 或 arm64 等通用指令集，这些指令集足以完成完整计算机的所有功能。eBPF 只提供了有限的指令集，这些指令集可以完成一部分内核的功能，但是没办法完整使用计算的所有功能。eBPF 指令使用 C 调用约定，它支持的辅助函数可以在 C 语言中直接调用。





eBPF 在内核运行时主要由五个模块组成：

1.eBPF辅助函数，提供一系列用于eBPF程序与内核其他模块进行交互的函数

2.eBPF验证器，用于确保eBPF程序的安全。验证器会将待执行的指令创建为一个有向无环图(DAG) 确保程序中不包含不可达指令，然后再模拟指令的执行过程，确保不会执行无效指令。

3.由11个64位寄存器，一个程序计数器和一个512字节的栈组成的存储模块。这个模块用于控制eBPF程序的执行。R0寄存器用于存储函数调用和eBPF程序的返回值，这代表函数调用最多只有一个返回值。R1-R5寄存器用于函数调用的参数，因此函数调用的参数最多不能超过5个，R10是一个只读寄存器，用于从栈中读取数据。

4.即时编译器，它将eBPF字节码编译成本地机器指令，以便更高效地在内核中执行。

5.BPF映射，用于提供大块的存储，这些存储可被用户空间程序用来进行访问，进而控制eBPF程序的运行状态。





```
478: kprobe  name hello_world  tag c3d700bdee3931e4  gpl
        loaded_at 2024-12-17T14:38:35+0800  uid 0
        xlated 528B  jited 360B  memlock 4096B  map_ids 131
        btf_id 112
```



一个完整的eBPF程序通常包含用户态和内核态两部分。用户态负责 eBPF 程序的加载，事件绑定以及 eBPF 程序运行结果汇总输出。内核态运行在 eBPF虚拟机中，负责定制和控制系统的运行状态。

eBPF 内部的内存空间只有寄存器和栈，如果要访问其他的内核空间或用户空间地址，需要借助 bpf_probe_read 这一系列的辅助函数。这些函数会进行安全性检查，并禁止缺页中断的发生。当eBPF程序需要大块存储时，不能像常规的内核代码直接分配内存，而是必须通过BPF映射(BPF Map)来完成。

### BPF 映射

BPF 映射用于提供大块的键值存储，这些存储可以被用户空间程序访问，以此来获取eBPF程序的运行状态，最多可以访问64个不同的BPF映射，并且不同的 eBPF程序也可以通过相同的BPF映射来共享它们的状态。



eBPF 程序可以根据功能和场景的不同划分为三类：

1.跟踪，从内核和程序的运行状态中提取跟踪信息，以此来了解当前系统发生了什么。

2.网络，对网络数据包进行过滤和处理，以此了解和控制网络数据包的收发过程。

3.安全，安全控制，BPF扩展等。



