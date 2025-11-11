---
title: SICP|构造数据抽象
date: 2025-10-26 10:54:00
tags: [计算机基础,读书笔记]
category: 计算机基础

---

### 前言

在[构造过程抽象](https://silentechoe.github.io/2025/09/28/%E6%9E%84%E9%80%A0%E8%BF%87%E7%A8%8B%E7%9A%84%E6%96%B9%E5%BC%8F/)一文中，我们讨论了**过程的抽象**这一思想。通过复合过程，我们得以提升程序设计时的概念层次，使设计更加模块化，并增强语言的表达能力。

不过，在那个阶段中，所有函数操作的都是最简单的数值类型，而简单的数据结构表达能力是不够的，许多程序在设计之初就是为了模拟真实世界中多层次的结构和关系。正如定义过程让我们能在更高层次上组织计算一样，同样的，能够构造复合数据的能力，也将使我们得以在此比语言提供的基本数据对象更高的概念层次上，处理与数据有关的各种问题。



### 什么是数据抽象？

SICP中数据抽象被解释为一种方法学，它可以将一个复合对象的使用，与该数据对象怎样由更基本的数据对象构造起来的细节隔离开。它的核心目的是为了**使用数据时，只关心“数据能干什么”，而不去关心它是如何实现的。**

> 数据抽象的基本思想，就是设法构造出一些使用复合数据对象的程序，使它们就像是在“抽象数据”上操作一样。我们的程序中使用数据的方式应该是这样的:除了完成当前工作所必要的东西之外，它们不对所用数据做任何多余的假设。与此同时，一种“具体”数据表示的定义，也应该与程序中使用数据的方式无关。在我们的系统中，这样两个部分之间的界面将是一组过程，称为“选择函数”和“构造函数”。

比如在程序里面操作一个坐标点位，如果不使用数据抽象，可能直接使用数组，Map，结构体等数据结构：

```go
p := [2]float64{3, 4}  
fmt.Println(p[0], p[1])
```

这样的数据结构会有一个问题：使用这个数据结构的开发者，必须知道每个下标代表的值是纵坐标还是横坐标，一旦数据结构发生变化，那么所用引用到该数据的地方都需要更改。

如果用数据抽象，不直接操作数组，而直接提供一组“接口”：

```go
type Point struct{ x, y float64 }

func MakePoint(x, y float64) Point { return Point{x, y} }
func X(p Point) float64           { return p.x }
func Y(p Point) float64           { return p.y }

p := MakePoint(3, 4)
fmt.Println(X(p), Y(p))
```

使用者无需知道内部是由结构体还是数组来构建，只需要把数据的构建和如何使用数据分开，调用者通过“构造函数”和“选择函数”来操作数据，那么就能自由地改变数据实现，无需更改上层的业务逻辑。

**抽象屏障**在表述中更像是一条分界线，用来把调用者和底层的实现细节隔离开。在上述的例子中，`MakePoint`就是构造函数，而`X`和`Y`是选择函数。这么做的好处在于：通过分界线来屏蔽细节，当有一天底层发生改变时上层的业务也无需改动。

```go
type Point map[string]float64

func MakePoint(x, y float64) Point { return Point{"x": x, "y": y} }
func X(p Point) float64            { return p["x"] }
func Y(p Point) float64            { return p["y"] }
```

使用Map替换struct，但是对**fmt.Println(X(p), Y(p))**函数来说无需改动，可以看到所有的改动都发生在屏障内部(MakePoint/X/Y)。

抽象屏障是为了隔离系统中不同的层次，每一层都可以使用数据的抽象程序与实现数据抽象的程序分开。

> 这里不是OOP中的抽象类和实现类，而是从使用的角度来说，用这种抽象的概念来“分隔数据的使用者与数据的实现”。
>
> 抽象屏障的意义不是为了方便，而是为了稳定和解耦。用接口来屏蔽结构。

表达方式的选择会对操作它的程序产生影响。这才是关键，如果后来表示的方式改变了，所有受影响的程序也需要随之改变。对于大型程序而言，这种工作非常耗时，而且代价极其高，除非在设计时就已经将依赖于表示的成分限制到很少的一些程序模块上。

借助数据抽象的思想就能设计出不会被数据表示的细节纠缠的程序，使程序能够保持很好的弹性，得以应用到不同的具体表示上。



#### 数据到底是什么？

通过过程组合形成更复杂的过程，通过将多个数据对象组合在一起形成复合数据对象，这种组合本身只是过程。在SICP中甚至有例子可以只用过程来表示数据。比如上述的`Point`结构体，通常情况下对外表述可能会说这个结构体就是数据。但是同样地可以用**闭包(过程)**来进行改写：

```go
func MakePoint(x, y float64) func(string) float64 {
	return func(op string) float64 {
		if op == "x" {
			return x
		}
		return y
	}
}

func X(p func(string) float64) float64 {
	return p("x")
}

func Y(p func(string) float64) float64 {
	return p("y")
}

p := MakePoint(3, 4)
fmt.Println(X(p), Y(p))
```

即使不使用高级编程语言内置的基础数据结构，也能通过函数来进行表达。调用放在执行`MakePoint`函数和`fmt.Println`时输出与结构体改写的并无区别。

从抽象的角度来看，一个数据对象的身份是由其支持的操作决定的；存储结构只是实现细节，可以随时替换。

> 不存在天然的数据，只有我们设计出的接口与行为。

SICP中实现了序对这一数据抽象，它是一种复合的结构，它会将两个对象组合成一个新的"对象"。用Go语言表达大概是这样的：

```go
type Pair func(func(interface{}, interface{}) interface{}) interface{}

func Cons(a, b interface{}) Pair {
	return func(f func(interface{}, interface{}) interface{}) interface{} {
		return f(a, b)
	}
}

func Car(p Pair) interface{} {
	return p(func(a, b interface{}) interface{} { return a })
}

func Cdr(p Pair) interface{} {
	return p(func(a, b interface{}) interface{} { return b })
}
```

这里的Pair没有使用Map，数组，结构体，它使用闭包。但是它在行为上完全等价于一个包含a和b的二元组。



### 层次性数据和闭包性质

复合数据不等于直接把两个数放在一起，而是为了把对象和关系组合起来，形成更复杂的层次结构。它表达的是如何用少量组合规则无限扩展表达能力。

> 通过复合数据，我们不仅能表示单个对象，还能表示对象之间的结构和关系

SICP中用"**序列**"组合成"链表"来作为构建层次性数据结构的示例：

```go
package main

import "fmt"

type Pair struct {
	Car interface{}
	Cdr interface{}
}

func Cons(a, b interface{}) *Pair {
	return &Pair{Car: a, Cdr: b}
}

func Car(p *Pair) interface{} {
	return p.Car
}

func Cdr(p *Pair) interface{} {
	return p.Cdr
}

func PrintList(list *Pair) {
	for list != nil {
		fmt.Println(Car(list))
		next := Cdr(list)
		if next == nil {
			break
		}
		list = next.(*Pair)
	}
}

func main() {
	l := Cons(1, Cons(2, Cons(3, nil)))
	PrintList(l)
}

$ 1 -> 2 -> 3 -> nil
```

一个Pair可以组合两个元素，两个元素可以继续嵌套成Pair。链表就是多重嵌套的序队，用相同的构造方式把数据一层一层堆叠起来就形成了“结构”，这就是**层次性**的雏形。通过简单的组合操作，可以无限嵌套构造更加复杂的结构，比如链表，图，树等。

```go
ree := Cons(1, Cons(Cons(2, Cons(3, nil)), Cons(4, nil)))

        (1)
       /   \
     (2 3)  (4)
```

SICP中谈论的“**闭包性质（Closure Property）**”可以视为：当一种组合操作的结果依然是同一种类型的元素时，可以说这个操作具有闭包性。

这意味着编程语言中只需要包含少量的基础数据结构，通过闭包的性质，**可以递归构造出任意复杂的数据结构。**在云原生中，Kubernetes 中的Deployment,StatefulSets本质上也是层层组合而成(Pod模版+ReplicaSet)，只要基础的规则稳定实现，就能不断嵌套成更复杂的结构，也无需改动底层的实现机制。

SICP中提到的“闭包性质”不是指编程语言中的闭包机制，而是数学意义上的“封闭性”——用一种组合方式产生的新数据对象，还可以继续用相同方式进行组合。比如JSON中一个对象嵌套另外一个对象，也是闭包性质体现。

> 如果有一种组合方式，它组合两个对象之后返回的结果，也可以被同样的方式再次组合，那么这种组合方式具有闭包性质。

不同于编程语言中的“闭包”(它指的是函数可以捕获它定义环境中的自由变量)，闭包性质关注的是数据结构和操作的自我封闭性，不是函数捕获环境的行为，这点要区分来看待。闭包函数是编程语言的机制，闭包性质是抽象设计的结构性特征，这两者经常与“递归”和“抽象”同时出现，所以容易混淆，这点要特别说明。



#### 序列作为接口

SICP中引入了**序列抽象**作为操作层的统一接口，这里“序列”不是指某一种具体的数据结构，而是一种**约定的抽象接口**，它指的是一类可以被遍历，组合，变换的数据结构。它不特指某种数据结构，只要满足接口的约定就行。

> 当我们拥有了闭包性质，就可能有许多种不同方式来表示同样的概念。但如果我们为这些结构约定一套标准接口，那么我们就能对所有这些不同表示采取相同的操作方式。

```go
package main

import (
    "fmt"
)

type Sequence[T any] interface {
    // Next 返回下一个元素和一个 bool 表示是否还有
    Next() (T, bool)
    // Map 返回一个新的 Sequence，其中每个元素经 f 变换
    Map(func(T) T) Sequence[T]
    // Filter 返回一个新的 Sequence，仅包含满足 p 的元素
    Filter(func(T) bool) Sequence[T]
}

type sliceSequence[T any] struct {
    data []T
    idx  int
}

func NewSliceSequence[T any](data []T) *sliceSequence[T] {
    return &sliceSequence[T]{data: data, idx: 0}
}

func (s *sliceSequence[T]) Next() (T, bool) {
    if s.idx >= len(s.data) {
        var zero T
        return zero, false
    }
    v := s.data[s.idx]
    s.idx++
    return v, true
}

func (s *sliceSequence[T]) Reset() {
    s.idx = 0
}

func (s *sliceSequence[T]) Map(f func(T) T) Sequence[T] {
    s.Reset()
    newData := make([]T, 0, len(s.data))
    for v, ok := s.Next(); ok; v, ok = s.Next() {
        newData = append(newData, f(v))
    }
    return NewSliceSequence(newData)
}

func (s *sliceSequence[T]) Filter(p func(T) bool) Sequence[T] {
    s.Reset()
    newData := make([]T, 0, len(s.data))
    for v, ok := s.Next(); ok; v, ok = s.Next() {
        if p(v) {
            newData = append(newData, v)
        }
    }
    return NewSliceSequence(newData)
}

func main() {
    ints := []int{1, 2, 3, 4, 5, 6}
    seq := NewSliceSequence(ints)

    // Map: 每个元素乘以2
    mapped := seq.Map(func(x int) int {
        return x * 2
    })

    // Filter: 只保留 > 6 的元素
    filtered := mapped.Filter(func(x int) bool {
        return x > 6
    })
  
    for v, ok := filtered.Next(); ok; v, ok = filtered.Next() {
        fmt.Println(v)  // 输出：8, 10, 12
    }
}
```

`Sequence[T]`是一个通用的接口，无论底层实现的是什么，只要实现了这个接口就能使用相同的操作。Map 和 Filter 返回的依然是`Sequence[T]`，这说明操作的结果依然是序列，使用者不需要关心底层数据是什么，只要使用`Next()` `Map()` `Filter()`就好了。

这种抽象的好处在于不用为每种数据结构分别编写具体的逻辑，Map、Filter、Reduce 这样的操作可以“统一”处理所有结构，结构的差异又可以被隐藏在抽象屏障后面，操作却又能保持一致。也就是说可以像操作数组一样去操作树，图，流等。

在现代的架构中使用接口来屏蔽实现类的方式被普遍应用，复杂性不会和结构绑定，而是转移到了更抽象的”操作层“(业务层或领域层Domain Layer)。



### 符号数据

> 迄今为止，我们用过的所有复合数据都是从数值出发构造起来的。在这一节我们要扩充语言的表述能力，引进把字符的串（字符串，也简称为串）作为数据的功能。

无论是**序对**还是**序列**在示例中，都是使用数字作为底层的原子数据，虽然通过这些已经能组合出复杂的数据结构，但想表达“真实的世界”还远远不够。字符串和数字一样能存放在变量中，被组合成复合的数据结构，能被函数操作……这意味着我们表达的内容会更丰富，比如名称，状态，标识符。C语言中没有字符串的基础数据结构，但是基于C语言开发的Redis中实现`String`并被被普遍使用，Json字符串被存储在中间件中更是被广泛使用。

> “符号数据”可以将名字，结构当作**数据本身**处理，但又不执行它们。

```go
(+ 1 2)  ; 这是一个表达式
'+       ; 这是一个符号
symbol := "add"  // 只是名字本身，不执行
```

站在更加宏观的角度上来看**符号数据**，高级的编程语言不止计算“数值”，还需要在底层来**构造解释器**。解释器必须要有“符号”这个概念，比如在进行数据运算时通常使用(+ - * /)来表达加减乘除，可 + 是一个符号，解释器的工作就是用于识别这些特殊的符号，找到对应的操作然后运用到配套的数值上，从而产生运算。

至此编程语言获得了**自我描述**的能力，可以编写程序来“解释”“生成”“修改”某一个表达式，这就是元编程的雏形。

解释器比较关键的点在于：

1.词法分析，把(+ 1 2)切分成最小的“词快” `"(", "+", "1", "2", ")"` 

2.语法分析，进一步把“词快”组织成树状结构（AST）：

```
Call
├─ operator: "+"
└─ args: [1, 2]
```

3.环境，一个“字典”，记录名字—值，比如 + 这个名字对应的是“加法函数”；x 对应 42。

4.求值器，按预定的规则执行：如果是数字就返回自己；如果是 (+ a b) 就先算出 a 和 b，再做加法；如果是变量就去环境查它的值。

解释器作为现在高级编程语言中的幕后功臣，一直是重要的部分，因为它能让代码变为数据，提供了宏、内置函数、语法糖这些可扩展能力。最终能用语言写出这个语言自己的解释器，也被称为自举。



### 抽象数据的多重表示

一个系统的程序设计通常情况下是由多人在相当一段时间内完成的，系统的需求不是永恒不变的，它会随着需求和时间而变化。在这种情况下要求所有开发者在**数据表示**的选择达成一致几乎是不可能的。

抽象屏障能很好地**隔离表示与使用**，但在一个真实的大型系统中，**“表示本身”也并非永远唯一**。

除了需要隔离表示与使用的数据抽象屏障外，我们还需要有抽象屏障去隔离不同的设计选择，允许不同的设计选择在同一个程序里共存。这段话本质其实在讲"**模块化设计**"的必要性，讨论的是在大型系统中如何通过"抽象屏障"实现模块化与演进，从概念上更进一步后讨论的不只是如何划分模块，而是在抽象层面上为模块与模块之间建立一个“隔离/协作”的机制。

SICP中使用处理“复数”作为例子，可以使用**直角坐标** (x, y)或**极坐标** (r, θ) 基于不同的数据结构计算返回复数的模长，在不抽象的情况下可以写成：

```go
func Magnitude(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
```

不同开发者会有不同的表达习惯，如果更偏向**极坐标**（r, θ）的方式，也能通过抽象屏障封装起来，这两种表示方式各有优势，但系统最终必须要面对的是——**它们会共存**。

即便我们很好地定义了抽象屏障，也无法避免**同一个概念在系统中出现多种实现**的现实。这就是 **“多重表示”**（multiple representations）问题出现的根源：

> 不同表示方式共存于同一系统中，系统必须设计机制让这些表示能够协同工作，而不让上层业务陷入混乱。

SICP提到使用**类型标签**，相当于通过标注让数据表明它的类型，在Go中可以使用interface 加上类型字段来模拟：

```go
type ComplexType string

const (
	Rectangular ComplexType = "rectangular"
	Polar       ComplexType = "polar"
)

type ComplexNumber struct {
	Type ComplexType
	X, Y float64 
	R, T float64
}

func main() {
	c1 := ComplexNumber{Type: Rectangular, X: 3, Y: 4}
	c2 := ComplexNumber{Type: Polar, R: 5, T: math.Pi / 4}
}
```

让数据携带"类型"标注是最简单的区分方式，这种方式简单易用，比如在某些框架支持多种数据库，在启动参数中标注一个带有“数据库类型”的参数就可以在多种数据库中实现切换。

```go
func Run(c Config)  {
  switch c.Type {
    case Mysql:
    			....
    case Mongo:
    			...
  }
}
```

通过`Run`函数可以在服务启动时自由切换数据库的使用，但是该函数是一个**通用型函数**，它无需调用使用的哪一种数据结构，它通过的是内部的Type来派生行为，如果未来还需要支持更多的数据库，只需要在这个函数中添加一个`case`。这种方式虽然简单方便，但是违背了"开闭原则"每次添加新的数据都要更改该函数。

在现代的架构系统中可以看到，一些开发者选择另外一种方式：把“操作”与“类型”的映射关系，存到一张表中。

```go
package main

import "fmt"

type DBType string

const (
	MySQL    DBType = "mysql"
	MongoDB  DBType = "mongo"
	Postgres DBType = "postgres"
)

type DBConfig struct {
	Type     DBType
	Host     string
	Port     int
	Username string
	Password string
}

// 数据库操作的统一接口
type DBHandler func(cfg DBConfig) error

// 分派表，存储“类型 → handler”的映射
var dbHandlerTable = map[DBType]DBHandler{}

// 注册 handler
func RegisterDBHandler(t DBType, f DBHandler) {
	dbHandlerTable[t] = f
}

// 通用调用入口
func ConnectDB(cfg DBConfig) error {
	if handler, ok := dbHandlerTable[cfg.Type]; ok {
		return handler(cfg)
	}
	return fmt.Errorf("no handler registered for db type: %s", cfg.Type)
}


func init() {
	RegisterDBHandler(MySQL, func(cfg DBConfig) error {
		fmt.Printf("[MySQL] Connecting to %s:%d as %s\n", cfg.Host, cfg.Port, cfg.Username)
		return nil
	})

	RegisterDBHandler(MongoDB, func(cfg DBConfig) error {
		fmt.Printf("[MongoDB] Connecting to %s:%d as %s\n", cfg.Host, cfg.Port, cfg.Username)
		return nil
	})

	RegisterDBHandler(Postgres, func(cfg DBConfig) error {
		fmt.Printf("[Postgres] Connecting to %s:%d as %s\n", cfg.Host, cfg.Port, cfg.Username)
		return nil
	})
}
```

没有把逻辑硬编码在通用函数内部，而是通过数据库的动态数据来动态查找并实现，在现代编程中有部分项目使用这种方式来灵活构建开源项目，比较出名的是Casbin，它将数据结构存放在数据库或Csv中来实现动态的权限判断。



### 包含通用型操作的系统

很多开发大型项目的工程师的目标就是想要构建一个足够“通用”的程序，这意味着系统可以适配的场景更多，灵活性高，迭代速度也更快。从细节的角度来说，对函数的要求也更高，比如在编写一个“加减乘除”通用的操作时，开发者希望这些操作既能算 1+2 也能算 1/2，需要支持复数(直角/极坐标)，在后续的迭代中可能还需要增加“矩阵”，”多项式“这些操作。

SICP给出的思路就是通过类型标签和操作表进行分派，可以把这个当作"插件系统"的雏形，通过“类型标签”和“操作表”让同一个**操作名**根据运行时的**参数类型组合**查询表找到对应的实现。它的优势是可以在不更改代码的情况下新增一个新类型或新操作，并支持**跨类型运算**与**类型提升**，以适应某些需要可插拔的场景。

这很容易联想到现代编程语言中的**泛型**，像加减乘除这种通用操作可以通过**泛型**来实现一样的效果，但是**泛型**的目标是让同一份代码复用与多种静态类型，它在编译期间完成类型检查与实例化，它强调的是由编译器保证的强类型安全。

如果基于Go语言的**泛形**来实现加法，泛型约束对内置的数值类型可以勉强实现，但对用户自定义的数值就无法灵活达成：

```go
import "golang.org/x/exp/constraints"

func Add[T constraints.Integer | constraints.Float](a, b T) T {
	return a + b
}
```

泛型的长处在于算法的实现，比如拓扑排序，集合操作等，在这种情况下它的性能不错，类型安全，调试起来也方便。使用SICP中实现分派，它的核心在于add(T,U) 能在运行时决定，并在插件中动态的注册而无需更改核心框架。

```go
type Value interface{ Tag() string }
type opFn func(a, b Value) (Value, error)

var opTable = map[string]opFn{}                     
func key(op string, a, b Value) string { return op + "|" + a.Tag() + "," + b.Tag() }
func register(op string, aTag, bTag string, f opFn)  { opTable[op+"|"+aTag+","+bTag] = f }

func ApplyAdd(a, b Value) (Value, error) {
	if f, ok := opTable[key("add", a, b)]; ok { return f(a, b) }
	return nil, fmt.Errorf("no method for add on (%s,%s)", a.Tag(), b.Tag())
}
```

不过这两种方式不是互斥的，现代架构中经常使用混合架构，它的核心思路是在外层用**操作表**开发一个派生的路径，在内层使用泛形把算法实现好：

```go
type Value interface{ Tag() string }
type Fn1 func(Value) error
var op1 = map[string]Fn1{}


func Register1(op, tag string, f Fn1) { op1[op+"|"+tag] = f }
func Apply1(op string, v Value) error {
	if f, ok := op1[op+"|"+v.Tag()]; ok { return f(v) }
	return fmt.Errorf("no %s for %s", op, v.Tag())
}

// 内层实现
type Num interface{ ~int | ~int64 | ~float64 }
func Clamp[T Num](x, lo, hi T) T {
	if x < lo { return lo }
	if x > hi { return hi }
	return x
}

type PVCSpec struct{  }
func (PVCSpec) Tag() string { return "storage.pvc" }

func installPVC() {
	Register1("validate", "storage.pvc", func(v Value) error {
		spec := v.(PVCSpec)
		_ = Clamp[int]()     // 这里用泛型做内核算法
		_ = Clamp[float64]() // 多个数值类型安全复用
		return nil
	})
}
```

通过这种方式可以实现PaaS的Trait，让多个资源可以热插不同的“特征”，同时内部又能享受**享受泛型带来的类型安全与性能**，外层的调度器也不用编写 `if` 或 `switch`就可以扩展新特征。



### 小结

SICP中数据抽象是一种思想，它不是简单告诉怎么使用数组，结构体这些编程语言的特性，而是通过这些基础例子的演化来反应一些设计的原则，比如提及**抽象屏障**是为了讲述数据与使用之间的分界线，让表示与使用彻底解耦，**闭包性质** 让我们能用少量基础元素递归构造出复杂结构，这是所有现代语言复合数据结构的基础。**符号数据** 扩展了语言的表达能力。**多重表示与分派机制** 是大型系统演进的必然，提供了模块化和灵活性。

在现代架构中，这些思想早已演化为各种成熟的技术体系：接口抽象、泛型、动态分派、插件系统、DSL、策略模式、Trait 机制……





### 参考资料

https://coolshell.cn/articles/21164.html
