/*
	访问者模式

它能将算法与其所作用的对象隔离开来。
访问者模式建议将新行为放入一个名为"访问者"的独立类中，而不是试图将其整合到已有类中。

"双分派"技巧。与其让客户端选择调用正确版本到方法，不如将选择权委派给作为参数传递给访问者的对象。
由于该对象知晓其自身的类，因此能更自然地在访问者中选出正确的方法。
它们会"接收"一个访问者并告诉其应执行的访问者方法。
*/
package main

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type visitor interface {
	visitForSquare(*square)

	visitForrectangle(*rectangle)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForrectangle(s *rectangle) {
	//Calculate are for rectangle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
	square := &square{side: 2}

	rectangle := &rectangle{l: 2, b: 3}
	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)

	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)

	rectangle.accept(middleCoordinates)
}
