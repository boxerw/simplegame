package foundation

type Hook interface {
	GetHookID() uint64
	GetPriority() int
}

type Event interface {
	AddHook(hook Hook)
	RemoveHook(hookID uint64)
	RangeHooks(fun func(hook Hook) bool)
}
