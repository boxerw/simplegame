package model

import (
	"github.com/nsf/termbox-go"
	"unsafe"
)

type ControllerEvent interface {
	OnControllerKeyPress(controller Controller, key termbox.Key, ch rune) bool
	OnControllerMousePress(controller Controller, key termbox.Key, posi Posi2D) bool
}

type ControllerEventHook struct {
}

func (hook *ControllerEventHook) GetHookID() uintptr {
	return uintptr(unsafe.Pointer(hook))
}

func (hook *ControllerEventHook) GetPriority() int {
	return 0
}

func (hook *ControllerEventHook) OnControllerKeyPress(controller Controller, key termbox.Key, ch rune) bool {
	return true
}

func (hook *ControllerEventHook) OnControllerMousePress(controller Controller, key termbox.Key, posi Posi2D) bool {
	return true
}
