package main

import (
	"fmt"
)

// Pair represents a simple cons cell (a . b)
type Pair struct {
	Car any
	Cdr any
}

// Cons creates a new pair
func Cons(a, b any) *Pair {
	return &Pair{Car: a, Cdr: b}
}

// Car returns the first element of a pair
func Car(p *Pair) any {
	return p.Car
}

// Cdr returns the second element of a pair
func Cdr(p *Pair) any {
	return p.Cdr
}

// ToList constructs a linked list from elements
func ToList(items ...any) *Pair {
	if len(items) == 0 {
		return nil
	}
	return Cons(items[0], ToList(items[1:]...))
}

// PrintList prints the contents of a list
func PrintList(p *Pair) {
	for p != nil {
		fmt.Print(Car(p), " ")
		cdr, ok := Cdr(p).(*Pair)
		if !ok {
			break
		}
		p = cdr
	}
	fmt.Println()
}

// ToMapList constructs a list of key-value pairs
func ToMapList(pairs ...[2]any) *Pair {
	if len(pairs) == 0 {
		return nil
	}
	kv := Cons(pairs[0][0], pairs[0][1])
	return Cons(kv, ToMapList(pairs[1:]...))
}

// PrintMapList prints a key-value pair list
func PrintMapList(p *Pair) {
	for p != nil {
		pair := Car(p).(*Pair)
		fmt.Printf("%v => %v\n", Car(pair), Cdr(pair))
		cdr, ok := Cdr(p).(*Pair)
		if !ok {
			break
		}
		p = cdr
	}
}

// Sequence abstracts list operations (abstract barrier)
type Sequence interface {
	First() any
	Rest() Sequence
	IsEmpty() bool
}

type List struct {
	Head any
	Tail *List
}

func (l *List) First() any {
	return l.Head
}

func (l *List) Rest() Sequence {
	if l.Tail == nil {
		return EmptyList{}
	}
	return l.Tail
}

func (l *List) IsEmpty() bool {
	return false
}

type EmptyList struct{}

func (EmptyList) First() any     { return nil }
func (EmptyList) Rest() Sequence { return EmptyList{} }
func (EmptyList) IsEmpty() bool  { return true }

// NewList builds a Sequence
func NewList(items ...any) Sequence {
	if len(items) == 0 {
		return EmptyList{}
	}
	next := NewList(items[1:]...)
	if next == nil {
		next = &List{}
	}
	return &List{Head: items[0], Tail: next.(*List)}
}

// PrintSequence prints a Sequence
func PrintSequence(seq Sequence) {
	for !seq.IsEmpty() {
		fmt.Print(seq.First(), " ")
		seq = seq.Rest()
	}
	fmt.Println()
}
