package main

import "fmt"

type GameBase interface {
	OnInit()
	OnBegin()
}

type LOL struct {
}

func (l LOL) OnInit() {
	fmt.Println("LOL OnInit")
}

func (l LOL) OnBegin() {
	fmt.Println("LOL OnBegin")
}

func main() {
	var gameBase GameBase
	gameBase = new(LOL)
	gameBase.OnInit()
}
