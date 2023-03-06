package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

// 构建一个限流器对象
// 第一个参数是 r Limit。代表每秒可以向 Token 桶中产生多少 token。Limit 实际上是 float64 的别名。
// 第二个参数是 b int。b 代表 Token 桶的容量大小。
// 令牌桶大小为1,以每秒10个Token的速率向桶中放置Token
// limiter := NewLimiter(10,1)

func main() {
	//代表每100ms往桶中放一个Token,本质上是一秒钟往桶里放10个
	limit := rate.Every(100 * time.Millisecond)
	limiter := rate.NewLimiter(limit, 100)

	// Limiter 提供了三类方法消费Token,每次消费一个Token,也可以一次性消费多个Token

	//// 等待获取到桶中的令牌
	//err := limiter.Wait(context.Background())
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	//
	//// 设置一秒的等待超时事件
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	//err = limiter.Wait(ctx)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}

	if limiter.AllowN(time.Now(), 2) {
		fmt.Println("event allowed")
	} else {
		fmt.Println("event not allowed")
	}

	r := limiter.Reserve()
	if !r.OK() {
		// Not allowed to act! Did you remember to set lim.burst to be > 0 ?
		return
	}
	time.Sleep(r.Delay())
	// 执行相关逻辑
	r.Cancel()
}
