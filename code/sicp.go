package main

type Pair struct {
	Car interface{}
	Cdr interface{}
}

func Cons(a, b interface{}) *Pair {
	return &Pair{Car: a, Cdr: b}
}

func Car(p Pair) interface{} {
	return p.Car
}

func Cdr(p Pair) interface{} {
	return p.Cdr
}

func main() {

}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func MakeRat(n, d int) *Pair {

	g := GCD(n, d)
	return Cons(n/g, d/g) // 返回一个 pair：最简分数
}
