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

func (gaming *GamingStage) Init(object core.Object, name string) {
	gaming.Scene = object.(shell.Scene)

	gaming.screen = gaming.GetEnvironment().GetValue("screen").(shell.Screen)

	gaming.controller = shell.NewController(object.GetEnvironment())
	gaming.controller.AddHook(gaming)

	{
		walls := shell.NewAtom(object.GetEnvironment(), core.NewComponentBundle("GameObj", &gameobj.Walls{}))
		gaming.AddChild(walls)
		gaming.wallsObj = walls.GetComponent("GameObj").(*gameobj.Walls)
	}

	{
		snake := shell.NewAtom(object.GetEnvironment(), core.NewComponentBundle("GameObj", &gameobj.Snake{
			Direction:    gameobj.SnakeDirection(rand.Intn(int(gameobj.SnakeDirection_Count))),
			MoveInterval: 800 * time.Millisecond,
			Length:       5,
		}))
		gaming.AddChild(snake)
		gaming.snakeObj = snake.GetComponent("GameObj").(*gameobj.Snake)

		size := gaming.screen.GetCanvasSize()
		gaming.snakeObj.SetPosi(core.Vec2{float32(rand.Intn(size.GetX())), float32(rand.Intn(size.GetY()))})
	}

}

func (gaming *GamingStage) Shut() {
	gaming.controller.Destroy()
	gaming.snakeObj.Destroy()
	gaming.wallsObj.Destroy()
	for _, fruit := range gaming.fruitObjList {
		fruit.Destroy()
	}
}

func (gaming *GamingStage) Update(frameCtx core.FrameContext) {
	gaming.controller.Update(frameCtx)

}

func (gaming *GamingStage) OnControllerKeyPress(controller shell.Controller, key termbox.Key, ch rune) bool {
	switch key {
	case termbox.KeyArrowUp:
		gaming.snakeObj.Direction = gameobj.SnakeDirection_Up
	case termbox.KeyArrowDown:
		gaming.snakeObj.Direction = gameobj.SnakeDirection_Down
	case termbox.KeyArrowLeft:
		gaming.snakeObj.Direction = gameobj.SnakeDirection_Left
	case termbox.KeyArrowRight:
		gaming.snakeObj.Direction = gameobj.SnakeDirection_Right
	}

	return true
}
