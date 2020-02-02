package main
import "fmt"
type innerS struct {
    in1 int
    in2 int
}
type outerS struct {
    b int
    c float32
    int // anonymous field
    innerS //anonymous field
}

type A struct{

	ax, ay int
}

type B struct{
	A
	bx, by float32
}


func main() {
	b := B{A{1,2},3.0,4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}