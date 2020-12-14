package core

type Entity interface {
	Destroy()
	GetZombie() bool
	GetContext() *Context
	GetUID() uint64
}

type _EntityBase struct {
	ctx    *Context
	uid    uint64
	entity Entity
	zombie bool
}

func (entBase *_EntityBase) initEntity(ctx *Context, entity Entity) {
	entBase.ctx = ctx
	entBase.uid = ctx.MakeUID()
	entBase.entity = entity
	entBase.zombie = false
	ctx.addEntity(entity)
}

func (entBase *_EntityBase) shutEntity() {
	entBase.ctx.removeEntity(entBase.uid)
	entBase.zombie = true
}

func (entBase *_EntityBase) GetZombie() bool {
	return entBase.zombie
}

func (entBase *_EntityBase) GetContext() *Context {
	return entBase.ctx
}

func (entBase *_EntityBase) GetUID() uint64 {
	return entBase.uid
}
