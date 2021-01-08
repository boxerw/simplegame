package foundation

type Object interface {
	Entity
	Update(frameCtx FrameContext)
	GetDisabled() bool
	AddComponent(cb ComponentBundle) bool
	RemoveComponent(name string)
	GetComponent(name string) Component
	RangeComponents(fun func(cb ComponentBundle) bool)
}

type ObjectWhole interface {
	Object
	ExecFunc(fun func())
}
