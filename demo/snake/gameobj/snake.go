package gameobj

import (
	"github.com/nsf/termbox-go"
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
	Color        termbox.Attribute

	moveTime time.Time
	maps     shell.VertexMaps
}

func (snake *Snake) Init(object core.Object, name string) {
	snake.Atom = object.(shell.Atom)
	snake.screen = snake.GetEnvironment().GetValue("screen").(shell.Screen)
	snake.Reset()
}

func (snake *Snake) Shut() {
}

func (snake *Snake) Update(frameCtx core.FrameContext) {
	if snake.Length <= 0 {
		return
	}

	var pos shell.Posi2D
	pos.FromVec(snake.GetPosi())

	delta := shell.Posi2D{pos.GetX() - snake.maps.List[0].Posi.GetX(), pos.GetY() - snake.maps.List[0].Posi.GetY()}

	for i := 0; i < len(snake.maps.List); i++ {
		v := &snake.maps.List[i]
		if i <= 0 {
			v.Posi = pos
		} else {
			v.Posi = shell.Posi2D{v.Posi.GetX() + delta.GetX(), v.Posi.GetY() + delta.GetY()}
		}
	}

	if time.Now().Sub(snake.moveTime) > snake.MoveInterval {
		switch snake.Direction {
		case SnakeDirection_Up:
			snake.SetPosiY(snake.GetPosiY() - 1)
		case SnakeDirection_Down:
			snake.SetPosiY(snake.GetPosiY() + 1)
		case SnakeDirection_Left:
			snake.SetPosiX(snake.GetPosiX() - 1)
		case SnakeDirection_Right:
			snake.SetPosiX(snake.GetPosiX() + 1)
		}

		snake.moveTime = time.Now()

		head := &snake.maps.List[0]
		head.Pixel.Ch = '*'
		head.Pixel.Fg = snake.Color

		newHead := shell.Vertex{}
		newHead.Pixel.Ch = '@'
		newHead.Pixel.Fg = termbox.AttrBlink | snake.Color
		newHead.Pixel.Transparent.Bg = true
		newHead.Posi.FromVec(snake.GetPosi())

		t := snake.maps.List[:len(snake.maps.List)-1]
		snake.maps.List = append([]shell.Vertex{}, newHead)
		snake.maps.List = append(snake.maps.List, t...)
	}

	snake.screen.DrawMaps(10, shell.Posi2D{}, &snake.maps)
}

func (snake *Snake) Reset() {
	if snake.Length <= 0 {
		panic("length invalid")
	}

	snake.moveTime = time.Now()

	snake.maps.List = make([]shell.Vertex, snake.Length)
	for i := 0; i < len(snake.maps.List); i++ {
		v := &snake.maps.List[i]

		if i <= 0 {
			v.Pixel.Ch = '@'
			v.Pixel.Fg = termbox.AttrBlink | snake.Color
		} else {
			switch snake.Direction {
			case SnakeDirection_Up:
				v.Posi.SetY(i)
			case SnakeDirection_Down:
				v.Posi.SetY(-i)
			case SnakeDirection_Left:
				v.Posi.SetX(i)
			case SnakeDirection_Right:
				v.Posi.SetX(-i)
			}
			v.Pixel.Ch = '*'
			v.Pixel.Fg = snake.Color
		}

		v.Pixel.Transparent.Bg = true
	}
}
