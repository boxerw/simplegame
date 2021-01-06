package foundation

type Child interface {
	GetParent() Object
}

type ChildWhole interface {
	Child
	SetParent(parent Object)
}
