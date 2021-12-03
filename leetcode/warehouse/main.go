package main

/* 	栈算法题
栈的特点是后入先出
*/

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead
//分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
	value []int
}

func Constructor() CQueue {
	return CQueue{
		value: make([]int, 0),
	}

}

func main() {

}

// 队列尾部插入整数
func (this *CQueue) AppendTail(value int) {

}
