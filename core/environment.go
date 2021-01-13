package core

import (
	"github.com/boxerw/simplegame/core/foundation"
	"sync"
	"sync/atomic"
)

type Environment = foundation.Environment

func NewEnvironment(startUID uint64) Environment {
	return &_Environment{
		uidMaker: startUID,
	}
}

type _Environment struct {
	uidMaker  uint64
	entityMap sync.Map
	valueMap  sync.Map
}

func (env *_Environment) MakeUID() uint64 {
	return atomic.AddUint64(&env.uidMaker, 1)
}

func (env *_Environment) GetEntity(uid uint64) (Entity, bool) {
	v, ok := env.entityMap.Load(uid)
	if !ok {
		return nil, false
	}
	return v.(Entity), true
}

func (env *_Environment) SetValue(name string, value interface{}) {
	env.valueMap.Store(name, value)
}

func (env *_Environment) GetValue(name string) interface{} {
	v, ok := env.valueMap.Load(name)
	if !ok {
		return nil
	}
	return v
}

func (env *_Environment) AddEntity(entity Entity) {
	if _, loaded := env.entityMap.LoadOrStore(entity.GetUID(), entity); loaded {
		panic("repeat add inheritor")
	}
}

func (env *_Environment) RemoveEntity(entUID uint64) {
	env.entityMap.Delete(entUID)
}
