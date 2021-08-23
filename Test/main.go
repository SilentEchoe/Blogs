package main

import "fmt"

type GameBase interface {
	OnInit()  // 初始化游戏
	OnBegin() // 游戏数据同步
}

type Game_LOL struct {
}

func (l *Game_LOL) OnInit() {
	fmt.Println("LOL OnInit")
}

func (l *Game_LOL) OnBegin() {
	fmt.Println("LOL OnBegin")
}

func main() {
	var gameBase GameBase
	gameBase = new(Game_LOL)
	gameBase.OnInit()
}
