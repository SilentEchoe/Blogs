---
title: Golang限速器
date: 2023-3-6 14:54:00
tags: [Go,学习笔记]
category: Go
---

### 限速器

在高并发的系统中,限流已经成为必不可少的功能，它成为提升服务稳定性非常重要的组件，可以用于限制请求速率，保护服务以免服务过载。

限流器有很多种实现方式，常见的限流算法有**固定窗口,滑动窗口,漏桶，令牌桶等**。Golang 官方提供的扩展库自带了限流算法的实现: golang.org/x/time/rate 使用的是令牌桶。



构建限流器

```go
// NewLimiter returns a new Limiter that allows events up to rate r and permits
// bursts of at most b tokens.
func NewLimiter(r Limit, b int) *Limiter {
	return &Limiter{
		limit: r,
		burst: b,
	}
}

// r Limit 设置限流器Limiter的Limit字段，代表每秒可以向Token桶中产生多少Token，Limit实际上是float64的别名
// b int 代表Token桶的容量大小,这代表:构建出一个大小为100的令牌桶，以每秒1个Token的速率向桶中放置Token
limiter := rate.NewLimiter(1, 100);
```

使用限流器

Limiter提供了三类方法供程序消费Token,可以每次消费一个消费,也可以每次消费多个Token

Wait/waitN (Wait 实际上就是 `WaitN(ctx,1)`)

```go
// 使用Wait方法消费Token时,如果此时桶内Token数组不足(小于N),那么Wait方法会组塞一段时间,直到Token满足条件
// Wait 方法存有一个 context参数，可以设置 context的 Deadline 或者 Timeout,来决定 Wait(等待)的最长时间
func (lim *Limiter) Wait(ctx context.Context) (err error)
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error)


func main() {
	//代表每100ms往桶中放一个Token,本质上是一秒钟往桶里放10个
	limit := rate.Every(100 * time.Millisecond)
	limiter := rate.NewLimiter(limit, 100)

	// 等待获取到桶中的令牌
	err := limiter.Wait(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 设置一秒的等待超时事件
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	err = limiter.Wait(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```



Allow/AllowN

Allow 实际上就是对 `AllowN(time.Now(),1)` 进行简化的函数。

```go
// AllowN 截止到某一个时刻，目前桶中的Token是否至少为n个,满足则返回true,同时从桶中消费n个Token。如果不满足则不消费Token,返回false
func (lim *Limiter) Allow() bool
func (lim *Limiter) AllowN(now time.Time, n int) bool

// 当前时间是否存在2个Token
if limiter.AllowN(time.Now(), 2) {
		fmt.Println("event allowed")
	} else {
		fmt.Println("event not allowed")
	}

```

Reserve/ReserveN

Reserve 相当于 `ReserveN(time.Now(), 1)`

```go
r := limiter.Reserve()
if !r.OK() {
    // Not allowed to act! Did you remember to set lim.burst to be > 0 ?
    return
}
// 执行Reserve后可以调用 Delay方法,这是一个时间类型，反映了需要等待的时间
time.Sleep(r.Delay())
// 执行相关逻辑
Work()

//如果不想等待,可以调用 Cancel 方法,它可以将Token归还
r.Cancel()
```



### 学习资料

[Golang官方限流器的用法详解](https://cloud.tencent.com/developer/article/1847918)

