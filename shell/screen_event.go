package shell

import "simplegame/core"

type ScreenEvent interface {
	OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool
}

type ScreenEventHook struct {
	core.HookModule
}

func (hook *ScreenEventHook) OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool {
	return true
}
