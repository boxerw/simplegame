package client

import (
	"github.com/boxerw/simplegame/core"
	"github.com/nsf/termbox-go"
)

type ControllerEvent interface {
	OnControllerKeyPress(controller Controller, key termbox.Key, ch rune) bool
	OnControllerMousePress(controller Controller, key termbox.Key, posi Posi2D) bool
}

type ControllerEventHook struct {
	core.HookModule
}

func (hook *ControllerEventHook) OnControllerKeyPress(controller Controller, key termbox.Key, ch rune) bool {
	return true
}

func (hook *ControllerEventHook) OnControllerMousePress(controller Controller, key termbox.Key, posi Posi2D) bool {
	return true
}
