package core

func NewScene(ctx *Context, logic ...Logic) *Scene {
	scene := &Scene{
		childMap: map[uint64]Object{},
	}
	scene.initObject(ctx, scene, logic...)
	return scene
}

type Scene struct {
	_ObjectBase
	childList []Object
	childMap  map[uint64]Object
}

func (scene *Scene) Destroy() {
	if scene.GetZombie() {
		return
	}

	for _, child := range scene.childList {
		child.Destroy()
	}

	scene.shutObject()
}

func (scene *Scene) Update() {
	if scene.GetZombie() {
		return
	}

	scene.updateObject()

	for _, child := range scene.childList {
		child.Update()
	}
}

func (scene *Scene) AddChild(object Object) bool {
	if scene.GetZombie() || object.GetZombie() {
		return false
	}

	child, ok := object.(_Child)
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
	child.setParent(scene)

	return true
}

func (scene *Scene) RemoveChild(entUID uint64) {
	if scene.GetZombie() {
		return
	}

	delete(scene.childMap, entUID)

	for i, child := range scene.childList {
		if entUID == child.GetUID() {
			scene.childList = append(scene.childList[:i], scene.childList[i+1:]...)
			child.(_Child).setParent(nil)
			break
		}
	}
}

func (scene *Scene) GetChild(entUID uint64) (Object, bool) {
	object, ok := scene.childMap[entUID]
	return object, ok
}

func (scene *Scene) GetChildList() []Object {
	return scene.childList
}
