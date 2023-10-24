---
title: Go 语言康复训练
date: 2021-7-20 22:23:00
tags: [Go,面试]
category: Go
---



# Go语言基础

## 数组

数组是由相同类型元素的集合组成的数据结构，计算机会为数组分配一块连续的内存来保存其中的元素。我们可以利用数组中元素的索引快速访问特定元素，常见的数组大多都是一维的线性数组，多维数组在数值和图形计算领域有比较常见的应用。

下例先以一维数组为例：

```go
// go 语言中的数组有两种创建方式

//1.显式制定数组大小
arr := [3]int{1,2,3}

//2.使用[...]T声明数组
arr := [...]int{1,2,3}

//3.声明数组，但不完全初始化值
var arry = [3]int{}
arry[0] = 1
```

上例前两种不同的声明会导致编译器做出完全不同的处理：

如果使用第一种`[3]T`，那么变量的类型在编译进行到**类型检查**阶段就会被提取出来，然后创建包含数组大小的结构体。

如果使用`[...]T`方式声明，编译器会先对数组的大小进行推导。但是要强调的是，[...]T 这种方式只是Go语言给我们提供的语法糖，不想计算数组中的元素时可以使用这种方法，这两种方法在运行期间得到的结果是完全相同的。

#### 数组堆栈分配

对于一个由字面量组成的数组，根据数组元素数量的不同，编译器会在负责初始化字面量时候有不同的优化：

1.当元素小于或等于 4 个时，会直接将数组中的元素放置在栈上；

2.当元素数量大于 4 个时，会将数组中的元素放置到静态区并在运行时取出；(变量在静态存储区初始化然后拷贝到栈上)

> 静态存储区：内存在程序编译的时候就已经分配好，这块内存在程序的整个运行期间都存在。它主要存放静态数据、全局数据和常量。

无论在栈上还是静态存储区，数组在内存中都是一连串的内存空间，通过数组开头的指针，元素数量以及元素类型占的空间大小表示数组。

在使用数组时，要特别注意**访问越界**的问题。编译器无法提前发生错误，这种错误会在Go语言运行时出现。





## 切片

Go语言中更常用的数组结构是切片，即动态数组。我们可以在切片中追加元素，切片会在容量不足时自动扩容。

```go
// 1.通过下标的方式获得数组或者切片的一部分
arr := [...]int{1, 2, 3, 4, 5}
slice := arr[0:3]
fmt.Println(slice)

#输出：[1 2 3]

slice := []int{1, 2, 3, 4, 5}
slice2 := slice[0:3]
fmt.Println(slice2)

#输出：[1 2 3]


// 2.使用字面量初始化新的切片
slice := []int{1, 2, 3, 4, 5}

//3.使用make关键字
slice := make([]int,10)

```

下标：使用下标初始化切片不会拷贝原数组或原切片中的数据，它只会创建一个指向原数组的切片结构体，**所以修改新切片的数组也会修改原切片**。

字面量：使用字面量，大部分工作会在编译期间完成。

关键字： 使用 make 关键字时，很多工作需要运行时的参与，调用方必须向 make 函数传入切片的大小及可选容量，这是为了确保

1.切片的大小和容量是否足够小

2.切片是否发生了逃逸，最终在堆中初始化。

如果切片非常大，运行时会直接在堆上初始化，如果切片不会发生逃逸并且非常小，例如小于等于4个元素，则直接在栈上或静态存储区创建数组。

> 大于32 KB 的对象会在堆中初始化。



#### 访问元素

使用`len`和`cap`可以获取切片的长度或容量。切片的操作基本都是在编译期间完成的，除了访问切片的长度，容量或其中的元素外，编译器也会将包含`range`关键字的遍历转换成形式更简单的循环。

```go
slice := make([]int,10)
fmt.Println(len(slice)) // 10
fmt.Println(cap(slice)) // 10
```



#### 追加和扩容

go 语言中切片食用 append 关键字向切片追加元素，在中间代码生成阶段会根据返回值是否会覆盖原变量，选择进入两种流程

1.第一种，会覆盖原切片

```go
slice := []int{1, 2, 3}
slice = append(slice, 1, 2, 3)
fmt.Println(slice) // [1 2 3 1 2 3]

```

下图来源《Go语言设计与实现》

