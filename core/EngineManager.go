package core

import (
	"MinecraftRemake/core/utils"
	"MinecraftRemake/game"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"time"
)

var frameTime = 1.0 / utils.FRAMERATE

type EngineManager struct {
	fps       int
	isRunning bool

	window     *WindowManager
	game       *game.Game
	mouseInput *MouseInput
}

func (engine *EngineManager) Fps() int {
	return engine.fps
}

func (engine *EngineManager) SetFps(fps int) {
	engine.fps = fps
}

func NewEngineManager(window *WindowManager, game *game.Game) EngineManager {
	return EngineManager{window: window, game: game, mouseInput: NewMouseInput()}
}

func (engine *EngineManager) Start() {
	engine.Init()
	if engine.isRunning {
		return
	}
	engine.run()
}

func (engine *EngineManager) Init() {
	engine.window.Init()
	engine.game.Init()
	engine.mouseInput.Init()
}

func (engine *EngineManager) run() {
	engine.isRunning = true
	frames := 0
	var frameCounter int64 = 0
	lastTime := time.Now().UnixNano()
	unprocessedTime := 0.0

	for engine.isRunning {
		isRender := false
		startTime := time.Now().UnixNano()
		passedTime := startTime - lastTime
		lastTime = startTime

		unprocessedTime += float64(passedTime) / utils.NANOSECOND
		frameCounter += passedTime

		engine.input()

		for unprocessedTime > frameTime {
			isRender = true
			unprocessedTime = 0
			if engine.window.window.ShouldClose() {
				engine.stop()
			}
			if frameCounter >= utils.NANOSECOND {
				engine.SetFps(frames)
				engine.window.SetTitle(fmt.Sprintf("%s FPS: %d", utils.TITLE, frames))
				frames = 0
				frameCounter = 0
			}
		}

		if isRender {
			engine.update()
			engine.render()
			frames++
		}
	}
	engine.cleanup()
}

func (engine *EngineManager) input() {
	engine.mouseInput.Input()
	engine.game.Input()
}

func (engine *EngineManager) render() {
	engine.window.Update()

}

func (engine *EngineManager) update() {
	engine.game.Update(engine.mouseInput)
}

func (engine *EngineManager) stop() {
	if !engine.isRunning {
		return
	}
	engine.isRunning = false
}

func (engine *EngineManager) cleanup() {
	engine.window.Cleanup()
	engine.game.Cleanup()
	glfw.Terminate()
}
