package core

import "simple/core/foundation"

type Transform2D = foundation.Transform2D

type Transform2DModule struct {
	posi  Vec2
	angle float32
}

func (transBase *Transform2DModule) GetPosi() Vec2 {
	return transBase.posi
}

func (transBase *Transform2DModule) SetPosi(pos Vec2) {
	transBase.posi = pos
}

func (transBase *Transform2DModule) GetPosiX() float32 {
	return transBase.posi[0]
}

func (transBase *Transform2DModule) SetPosiX(v float32) {
	transBase.posi[0] = v
}

func (transBase *Transform2DModule) GetPosiY() float32 {
	return transBase.posi[1]
}

func (transBase *Transform2DModule) SetPosiY(v float32) {
	transBase.posi[1] = v
}

func (transBase *Transform2DModule) GetAngle() float32 {
	return transBase.angle
}

func (transBase *Transform2DModule) SetAngle(angle float32) {
	transBase.angle = angle
}
