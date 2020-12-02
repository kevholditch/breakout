package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BallPhysicsSystem(t *testing.T) {
	testCases := []struct {
		name                 string
		blocks               []*BlockEntity
		playerState          *components.PlayerStateComponent
		playerQuad           *components.Quad
		playingSpace         PlayingSpace
		ball                 *BallPhysicsComponent
		expectedBallPosition mgl32.Vec2
		expectedBallSpeed    mgl32.Vec2
	}{
		{
			name:                 "during kick off ball is stuck to bat",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Kickoff},
			ball:                 NewBallPhysicsComponent([2]float32{0, 0}, components.NewCircleWithColour(300, 300, 12, colourWhite)),
			expectedBallPosition: [2]float32{125, 74},
			expectedBallSpeed:    [2]float32{0, 0},
		},
		{
			name:                 "ball going left and hits left side of screen",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{-1, 0}, components.NewCircleWithColour(1, 300, 12, colourWhite)),
			expectedBallPosition: [2]float32{0, 300},
			expectedBallSpeed:    [2]float32{1, 0},
		},
		{
			name:                 "ball going right and hits right side of screen",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{1, 0}, components.NewCircleWithColour(799, 300, 12, colourWhite)),
			expectedBallPosition: [2]float32{800, 300},
			expectedBallSpeed:    [2]float32{-1, 0},
		},
		{
			name:                 "ball hits left most edge of players bat",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, -1}, components.NewCircleWithColour(100, 71, 10, colourWhite)),
			expectedBallPosition: [2]float32{100, 70},
			expectedBallSpeed:    [2]float32{0, 1},
		},
		{
			name:                 "ball hits right most edge of players bat",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, -1}, components.NewCircleWithColour(200, 71, 10, colourWhite)),
			expectedBallPosition: [2]float32{200, 70},
			expectedBallSpeed:    [2]float32{0, 1},
		},
		{
			name:                 "ball goes past left side of players bat",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, -1}, components.NewCircleWithColour(90, 71, 10, colourWhite)),
			expectedBallPosition: [2]float32{90, 70},
			expectedBallSpeed:    [2]float32{0, -1},
		},
		{
			name:                 "ball goes past right side of players bat",
			blocks:               []*BlockEntity{},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, -1}, components.NewCircleWithColour(210, 71, 10, colourWhite)),
			expectedBallPosition: [2]float32{210, 70},
			expectedBallSpeed:    [2]float32{0, -1},
		},

		// hit blocks
		{
			name: "ball hits top of block going straight down",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, -1}, components.NewCircleWithColour(550, 561, 10, colourWhite)),
			expectedBallPosition: [2]float32{550, 560},
			expectedBallSpeed:    [2]float32{0, 1},
		},
		{
			name: "ball hits top of block going down",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{10, -1}, components.NewCircleWithColour(550, 561, 10, colourWhite)),
			expectedBallPosition: [2]float32{560, 560},
			expectedBallSpeed:    [2]float32{10, 1},
		},
		{
			name: "ball hits right of block going dead right",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{1, 0}, components.NewCircleWithColour(489, 525, 10, colourWhite)),
			expectedBallPosition: [2]float32{490, 525},
			expectedBallSpeed:    [2]float32{-1, 0},
		},
		{
			name: "ball hits right of block going right",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{1, 10}, components.NewCircleWithColour(489, 525, 10, colourWhite)),
			expectedBallPosition: [2]float32{490, 535},
			expectedBallSpeed:    [2]float32{-1, 10},
		},
		{
			name: "ball hits bottom of block going straight up",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{0, 1}, components.NewCircleWithColour(550, 489, 10, colourWhite)),
			expectedBallPosition: [2]float32{550, 490},
			expectedBallSpeed:    [2]float32{0, -1},
		},
		{
			name: "ball hits bottom of block going up",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{10, 1}, components.NewCircleWithColour(550, 489, 10, colourWhite)),
			expectedBallPosition: [2]float32{560, 490},
			expectedBallSpeed:    [2]float32{10, -1},
		},
		{
			name: "ball hits left of block going dead left",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{-1, 0}, components.NewCircleWithColour(611, 525, 10, colourWhite)),
			expectedBallPosition: [2]float32{610, 525},
			expectedBallSpeed:    [2]float32{1, 0},
		},
		{
			name: "ball hits left of block going left",
			blocks: []*BlockEntity{
				NewBlockEntity(components.NewRenderComponent(components.NewQuadWithColour(500, 500, 100, 50, colourWhite))),
			},
			playingSpace:         NewPlayingSpace(800, 600),
			playerQuad:           components.NewQuadWithColour(100, 40, 100, 20, colourCoral),
			playerState:          &components.PlayerStateComponent{State: components.Playing},
			ball:                 NewBallPhysicsComponent([2]float32{-1, 10}, components.NewCircleWithColour(611, 525, 10, colourWhite)),
			expectedBallPosition: [2]float32{610, 535},
			expectedBallSpeed:    [2]float32{1, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			system := NewBallPhysicsSystem(tc.playerQuad, tc.playerState, tc.playingSpace, tc.ball)
			system.New(&ecs.World{})

			for _, b := range tc.blocks {
				system.Add(b.GetBasicEntity(), b.RenderComponent)
			}

			system.Update(1)
			assert.Equal(t, tc.expectedBallSpeed, tc.ball.Speed)
			assert.Equal(t, tc.expectedBallPosition, tc.ball.Circle.Position)

		})
	}

}
