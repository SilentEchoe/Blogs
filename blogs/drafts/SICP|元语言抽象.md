---
title: SICP|元语言抽象
date: 2026-01-06 17:20:00
tags: [计算机基础,读书笔记]
category: 计算机基础
---

元语言抽象(Metalinguistic Abstraction)指的就是设计和实现新的领域专用语言(DSL),这使工程不仅可以使用现有的高级编程语言，还可以通过编写**解释器**来创造新语言，比如Go语言最初的解释器是C语言写的，在经历数个版本后才使用Go语言实现**自举**。

解释器本质上就是一个接受“程序”(表达式)作为输入并执行它的过程。编程语言本身也可以被程序操控和实现，它定义了某种语言中表达式的含义。

SICP中使用 Scheme 实现了 Scheme 的解释器，这是为了展示**语言与解释器的关系**。一种语言的语法和语义完全由其解释器决定，如果用另外一门语言编写一个解释器来执行某种语言的程序，那么这个解释器本身就是对改语言的一种定义。





### 解释器的基本结构

解释器的核心在于两个互相递归调用的过程: eval(求值)和apply(应用)，eval在某个环境中计算一个表达式的值，apply则将一个过程作用于一组参数值。这两个过程构成了解释器的基本工作循环：eval 将表达式的求值转换为对过程的调用，apply 在调用过程中产生新的表达式需要求值，通过互相递归调用直到得出最终结果(如原始数据类型的值或基本过程的执行结果)。

在任何解释器中处理不同的表达式逻辑都不相同，比如表达式的值就是它自身，1就是1；变量表达式的值需要在环境中查找。

> 环境是一系列**帧**（frame）的集合，每个帧包含一些符号到值的绑定，并可能有指向外围环境的链接。如果在当前环境中找不到变量，会沿着链寻找全局环境，找不到则报未绑定错误。

复合表达式，比如函数调用则需要先计算运算符和操作数，再将运算符得到的过程应用与操作数的值。这进一步引出了**应用(apply)**:如果有一个过程(函数)和参数值列表，那么apply负责执行这个过程，假设它是最基础的加法运算，apply就会对参数求和；如果过程是用户定义的函数，apply则会创建新环境将参数值绑定到函数的形参，再执行函数体。

条件表达式，比如If,eval 对 if 条件语句会先求值，根据结果真/假选择性地求值后续分支。

Lambda(匿名函数)， 解释器不会立即执行匿名函数中的主体内容，而是创建一个**过程对象(闭包)**，记录参数列表，函数体以及当前环境作为闭包的“外围环境”。在 Scheme 实现中，这可以用数据结构封装三个部分。之后eval返回这个过程对象，以便在后面应用。

```go
import "fmt"

// 表达式类型定义：数字、符号、以及由多个子表达式构成的列表
type Number float64
type Symbol string
type Expr interface{}      
type List []Expr


type Value interface{}      

// 环境：保存变量绑定的结构，支持嵌套（链表结构）
type Env struct {
    vars  map[string]Value  // 当前环境的名字->值映射
    outer *Env              // 指向外层环境（用于实现嵌套的词法作用域）
}

// 在环境中查找变量的值
func (env *Env) Lookup(name string) Value {
    if val, ok := env.vars[name]; ok {
        return val  // 当前环境找到变量
    }
    if env.outer != nil {
        return env.outer.Lookup(name)  // 在外层环境递归查找
    }
    panic("未绑定的变量: " + name)
}

// 创建新环境（可指定外层环境）
func NewEnv(outer *Env) *Env {
    return &Env{vars: make(map[string]Value), outer: outer}
}

// Eval: 在给定环境中求值一个表达式
func Eval(expr Expr, env *Env) Value {
    switch e := expr.(type) {
    case Number:
        // 数字字面量，直接返回其值
        return float64(e)
    case Symbol:
        // 符号（变量），从环境中查询其值
        return env.Lookup(string(e))
    case List:
        if len(e) == 0 {
            panic("空表达式无法求值")
        }
        // 非空列表表达式，抽取运算符及参数子表达式
        opExpr := e[0]
        argExprs := e[1:]
        // 特殊形式（如if、define等）将在后续小节处理，这里假定不是特殊形式，即为函数调用
        // 1. 先求值运算符得到过程值
        proc := Eval(opExpr, env)
        // 2. 依次求值参数表达式得到参数值切片
        var argVals []Value
        for _, arg := range argExprs {
            argVals = append(argVals, Eval(arg, env))
        }
        // 3. 调用 Apply，将过程应用于参数值
        return Apply(proc, argVals, env)
    default:
        panic("不支持的表达式类型")
    }
}

// Apply: 将过程proc应用于参数列表args
func Apply(proc Value, args []Value, env *Env) Value {
    switch fn := proc.(type) {
    case func([]Value) Value:
        // 原生过程（基本过程）：直接调用Go函数执行
        return fn(args)
    case *Proc: 
        // 复合过程（用户定义的函数），见4.1.3节详述
        // 这里简单演示其结构：创建新环境，绑定形参并执行函数体
        newEnv := NewEnv(fn.env)              
        // 将参数值绑定到形参
        for i, paramName := range fn.params {
            newEnv.vars[paramName] = args[i]
        }
        // 在新环境中求值函数体（支持函数体是单一表达式）
        return Eval(fn.body, newEnv)
    default:
        panic("无法应用的类型")
    }
}
```

在上述代码中，eval函数按照表达式的类型执行不同逻辑，如果是数字直接返回，如果是符号查环境，如果是列表则视为一次函数调用。对于一般的函数调用，eval 首先递归求值运算符子表达式得到一个过程对象，然后求值所有参数子表达式得到参数值列表，最后调用 apply 执行过程。如果是基本过程则直接调用得到结果；如果是一个复合过程，apply 会创建新环境绑定参数并递归调用 eval执行函数体。整个过程往复循环，这就是解释器工作的核心。



### 作为数据的表达式

在Lisp系列语言中，有一个重要思想是**程序代码本身可以想数据结构一样来表示和操作**。表达式是数据，所以程序能读取，构造和执行表示为数据的代码。

为了让解释器处理输出的程序，首先需要一种新的方式在解释器中表示程序的结构，比如(1+2)这种算术表达式可以用一种数据表示方式存储在内存中。Go语言中也需要定义数据类型来表示各种表达式构成的树形结构，这样 eval 函数才能遍历和识别这些表达式。

通过将代码表示为数据，解释器就能像处理普通数据一样解释和执行这些代码。

```go
var globalEnv = NewEnv(nil)

// 将加法过程 "+" 加入环境：实现为求和的Go函数
globalEnv.vars["+"] = func(args []Value) Value {
    sum := 0.0
    for _, arg := range args {
        sum += arg.(float64)
    }
    return sum
}

// 将 "=" 过程加入环境：判断两个参数是否相等
globalEnv.vars["="] = func(args []Value) Value {
    if len(args) != 2 {
        panic("= 需要2个参数")
    }
    return args[0] == args[1]
}

// 构造表达式 (+ 1 2) 对应的 Go 数据结构
expr := List{Symbol("+"), Number(1), Number(2)}

// 在 globalEnv 环境中求值该表达式
result := Eval(expr, globalEnv)
fmt.Println(result)  // 输出结果: 3
```

