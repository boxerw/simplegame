package model

import (
	"math"
	"simplegame/core/foundation"
)

type Posi2D [2]int

func (posi *Posi2D) GetX() int {
	return (*posi)[0]
}

func (posi *Posi2D) SetX(v int) {
	(*posi)[0] = v
}

func (posi *Posi2D) GetY() int {
	return (*posi)[1]
}

func (posi *Posi2D) SetY(v int) {
	(*posi)[1] = v
}

func (posi *Posi2D) FromVec(vec foundation.Vec2) {
	(*posi)[0] = int(math.Round(float64(vec[0])))
	(*posi)[1] = int(math.Round(float64(vec[1])))
}

func (posi *Posi2D) ToVec() foundation.Vec2 {
	return foundation.Vec2{float32(posi.GetX()), float32(posi.GetY())}
}
