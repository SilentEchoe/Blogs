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

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)

}
