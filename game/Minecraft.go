package game

import (
	"MinecraftRemake/Interface"
	"fmt"
)

// Game 游戏 继承 ILogic
type Game struct {
	window Interface.IManager
}

func NewGame() *Game {
	return &Game{}
}

func (g Game) Init() {
	//TODO implement me
	fmt.Println("game init ...")
}

func (g Game) Input() {
	//TODO implement me
}

func (g Game) Update(mouseInput Interface.IMouseInput) {
	//TODO implement me
	fmt.Printf("x: %f y: %f\n", mouseInput.DisplVec().X(), mouseInput.DisplVec().Y())
}

func (g Game) Render() {
	//TODO implement me
	fmt.Println("game Render ...")
}

func (g Game) Cleanup() {
	//TODO implement me
	fmt.Println("game Cleanup ...")
}
