package logic

import (
	"github.com/nsf/termbox-go"
	"simplegame/client"
	"simplegame/core"
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
	client.Atom
	screen    client.Screen
	moveTime  time.Time
	bodyMaps  client.VertexMaps
	noInitPos bool

	Direction    SnakeDirection
	MoveInterval time.Duration
	Length       int
	Color        termbox.Attribute
}

func (snake *Snake) Init(object core.Object, name string) {
	snake.Atom = object.(client.Atom)
	snake.screen = snake.GetEnvironment().GetValue("screen").(client.Screen)
	snake.Reset()
}

func (snake *Snake) Shut() {
}

func (snake *Snake) Update(frameCtx core.FrameContext) {
	if snake.Length <= 0 {
		return
	}

	var pos client.Posi2D
	pos.FromVec(snake.GetPosi())

	delta := client.Posi2D{pos.GetX() - snake.bodyMaps.List[0].Posi.GetX(), pos.GetY() - snake.bodyMaps.List[0].Posi.GetY()}

	for i := 0; i < len(snake.bodyMaps.List); i++ {
		v := &snake.bodyMaps.List[i]
		if i <= 0 {
			v.Posi = pos
		} else {
			v.Posi = client.Posi2D{v.Posi.GetX() + delta.GetX(), v.Posi.GetY() + delta.GetY()}
		}
	}

	snake.noInitPos = false

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

		head := &snake.bodyMaps.List[0]
		head.Pixel.Ch = '*'
		head.Pixel.Fg = snake.Color

		newHead := client.Vertex{}
		newHead.Pixel.Ch = '@'
		newHead.Pixel.Fg = snake.Color
		newHead.Pixel.Transparent.Bg = true
		newHead.Posi.FromVec(snake.GetPosi())

		t := snake.bodyMaps.List[:len(snake.bodyMaps.List)-1]
		snake.bodyMaps.List = append([]client.Vertex{}, newHead)
		snake.bodyMaps.List = append(snake.bodyMaps.List, t...)
	}

	snake.screen.DrawMaps(10, client.Posi2D{}, &snake.bodyMaps)
}

func (snake *Snake) Reset() {
	if snake.Length <= 0 {
		panic("length invalid")
	}

	snake.moveTime = time.Now()
	snake.noInitPos = true

	snake.bodyMaps.List = make([]client.Vertex, snake.Length)
	for i := 0; i < len(snake.bodyMaps.List); i++ {
		v := &snake.bodyMaps.List[i]

		if i <= 0 {
			v.Pixel.Ch = '@'
			v.Pixel.Fg = snake.Color
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

func (snake *Snake) ExtendBody(num int) {
	for i := 0; i < num; i++ {
		v := client.Vertex{}
		v.Pixel.Transparent.Fg = true
		v.Pixel.Transparent.Bg = true
		v.Pixel.Transparent.Ch = true
		v.Pixel.Transparent.Attr = true
		v.Posi = snake.bodyMaps.List[len(snake.bodyMaps.List)-1].Posi

		snake.bodyMaps.List = append(snake.bodyMaps.List, v)
	}

	snake.Length = len(snake.bodyMaps.List)
}

func (snake *Snake) RangeBody(fun func(posi client.Posi2D) bool) {
	if fun == nil || snake.noInitPos {
		return
	}

	for i, v := range snake.bodyMaps.List {
		if i > 0 {
			if !fun(v.Posi) {
				return
			}
		}
	}
}
