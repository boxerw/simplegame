package logic

import (
	"github.com/boxerw/simplegame/client"
	"github.com/boxerw/simplegame/core"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type GamingStage struct {
	client.Scene
	client.ControllerEventHook
	screen     client.Screen
	controller client.Controller
	snakeObj   *Snake
	itemsObj   *Items
}

func (gamingStage *GamingStage) Init(object core.Object, name string) {
	gamingStage.Scene = object.(client.Scene)

	gamingStage.screen = gamingStage.GetEnvironment().GetValue("screen").(client.Screen)
	gamingStage.screen.SetCanvasFGBG(termbox.ColorWhite, termbox.ColorBlack)

	gamingStage.controller = client.NewController(object.GetEnvironment())
	gamingStage.controller.AddHook(gamingStage)

	gamingStage.Reset()
}

func (gamingStage *GamingStage) Shut() {
	gamingStage.controller.Destroy()

	if gamingStage.snakeObj != nil {
		gamingStage.snakeObj.Destroy()
	}

	if gamingStage.itemsObj != nil {
		gamingStage.itemsObj.Destroy()
	}
}

func (gamingStage *GamingStage) Update(frameCtx core.FrameContext) {
	gamingStage.controller.Update(frameCtx)

	tips := "按'q'键返回，按'w'键加速，按's'键减速"

	size := gamingStage.screen.GetCanvasSize()
	pos := client.Posi2D{size.GetX()/2 - client.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

	gamingStage.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)

	if _, b := gamingStage.CheckBlocked(); b {
		gamingStage.Reset()
	}

	if pos, b := gamingStage.CheckEatFruits(); b {
		gamingStage.snakeObj.ExtendBody(5)
		gamingStage.itemsObj.RemoveFruit(pos)
		gamingStage.itemsObj.AddFruit(1)
		gamingStage.itemsObj.AddWall(1)
	}
}

func (gamingStage *GamingStage) OnControllerKeyPress(controller client.Controller, key termbox.Key, ch rune) bool {
	if gamingStage.snakeObj == nil {
		return true
	}

	switch key {
	case termbox.KeyArrowUp:
		if gamingStage.snakeObj.Direction != SnakeDirection_Down {
			gamingStage.snakeObj.Direction = SnakeDirection_Up
		}
	case termbox.KeyArrowDown:
		if gamingStage.snakeObj.Direction != SnakeDirection_Up {
			gamingStage.snakeObj.Direction = SnakeDirection_Down
		}
	case termbox.KeyArrowLeft:
		if gamingStage.snakeObj.Direction != SnakeDirection_Right {
			gamingStage.snakeObj.Direction = SnakeDirection_Left
		}
	case termbox.KeyArrowRight:
		if gamingStage.snakeObj.Direction != SnakeDirection_Left {
			gamingStage.snakeObj.Direction = SnakeDirection_Right
		}
	}

	switch ch {
	case 'w':
		if gamingStage.snakeObj.MoveInterval > 100*time.Millisecond {
			gamingStage.snakeObj.MoveInterval -= 100 * time.Millisecond
		}
	case 's':
		if gamingStage.snakeObj.MoveInterval < 1500*time.Millisecond {
			gamingStage.snakeObj.MoveInterval += 100 * time.Millisecond
		}
	}

	return true
}

func (gamingStage *GamingStage) Reset() {
	gamingStage.screen.SetCanvasFGBG(termbox.ColorWhite, termbox.ColorBlack)

	if gamingStage.snakeObj == nil {
		snake := client.NewAtom(gamingStage.GetEnvironment(), core.NewComponentBundle("GameObj", &Snake{
			Direction:    SnakeDirection(rand.Intn(int(SnakeDirection_Count))),
			MoveInterval: 800 * time.Millisecond,
			Length:       10,
			Color:        termbox.ColorLightGray,
		}))
		gamingStage.AddChild(snake)

		snakeObj := snake.GetComponent("GameObj").(*Snake)

		size := gamingStage.screen.GetCanvasSize()
		snakeObj.SetPosi(client.Vec2{float32(rand.Intn(size.GetX())), float32(rand.Intn(size.GetY()))})

		gamingStage.snakeObj = snakeObj
	} else {
		gamingStage.snakeObj.Direction = SnakeDirection(rand.Intn(int(SnakeDirection_Count)))
		gamingStage.snakeObj.MoveInterval = 800 * time.Millisecond
		gamingStage.snakeObj.Length = 10

		size := gamingStage.screen.GetCanvasSize()
		gamingStage.snakeObj.SetPosi(client.Vec2{float32(rand.Intn(size.GetX())), float32(rand.Intn(size.GetY()))})

		gamingStage.snakeObj.Reset()
	}

	if gamingStage.itemsObj == nil {
		walls := client.NewAtom(gamingStage.GetEnvironment(), core.NewComponentBundle("GameObj", &Items{
			WallColor:  termbox.ColorRed,
			WallNum:    1,
			FruitColor: termbox.ColorGreen,
			FruitNum:   1,
		}))
		gamingStage.AddChild(walls)

		itemsObj := walls.GetComponent("GameObj").(*Items)

		gamingStage.itemsObj = itemsObj
	} else {
		gamingStage.itemsObj.WallNum = 1
		gamingStage.itemsObj.FruitNum = 1
		gamingStage.itemsObj.Reset()
	}
}

func (gamingStage *GamingStage) CheckBlocked() (client.Posi2D, bool) {
	if gamingStage.snakeObj == nil {
		return client.Posi2D{}, false
	}

	var pos client.Posi2D
	pos.FromVec(gamingStage.snakeObj.GetPosi())

	size := gamingStage.screen.GetCanvasSize()

	if pos.GetX() < 0 || pos.GetX() >= size.GetX() {
		return pos, true
	}

	if pos.GetY() < 0 || pos.GetY() >= size.GetY() {
		return pos, true
	}

	block := false

	gamingStage.snakeObj.RangeBody(func(blockPos client.Posi2D) bool {
		if blockPos == pos {
			block = true
			return false
		}
		return true
	})

	if gamingStage.itemsObj != nil {
		gamingStage.itemsObj.RangeWalls(func(blockPos client.Posi2D) bool {
			if blockPos == pos {
				block = true
				return false
			}
			return true
		})
	}

	return pos, block
}

func (gamingStage *GamingStage) CheckEatFruits() (client.Posi2D, bool) {
	if gamingStage.snakeObj == nil {
		return client.Posi2D{}, false
	}

	var pos client.Posi2D
	pos.FromVec(gamingStage.snakeObj.GetPosi())

	eat := false

	gamingStage.itemsObj.RangeFruit(func(posi client.Posi2D) bool {
		if pos == posi {
			eat = true
			return false
		}
		return true
	})

	return pos, eat
}
