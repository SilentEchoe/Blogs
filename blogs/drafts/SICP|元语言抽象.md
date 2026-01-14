---
title: SICP|元语言抽象
date: 2026-01-06 17:20:00
tags: [计算机基础,读书笔记]
category: 计算机基础
---

> 无论哪一种确定的程序设计语言，都不足以满足我们的需要。我们必须经常转向新的语言，以便更有效地表述自己的想法。建立新语言是在工程设计中控制复杂性的一种威力强大的策略，我们常常能通过采用一种新语言而提升处理复杂问题的能力，因为新的语言可能使我们以不同的路径，使用不同的原语、不同的组合方法和抽象方法，去描述（因此也是去思考）面对的问题，这些都可以是为了处理手头的问题而专门打造的。

元语言抽象(Metalinguistic Abstraction)到创造新语言，这在工程设计中扮演着重要角色。DSL(领域专用语言) 是这种能力在工程中的一种常见产物，这使工程师不仅可以使用现有的高级编程语言，还可以通过编写**解释器**来创造新语言，比如Go语言最初的解释器是C语言写的，在经历数个版本后才使用Go语言实现**自举**。

解释器本质上就是一个接受“程序”(表达式)作为输入并执行它的过程。编程语言本身也可以被程序操控和实现，它定义了某种语言中表达式的含义。

SICP中使用 Scheme 实现了 Scheme 的解释器，这是为了展示**语言与解释器的关系**。解释器（或编译器）是对语言语义的一种可执行定义；不同实现必须遵循语言规范，从而保持语义一致。如果用另外一门语言编写一个解释器来执行某种语言的程序，那么这个解释器本身就是对改语言的一种定义。



### 解释器的基本结构

解释器的核心在于两个互相递归调用的过程: eval(求值)和apply(应用)，eval在某个环境中计算一个表达式的值，apply则将一个过程作用于一组参数值。这两个过程构成了解释器的基本工作循环：eval 将表达式的求值转换为对过程的调用，apply 在调用过程中产生新的表达式需要求值，通过互相递归调用直到得出最终结果(如原始数据类型的值或基本过程的执行结果)。

在任何解释器中处理不同的表达式逻辑都不相同，比如自求值表达式（如数字、布尔值）其值就是它自身，1就是1；

> 环境是一系列**帧**（frame）的集合，每个帧包含一些符号到值的绑定，并可能有指向外围环境的链接。如果在当前环境中找不到变量，会沿着链寻找全局环境，找不到则报未绑定错误。

复合表达式，比如函数调用则需要先计算运算符和操作数，再将运算符得到的过程应用与操作数的值。这进一步引出了**应用(apply)**:如果有一个过程(函数)和参数值列表，那么apply负责执行这个过程，假设它是最基础的加法运算，apply就会对参数求和；如果过程是用户定义的函数，apply则会创建新环境将参数值绑定到函数的形参，再执行函数体。

条件表达式，比如If,eval 对 if 条件语句会先求值，根据结果真/假选择性地求值后续分支。

Lambda(匿名函数)， 解释器不会立即执行匿名函数中的主体内容，而是创建一个**过程对象(闭包)**，记录参数列表，函数体以及当前环境作为闭包的“外围环境”。在 Scheme 实现中，这可以用数据结构封装三个部分。之后eval返回这个过程对象，以便在后面应用。

```go
import "fmt"

type Number float64
type Symbol string
type Expr interface{}      
type List []Expr


type Value interface{}     

type Proc struct {
    params []string
    body   Expr
    env    *Env 
}

// 环境：保存变量绑定的结构，支持嵌套（链表结构）
type Env struct {
    vars  map[string]Value  
    outer *Env             //外层环境
}

// 在环境中查找变量的值
func (env *Env) Lookup(name string) Value {
    if val, ok := env.vars[name]; ok {
        return val  
    }
    if env.outer != nil {
        return env.outer.Lookup(name)  
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
        return float64(e)
    case Symbol:
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
func Apply(proc Value, args []Value) Value {
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

globalEnv.vars["+"] = func(args []Value) Value {
    sum := 0.0
    for _, arg := range args {
        sum += arg.(float64)
    }
    return sum
}

globalEnv.vars["="] = func(args []Value) Value {
    if len(args) != 2 {
        panic("= 需要2个参数")
    }
    return args[0] == args[1]
}

expr := List{Symbol("+"), Number(1), Number(2)}

// 在 globalEnv 环境中求值该表达式
result := Eval(expr, globalEnv)
fmt.Println(result)
$ 输出结果: 3
```

