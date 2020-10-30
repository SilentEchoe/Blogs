package main

type MinStack struct {
	task     []int
	miniTask []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		task:     nil,
		miniTask: nil,
	}
}

func (this *MinStack) Push(x int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {
	return 0
}

func (this *MinStack) GetMin() int {
	return 0
}