![image-20231021003500649](https://raw.githubusercontent.com/AnAnonymousFriend/images/main/image-20231021003500649.png)

上图可以看到，当切片追加元素时如果容量不足，则会创建一个新切片并将旧切片与追加元素放入一个新的切片。如果我们选择覆盖原有的变量，就不需要担心切片发生拷贝影响性能，Go语言编译器会对这种常见的情况做出优化。

扩容策略：

1. 如果期望容量大于当前容量的两倍就会使用期望容量；
2. 如果当前切片的长度小于 1024 就会将容量翻倍；
3. 如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；

> 内存对齐：Go语言会将待申请的内存向上取整，让数组中的整数可以提高内存的分配效率并减少碎片。



#### 拷贝切片

```go
// 使用 copy(a,b) 的形式对切片进行拷贝

slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
fmt.Println(slice2)

```

大切片上执行拷贝操作时一定要注意对性能的影响，因为整块拷贝内存会占用非常多的资源。



## 哈希

数组用于表示元素的序列，哈希则表示的是键值对之间映射的关系。想要实现一个性能优异的哈希表，需要注意两个关键点——哈希函数和冲突解决办法。

```go
// 字面量初始化
hash := map[string]int{"a": 1}
fmt.Println(hash)

//make 初始化
hash := make(map[string]int)
fmt.Println(hash)
```

当创建的哈希被分配到栈上并且其容器小于`BUCKETSIZE = 8` 时，Go 语言在编译阶段会对小容量的哈希做优化。



#### 读取方式

```go
_ = hash[key]

for k,v := range hash{
  // k,v
}
```

上述两种方式读取哈希表的数据使用的函数和底层原理完全不同，前者需要知道哈希的键，后者遍历哈希表中的全部键值对，访问数据时候不需要知道哈希的键。



#### 实现哈希

**开放寻址法**是一种在哈希表中解决哈希碰撞的方法，这种方法的核心思想是：依次探测和比较数组中的元素以判断目标键值对是否存在于哈希表中。它底层实现哈希表的数据结构就是数组。

它的实现方式非常简单，假设我们有一个长度为5的数组，当我们在哈希表中新增一个键：key3，它会依次遍历整个数组，从[0]到[4]探测，直到找到目标键值或空闲内存。

> 开放寻址法中对性能影响最大的是**装载因子**，它是数组中元素的数量与数组大小的比值。随着装载因子的增加，线性探测的平均用时就会逐渐增加，这会影响哈希表的读写性能。当装载率超过 70% 之后，哈希表的性能就会急剧下降，而一旦装载率达到 100%，整个哈希表就会完全失效，这时查找和插入任意元素的时间复杂度都是 𝑂(𝑛) 的，这时需要遍历数组中的全部元素，所以在实现哈希表时一定要关注装载因子的变化。



**拉链法**是哈希表最常见的实现方式，大多数的编程语言都用拉链法实现哈希表。它的好处就是查找的长度比较短，各个用于存储节点的内存都是动态申请的，可以节约比较多的存储空间。

它的底层实现方式使用数组加链表，有一些编程语言会在拉链法的哈希中引入红黑树以优化性能，拉链法会使用链表数组作为哈希底层的数据结构，我们可以将它看成可以扩展的二维数组。



## 字符串

字符串其实是一片连续的内存空间，可以将它理解成一个由字符串组成的数组。

C 语言中的字符串使用字符数组 `char[]` 表示。数组会占用一片连续的内存空间，而内存空间存储的字节共同组成了字符串，Go 语言中的字符串只是一个只读的字节数组。

只读代表字符串只会分配到只读到内存空间，Go语言不支持直接修改 string 类型变量的内存空间，但是仍然可以通过 string 和 []byte 类型之间反复转换实现修改这一目的:

1. 先将这段内存拷贝到堆或者栈上；
2. 将变量的类型转换成 `[]byte` 后并修改字节数据；
3. 将修改后的字节数组转换回 `string`；

Go 语言的字符串可以作为哈希的键，所以如果哈希的键是可变的，不仅会增加哈希实现的复杂度，还可能会影响哈希的比较。

Go 语言中字符串和切片的结构体比较：字符串只少了一个表示容量的 Cap 字段。

```go
type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

在谈论切片的时候我们提到过一个代码例子，稍微更改代码示例后

```go
package main

import "fmt"

func main() {
	var data = "123"
	doWork(data)
	fmt.Println(data)
}

func doWork(data string) {
	data = "99"
}
```

这段代码并不会按照切片代码逻辑执行，不会输出 “99”，而是依然输出 “123”。这是因为 **字符串虽然和切片的结构体高度相似，但是字符串是一个只读的切片类型。所有在字符串上的写入操作都是通过拷贝实现的。**

正常情况下，运行时会调用 `copy`将多个字符串拷贝到目标字符串所在的内存空间。新的字符串是一片新的内存空间，与原来的字符串没有任何关联，所以要注意，如果拼接的字符串非常大，拷贝带来的性能损失是无法忽略的。遇到需要极致性能的场景一定要尽量减少类型转换的次数。



## 函数

函数在Go语言中算是一等公民。在Go语言函数中使用栈传递参数和返回值，这种方式能够降低实现的复杂度并支持多返回值，但是牺牲了函数调用的性能：

1.不需要考虑超过寄存器数量的参数应该如何传递

2.不需要考虑不同架构上的寄存器差异

3.函数入参和出参的内存空间需要在栈上进行分配

Go 语言使用栈作为参数和返回值传递的方法是综合考虑后的设计，选择这种设计意味着编译器会更加简单、更容易维护。

Go 语言在传递参数时使用传值还是传引用类型也是需要注意的点，不同的方式会影响在函数修改入参时是否会影响调用方看到的数据：

传值：函数调用时会对参数进行拷贝，被调用放和调用方两者持有不相关的两份数据；

```go
func main() {
	x := 1
	my_func(x)
	fmt.Println("x 值为:",x)
}

func my_func(i int) {
	i = 2
	fmt.Println("函数内调用：", i)
}

//输出：
函数内调用： 2
x 值为: 1
```

**Go 语言的整型和数组类型都是值传递的**，也就是在调用函数时会对内容进行拷贝。



传引用：函数调用时会传递参数的指针，被调用方和调用方两者持有相同的数据，任意一方做出的修改都会影响另一方

```go
func main() {
	x := 1
	my_func(&x)
	fmt.Println("x 值为:", x)
}

func my_func(i *int) {
	*i = 2
	fmt.Println("函数内调用：", *i)
}

//输出
函数内调用： 2
x 值为: 2
```

上述传引用的函数也侧面表明了：无论传递基础类型，结构体还是指针，都会对传递的参数进行拷贝。将指针作为参数传入某个函数时，函数内部会复制指针，也就是会同时出现两个指针指向原有的内存空间，所以 Go 语言中传指针也是传值。



## 接口

接口的本质是引入一个新的中间层，调用方可以通过接口与具体实现分离，解除上下游的耦合，上层模块不需要依赖下层的具体模块，只需要依赖一个约定好的接口。

接口还可以帮助我们隐藏底层实现，减少关注点。在计算机科学中，接口是比较抽象的概念，但是编程语言中接口的概念更具体。

Go 语言中的接口是隐式实现，只要实现了接口里面的方法就相当于实现了接口。Go语言只会在传递参数，返回参数以及变量赋值才会对某个类型是否实现接口进行检查

```go
// 声明一个接口
type error interface {
	Error() string
}

// 声明一个结构体
type RPCError struct {
	Code    int
	Message string
}

// 实现接口
func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

// 调用接口
func main() {
	var rpcErr error = &RPCError{
		Code:    404,
		Message: "Not Find",
	}
	fmt.Println(rpcErr.Error())
}
```



Go语言中，接口一种是带有一组方法的接口，另一种是不带任何方法的`interface{}`。后者在Go语言中很常见，所以在实现时使用了特殊类型。要注意 `interface{}`类型**不是任意类型**。如果我们将类型转换成了 `interface{}` 类型，变量在运行期间的类型也会发生变化，获取变量类型时会得到 `interface{}`。

在《Go语言设计与实现》中有一个例子

```go
package main

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	return v == nil
}

