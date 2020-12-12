package game

import (
	"fmt"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"

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

	world := ecs.NewWorld()
	player := NewPlayer()
	ball := NewBall()
	playingSpace := NewPlayingSpace(width, height)

	world.AddEntity(player)
	world.AddEntity(ball)

	gameState := NewGameState()

	world.AddSystem(NewQuadRenderSystem(NewWindowSize(width, height)))
	world.AddSystem(NewCircleRenderSystem(NewWindowSize(width, height)))
	world.AddSystem(NewLateralMovementSystem(playingSpace))
	world.AddSystem(NewPlayerInputSystem(w.OnKeyPress, gameState))

	levelSystem := NewLevelSystem(playingSpace)
	world.AddSystem(levelSystem)

	world.AddSystem(NewBallPhysicsSystem(
		player.Component(components.IsPositioned).(*components.PositionedComponent),
		player.Component(components.HasDimensions).(*components.DimensionComponent),
		playingSpace,
		levelSystem,
		gameState))

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
