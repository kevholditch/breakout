package game

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"github.com/liamg/ecs"
)

type circleEntity struct {
	base       *ecs.Entity
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
	colour     *components.ColouredComponent
	circle     *components.CircleComponent
}

type CircleRenderSystem struct {
	circles          []circleEntity
	buffer           []float32
	vertexBuffer     *render.VertexBuffer
	vertexArray      *render.VertexArray
	program          *render.Program
	projectionMatrix mgl32.Mat4
}

func NewCircleRenderSystem(windowSize WindowSize) *CircleRenderSystem {
	return &CircleRenderSystem{circles: []circleEntity{},
		buffer:           []float32{},
		program:          primitives.NewCircleShaderProgramOrPanic(),
		projectionMatrix: mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0),
	}

}

func (cr *CircleRenderSystem) Update(float32) {
	cr.vertexArray.Bind()
	cr.program.Bind()
	for _, c := range cr.circles {
		mvp := cr.projectionMatrix.Mul4(mgl32.Ident4().Mul4(mgl32.Translate3D(c.position.X, c.position.Y, 0)))
		cr.program.SetUniformMat4f("u_MVP", mvp)
		cr.program.SetUniformVec4("u_Colour", c.colour.Colour.ToVec4())

		gl.DrawArrays(gl.TRIANGLE_FAN, 0, cr.vertexArray.GetBufferCount())

	}
}

func (cr *CircleRenderSystem) Add(entity *ecs.Entity) {

	circle := entity.Component(components.IsCircle).(*components.CircleComponent)
	cr.circles = append(cr.circles, circleEntity{
		base:       entity,
		position:   entity.Component(components.IsPositioned).(*components.PositionedComponent),
		dimensions: entity.Component(components.HasDimensions).(*components.DimensionComponent),
		colour:     entity.Component(components.IsColoured).(*components.ColouredComponent),
		circle:     circle,
	})

	cr.buffer = append(cr.buffer, circle.Buffer...)

	cr.vertexArray = render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(cr.buffer),
			render.NewVertexBufferLayout().AddLayoutFloats(2))
}

func (cr *CircleRenderSystem) Remove(entity *ecs.Entity) {
	var del = -1
	for index, e := range cr.circles {
		if e.base.ID() == entity.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		cr.circles = append(cr.circles[:del], cr.circles[del+1:]...)
	}
}

func (cr *CircleRenderSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.IsColoured,
		components.IsPositioned,
		components.HasDimensions,
		components.IsCircle,
	}
}
