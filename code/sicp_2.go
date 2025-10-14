package main

import "fmt"

type Pair struct {
	Car interface{}
	Cdr interface{}
}

func Cons(a, b interface{}) *Pair {
	return &Pair{Car: a, Cdr: b}
}

func Car(p *Pair) interface{} {
	return p.Car
}

func Cdr(p *Pair) interface{} {
	return p.Cdr
}

func PrintList(list *Pair) {
	for list != nil {
		fmt.Println(Car(list))
		next := Cdr(list)
		if next == nil {
			break
		}
		list = next.(*Pair)
	}
}

func main() {
	l := Cons(1, Cons(2, Cons(3, nil)))
	PrintList(l)
}
