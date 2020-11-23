package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testMoveEntity struct {
	ecs.BasicEntity
	moveComponent *MoveComponent
}

func Test_PlayerMovementSystem(t *testing.T) {

	testCases := []struct {
		name             string
		testEntity       testMoveEntity
		expectedPosition mgl32.Vec4
		expectedSpeed    float32
	}{
		{
			name: "zero movement",
			testEntity: testMoveEntity{
				ecs.NewBasic(),
				NewMoveComponent(render.NewQuad(10, 20, 100, 50, 0, 0, 0, 0), 0),
			},
			expectedPosition: [4]float32{10, 20, 110, 70},
			expectedSpeed:    0,
		},
		{
			name: "negative speed moves left",
			testEntity: testMoveEntity{
				ecs.NewBasic(),
				NewMoveComponent(render.NewQuad(10, 20, 100, 50, 0, 0, 0, 0), -10),
			},
			expectedPosition: [4]float32{0, 20, 100, 70},
			expectedSpeed:    -10,
		},
		{
			name: "positive speed moves right",
			testEntity: testMoveEntity{
				ecs.NewBasic(),
				NewMoveComponent(render.NewQuad(10, 20, 100, 50, 0, 0, 0, 0), 10),
			},
			expectedPosition: [4]float32{20, 20, 120, 70},
			expectedSpeed:    10,
		},
		{
			name: "negative speed does does not go below zero",
			testEntity: testMoveEntity{
				ecs.NewBasic(),
				NewMoveComponent(render.NewQuad(0, 20, 100, 50, 0, 0, 0, 0), -10),
			},
			expectedPosition: [4]float32{0, 20, 100, 70},
			expectedSpeed:    0,
		},
		{
			name: "positive speed does does not go beyond screen",
			testEntity: testMoveEntity{
				ecs.NewBasic(),
				NewMoveComponent(render.NewQuad(695, 20, 100, 50, 0, 0, 0, 0), 10),
			},
			expectedPosition: [4]float32{700, 20, 800, 70},
			expectedSpeed:    0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			system := NewPlayerMovementSystem(800, 600)
			system.Add(testCase.testEntity.BasicEntity, testCase.testEntity.moveComponent)
			system.Update(1)

			assert.Equal(t, testCase.expectedPosition, testCase.testEntity.moveComponent.Quad.Position)
			assert.Equal(t, testCase.expectedSpeed, testCase.testEntity.moveComponent.Speed)
		})

	}

}
