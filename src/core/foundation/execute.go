package foundation

import "sync"

type Execute interface {
	Start() *sync.WaitGroup
	Run()
	Shut()
	RangeObjects(fun func(object Object) bool)
}
