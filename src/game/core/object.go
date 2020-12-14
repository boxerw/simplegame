package core

type Object interface {
	Entity
	AddLogic(logic Logic)
	RemoveLogic(name string)
	GetLogic(name string) Logic
	GetLogicList() []Logic
	Update()
}

type _ObjectBase struct {
	_EntityBase
	object    Object
	logicList []Logic
}

func (objBase *_ObjectBase) initObject(ctx *Context, object Object, logic ...Logic) {
	objBase.initEntity(ctx, object)
	objBase.object = object
	objBase.logicList = logic

	for _, logic := range objBase.logicList {
		logic.Init(object)
	}
}

func (objBase *_ObjectBase) shutObject() {
	if objBase.GetZombie() {
		return
	}

	for _, logic := range objBase.logicList {
		logic.Shut()
	}

	objBase.shutEntity()
}

func (objBase *_ObjectBase) updateObject() {
	if objBase.GetZombie() {
		return
	}

	for _, logic := range objBase.logicList {
		logic.Update()
	}
}

func (objBase *_ObjectBase) AddLogic(logic Logic) {
	if objBase.GetZombie() {
		return
	}

	objBase.logicList = append(objBase.logicList, logic)

	logic.Init(objBase.object)
}

func (objBase *_ObjectBase) RemoveLogic(name string) {
	if objBase.GetZombie() {
		return
	}

	for i, logic := range objBase.logicList {
		if logic.Name() == name {
			objBase.logicList = append(objBase.logicList[:i], objBase.logicList[i+1:]...)
			logic.Shut()
		}
	}
}

func (objBase *_ObjectBase) GetLogic(name string) Logic {
	for _, logic := range objBase.logicList {
		if logic.Name() == name {
			return logic
		}
	}
	return nil
}

func (objBase *_ObjectBase) GetLogicList() []Logic {
	return objBase.logicList
}
