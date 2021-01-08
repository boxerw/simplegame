package core

import (
	"context"
	"simplegame/core/foundation"
	"sync"
	"time"
)

type Execute = foundation.Execute

func NewExecute(fps int32, sync bool, objects ...Object) Execute {
	return &_Execute{
		fps:     fps,
		sync:    sync,
		objects: objects,
	}
}

type _Execute struct {
	fps     int32
	sync    bool
	objects []Object
	mutex   sync.Mutex
	running bool
	cancel  context.CancelFunc
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

		frameCtx := &_FrameContext{
			fixedFPS: exec.fps,
		}
		statBeginTime := time.Now()
		statFrames := 0

		ticker := time.NewTicker(time.Second / time.Duration(exec.fps))

		for frameCtx.framesCount = 0; ; frameCtx.framesCount++ {
			now := time.Now()

			statInterval := now.Sub(statBeginTime).Seconds()
			if statInterval >= 1 {
				frameCtx.statFPS = float32(float64(statFrames) / statInterval)
				statBeginTime = now
				statFrames = 0
			}

			frameCtx.frameBeginTime = now

			select {
			case <-ticker.C:
				for _, object := range objectList {
					object.Update(frameCtx)
				}
			case <-ctx.Done():
				return
			}

			frameCtx.lastFrameElapseTime = time.Now().Sub(frameCtx.frameBeginTime)
			statFrames++
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

func NewQuickExecute(totalFrames int32, sync bool, objects ...Object) Execute {
	return &_QuickExecute{
		totalFrames: totalFrames,
		sync:        sync,
		objects:     objects,
	}
}

type _QuickExecute struct {
	totalFrames int32
	sync        bool
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

		frameCtx := &_FrameContext{
			totalFrames: exec.totalFrames,
		}

		for i := int32(0); i < exec.totalFrames; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				frameCtx.frameBeginTime = time.Now()

				for _, object := range objectList {
					object.Update(frameCtx)
				}

				frameCtx.lastFrameElapseTime = time.Now().Sub(frameCtx.frameBeginTime)
				frameCtx.framesCount++
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

type FrameContext = foundation.FrameContext

type _FrameContext struct {
	DataModule
	fixedFPS, totalFrames int32
	statFPS               float32
	framesCount           uint64
	frameBeginTime        time.Time
	lastFrameElapseTime   time.Duration
}

func (frameCtx *_FrameContext) GetFixedFPS() int32 {
	return frameCtx.fixedFPS
}

func (frameCtx *_FrameContext) GetFPS() float32 {
	return frameCtx.statFPS
}

func (frameCtx *_FrameContext) GetTotalFrames() int32 {
	return frameCtx.totalFrames
}

func (frameCtx *_FrameContext) GetFramesCount() uint64 {
	return frameCtx.framesCount
}

func (frameCtx *_FrameContext) GetCurFrameBeginTime() time.Time {
	return frameCtx.frameBeginTime
}

func (frameCtx *_FrameContext) GetLastFrameElapseTime() time.Duration {
	return frameCtx.lastFrameElapseTime
}
