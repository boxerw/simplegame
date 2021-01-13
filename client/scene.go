package client

import (
	"github.com/boxerw/simplegame/core"
	"github.com/boxerw/simplegame/core/foundation"
)

type Scene interface {
	core.Object
	core.Container
}

func NewScene(env core.Environment, components ...core.ComponentBundle) Scene {
	scene := &_Scene{}
	scene.Init(env, scene, func() {
		scene.childMap = map[uint64]core.Object{}
	}, components...)
	return scene
}

type _Scene struct {
	core.ObjectModule
	childList []core.Object
	childMap  map[uint64]core.Object
}

func (scene *_Scene) Destroy() {
	scene.Shut(func() {
		for _, child := range scene.childList {
			child.Destroy()
		}
	})
}

func (scene *_Scene) Update(frameCtx core.FrameContext) {
	if scene.GetDisabled() {
		return
	}

	scene.ObjectModule.Update(frameCtx)

	for _, child := range scene.childList {
		child.Update(frameCtx)
	}
}

func (scene *_Scene) AddChild(object core.Object) bool {
	if scene.GetEnvironment() != object.GetEnvironment() {
		panic("env invalid")
	}

	if scene.GetDestroyed() || object.GetDisabled() {
		return false
	}

	child, ok := object.(core.Child)
	if !ok {
		return false
	}

	if child.GetParent() != nil {
		return false
	}

	if _, ok := scene.childMap[object.GetUID()]; ok {
		return false
	}

	scene.childMap[object.GetUID()] = object
	scene.childList = append(scene.childList, object)
	child.(foundation.ChildWhole).SetParent(scene)

	return true
}

func (scene *_Scene) RemoveChild(entUID uint64) {
	if scene.GetDestroyed() {
		return
	}

	if _, ok := scene.childMap[entUID]; !ok {
		return
	}

	delete(scene.childMap, entUID)

	for i, child := range scene.childList {
		if entUID == child.GetUID() {
			scene.childList = append(scene.childList[:i], scene.childList[i+1:]...)
			child.(foundation.ChildWhole).SetParent(nil)
			break
		}
	}
}

func (scene *_Scene) GetChild(entUID uint64) (core.Object, bool) {
	object, ok := scene.childMap[entUID]
	return object, ok
}

func (scene *_Scene) RangeChildren(fun func(child core.Object) bool) {
	if fun == nil {
		return
	}

	for _, child := range scene.childList {
		if !fun(child) {
			return
		}
	}
}
