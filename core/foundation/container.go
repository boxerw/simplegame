package foundation

type Container interface {
	AddChild(child Object) bool
	RemoveChild(entUID uint64)
	GetChild(entUID uint64) (Object, bool)
	RangeChildren(fun func(child Object) bool)
}