List 这个新的结构体是描述了一个列表表达式，表示想要计算的信息，这个表达式作为Go数据直接传递给 eval 函数。eval 会识别出这是一个列表表达式，它会按照预设的逻辑先识别出运算符是 Symbol("+") 然后在 globalEnv 中找到符号 + 对应的过程，接着求值两个参数 Number(1) 和 Number(2) 得到 1.0 和 2.0，最后调用加法函数将它们相加。

整个过程中，表达式始终以数据结构形式存在，解释器通过检查和拆解这些结构来理解代码的含义。这个例子进一步说明了解释器将代码看作普通的数据来处理，Symbol 和 Number是数据，List将它们组合起来来表示复合表达式。这种表示让解析和执行变得统一。

在实现中，甚至可以编写代码去动态地构造或修改这些表达式数据，再交给解释器运行——这就是元编程和宏系统的基础思想。



### 环境模型

环境是解释器执行过程中用来跟踪**变量名和它们对应值**的机制。在SICP中，环境模型将环境定义为一个**链式的帧结构**，每个帧包含一组变量绑定（名字-值对），同时指向一个**外层环境。当解释器需要查找一个变量的值时，会首先在当前环境帧查找，如果找不到就沿着链向外层查找，直到找到该变量或达到最外层环境 。这种结构自然地支持了**词法作用域，即函数可以访问其定义所在环境中的变量。环境模型是理解函数闭包（函数携带引用环境）和可变状态（赋值如何影响环境）的关键。

> SICP中的**环境(**Environment**)**可以理解为**可叠加的字典**。每一层被视为一个**Frame(框架)**，本质上就是存放一组“名字-值”的字典，这些字典按层级连成链。检索时就会一层一层寻址，当前层找不到时就会去外层寻找，一直找到“全局层”。可以想象站在迷宫的最中心，从内一层一层向外探寻。

```go
// Env 增加一个方法：为已有变量赋新值（若未定义则抛错）
func (env *Env) Set(name string, val Value) {
    if _, ok := env.vars[name]; ok {
        env.vars[name] = val
    } else if env.outer != nil {
        env.outer.Set(name, val) 
    } else {
        panic("赋值出错，变量未定义: " + name)
    }
}

// 在全局环境中定义一个变量 x
globalEnv.vars["x"] = 10.0
fmt.Println(globalEnv.Lookup("x"))  // 输出: 10

// 创建一个嵌套环境，外层指向 globalEnv
localEnv := NewEnv(globalEnv)
// 在嵌套环境中新建变量 y
localEnv.vars["y"] = 5.0


fmt.Println(localEnv.Lookup("y"))
$ 输出: 5
fmt.Println(localEnv.Lookup("x")) 
$ 输出: 10
localEnv.Set("x", 20.0)
fmt.Println(globalEnv.Lookup("x"))
$ 输出: 20 
```

上述代码展示了环境的层级关系和作用域规则：创建了一个 localEnv 指向 globalEnv，此时的localEnv相当于一个函数调用时产生的局部环境，其中定义的新变量 y 只存于 localEnv 中。当在 localEnv 中查找 x 时，由于当前层没有 x，会自动转向它的外层 globalEnv 找到已定义的 x。这说明 localEnv 可以使用它外部环境中的定义。localEnv.Set("x", 20.0) 来给外层的 x 变量赋新值，这也表明赋值操作会沿着环境链找变量。一旦找到，就在那个环境帧中更新它。结果，全局的 x 从10变为20。

通过环境机制可以让解释器正确处理变量名解析和函数闭包。在原书中环境作为帧链表，Lookup 沿链搜索，NewEnv 创建新帧指向外层，Set 沿链赋值。同时强调**闭包**要包含一个指向它创建时环境的引用，以保证在执行时能找到非局部变量。



### 求值器的运行过程

将一个程序交给解释器时，eval-apply 循环会层层展开计算过程。在这个过程中，递归非常重要：解释器会先求值最内层的表达式，然后逐步完成外层的运算。比如表达式：`(+ 1 (* 2 3))` 解释器会先识别这是一个加法调用，但是它必须要先算出`(* 2 3)`的值，所以它递归调用自身去计算乘法表达式。在计算乘法时又要先拿到 2 和 3 的值，因为它们是原子数据，所以可以直接得到。然后进行乘法得到 6；返回加法表达式，再取出先前得到的 6 和 1 相加最终得到结果 7.

整个求值过程就是不断递归计算子表达式，然后回溯应用运算符的过程。这类似数学上对表达式求值的过程，只不过是通过解释器的实现来自动完成。

