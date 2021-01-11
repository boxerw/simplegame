package gameobj

import (
	"simplegame/core"
	"simplegame/shell"
)

type Fruit struct {
	shell.Atom
}

func (fruit *Fruit) Init(object core.Object, name string) {
	fruit.Atom = object.(shell.Atom)
}

func (fruit *Fruit) Shut() {

}

func (fruit *Fruit) Update(frameCtx core.FrameContext) {

}
