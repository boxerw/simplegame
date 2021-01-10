package shell

import "simplegame/core"

type ScreenEvent interface {
	OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool
	OnBeginDrawing(screen Screen, drawCache DrawCache) bool
	OnEndDrawing(screen Screen) bool
}

type ScreenEventHook struct {
	core.HookModule
}

func (hook *ScreenEventHook) OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool {
	return true
}

func (hook *ScreenEventHook) OnBeginDrawing(screen Screen, drawCache DrawCache) bool {
	return true
}

func (hook *ScreenEventHook) OnEndDrawing(screen Screen) bool {
	return true
}
