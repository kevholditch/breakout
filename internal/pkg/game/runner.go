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

	barHeight := 50

	world := ecs.World{}
	player := NewPlayer()
	ball := NewBall()
	hud := NewHud(barHeight, WindowDimensions{
		Width:  width,
		Height: height,
	})

	playingSpace := NewPlayingSpace(width, height-barHeight)
	world.AddSystem(NewFrameRateSystem())
	renderSystem := NewRenderSystem(NewWindowSize(width, height))
	renderSystem.Add(&player.BasicEntity, player.RenderComponent)
	renderSystem.Add(&ball.BasicEntity, ball.RenderComponent)
	renderSystem.Add(&hud.BasicEntity, hud.RenderComponent)

	world.AddSystem(renderSystem)
	world.AddSystem(NewPlayerMovementSystem(playingSpace).Add(&player.BasicEntity, player.MoveComponent))
	world.AddSystem(NewPlayerInputSystem(w.OnKeyPress, player.MoveComponent, player.StateComponent, ball.BallPhysicsComponent))
	world.AddSystem(NewLevelSystem(playingSpace))

	world.AddSystem(NewBallPhysicsSystem(player.RenderComponent.Quad, player.StateComponent, playingSpace).
		Add(&ball.BasicEntity, ball.BallPhysicsComponent))

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
