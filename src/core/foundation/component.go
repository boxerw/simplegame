package foundation

type Component interface {
	Init(object Object, name string)
	Shut()
	Update()
}

type ComponentBundle struct {
	Name      string
	Component Component
}
