package logic

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"simplegame/core"
	"simplegame/shell"
)

type Items struct {
	shell.Atom
	screen              shell.Screen
	wallMaps, fruitMaps shell.VertexMaps

	WallNum    int
	WallColor  termbox.Attribute
	FruitNum   int
	FruitColor termbox.Attribute
}

func (items *Items) Init(object core.Object, name string) {
	items.Atom = object.(shell.Atom)
	items.screen = items.GetEnvironment().GetValue("screen").(shell.Screen)
	items.Reset()
}

func (items *Items) Shut() {
}

func (items *Items) Update(frameCtx core.FrameContext) {
	items.screen.DrawMaps(10, shell.Posi2D{}, &items.wallMaps)
	items.screen.DrawMaps(10, shell.Posi2D{}, &items.fruitMaps)
}

func (items *Items) Reset() {
	items.wallMaps.List = nil
	items.fruitMaps.List = nil
	items.AddWall(items.WallNum)
	items.AddFruit(items.FruitNum)
}

func (items *Items) AddWall(num int) {
	size := items.screen.GetCanvasSize()

	for i := 0; i < num; i++ {
		pos := shell.Posi2D{rand.Intn(size.GetX()), rand.Intn(size.GetY())}

		if func() bool {
			for _, v := range items.fruitMaps.List {
				if v.Posi == pos {
					return true
				}
			}
			return false
		}() {
			continue
		}

		items.wallMaps.List = append(items.wallMaps.List, shell.Vertex{
			Posi: pos,
			Pixel: shell.Pixel{
				Ch: '#',
				Fg: items.WallColor,
				Transparent: shell.Transparent{
					Bg: true,
				},
			},
		})
	}

	items.WallNum = len(items.wallMaps.List)
}

func (items *Items) RemoveWall(pos shell.Posi2D) {
	for i, v := range items.wallMaps.List {
		if v.Posi == pos {
			items.wallMaps.List = append(items.wallMaps.List[:i], items.wallMaps.List[i+1:]...)
			break
		}
	}
	items.WallNum = len(items.wallMaps.List)
}

func (items *Items) RangeWalls(fun func(posi shell.Posi2D) bool) {
	if fun == nil {
		return
	}

	for _, v := range items.wallMaps.List {
		if !fun(v.Posi) {
			return
		}
	}
}

func (items *Items) AddFruit(num int) {
	size := items.screen.GetCanvasSize()

	for i := 0; i < num; i++ {
		pos := shell.Posi2D{rand.Intn(size.GetX()), rand.Intn(size.GetY())}

		if func() bool {
			for _, v := range items.wallMaps.List {
				if v.Posi == pos {
					return true
				}
			}
			return false
		}() {
			items.RemoveWall(pos)
		}

		items.fruitMaps.List = append(items.fruitMaps.List, shell.Vertex{
			Posi: pos,
			Pixel: shell.Pixel{
				Ch: '$',
				Fg: items.FruitColor,
				Transparent: shell.Transparent{
					Bg: true,
				},
			},
		})
	}

	items.FruitNum = len(items.fruitMaps.List)
}

func (items *Items) RemoveFruit(pos shell.Posi2D) {
	for i, v := range items.fruitMaps.List {
		if v.Posi == pos {
			items.fruitMaps.List = append(items.fruitMaps.List[:i], items.fruitMaps.List[i+1:]...)
			break
		}
	}
	items.FruitNum = len(items.fruitMaps.List)
}

func (items *Items) RangeFruit(fun func(posi shell.Posi2D) bool) {
	if fun == nil {
		return
	}

	for _, v := range items.fruitMaps.List {
		if !fun(v.Posi) {
			return
		}
	}
}
