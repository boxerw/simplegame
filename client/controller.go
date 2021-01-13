package client

import (
	"github.com/boxerw/simplegame/core"
	"github.com/nsf/termbox-go"
	"unsafe"
)

type Controller interface {
	core.Object
	core.Event
}

func NewController(env core.Environment, components ...core.ComponentBundle) Controller {
	controller := &_Controller{}
	controller.ObjectModule.Init(env, controller, func() {
		if !termboxEx.IsInit() {
			panic("termbox not init, please new screen first")
		}

		controller.inputChan = make(chan termbox.Event, 100)
		termboxEx.AddInputEventHook(unsafe.Pointer(controller), controller.inputChan)

	}, components...)
	return controller
}

type _Controller struct {
	core.ObjectModule
	core.EventModule
	inputChan chan termbox.Event
}

func (controller *_Controller) Destroy() {
	controller.ObjectModule.Shut(func() {
		termboxEx.RemoveInputEventHook(unsafe.Pointer(controller))
		close(controller.inputChan)
	})
}

func (controller *_Controller) Update(frameCtx core.FrameContext) {
	if controller.GetDisabled() {
		return
	}

	controller.ObjectModule.Update(frameCtx)

	func() {
		for {
			select {
			case ev := <-controller.inputChan:
				switch ev.Type {
				case termbox.EventKey:
					controller.RangeHooks(func(hook core.Hook) bool {
						return controller.ExecFunc(func() bool {
							if hook, ok := hook.(ControllerEvent); ok {
								return hook.OnControllerKeyPress(controller, ev.Key, ev.Ch)
							}
							return true
						})
					})
				case termbox.EventMouse:
					controller.RangeHooks(func(hook core.Hook) bool {
						return controller.ExecFunc(func() bool {
							if hook, ok := hook.(ControllerEvent); ok {
								return hook.OnControllerMousePress(controller, ev.Key, Posi2D{ev.MouseX, ev.MouseY})
							}
							return true
						})
					})
				}
			default:
				return
			}
		}
	}()
}
