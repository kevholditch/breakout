package game

import (
	"fmt"
	"github.com/kevholditch/breakout/internal/pkg/ecs"

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

	world := ecs.NewWorld()
	player := NewPlayer()
	ball := NewBall()
	//hud := NewHud(barHeight, WindowDimensions{
	//	Width:  width,
	//	Height: height,
	//})
	//
	playingSpace := NewPlayingSpace(width, height-barHeight)
	//world.AddSystem(NewFrameRateSystem())

	world.AddEntity(player)
	world.AddEntity(ball)

	levelFactory := NewLevelFactory(playingSpace)
	blocks := levelFactory.NewLevel()
	for _, block := range blocks {
		world.AddEntity(block)
	}

	//	renderSystem.Add(&player.BasicEntity, player.RenderComponent)
	//renderSystem.Add(&ball.BasicEntity, ball.RenderComponent)
	//renderSystem.Add(&hud.BasicEntity, hud.RenderComponent)

	world.AddSystem(NewQuadRenderSystem(NewWindowSize(width, height)))
	world.AddSystem(NewCircleRenderSystem(NewWindowSize(width, height)))
	world.AddSystem(NewLateralMovementSystem(playingSpace))
	world.AddSystem(NewPlayerInputSystem(w.OnKeyPress))
	//world.AddSystem(NewLevelSystem(playingSpace))
	//
	//world.AddSystem(NewBallPhysicsSystem(player.RenderComponent.Quad, player.StateComponent, playingSpace, ball.BallPhysicsComponent))

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
