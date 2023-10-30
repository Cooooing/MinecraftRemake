package core

import (
	"MinecraftRemake/core/utils"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

type WindowManager struct {
	window *glfw.Window
	vSync  bool // 垂直同步
	resize bool // 改变窗口大小
	Title  string
	Width  int
	Height int
}

func NewWindowManager(title string, width int, height int, vSync bool) *WindowManager {
	return &WindowManager{vSync: vSync, Title: title, Width: width, Height: height}
}

func (w *WindowManager) SetTitle(Title string) {
	w.Title = Title
	w.window.SetTitle(Title)
}

func (w *WindowManager) Init() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	//defer glfw.Terminate()

	glfw.DefaultWindowHints()
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	maximised := false // 最大化
	if w.Width == 0 || w.Height == 0 {
		w.Width = utils.WIDTH
		w.Height = utils.HEIGHT
		glfw.WindowHint(glfw.Maximized, glfw.True)
		maximised = true
	}
	if w.Title == "" {
		w.Title = utils.TITLE
	}

	window, err := glfw.CreateWindow(w.Width, w.Height, w.Title, nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}
	w.window = window

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape && action == glfw.Release {
			w.SetShouldClose(true)
		}
	})

	if maximised {
		window.Maximize()
	} else {
		// 获取显示器信息
		videoMode := glfw.GetPrimaryMonitor().GetVideoMode()
		// 设置窗口位置 在显示器上居中
		width, height := window.GetSize()
		window.SetPos((videoMode.Width-width)/2, (videoMode.Height-height)/2)
	}

	window.MakeContextCurrent()
	window.Show()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.ClearColor(0, 0, 0, 0)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.STENCIL_TEST)
	// 面剔除
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
}

func (w *WindowManager) Update() {
	// 交换缓冲区以更新显示的图像
	w.window.SwapBuffers()
	// 处理窗口中的事件
	/*
		事件会被保存在队列中，等待被处理。调用此函数，将队列中的事件逐个处理（回调函数）
	*/
	glfw.PollEvents()
}

func (w *WindowManager) Cleanup() {
	w.window.Destroy()
}
