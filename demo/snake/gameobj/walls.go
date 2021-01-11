package gameobj

import (
	"simplegame/core"
	"simplegame/shell"
)

type Walls struct {
	shell.Atom
	Size shell.Posi2D
	maps shell.Vertex
}

func (walls *Walls) Init(object core.Object, name string) {
	walls.Atom = object.(shell.Atom)
}

func (walls *Walls) Shut() {
}

func (walls *Walls) Update(frameCtx core.FrameContext) {
}

func (walls *Walls) RandWall(num int) {
	for i := 0; i < num; i++ {
		walls.wallList = append()
	}
}
