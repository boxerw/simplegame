package core

import (
	"simple/core/foundation"
	"strings"
)

type Object = foundation.Object

type _DestroyMark uint8

const (
	_DestroyMark_Off _DestroyMark = iota
	_DestroyMark_On
	_DestroyMark_Do
)

type ObjectModule struct {
	EntityModule
	components  []ComponentBundle
	destroyMark _DestroyMark
}

func (objModule *ObjectModule) Init(env Environment, object Object, initFun func(), components ...ComponentBundle) {
	if env == nil || object == nil {
		panic("nil env or object")
	}

	objModule.EntityModule.Init(env, object)
	objModule.components = components

	if initFun != nil {
		initFun()
	}

	objModule.ExecFunc(func() bool {
		for _, cb := range objModule.components {
			if cb.Component != nil {
				cb.Component.Init(object, cb.Name)
			}
		}
		return true
	})
}

func (objModule *ObjectModule) Shut(shutFun func()) {
	if objModule.GetDisabled() {
		return
	}

	if objModule.destroyMark >= _DestroyMark_On {
		objModule.setDestroyMark(_DestroyMark_Do)
		return
	}

	objModule.setDestroyMark(_DestroyMark_Do)

	for _, cb := range objModule.components {
		if cb.Component != nil {
			cb.Component.Shut()
		}
	}

	if shutFun != nil {
		shutFun()
	}

	objModule.EntityModule.Shut()
}

func (objModule *ObjectModule) Destroy() {
	objModule.Shut(nil)
}

func (objModule *ObjectModule) Update() {
	if objModule.GetDisabled() {
		return
	}

	for _, cb := range objModule.components {
		objModule.ExecFunc(func() bool {
			if cb.Component != nil {
				cb.Component.Update()
			}
			return true
		})
	}
}

func (objModule *ObjectModule) GetDisabled() bool {
	return objModule.GetDestroyed() || objModule.destroyMark >= _DestroyMark_Do
}

func (objModule *ObjectModule) AddComponent(cb ComponentBundle) bool {
	if objModule.GetDisabled() {
		return false
	}

	for _, oldCb := range objModule.components {
		if strings.ToLower(oldCb.Name) == strings.ToLower(cb.Name) {
			return false
		}
	}

	objModule.components = append(objModule.components, cb)

	objModule.ExecFunc(func() bool {
		if cb.Component != nil {
			cb.Component.Init(objModule.GetInheritor().(Object), cb.Name)
		}
		return true
	})

	return true
}

func (objModule *ObjectModule) RemoveComponent(name string) {
	if objModule.GetDisabled() {
		return
	}

	for i, cb := range objModule.components {
		if strings.ToLower(cb.Name) == strings.ToLower(name) {
			objModule.components = append(objModule.components[:i], objModule.components[i+1:]...)
			objModule.ExecFunc(func() bool {
				if cb.Component != nil {
					cb.Component.Shut()
				}
				return true
			})
			break
		}
	}
}

func (objModule *ObjectModule) GetComponent(name string) Component {
	for _, cb := range objModule.components {
		if strings.ToLower(cb.Name) == strings.ToLower(name) {
			return cb.Component
		}
	}
	return nil
}

func (objModule *ObjectModule) RangeComponents(fun func(cb ComponentBundle) bool) {
	if fun == nil {
		return
	}

	for _, cb := range objModule.components {
		if !fun(cb) {
			return
		}
	}
}

func (objModule *ObjectModule) ExecFunc(fun func() bool) bool {
	if objModule.GetDisabled() || fun == nil {
		return false
	}

	objModule.setDestroyMark(_DestroyMark_On)

	rv := fun()

	if objModule.setDestroyMark(_DestroyMark_Off) == _DestroyMark_Do {
		objModule.Destroy()
	}

	return rv
}

func (objModule *ObjectModule) setDestroyMark(v _DestroyMark) _DestroyMark {
	old := objModule.destroyMark
	objModule.destroyMark = v
	return old
}