func main() {
	var s *TestStruct
	fmt.Println(s == nil)      // #=> true
	fmt.Println(NilOrNot(s))   // #=> false
}

$ go run main.go
true
false
```

可以看到上述输出两个不同的结果，这是因为调用 NilOrNot 函数时发生了**隐式的类型转换**，除了向方法传入参数之外，变量的赋值也会触发隐式类型转换。在类型转换时，`*TestStruct` 类型会转换成 `interface{}` 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 `TestStruct`，所以转换后的变量与 `nil` 不相等。

使用结构体实现接口带来的开销会大于使用指针实现，而动态派发在结构体上的表现非常差，这也提醒我们应当尽量避免使用结构体类型实现接口。

使用结构体带来的巨大性能差异不只是接口带来的问题，带来性能问题主要因为 Go 语言在函数调用时是传值的，动态派发的过程只是放大了参数拷贝带来的影响。



## For 和 range

对于数组和切片来说，Go 语言有三种不同的遍历方式，这三种不同的遍历方式分别对应着代码中的不同条件

1. 分析遍历数组和切片清空元素的情况；

   Go 语言会直接使用内置函数(runtime函数)清空目标数组内存空间中的全部数据，并在执行完成后更新遍历数组的索引。

2. 分析使用 `for range a {}` 遍历数组和切片，不关心索引和数据的情况；

3. 分析使用 `for i := range a {}` 遍历数组和切片，只关心索引的情况；

4. 分析使用 `for i, elem := range a {}` 遍历数组和切片，关心索引和数据的情况；



如果同时遍历索引和元素的range循环时，Go会额外创建一个新的变量存储切片中的元素，循环中使用的这个变量会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝。

```go
// 错误写法
func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}



