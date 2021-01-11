package logic

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"simplegame/core"
	"simplegame/demo/snake/gameobj"
	"simplegame/shell"
	"time"
)

type GamingStage struct {
	shell.Scene
	shell.ControllerEventHook
	screen       shell.Screen
	controller   shell.Controller
	snakeObj     *gameobj.Snake
	wallsObj     *gameobj.Walls
	fruitObjList []*gameobj.Fruit
}

func (gamingStage *GamingStage) Init(object core.Object, name string) {
	gamingStage.Scene = object.(shell.Scene)

	gamingStage.screen = gamingStage.GetEnvironment().GetValue("screen").(shell.Screen)

	gamingStage.controller = shell.NewController(object.GetEnvironment())
	gamingStage.controller.AddHook(gamingStage)

	gamingStage.Reset()
}

func (gamingStage *GamingStage) Shut() {
	gamingStage.controller.Destroy()

	if gamingStage.snakeObj != nil {
		gamingStage.snakeObj.Destroy()
	}

	if gamingStage.wallsObj != nil {
		gamingStage.wallsObj.Destroy()
	}

	for _, fruit := range gamingStage.fruitObjList {
		fruit.Destroy()
	}
}

func (gamingStage *GamingStage) Update(frameCtx core.FrameContext) {
	gamingStage.controller.Update(frameCtx)

	tips := "按'q'键返回，按'w'键加速，按's'键减速"

	size := gamingStage.screen.GetCanvasSize()
	pos := shell.Posi2D{size.GetX()/2 - shell.StringWidth(tips)/2, int(float32(size.GetY()) * 0.8)}

	gamingStage.screen.DrawText(50, pos, tips, termbox.AttrBlink|termbox.ColorLightGray, termbox.ColorBlue)
}

func (gamingStage *GamingStage) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
	if gamingStage.snakeObj == nil {
		return true
	}

	switch key {
	case termbox.KeyArrowUp:
		if gamingStage.snakeObj.Direction != gameobj.SnakeDirection_Down {
			gamingStage.snakeObj.Direction = gameobj.SnakeDirection_Up
		}
	case termbox.KeyArrowDown:
		if gamingStage.snakeObj.Direction != gameobj.SnakeDirection_Up {
			gamingStage.snakeObj.Direction = gameobj.SnakeDirection_Down
		}
	case termbox.KeyArrowLeft:
		if gamingStage.snakeObj.Direction != gameobj.SnakeDirection_Right {
			gamingStage.snakeObj.Direction = gameobj.SnakeDirection_Left
		}
	case termbox.KeyArrowRight:
		if gamingStage.snakeObj.Direction != gameobj.SnakeDirection_Left {
			gamingStage.snakeObj.Direction = gameobj.SnakeDirection_Right
		}
	}

	switch ch {
	case 'w':
		if gamingStage.snakeObj.MoveInterval > 100*time.Millisecond {
			gamingStage.snakeObj.MoveInterval -= 100 * time.Millisecond
		}
	case 's':
		gamingStage.snakeObj.MoveInterval += 100 * time.Millisecond
	}

	return true
}

func (gamingStage *GamingStage) Reset() {
	if gamingStage.snakeObj != nil {
		gamingStage.snakeObj.Destroy()
		gamingStage.snakeObj = nil
	}

	if gamingStage.wallsObj != nil {
		gamingStage.wallsObj.Destroy()
		gamingStage.snakeObj = nil
	}

	for _, fruit := range gamingStage.fruitObjList {
		fruit.Destroy()
	}
	gamingStage.fruitObjList = nil

	color := termbox.Attribute(rand.Intn(int(termbox.ColorLightGray + 1)))

	fg, bg := gamingStage.screen.GetCanvasFGBG()
	if color == fg {
		color++
	}
	if color == bg {
		color++
	}
	color %= termbox.ColorLightGray + 1

	snake := shell.NewAtom(gamingStage.GetEnvironment(), core.NewComponentBundle("GameObj", &gameobj.Snake{
		Direction:    gameobj.SnakeDirection(rand.Intn(int(gameobj.SnakeDirection_Count))),
		MoveInterval: 800 * time.Millisecond,
		Length:       10,
		Color:        color,
	}))
	gamingStage.AddChild(snake)

	snakeObj := snake.GetComponent("GameObj").(*gameobj.Snake)

	size := gamingStage.screen.GetCanvasSize()
	snakeObj.SetPosi(core.Vec2{float32(rand.Intn(size.GetX())), float32(rand.Intn(size.GetY()))})

	gamingStage.snakeObj = snakeObj
}
