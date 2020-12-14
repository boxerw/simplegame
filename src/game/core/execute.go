package core

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func NewExecute(secondFrames int64, sync bool, ctx *Context, objectList ...Object) *Execute {
	return &Execute{
		ctx:          ctx,
		secondFrames: secondFrames,
		sync:         sync,
		objectList:   objectList,
	}
}

type Execute struct {
	ctx          *Context
	secondFrames int64
	sync         bool
	objectList   []Object
}

func (exec *Execute) Run() {
	var wg sync.WaitGroup

	frameUpdateFunc := func(objectList ...Object) {
		defer func() {
			for _, object := range objectList {
				object.Destroy()
			}
			wg.Done()
		}()

		endC := make(chan os.Signal, 1)
		signal.Notify(endC, os.Interrupt, os.Kill, syscall.SIGTERM)

		ticker := time.NewTicker(time.Second / time.Duration(exec.secondFrames))
		for {
			select {
			case <-ticker.C:
				for _, object := range objectList {
					object.Update()
				}
			case <-endC:
				return
			}
		}
	}

	if exec.sync {
		wg.Add(1)
		go frameUpdateFunc(exec.objectList...)
	} else {
		for _, object := range exec.objectList {
			wg.Add(1)
			go frameUpdateFunc(object)
		}
	}

	wg.Wait()
}
