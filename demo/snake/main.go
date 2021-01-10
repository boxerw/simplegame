package main

import (
	"simplegame/core"
	"simplegame/demo/snake/logic"
	"simplegame/shell"
)

func main() {
	env := core.NewEnvironment(0)

	screen := shell.NewScreen(env, core.NewComponentBundle("ShowInfo", &logic.ShowInfo{}))
	defer screen.Destroy()
	env.SetValue("screen", screen)

	scene := shell.NewScene(env, core.NewComponentBundle("MainFlow", &logic.MainFlow{}))
	defer scene.Destroy()
	env.SetValue("scene", scene)

	exec := core.NewExecute(30, true, screen, scene)
	env.SetValue("execute", exec)

	defer exec.Start().Wait()
}
