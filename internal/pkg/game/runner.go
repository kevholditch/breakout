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

	inc := float32(1000)

	world := ecs.World{}
	world.AddSystem(NewFrameRateSystem())

	player := NewPlayer()
	renderSystem := NewRenderSystem(width, height).
		Add(player.BasicEntity, player.RenderComponent)

	err = renderSystem.Initialise()
	if err != nil {
		return err
	}
	world.AddSystem(renderSystem)
	world.AddSystem(NewMovementSystem().Add(player.BasicEntity, player.MoveComponent))

	w.OnKeyPress(func(key int) {
		switch key {
		case 263:
			if player.MoveComponent.Speed > 0 {
				player.MoveComponent.Speed = 0
			} else {
				player.MoveComponent.Speed -= inc
			}
		case 262:
			if player.MoveComponent.Speed < 0 {
				player.MoveComponent.Speed = 0
			} else {
				player.MoveComponent.Speed += inc
			}
		}
	}, func(key int) {

	})

	last := time.Now()

	for !w.ShouldClose() {

		// 100 ms elasped
		// 0.1

		elapsed := time.Now().Sub(last)
		last = time.Now()
		world.Update(float32(elapsed.Milliseconds()) / float32(1000))

		w.SwapBuffers()
		w.PollEvents()

	}

	return nil
}