```go
// 修改 Eval 函数以打印调试信息
func Eval(expr Expr, env *Env) Value {
    fmt.Printf("%s求值：%v\n", currentIndent(), expr)  
    evalDepth++                                      
    defer func() { evalDepth-- }()                   

    switch e := expr.(type) {
    ...
    case List:
        ...
        proc := Eval(e[0], env)
        var argVals []Value
        for _, arg := range e[1:] {
            argVals = append(argVals, Eval(arg, env))
        }
        result := Apply(proc, argVals, env)
        fmt.Printf("%s完成：%v => %v\n", currentIndent(), expr, result)  
        return result
    }
     ...
}

// 工具函数：根据当前递归深度返回相应个数的缩进
var evalDepth = 0
func currentIndent() string {
    return fmt.Sprintf("%s", strings.Repeat("  ", evalDepth))
}

// 测试对复合表达式的求值过程
expr := List{Symbol("+"), Number(2), List{Symbol("*"), Number(3), Number(4)}}
result := Eval(expr, globalEnv)
fmt.Println("最终结果:", result)
```

SICP中提到：可以把程序看作一台机器的描述，而解释器就是读取这个描述并按部就班执行的机器。



### 变量、赋值与副作用

赋值语句会改变变量在其环境中的绑定，使之指向一个新值。赋值不直接产生计算结果，但是却改变了**状态**。而副作用是指程序执行时除了返回值之外对状态的改变，比如修改了变量值或输出了信息。引入可变变量会使程序的行为不再只由输入决定，也与执行顺序相关，这就是副作用带来的复杂性。

引入赋值后，就有**状态**的概念了：程序执行的后续部分可能受到之前赋值改变的环境的影响。例如，连加两次同一变量，第一次赋值会影响第二次的结果。这就需要工程师理解执行顺序和可变性，同时也为解释器带来了需要考虑的额外语义（如需要支持顺序执行，避免重新求值时状态变化产生干扰）。

```go
func Eval(expr Expr, env *Env) Value {
    ...
    case List:
        firstSym, _ := e[0].(Symbol)
        switch string(firstSym) {
        case "set!":
        
            // 模拟SICP 中的赋值
            varName := string(e[1].(Symbol))
            newVal := Eval(e[2], env)      // 计算赋值的值
            env.Set(varName, newVal)       // 更新环境中的变量绑定
            return newVal                 // 返回赋值后的值（或可返回特殊标记）
       
        default:
            ...
        }
    }
}


globalEnv.vars["x"] = 1.0

// 执行赋值表达式 (set! x (+ x 5))
expr := List{Symbol("set!"), Symbol("x"),
             List{Symbol("+"), Symbol("x"), Number(5)}}
result := Eval(expr, globalEnv)
fmt.Println("赋值结果:", result)
fmt.Println("x的新值:", globalEnv.Lookup("x"))
$ 输出：赋值结果: 6 ; x的新值: 6
```

赋值引入了**状态改变**：执行赋值前后，同样的变量 x 有不同的值。如果重复执行 (set! x (+ x 5))，结果会依次改变 x 的值为 11、16 等。副作用也意味着**执行顺序**变得重要：如果交换两条赋值语句，程序结果可能不同。在构造解释器时要正确维护环境状态，让后续求值看到的是最新的变量绑定。



### 块结构

块结构指的是**具有局部作用域的代码块**。比如在函数内部定义局部变量或内部函数。很多高级编程语言允许在一个作用内声明新的变量，这些变量只在该作用域（块）中可见。块结构体现了**词法作用域**的一种组织方式，使工程师能封装内部实现细节。

对于解释器来说，支持块结构需要特别处理内部定义。如果遇到内部定义时，需要把这些定义的名字加入当前的环境帧，这样在后续同个函数体中，其他的表达式才能访问这些局部定义。同时还要保证定义只影响局部环境而不污染全局。

SICP中为了实现上述，提出了两种实现方式：1.求值时顺序处理定义，当进入一个过程体环境时，先顺序求值所有内部定义的右侧表达式，将名字绑定，再继续执行函数体剩余部分。这相当于在进入块时**执行一系列赋值**。2.语法分析和执行分离，在进入过程体时，预先扫描整个块找出所有内部定义的名字，在环境中占位，再执行体。类似编译器在编译阶段分配局部变量，然后运行时直接使用的思路。这种方式避免了每次执行都重复分析内部定义，提高效率。



### 构造求值器

大部分编程语言都使用立即求值的策略：**调用函数时先把所有参数表达式求值，再进入函数。**但在SICP中提出了另外一种策略：惰性求值。意味着推迟计算参数，直到真正需要其值踩计算，通过这样的策略可以处理一些无限结构，也可以避免不必要的计算。

要实现惰性求值，这要求解释器在做函数调用时不再立即求值参数表达式，而是将参数转换为一种惰性对象封装起来，延迟它的计算。当真正需要这个参数时再对惰性对象进行求值计算，并把结果缓存起来。

