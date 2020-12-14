package core

type Logic interface {
	Name() string
	Init(object Object)
	Shut()
	Update()
}
