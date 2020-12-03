package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BallPhysicsSystem(t *testing.T) {
	testCases := []struct {
		name                 string
		blocks               []*ecs.Entity
		gameState            *GameState
		playerPosition       *components.PositionedComponent
		playerDimensions     *components.DimensionComponent
		ballPosition         *components.PositionedComponent
		ballDimensions       *components.DimensionComponent
		ballSpeed            *components.SpeedComponent
		expectedBallPosition mgl32.Vec2
		expectedBallSpeed    mgl32.Vec2
	}{
		{
			name:                 "ball going left and hits left side of screen",
			blocks:               []*ecs.Entity{},
			playerPosition:       components.NewPositionedComponent(100, 40),
			playerDimensions:     components.NewDimensionsComponent(100, 20),
			gameState:            NewGameState(),
			ballSpeed:            components.NewSpeedComponent([2]float32{-1, 0}),
			ballPosition:         components.NewPositionedComponent(1, 300),
			ballDimensions:       components.NewDimensionsComponent(10, 10),
			expectedBallPosition: [2]float32{0, 300},
			expectedBallSpeed:    [2]float32{1, 0},
		},
		//{
		//	name:                 "ball going left and hits left side of screen",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{-1, 0}, primitives.NewCircleWithColour(1, 300, 12, colourWhite)),
		//	expectedBallPosition: [2]float32{0, 300},
		//	expectedBallSpeed:    [2]float32{1, 0},
		//},
		//{
		//	name:                 "ball going right and hits right side of screen",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{1, 0}, primitives.NewCircleWithColour(799, 300, 12, colourWhite)),
		//	expectedBallPosition: [2]float32{800, 300},
		//	expectedBallSpeed:    [2]float32{-1, 0},
		//},
		//{
		//	name:                 "ball hits left most edge of players bat",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, -1}, primitives.NewCircleWithColour(100, 71, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{100, 70},
		//	expectedBallSpeed:    [2]float32{0, 1},
		//},
		//{
		//	name:                 "ball hits right most edge of players bat",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, -1}, primitives.NewCircleWithColour(200, 71, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{200, 70},
		//	expectedBallSpeed:    [2]float32{0, 1},
		//},
		//{
		//	name:                 "ball goes past left side of players bat",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, -1}, primitives.NewCircleWithColour(90, 71, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{90, 70},
		//	expectedBallSpeed:    [2]float32{0, -1},
		//},
		//{
		//	name:                 "ball goes past right side of players bat",
		//	blocks:               []*BlockEntity{},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, -1}, primitives.NewCircleWithColour(210, 71, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{210, 70},
		//	expectedBallSpeed:    [2]float32{0, -1},
		//},
		//
		//// hit blocks
		//{
		//	name: "ball hits top of block going straight down",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, -1}, primitives.NewCircleWithColour(550, 561, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{550, 560},
		//	expectedBallSpeed:    [2]float32{0, 1},
		//},
		//{
		//	name: "ball hits top of block going down",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{10, -1}, primitives.NewCircleWithColour(550, 561, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{560, 560},
		//	expectedBallSpeed:    [2]float32{10, 1},
		//},
		//{
		//	name: "ball hits right of block going dead right",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{1, 0}, primitives.NewCircleWithColour(489, 525, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{490, 525},
		//	expectedBallSpeed:    [2]float32{-1, 0},
		//},
		//{
		//	name: "ball hits right of block going right",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{1, 10}, primitives.NewCircleWithColour(489, 525, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{490, 535},
		//	expectedBallSpeed:    [2]float32{-1, 10},
		//},
		//{
		//	name: "ball hits bottom of block going straight up",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{0, 1}, primitives.NewCircleWithColour(550, 489, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{550, 490},
		//	expectedBallSpeed:    [2]float32{0, -1},
		//},
		//{
		//	name: "ball hits bottom of block going up",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{10, 1}, primitives.NewCircleWithColour(550, 489, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{560, 490},
		//	expectedBallSpeed:    [2]float32{10, -1},
		//},
		//{
		//	name: "ball hits left of block going dead left",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{-1, 0}, primitives.NewCircleWithColour(611, 525, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{610, 525},
		//	expectedBallSpeed:    [2]float32{1, 0},
		//},
		//{
		//	name: "ball hits left of block going left",
		//	blocks: []*BlockEntity{
		//		NewBlockEntity(components.NewRenderComponent(primitives.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
		//	},
		//	playingSpace:         NewPlayingSpace(800, 600),
		//	playerQuad:           primitives.NewQuadWithColour(100, 40, 100, 20, colourCoral),
		//	playerState:          &components.PlayerStateComponent{State: components.Playing},
		//	ball:                 NewBallPhysicsComponent([2]float32{-1, 10}, primitives.NewCircleWithColour(611, 525, 10, colourWhite)),
		//	expectedBallPosition: [2]float32{610, 535},
		//	expectedBallSpeed:    [2]float32{1, 10},
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			levelSystem := NewLevelSystem()
			system := NewBallPhysicsSystem(tc.playerPosition, tc.playerDimensions, NewPlayingSpace(800, 600), levelSystem, tc.gameState)
			system.New(&ecs.World{})

			ball := ecs.NewEntity()
			ball.Add(tc.ballDimensions)
			ball.Add(tc.ballPosition)
			ball.Add(tc.ballSpeed)
			ball.Add(components.NewBallPhysicsComponent())
			ball.Add(components.NewCircleComponent(10))

			for _, b := range tc.blocks {
				levelSystem.Add(b)
			}

			system.Update(1)
			assert.Equal(t, tc.expectedBallSpeed, tc.ballSpeed.Speed)
			assert.Equal(t, tc.expectedBallPosition[0], tc.ballPosition.X)
			assert.Equal(t, tc.expectedBallPosition[1], tc.ballPosition.Y)

		})
	}

}
