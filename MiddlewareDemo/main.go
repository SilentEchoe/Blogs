package main

import "fmt"

/*
	中间件原理Demo
*/

type Context struct {
	Handlers []func(*Context)
	index    int8
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.Handlers)) {
		c.Handlers[c.index](c)
		c.index++
	}

}

func main() {
	c := &Context{}
	c.Handlers = make([]func(ctx *Context), 0)

	// 注册中间件
	c.Handlers = append(c.Handlers, m1)
	c.Handlers = append(c.Handlers, m2)
	c.Handlers = append(c.Handlers, m3)

	c.Handlers = append(c.Handlers, action)

	c.Handlers[0](c)
	c.Next()
}

func action(c *Context) {
	fmt.Println("main handler")
}

func m1(c *Context) {
	fmt.Println("m1 start")
}

func m2(c *Context) {
	fmt.Println("m2 start")
	c.Next()
	fmt.Println("m2 end")
}

func m3(c *Context) {
	c.Next()
	fmt.Println("m3 end")
}
