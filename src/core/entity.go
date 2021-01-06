package core

import "simple/core/foundation"

type Entity = foundation.Entity

type EntityModule struct {
	env       Environment
	uid       uint64
	inheritor Entity
	destroyed bool
}

func (entModule *EntityModule) Init(env Environment, inheritor Entity) {
	if env == nil || inheritor == nil {
		panic("nil env or inheritor")
	}

	entModule.env = env
	entModule.uid = env.MakeUID()
	entModule.inheritor = inheritor
	entModule.destroyed = false
	entModule.env.(foundation.EnvironmentWhole).AddEntity(inheritor)
}

func (entModule *EntityModule) Shut() {
	if entModule.GetDestroyed() {
		return
	}

	entModule.env.(foundation.EnvironmentWhole).RemoveEntity(entModule.uid)
	entModule.destroyed = true
}

func (entModule *EntityModule) Destroy() {
	entModule.Shut()
}

func (entModule *EntityModule) GetDestroyed() bool {
	return entModule.destroyed
}

func (entModule *EntityModule) GetEnvironment() Environment {
	return entModule.env
}

func (entModule *EntityModule) GetUID() uint64 {
	return entModule.uid
}

func (entModule *EntityModule) GetInheritor() Entity {
	return entModule.inheritor
}
