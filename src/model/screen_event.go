package model

import "unsafe"

type ScreenEvent interface {
	OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool
}

type ScreenEventHook struct {
}

func (hook *ScreenEventHook) GetHookID() uintptr {
	return uintptr(unsafe.Pointer(hook))
}

func (hook *ScreenEventHook) GetPriority() int {
	return 0
}

func (hook *ScreenEventHook) OnScreenCanvasSizeChange(screen Screen, oldSize, newSize Posi2D) bool {
	return true
}
