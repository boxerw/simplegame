package core

import (
	"github.com/boxerw/simplegame/core/foundation"
	"sort"
	"unsafe"
)

type Hook = foundation.Hook

type HookModule struct{}

func (hookModule *HookModule) GetHookID() uint64 {
	return uint64(uintptr(unsafe.Pointer(hookModule)))
}

func (hookModule *HookModule) GetPriority() int {
	return 0
}

type Event = foundation.Event

type EventModule struct {
	hooks []Hook
}

func (eventModule *EventModule) AddHook(hook Hook) {
	if hook == nil {
		return
	}

	for _, t := range eventModule.hooks {
		if t.GetHookID() == hook.GetHookID() {
			return
		}
	}

	eventModule.hooks = append(eventModule.hooks, hook)

	sort.SliceStable(eventModule.hooks, func(i, j int) bool {
		return eventModule.hooks[i].GetPriority() < eventModule.hooks[j].GetPriority()
	})
}

func (eventModule *EventModule) RemoveHook(hookID uint64) {
	for i, hook := range eventModule.hooks {
		if hook.GetHookID() == hookID {
			eventModule.hooks = append(eventModule.hooks[:i], eventModule.hooks[i+1:]...)
			return
		}
	}
}

func (eventModule *EventModule) RangeHooks(fun func(hook Hook) bool) {
	if fun == nil {
		return
	}

	for _, hook := range eventModule.hooks {
		if !fun(hook) {
			return
		}
	}
}