```go
// 严格求值的条件函数：参数在调用前就被求值
func StrictIf(cond bool, a int, b int) int {
    if cond {
        return a
    } else {
        return b
    }
}

// 惰性求值的条件函数：使用函数闭包延迟参数计算
func LazyIf(cond bool, a func() int, b func() int) int {
    if cond {
        return a() 
    } else {
        return b()
    }
}

// 测试 StrictIf 和 LazyIf
val1 := StrictIf(true, 1, 1/0) 
val2 := LazyIf(true,
               func() int { return 1 },
               func() int { return 1/0 })
fmt.Println("LazyIf结果:", val2)
$  输出 LazyIf结果: 1
```

这里只是用一个例子来说明惰性求值的效果，StrictIf(true, 1, 1/0) 会造成 Go panic，因为 1/0在传入函数前就被求值（除零异常）。而 LazyIf(true, func(){return 1}, func(){return 1/0}) 则能够正常返回1，因为cond为真，所以只调用了第一个函数 a() 返回1，第二个函数 b()根本没有被执行。使用这种手动延迟计算的方式模拟了惰性求值。

原书中通过更改元循环求值器，实现了正常序(惰性)求值，它的核心就是将过程应用时对参数求值的动作换成了创建Thunk。包括缓存的具体实现，但在这里不再赘述。



### 非确定性求值

非确定性求值允许一个表达式有不止一个可能的值。程序在某处可以非确定性地，在多个选项中选择一个继续执行。计算的过程也不是线性的，而是带有探索和回溯（backtracking）：求值器会在需要时搜索所有可能的选择，直到找到满足条件的结果或者发现无解。

这种方式可以很好地解决搜索类的问题，而且因为解释器实现了它，意味着在底层的解释器会尝试各种组合。

在具体的实现上，amb (*ambiguous*) 可以被视为产生“选择点”的操作：它从一组给定的候选值里选一个继续执行。如果后续计算发现走不通（不满足条件），解释器会**回溯**到这个选择点，尝试下一个备选值，如此直到找到解或耗尽选项。为了实现回溯，求值器需要保存执行的分叉点和状态，并在失败时恢复。这通常可以用高级控制结构如**continuation（续延）**或者回溯算法来完成。

SICP中直接修改了解释器，让每个计算分支“携带”“成功继续”和“失败继续”两个过程，当一个分支走不通时调用失败继续跳回先前的选择点。

```go
func choose(options []int, f func(int) bool) bool {
    for _, opt := range options {
        if f(opt) {
            return true 
        }
    }
    return false 
}

func main() {
    var solX, solY int
    found := choose([]int{1, 2, 3, 7, 8}, func(x int) bool {
        return choose([]int{2, 5, 8}, func(y int) bool {
            if x+y == 10 {
                solX, solY = x, y
                return true 
            }
            return false  
        })
    })
    if found {
        fmt.Printf("找到一组解: x=%d, y=%d\n", solX, solY)
    } else {
        fmt.Println("没有找到解")
    }
}
```

这点很像图算法，虽然Go中没有continuation(延续)，但是可以通过递归来穷举。在上述例子中函数choose会尝试将列表中的每个选项opt送入f，如果f(opt)返回真就立即停止并返回真，否则继续。

用了两个嵌套的choose来模拟两个amb选择点：外层为x的选择，内层为y的选择。f内部检查条件 if x+y==10，一旦条件满足，就保存解并返回真来停止选择。这个结构正对应上面的伪代码：choose相当于amb操作符，而条件判断相当于require（或称 assert）失败就触发回溯。

需要强调的是，上面的 choose 更像“存在性搜索（find first solution）”：它返回的是“是否存在一组满足条件的解”，一旦找到一组解就立刻停止，并不会继续枚举更多解。

而 SICP 的 amb 是把“选择点”和“回溯机制”内建到求值器里：每个分支不仅携带“成功后怎么继续”的逻辑（success continuation），还携带“失败后回到哪里、继续尝试下一个分支”的逻辑（fail continuation）。因此 amb 天然支持在找到一个解之后，继续调用 fail continuation 回溯，从而枚举第二个、第三个解。这也是 amb 更像一种语言原语（由解释器统一管理搜索树）而不仅仅是一个普通的递归穷举函数。

当然这个例子并不十分严谨，如果要集成在解释器中需要做更加通用细致的处理。原书中使用更强大的技术，如continuation-passing style 让求值器自动回溯，但是这里只是作为一个了解思维的入口。

非确定性求值器赋予了语言**自动搜索**能力，使表达程序的方式更接近逻辑推理而非过程指令。通过非确定性求值，只需要列出可能值集合和约束条件，求值器就会输出令条件为真的值组合。
