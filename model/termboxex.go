package model

import (
	"github.com/nsf/termbox-go"
	"sync"
)

var termboxEx = _TermboxEx{
	inputHookMap: map[uintptr]chan termbox.Event{},
}

type _TermboxEx struct {
	mutex        sync.Mutex
	inputHookMap map[uintptr]chan termbox.Event
}

func (ex *_TermboxEx) Init() error {
	ex.mutex.Lock()
	defer ex.mutex.Unlock()

	if termbox.IsInit {
		return nil
	}

	if err := termbox.Init(); err != nil {
		return err
	}

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				fallthrough
			case termbox.EventResize:
				fallthrough
			case termbox.EventMouse:
				func() {
					ex.mutex.Lock()
					defer ex.mutex.Unlock()
					for _, inputChan := range ex.inputHookMap {
						select {
						case inputChan <- ev:
						default:
						}
					}
				}()
			case termbox.EventInterrupt:
				return
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}()

	return nil
}

func (ex *_TermboxEx) Shut() {
	ex.mutex.Lock()
	defer ex.mutex.Unlock()

	if !termbox.IsInit {
		return
	}

	termbox.Close()
}

func (ex *_TermboxEx) IsInit() bool {
	ex.mutex.Lock()
	defer ex.mutex.Unlock()

	return termbox.IsInit
}

func (ex *_TermboxEx) AddInputEventHook(id uintptr, inputChan chan termbox.Event) {
	ex.mutex.Lock()
	defer ex.mutex.Unlock()

	ex.inputHookMap[id] = inputChan
}

func (ex *_TermboxEx) RemoveInputEventHook(id uintptr) {
	ex.mutex.Lock()
	defer ex.mutex.Unlock()

	delete(ex.inputHookMap, id)
}
