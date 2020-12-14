package core

import (
	"sync"
)

func NewContext() *Context {
	return &Context{
		entityMap: map[uint64]Entity{},
		valueMap:  map[string]interface{}{},
	}
}

type Context struct {
	mutex     sync.Mutex
	uidMaker  uint64
	entityMap map[uint64]Entity
	valueMap  map[string]interface{}
}

func (ctx *Context) MakeUID() uint64 {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	for {
		ctx.uidMaker++

		if ctx.uidMaker == 0 {
			continue
		}

		_, ok := ctx.entityMap[ctx.uidMaker]
		if !ok {
			break
		}
	}

	return ctx.uidMaker
}

func (ctx *Context) addEntity(entity Entity) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	if _, ok := ctx.entityMap[entity.GetUID()]; ok {
		panic("repeat add entity")
	}

	ctx.entityMap[entity.GetUID()] = entity
}

func (ctx *Context) removeEntity(entUID uint64) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	delete(ctx.entityMap, entUID)
}

func (ctx *Context) GetEntity(uid uint64) (Entity, bool) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	entity, ok := ctx.entityMap[uid]
	return entity, ok
}

func (ctx *Context) SetValue(name string, value interface{}) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	ctx.valueMap[name] = value
}

func (ctx *Context) GetValue(name string) interface{} {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	return ctx.valueMap[name]
}
