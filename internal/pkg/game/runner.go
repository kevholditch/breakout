package game

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"runtime"
	"time"

	"github.com/kevholditch/breakout/internal/pkg/render"
)

const (
	width, height = 1024, 768
	version       = "v0.0.1"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func Run() error {

	cleanUp := render.Initialise()
	defer cleanUp()

	w, err := render.NewWindow(render.Config{
		MajorVersion: 3,
		MinorVersion: 2,
		Width:        width,
		Height:       height,
		Title:        fmt.Sprintf("Breakout %s- Kevin Holditch", version),
		SwapInterval: 1,
	})

	if err != nil {
		return err
	}

	world := ecs.World{}
	player := NewPlayer()

	world.AddSystem(NewFrameRateSystem())
	world.AddSystem(NewRenderSystem(width, height).Add(player.BasicEntity, player.RenderComponent))
	world.AddSystem(NewPlayerMovementSystem(width, height).Add(&player.BasicEntity, player.MoveComponent))
	world.AddSystem(NewPlayerInputSystem(w.OnKeyPress).Add(&player.BasicEntity, player.MoveComponent))

	last := time.Now()

	for !w.ShouldClose() {
		elapsed := time.Now().Sub(last)
		last = time.Now()
		world.Update(float32(elapsed.Milliseconds()))

		w.SwapBuffers()
		w.PollEvents()
	}

	return nil
}
