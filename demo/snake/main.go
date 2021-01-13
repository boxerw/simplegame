package main

import (
	"github.com/boxerw/simplegame/client"
	"github.com/boxerw/simplegame/core"
	"github.com/boxerw/simplegame/demo/snake/logic"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	env := core.NewEnvironment(0)

	screen := client.NewScreen(env, core.NewComponentBundle("ShowInfo", &logic.ShowInfo{}))
	defer screen.Destroy()
	env.SetValue("screen", screen)

	scene := client.NewScene(env, core.NewComponentBundle("MainFlow", &logic.MainFlow{}))
	defer scene.Destroy()
	env.SetValue("scene", scene)

	exec := core.NewExecute(30, true, screen, scene)
	env.SetValue("execute", exec)

	defer exec.Start().Wait()
}
