package foundation

type Transform2D interface {
	GetPosi() Vec2
	SetPosi(pos Vec2)
	GetPosiX() float32
	SetPosiX(v float32)
	GetPosiY() float32
	SetPosiY(v float32)
	GetAngle() float32
	SetAngle(angle float32)
}
