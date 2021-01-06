package foundation

type Hook interface {
	GetHookID() uintptr
	GetPriority() int
}

type Event interface {
	AddHook(hook Hook)
	RemoveHook(hookID uintptr)
	RangeHooks(fun func(hook Hook) bool)
}
