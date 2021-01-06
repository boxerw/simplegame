package core

import (
	"context"
	"simplegame/core/foundation"
	"sync"
	"time"
)

type Execute = foundation.Execute

func NewExecute(secondFrames int, sync bool, env Environment, objects ...Object) Execute {
	return &_Execute{
		secondFrames: secondFrames,
		sync:         sync,
		env:          env,
		objects:      objects,
	}
}

type _Execute struct {
	secondFrames int
	sync         bool
	env          Environment
	objects      []Object
	mutex        sync.Mutex
	running      bool
	cancel       context.CancelFunc
}

func (exec *_Execute) Start() *sync.WaitGroup {
	exec.mutex.Lock()
	defer exec.mutex.Unlock()

	if exec.running {
		panic("repeat start")
	}
	exec.running = true

	var ctx context.Context
	ctx, exec.cancel = context.WithCancel(context.Background())

	var wg sync.WaitGroup

	frameUpdateFunc := func(objectList ...Object) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second / time.Duration(exec.secondFrames))
		for {
			select {
			case <-ticker.C:
				for _, object := range objectList {
					object.Update()
				}
			case <-ctx.Done():
				return
			}
		}
	}

	if exec.sync {
		wg.Add(1)
		go frameUpdateFunc(exec.objects...)
	} else {
		for _, object := range exec.objects {
			wg.Add(1)
			go frameUpdateFunc(object)
		}
	}

	return &wg
}

func (exec *_Execute) Run() {
	exec.Start().Wait()
}

func (exec *_Execute) Shut() {
	exec.mutex.Lock()
	defer exec.mutex.Unlock()

	exec.cancel()
	exec.running = false
}

func (exec *_Execute) RangeObjects(fun func(object Object) bool) {
	if fun == nil {
		return
	}

	for _, object := range exec.objects {
		if !fun(object) {
			return
		}
	}
}

func NewQuickExecute(totalFrames int, sync bool, env Environment, objects ...Object) Execute {
	return &_QuickExecute{
		totalFrames: totalFrames,
		sync:        sync,
		env:         env,
		objects:     objects,
	}
}

type _QuickExecute struct {
	totalFrames int
	sync        bool
	env         Environment
	objects     []Object
	mutex       sync.Mutex
	running     bool
	cancel      context.CancelFunc
}

func (exec *_QuickExecute) Start() *sync.WaitGroup {
	exec.mutex.Lock()
	defer exec.mutex.Unlock()

	if exec.running {
		panic("repeat start")
	}
	exec.running = true

	var ctx context.Context
	ctx, exec.cancel = context.WithCancel(context.Background())

	var wg sync.WaitGroup

	frameUpdateFunc := func(objectList ...Object) {
		defer wg.Done()

		for i := 0; i < exec.totalFrames; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				for _, object := range objectList {
					object.Update()
				}
			}
		}
	}

	if exec.sync {
		wg.Add(1)
		go frameUpdateFunc(exec.objects...)
	} else {
		for _, object := range exec.objects {
			wg.Add(1)
			go frameUpdateFunc(object)
		}
	}

	return &wg
}

func (exec *_QuickExecute) Run() {
	exec.Start().Wait()
}

func (exec *_QuickExecute) Shut() {
	exec.mutex.Lock()
	defer exec.mutex.Unlock()

	exec.cancel()
	exec.running = false
}

func (exec *_QuickExecute) RangeObjects(fun func(object Object) bool) {
	if fun == nil {
		return
	}

	for _, object := range exec.objects {
		if !fun(object) {
			return
		}
	}
}
