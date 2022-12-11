/*
	责任链模式

运行将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或者传递给链上的下个处理者。
责任链会将特定行为转换为"处理者"的独立对象。
优势：
可以控制请求处理的顺序。
单一职责原则。可以对发起操作和执行操作结耦。
开闭原则。可以在不更改现有代码对情况下新增处理者。
*/
package main

import "fmt"

// Department 处理者接口（某部门
type Department interface {
	execute(*Requester)
	setNext(Department) //下一个请求部门
}

// Requester 请求者
type Requester struct {
	name               string
	DepartmentOneDone  bool
	DepartmentTwoDone  bool
	DepartmentLastDone bool
}

// DepartmentOne 具体处理部门一
type DepartmentOne struct {
	next Department
}

// 部门一,实现接口
func (d *DepartmentOne) execute(r *Requester) {
	if r.DepartmentOneDone {
		fmt.Println("部门一已经处理")
		d.next.execute(r)
		return
	}
	fmt.Println("部门一已接待")
	r.DepartmentOneDone = true
	d.next.execute(r)
}

func (d *DepartmentOne) setNext(next Department) {
	d.next = next
}

// DepartmentTwo 具体处理部门二
type DepartmentTwo struct {
	next Department
}

// 部门二,实现接口
func (d *DepartmentTwo) execute(r *Requester) {
	if r.DepartmentTwoDone {
		fmt.Println("部门二已经处理")
		d.next.execute(r)
		return
	}
	fmt.Println("部门二已接待")
	r.DepartmentTwoDone = true
	d.next.execute(r)
}

func (d *DepartmentTwo) setNext(next Department) {
	d.next = next
}

// DepartmentLast 具体处理,最后一个部门
type DepartmentLast struct {
	next Department
}

// 部门二,实现接口
func (d *DepartmentLast) execute(r *Requester) {

	if r.DepartmentLastDone {
		fmt.Println("最后一个部门已经处理")
	}
	fmt.Println("最后一个部门已接待")

}

func (d *DepartmentLast) setNext(next Department) {
	d.next = next
}

func main() {
	last := &DepartmentLast{}

	//将最后一个放在two 后面
	two := &DepartmentTwo{}
	two.setNext(last)

	//将最后一个放在two 后面
	one := &DepartmentOne{}
	one.setNext(two)

	r := &Requester{name: "abc"}
	one.execute(r)
}
