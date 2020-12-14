package core

type Transform interface {
	GetPosi() Vec3
	SetPosi(pos Vec3)
	GetAngle() float32
	SetAngle(angle float32)
}

type _TransformBase struct {
	posi  Vec3
	angle float32
}

func (transBase *_TransformBase) GetPosi() Vec3 {
	return transBase.posi
}

func (transBase *_TransformBase) SetPosi(pos Vec3) {
	transBase.posi = pos
}

func (transBase *_TransformBase) GetAngle() float32 {
	return transBase.angle
}

func (transBase *_TransformBase) SetAngle(angle float32) {
	transBase.angle = angle
}
