package foundation

import (
	"sync"
	"time"
)

type Execute interface {
	Start() *sync.WaitGroup
	Run()
	Shut()
	RangeObjects(fun func(object Object) bool)
}

type FrameContext interface {
	Data
	GetFixedFPS() int32
	GetFPS() float32
	GetTotalFrames() int32
	GetFrames() uint64
	GetCurFrameBeginTime() time.Time
	GetLastFrameElapseTime() time.Duration
}
