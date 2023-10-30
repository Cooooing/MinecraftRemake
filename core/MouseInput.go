package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
)

type MouseInput struct {
	previousPos mgl64.Vec2
	currentPos  mgl64.Vec2
	displVec    *mgl32.Vec2

	inWindow         bool
	leftButtonPress  bool
	rightButtonPress bool
}

func NewMouseInput() *MouseInput {
	return &MouseInput{}
}

func (m *MouseInput) DisplVec() *mgl32.Vec2 {
	return m.displVec
}

func (m *MouseInput) SetDisplVec(displVec *mgl32.Vec2) {
	m.displVec = displVec
}

func (m *MouseInput) PreviousPos() mgl64.Vec2 {
	return m.previousPos
}

func (m *MouseInput) SetPreviousPos(previousPos mgl64.Vec2) {
	m.previousPos = previousPos
}

func (m *MouseInput) CurrentPos() mgl64.Vec2 {
	return m.currentPos
}

func (m *MouseInput) SetCurrentPos(currentPos mgl64.Vec2) {
	m.currentPos = currentPos
}

func (m *MouseInput) InWindow() bool {
	return m.inWindow
}

func (m *MouseInput) SetInWindow(inWindow bool) {
	m.inWindow = inWindow
}

func (m *MouseInput) LeftButtonPress() bool {
	return m.leftButtonPress
}

func (m *MouseInput) SetLeftButtonPress(leftButtonPress bool) {
	m.leftButtonPress = leftButtonPress
}

func (m *MouseInput) RightButtonPress() bool {
	return m.rightButtonPress
}

func (m *MouseInput) SetRightButtonPress(rightButtonPress bool) {
	m.rightButtonPress = rightButtonPress
}

func (m *MouseInput) Init() {
	Window.window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		m.currentPos = mgl64.Vec2{x, y}
	})
	Window.window.SetCursorEnterCallback(func(w *glfw.Window, entered bool) {
		m.inWindow = entered
	})
	Window.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		m.leftButtonPress = button == glfw.MouseButtonLeft && action == glfw.Press
		m.rightButtonPress = button == glfw.MouseButtonRight && action == glfw.Press
	})
}

func (m *MouseInput) Input() {
	m.SetDisplVec(&mgl32.Vec2{0, 0})
	if m.previousPos.X() > 0 && m.currentPos.Y() > 0 && m.inWindow {
		x := m.currentPos.X() - m.previousPos.X()
		y := m.currentPos.Y() - m.previousPos.Y()
		rotateX := x != 0
		rotateY := y != 0
		if rotateX {
			m.displVec[1] = float32(x)
		}
		if rotateY {
			m.displVec[0] = float32(y)
		}
	}
	m.previousPos = m.currentPos
}
