package Interface

// ILogic 逻辑接口
type ILogic interface {
	Init()
	Input()
	Update(input *IMouseInput)
	Render()
	Cleanup()
}
