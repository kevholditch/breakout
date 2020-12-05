package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BallPhysicsSystemScreenBoundaries(t *testing.T) {
	testCases := []struct {
		name string
		ball struct {
			X      float32
			Y      float32
			SpeedX float32
			SpeedY float32
		}
		expected struct {
			X      float32
			Y      float32
			SpeedX float32
			SpeedY float32
		}
	}{
		{
			name: "ball going left and hits left side of screen",
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 10, Y: 300, SpeedX: -1, SpeedY: 0},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 10, Y: 300, SpeedX: 1, SpeedY: 0},
		},
		{
			name: "ball going right and hits right side of screen",
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 790, Y: 300, SpeedX: 1, SpeedY: 0},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 790, Y: 300, SpeedX: -1, SpeedY: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			system := NewBallPhysicsSystem(
				components.NewPositionedComponent(100, 40),
				components.NewDimensionsComponent(1, 1),
				NewPlayingSpace(800, 600),
				NewLevelSystem(),
				NewGameState())
			system.New(&ecs.World{})

			ball := ecs.NewEntity()
			position := components.NewPositionedComponent(tc.ball.X, tc.ball.Y)
			ball.Add(position)
			speed := components.NewSpeedComponent([2]float32{tc.ball.SpeedX, tc.ball.SpeedY})
			ball.Add(speed)
			ball.Add(components.NewDimensionsComponent(10, 10))
			ball.Add(components.NewBallPhysicsComponent())
			ball.Add(components.NewCircleComponent(10))

			system.Add(ball)

			system.Update(1)
			assert.Equal(t, components.NewSpeedComponent([2]float32{tc.expected.SpeedX, tc.expected.SpeedY}), speed)
			assert.Equal(t, tc.expected.X, position.X)
			assert.Equal(t, tc.expected.Y, position.Y)

		})
	}
}

func Test_BallPhysicsSystemHittingPlayer(t *testing.T) {
	testCases := []struct {
		name   string
		player struct {
			X      float32
			Y      float32
			Width  float32
			Height float32
		}
		ball struct {
			X      float32
			Y      float32
			SpeedX float32
			SpeedY float32
		}
		expected struct {
			X      float32
			Y      float32
			SpeedX float32
			SpeedY float32
		}
	}{
		{
			name: "ball hits left most edge of players bat whilst going down",
			player: struct {
				X      float32
				Y      float32
				Width  float32
				Height float32
			}{X: 100, Y: 40, Width: 100, Height: 20},
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 100, Y: 71, SpeedX: 0, SpeedY: -1},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 100, Y: 70, SpeedX: 0, SpeedY: 1},
		},
		{
			name: "ball hits right most edge of players bat whilst going down",
			player: struct {
				X      float32
				Y      float32
				Width  float32
				Height float32
			}{X: 100, Y: 40, Width: 100, Height: 20},
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 200, Y: 71, SpeedX: 0, SpeedY: -1},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 200, Y: 70, SpeedX: 0, SpeedY: 1},
		},
		{
			name: "ball goes past left side of players bat whilst going down",
			player: struct {
				X      float32
				Y      float32
				Width  float32
				Height float32
			}{X: 100, Y: 40, Width: 100, Height: 20},
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 90, Y: 71, SpeedX: 0, SpeedY: -1},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 90, Y: 70, SpeedX: 0, SpeedY: -1},
		},
		{
			name: "ball goes past right side of players bat whilst going down",
			player: struct {
				X      float32
				Y      float32
				Width  float32
				Height float32
			}{X: 100, Y: 40, Width: 100, Height: 20},
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 210, Y: 71, SpeedX: 0, SpeedY: -1},
			expected: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 210, Y: 70, SpeedX: 0, SpeedY: -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			system := NewBallPhysicsSystem(
				components.NewPositionedComponent(tc.player.X, tc.player.Y),
				components.NewDimensionsComponent(tc.player.Width, tc.player.Height),
				NewPlayingSpace(800, 600),
				NewLevelSystem(),
				NewGameState())
			system.New(&ecs.World{})

			ball := ecs.NewEntity()
			position := components.NewPositionedComponent(tc.ball.X, tc.ball.Y)
			ball.Add(position)
			speed := components.NewSpeedComponent([2]float32{tc.ball.SpeedX, tc.ball.SpeedY})
			ball.Add(speed)
			ball.Add(components.NewDimensionsComponent(10, 10))
			ball.Add(components.NewBallPhysicsComponent())
			ball.Add(components.NewCircleComponent(10))

			system.Add(ball)

			system.Update(1)
			assert.Equal(t, components.NewSpeedComponent([2]float32{tc.expected.SpeedX, tc.expected.SpeedY}), speed)
			assert.Equal(t, tc.expected.X, position.X)
			assert.Equal(t, tc.expected.Y, position.Y)

		})
	}

}

func Test_BallPhysicsSystemHittingBlock(t *testing.T) {
	testCases := []struct {
		name  string
		block struct {
			X      float32
			Y      float32
			Width  float32
			Height float32
		}
		ball struct {
			X      float32
			Y      float32
			SpeedX float32
			SpeedY float32
		}
		expected struct {
			X        float32
			Y        float32
			SpeedX   float32
			SpeedY   float32
			BlockHit int
		}
	}{
		{
			name: "ball hits top of block going straight down",
			block: struct {
				X      float32
				Y      float32
				Width  float32
				Height float32
			}{X: 400, Y: 400, Width: 100, Height: 100},
			ball: struct {
				X      float32
				Y      float32
				SpeedX float32
				SpeedY float32
			}{X: 450, Y: 411, SpeedX: 0, SpeedY: -1},
			expected: struct {
				X        float32
				Y        float32
				SpeedX   float32
				SpeedY   float32
				BlockHit int
			}{X: 450, Y: 410, SpeedX: 0, SpeedY: 1, BlockHit: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			levelSystem := NewLevelSystem()
			block := ecs.NewEntity()
			block.Add(components.NewPositionedComponent(tc.block.X, tc.block.Y))
			block.Add(components.NewDimensionsComponent(tc.block.Width, tc.block.Height))
			block.Add(components.NewBlockComponent())
			levelSystem.Add(block)

			system := NewBallPhysicsSystem(
				components.NewPositionedComponent(100, 40),
				components.NewDimensionsComponent(1, 1),
				NewPlayingSpace(800, 600),
				levelSystem,
				NewGameState())
			system.New(&ecs.World{})

			ball := ecs.NewEntity()
			position := components.NewPositionedComponent(tc.ball.X, tc.ball.Y)
			ball.Add(position)
			speed := components.NewSpeedComponent([2]float32{tc.ball.SpeedX, tc.ball.SpeedY})
			ball.Add(speed)
			ball.Add(components.NewDimensionsComponent(10, 10))
			ball.Add(components.NewBallPhysicsComponent())
			ball.Add(components.NewCircleComponent(10))

			system.Add(ball)

			system.Update(1)
			assert.Equal(t, components.NewSpeedComponent([2]float32{tc.expected.SpeedX, tc.expected.SpeedY}), speed)
			assert.Equal(t, tc.expected.X, position.X)
			assert.Equal(t, tc.expected.Y, position.Y)
			assert.Equal(t, tc.expected.BlockHit, len(levelSystem.GetBlocks()))

		})
	}

}
