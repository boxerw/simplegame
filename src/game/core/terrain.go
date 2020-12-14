package core

func NewTerrain(ctx *Context, logic ...Logic) *Terrain {
	terrain := &Terrain{}
	terrain.initObject(ctx, terrain, logic...)
	return terrain
}

type Terrain struct {
	_ObjectBase
}

func (terrain *Terrain) Destroy() {
	terrain.shutObject()
}

func (terrain *Terrain) Update() {
	terrain.updateObject()
}
