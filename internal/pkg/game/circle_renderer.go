package game

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type CircleRenderer struct {
	circles []struct {
		id     uint64
		circle *Circle
	}
	buffer           []float32
	vertexBuffer     *render.VertexBuffer
	vertexArray      *render.VertexArray
	program          *render.Program
	projectionMatrix mgl32.Mat4
}

func NewCircleRenderer(projectionMatrix mgl32.Mat4) *CircleRenderer {
	return &CircleRenderer{circles: []struct {
		id     uint64
		circle *Circle
	}{},
		buffer:           []float32{},
		program:          NewCircleShaderProgramOrPanic(),
		projectionMatrix: projectionMatrix,
	}

}

func (cr *CircleRenderer) Add(id uint64, circle *Circle) {
	cr.circles = append(cr.circles,
		struct {
			id     uint64
			circle *Circle
		}{
			id:     id,
			circle: circle,
		})

	cr.buffer = append(cr.buffer, circle.ToBuffer()...)

	cr.vertexArray = render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(cr.buffer),
			render.NewVertexBufferLayout().AddLayoutFloats(2))
}

func (cr *CircleRenderer) Render() {

	cr.vertexArray.Bind()
	cr.program.Bind()
	for _, c := range cr.circles {
		mvp := cr.projectionMatrix.Mul4(mgl32.Ident4().Mul4(mgl32.Translate3D(c.circle.Position.X(), c.circle.Position.Y(), 0)))
		cr.program.SetUniformMat4f("u_MVP", mvp)
		cr.program.SetUniformVec4("u_Colour", c.circle.Colour)

		gl.DrawArrays(gl.TRIANGLE_FAN, 0, cr.vertexArray.GetBufferCount())

	}
}

func (cr *CircleRenderer) Remove(id uint64) {
	var del = -1
	for index, e := range cr.circles {
		if id == e.id {
			del = index
			break
		}
	}
	if del >= 0 {
		cr.circles = append(cr.circles[:del], cr.circles[del+1:]...)
	}
}
