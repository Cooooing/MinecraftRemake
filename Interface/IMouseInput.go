package Interface

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
)

type IMouseInput interface {
	PreviousPos() mgl64.Vec2
	SetPreviousPos(previousPos mgl64.Vec2)
	CurrentPos() mgl64.Vec2
	SetCurrentPos(currentPos mgl64.Vec2)
	DisplVec() *mgl32.Vec2
	SetDisplVec(displVec *mgl32.Vec2)
	InWindow() bool
	SetInWindow(inWindow bool)
	LeftButtonPress() bool
	SetLeftButtonPress(leftButtonPress bool)
	RightButtonPress() bool
	SetRightButtonPress(rightButtonPress bool)
}
