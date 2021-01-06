package main

import (
	"simplegame/core"
	"simplegame/demo/snake/logic"
	"simplegame/model"
)

func main() {
	env := core.NewEnvironment()

	screen := model.NewScreen(env, core.NewComponentBundle("ShowInfo", &logic.ShowInfo{}))
	defer screen.Destroy()
	env.SetValue("screen", screen)

	scene := model.NewScene(env, core.NewComponentBundle("MainFlow", &logic.MainFlow{}))
	defer scene.Destroy()
	env.SetValue("mainScene", scene)

	exec := core.NewExecute(30, true, env, screen, scene)
	env.SetValue("execute", exec)

	defer exec.Start().Wait()
}