// 正确写法
func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
```



## Select

Go 语言中的 `select` 也能够让 Goroutine 同时等待多个 Channel 可读或者可写，在多个文件或者 Channel状态改变之前，`select` 会一直阻塞当前线程或者 Goroutine。

`select` 是与 `switch` 相似的控制结构，与 `switch` 不同的是，`select` 中虽然也有多个 `case`，但是这些 `case` 中的表达式必须都是 Channel 的收发操作。

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
```

1. `select` 能在 Channel 上进行非阻塞的收发操作；
2. `select` 在遇到多个 Channel 同时响应时，会随机执行一种情况；



### 非阻塞的收发

在通常情况下，`select` 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。但是如果 `select` 控制结构中包含 `default` 语句，那么这个 `select` 语句在执行时会遇到以下两种情况：

1. 当存在可以收发的 Channel 时，直接处理该 Channel 对应的 `case`；
2. 当不存在可以收发的 Channel 时，执行 `default` 中的语句；

```
func main() {
	ch := make(chan int)
	select {
	case i := <-ch:
		fmt.Println("i:", i)
	default:
		fmt.Println("default")
	}
}
```

随机执行

`select` 在遇到多个 `<-ch` 同时满足可读或者可写条件时会随机选择一个 `case` 执行其中的代码。

```go
func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

//输出:
case2
case1
case2
case2
case1
case2
...
```

两个 `case` 都是同时满足执行条件的，如果我们按照顺序依次判断，那么后面的条件永远都会得不到执行，而随机的引入就是为了避免饥饿问题的发生。

第一种情况：`select` 不存在任何的 `case`，空的 `select` 语句会直接阻塞当前 Goroutine，导致 Goroutine 进入无法被唤醒的永久休眠状态。

第二种情况：`select` 只存在一个 `case`，编译器会将`select` 语句改写为 if 条件语句。当 `case` 中的 Channel 是空指针时，会直接挂起当前 Goroutine 并陷入永久休眠。

第三种情况：`select` 存在两个 `case`，其中一个 `case` 是 `default`。编译器认为这是一次非阻塞的收发操作，该函数会将 `case` 中的所有 Channel 都转换成指向 Channel 的地址。

第四种情况：`select` 存在多个 `case`。编译器会编译成多个 if 语句执行对应 case 的代码。



## Defer

Go 语言的 `defer` 会在当前函数返回前执行传入的函数，它会经常被用于关闭文件描述符、关闭数据库连接以及解锁资源。

使用`defer`一般是在函数调用结束后完成一些收尾工作。

### 作用域

向 `defer` 关键字传入的函数会在函数返回之前运行。假设我们在 `for` 循环中多次调用 `defer` 关键字：

可以看到下例代码输出是倒序，可以把`defer`的执行顺序看成一个出栈的顺序，即最后加入栈的最先出。是这里要注意，如果函数中包含 return ，会先执行 return ，再执行 defer 。如果函数中包含 **panic** 函数，那么会先执行 defer 函数，最后再执行 panic 函数。

