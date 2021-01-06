package main

import (
	"simplegame/core"
	"simplegame/demo/snake/logic"
	"simplegame/model"
)

func main() {
	env := core.NewEnvironment()

	screen := model.NewScreen(env)
	env.SetValue("screen", screen)

	screen.AddComponent(core.NewComponentBundle("ShowInfo", &logic.ShowInfo{}))

	deviceExec := core.NewExecute(60, false, env, screen)
	env.SetValue("deviceExec", deviceExec)

	wg := deviceExec.Start()
	defer wg.Wait()
}
