package gameobj

import (
	"simplegame/core"
	"simplegame/shell"
	"time"
)

type SnakeDirection uint32

const (
	SnakeDirection_Up SnakeDirection = iota
	SnakeDirection_Down
	SnakeDirection_Left
	SnakeDirection_Right
	SnakeDirection_Count
)

type Snake struct {
	shell.Atom
	screen shell.Screen

	Direction    SnakeDirection
	MoveInterval time.Duration
	Length       int

	moveTime time.Time
	maps     shell.VertexMaps
}

func (snake *Snake) Init(object core.Object, name string) {
	snake.Atom = object.(shell.Atom)
	snake.screen = snake.GetEnvironment().GetValue("screen").(shell.Screen)
	snake.moveTime = time.Now()
	snake.Reset()
}

func (snake *Snake) Shut() {
}

func (snake *Snake) Update(frameCtx core.FrameContext) {
	if time.Now().Sub(snake.moveTime) > snake.MoveInterval {

	}

	shell.Posi2D{}

	snake.screen.DrawMaps(10)
}

func (snake *Snake) Reset() {
	for i := 0; i < snake.Length; i++ {

	}
}