```go
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}


// 输出
4
3
2
1
0


func main() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Println("结束")
			return
		}
		defer fmt.Println(i)
	}
}

// 输出
结束
3
2
1
0
```

同时`defer` 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用。

`defer`关键字会拷贝函数中引用外部参数，在调用`defer`关键字的时候就会进行计算(defer也继承了函数调用传值的特性)。

默认情况下Go语言中defer结构体都会在堆上分配，分配在堆上的方案是一个保底方案。但是除了分配的位置不同，本质上没有什么不同，除了分配在栈上可以节约额外开销。



**执行顺序**

一个函数中，多个 defer 的执行顺序为 “后进先出”，但是这里要注意，如果函数中包含 return ，会先执行 return ，再执行 defer 。如果函数中包含 **panic** 函数，那么会先执行 defer 函数，最后再执行 panic 函数。

**defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体。**



## panic 和 recover

`panic` 能够改变程序的控制流，调用 `panic` 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 `defer`；

`recover` 可以中止 `panic` 造成的程序崩溃。它是一个只能在 `defer` 中发挥作用的函数，在其他作用域中调用不会发挥作用；

也就是说: panic 只会触发当前 Goroutine 的 defer ，而 recover 只有在defer 中调用才会生效

panic 允许在 defer 中嵌套多次调用。

```go
func main() {
	defer println("in main")

	go func() {
		defer println("in goroutine")
		panic("error !")
	}()

	time.Sleep(1 * time.Second)
}

// 输出
in goroutine
panic: error !

```

要注意：main函数中的defer 语句没有执行，执行的只有当前Goroutine 中的 `defer`。多个 Goroutine 之间没有太多的关联，一个 Goroutine 在 `panic` 时也不应该执行其他 Goroutine 的延迟函数。

`recover` 只有在发生 `panic` 之后调用才会生效：

```go
// 必须要先声明defer，否则不能捕获到panic异常
defer func() { 
		if err := recover(); err != nil {
			fmt.Println("err info:", err) // 这里的err其实就是panic传入的内容
		}
	}()
panic("异常信息")
```

在 Goroutine 中使用 recover 和 panic

```go
func main() {
	go test()
	fmt.Println("in main")
	time.Sleep(2 * time.Second)
}

func test() {
	defer func() { 
		if err := recover(); err != nil {
			fmt.Println("err info:", err) 
		}
	}()

	panic("test 异常信息")
}
```





# 并发编程

## 上下文Conetext

Context 是Go 语言中独特的设计。它用来设置截止日期，同步信号，传递请求相关值的结构体。

