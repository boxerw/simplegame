package foundation

type Component interface {
	Init(object Object, name string)
	Shut()
	Update(frameCtx FrameContext)
}

type ComponentBundle struct {
	Name      string
	Component Component
}
