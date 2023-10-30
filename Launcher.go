package main

import (
	"MinecraftRemake/core"
	"MinecraftRemake/game"
	_ "embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	_ "image/png"
	"log"
)

func main() {

	log.Println(glfw.GetVersion())

	engine := core.NewEngineManager(core.NewWindowManager("", 960, 640, false), game.NewGame())
	engine.Start()

}