在 Goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费是 [`context.Context`](https://draveness.me/golang/tree/context.Context) 的最大作用。Go 服务的每一个请求都是通过单独的 Goroutine 处理的[2](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/#fn:2)，HTTP/RPC 请求的处理器会启动新的 Goroutine 访问数据库和其他服务。

我们可能会创建多个 Goroutine 来处理一次请求，而 [`context.Context`](https://draveness.me/golang/tree/context.Context) 的作用是在不同 Goroutine 之间同步请求特定数据、取消信号以及处理请求的截止日期。

每一个 [`context.Context`](https://draveness.me/golang/tree/context.Context) 都会从最顶层的 Goroutine 一层一层传递到最下层。[`context.Context`](https://draveness.me/golang/tree/context.Context) 可以在上层 Goroutine 执行出现错误时，将信号及时同步给下层。

这样设计的好处就是：如果最上层的Goroutine出现某些原因执行失败了，可以通过 `context.Context`在下层及时停掉无用的工作以减少额外资源的消耗。

多个Goroutine同时订阅 ctx.Done()管道中的消息，一旦接受道取消信号就立即停止当前正在执行的工作。

```go
func main() {
	ctx, canel := context.WithTimeout(context.Background(), 1*time.Second) // 设置一个超时的上下文
	defer canel()

  // 设置子任务超时时间
	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
```

从源代码来看，[`context.Background`](https://draveness.me/golang/tree/context.Background) 和 [`context.TODO`](https://draveness.me/golang/tree/context.TODO) 也只是互为别名，没有太大的差别，只是在使用和语义上稍有不同：

- [`context.Background`](https://draveness.me/golang/tree/context.Background) 是上下文的默认值，所有其他的上下文都应该从它衍生出来；
- [`context.TODO`](https://draveness.me/golang/tree/context.TODO) 应该仅在不确定应该使用哪种上下文时使用；

在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 [`context.Background`](https://draveness.me/golang/tree/context.Background) 作为起始的上下文向下传递。

Go 语言中的 [`context.Context`](https://draveness.me/golang/tree/context.Context) 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用，虽然它也有传值的功能，但是这个功能我们还是很少用到。

在真正使用传值的功能时我们也应该非常谨慎，使用 [`context.Context`](https://draveness.me/golang/tree/context.Context) 传递请求的所有参数一种非常差的设计，比较常见的使用场景是传递请求对应用户的认证令牌以及用于进行分布式追踪的请求 ID。



## Channel





## Sync.Mutex 

互斥锁是并发控制的一个基本手段，是为了避免竞争而建立的一种并发控制机制。当一个公共变量被多个Goroutine所访问，为了避免并发访问导致意想不到的结果，使用互斥锁让公共变量只能同时由一个线程持有。

当一个变量被某个线程持有时，其他线程如果想访问这个变量，会访问失败或等待。直到持有这个变量的线程释放**锁**，其他线程才有机会获取这个变量。

Mutex 是使用最广泛的同步原语，所以我们从互斥锁开始，再到读写锁，并发编排等。在Go标准库中 sync 提供锁等一系列同步原语。

```go
func main() {
	var wg sync.WaitGroup
	count := 0
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			count++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

//输出：
9347
```

上述代码中使用了多个协程访问同一个变量，可以看到输出结果是9347，这并不是我们想要的。如果使用`mu.Lock()`和 `mu.unLock()` 来安全访问公共变量：

```go
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

//输出：
10000
```





还有一种方式是：Mutex 嵌入







# 面试题

## **Go 语言中 new 和 make 的区别**

new 和 make 都是 分配内存的原语。new 只分配内存但并不初始化内存，而 make 用于 slice , map 和 channel 的初始化。

slice , map , channel 类型属于引用类型，go 会给引用类型初始化为 nil , 所以 make 不仅可以开辟一个内存，还能给找个内存的类型初始化其零值。

make 只能用来分配及初始化类型为 slice, map , channel 的数据。new 可以分配任意类型的数据。

make 返回的还是引用类型本身；而 new 返回的是指向类型的指针。



## **数组和切片的区别**

数组类型的值的长度必须在声明的时候给定，并且之后不会再改变。

切片可以自动扩容，我们可以将切片理解成一片连续的内存空间加上长度与容量的标识。

**切片引入了一个抽象层，提供了对数组中部分连续的片段引用，**



## **数组相比切片有什么优势**

**可比较**：数组是固定长度，它们之间是可以比较的，数组是**值对象**。切片不可以直接比较，也不能用于判断。数组可以作为 map 的 **键**（key）, 而切片不行。

**编译安全**：数组可以提供更高的编译时安全，可以在编译时检查索引范围。

**规划内存布局**：更好控制内存布局，因为不能直接在带有切片的结构中分配空间，所以可以使用数组来解决。

**访问速度**：其访问（单个）数组元素比访问切片元素更高效，时间复杂度是 O (1)

更多细节：https://eddycjy.com/posts/go/go-array-slice/



## 切片会输出什么结果？

```go
package main

import "fmt"

func main() {
	var data = make([]int, 3, 3)
	doWork(data)
	fmt.Println(data)
}

func doWork(data []int) {
	data = append(data, 1)
	data[0] = 1
}
```

这里要注意两个点：

1.Go 语言函数传值，**无论是传递基本类型，结构体还是指针，都会对传递的参数进行拷贝。**

2.切片的扩容机制

这里可以先看一下切片的数据结构:

```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

从切片的数据结构可以看出，Data 是一个内置的指针类型，可直接更改引用地址的参数。所以虽然函数使用值传递，但是在函数内部执行 date[0] = 1 ，外部的切片数据也会发生变化。

但是 Len 和 Cap 是 int 类型，这代表，函数内部更改不会影响到函数外的数据源。

回到题目本身，data 初始化后是 [0,0,0] 值传递到 doWork 函数后进行扩容，数据发生变化 data 为 [0,0,0,0] 后修改切片的第一个元素 [1,0,0,0]，但是这是在 doWork 函数中内的数据变化，并不会影响到 main 函数中 data 的值。

切片的扩容是为切片分配新的内存空间并复制原切片中元素的过程。如果切片中的元素不是指针类型，那么会将原数组内存中的内容复制到新申请的内容中，这将最终会返回一个新切片，并覆盖原切片。

> 所以在使用 append 函数对切片进行扩容后，需要一个变量去接受它的新切片。



> **遇到大切片扩容或复制的时候，可能会引发大规模的内存复制，一定要减少类似的操作以避免影响程序的性能。**





## **Map的线程安全**

Go 内建的 map 对象不是线程安全的，并发读写的时候运行时会有检查，遇到并发问题就会导致 panic 。

解决 Map 的线程安全有多个方案：1.互斥锁 2.读写锁 [3.Sync.Map](http://3.Sync.Map) 4.分片加锁

前两个方案不用过多赘述，可以重点谈论一下后两个方案。

Go 内建的 map 类型不是线程安全的，而 [Sync.Map](http://sync.Map) 并不是来替换内建的 map 类型的，它只能被应用在一些特殊场景内

1.只会增长的缓存系统中，一个 key 只写入一次而被读很多次。

2.多个 goroutine 为不相交的键读，写 和 重写键值对。

优点：

1.空间换时间。通过冗余的两个数据结构（只读的 read 字段，可写的 dirty ）,来减少加锁对性能的影响。对只读字段（read）的操作不需要加锁。

2.优先从 read 字段读取，更新，删除，因为对read字段的读取不需要锁。

3.动态调整。miss 次数过多，将 dirty 数据提升为read，避免总是从 dirty 中加锁读取。

4.double-checking。加锁之后还要再检查 read 字段，确定真的不存在才操作 dirty 字段。

5.延迟删除。删除一个键值只是打标记，只有在提升dirty 字段为 read 字段的时候才清理删除的数据。

**分片加锁 可看**：https://github.com/orcaman/concurrent-map





## **反射**

Go 语言中反射的第一法则：**我们能将 Go 语言的 interface{} 变量转换成反射对象。因为函数的调用都是值传递，所以变量类型在底层函数调用时进行类型转换。所以会从基本类型转换到 interface{}**

第二法则：我们可以从反射对象获取 interface{} 变量。

第三法则：我们得到的反射对象跟原对象没有任何关系，那么直接修改反射对象无法改变原变量，程序为了防止错误就会崩溃。











## **Channel**

先从 Channel 读取数据的 Goroutine 会先接收到数据

先向 Channel 发送数据的 Goroutine 会得到先发送数据的权利

Channel 在运行时使用 runtime.hchan 结构体：

```go
type hchan struct {
	qcount   uint   // Channel 中的元素个数
	dataqsiz uint   // Channel 中的循环队列的长度
	buf      unsafe.Pointer  // Channel 的缓冲区数据指针
	elemsize uint16 // Channel 能够手法的元素大小  
	closed   uint32
	elemtype *_type // Channel 能够手法的元素类型
	sendx    uint   // Channel 的发送操作处理到的位置       
	recvx    uint   // Channel 的接收操作处理到的位置
	recvq    waitq  // 存储当前 Channel 由于缓冲区空间不足而阻塞的 Goroutine 列表
	sendq    waitq

	lock mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
```

Channel 是一个用于同步和通信的有锁队列。

> 向一个已经关闭的 Channel 发送数据时，会报告错误并中止程序。 向一个已经关闭的 Channel （无缓存）读数据时，会读取到零值。 向一个已经关闭的 Channel （有缓存） 读取数据时，会读取通道里面的剩余值。剩余值读取完后会读到零值。



**Goroutine 的泄露**

如果启动了一个 goroutine ，但是没有符合预期地退出，直到程序结束，此 goroutine 才退出，这种情况叫做 goroutine 泄露。

一般泄露是因为 Channel 操作阻塞导致整个 goroutine 一直阻塞等待或 goroutine 里有死循环。



> 共享资源的并发访问使用传统并发原语 复杂的任务编排和消息传递使用 Channel 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond 简单等待所有任务的完成用 WaitGroup ，也有 Channel 的推崇者用 Channel，都可以使用 需要和 Select 语句结合，使用 Channel 需要和超时配合时，使用 Channel 和 Context



## 学习资料 

《Go 并发编程实战》

《Go语言高性能编程》

《Go 语言设计与实现》