package game

import (
	"simple/game/core"
	"simple/game/logic"
)

func Run() {
	ctx := core.NewContext()

	exec := core.NewExecute(
		30,
		false,
		ctx,
		core.NewScene(ctx, &logic.LogicMainScene{}),
		core.NewScreen(ctx),
	)

	exec.Run()
}
