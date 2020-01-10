package model

import "fmt"

// 其它包不能直接访问..
type person struct {
	Name string
	age  int
	sal  float64
}

// NewPerson 包含
// *person 返回类型
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	hypot(1, 2, hp)
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员给一个默认值
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
