package game

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type QuadRenderer struct {
	quads []struct {
		id       uuid.UUID
		quad     *components.Quad
		position *components.PositionedComponent
	}
	buffer           []float32
	vertexBuffer     *render.VertexBuffer
	vertexArray      *render.VertexArray
	program          *render.Program
	projectionMatrix mgl32.Mat4
	generator        *TriangleIndexBufferGenerator
}

func NewQuadRenderer(projectionMatrix mgl32.Mat4) *QuadRenderer {
	return &QuadRenderer{quads: []struct {
		id       uuid.UUID
		quad     *components.Quad
		position *components.PositionedComponent
	}{},
		buffer:           []float32{},
		program:          components.NewQuadShaderProgramOrPanic(),
		projectionMatrix: projectionMatrix,
		generator:        NewTriangleIndexBufferGenerator(),
	}
}

func (qr *QuadRenderer) Add(id uuid.UUID, quad *components.Quad, position *components.PositionedComponent) {
	qr.quads = append(qr.quads,
		struct {
			id       uuid.UUID
			quad     *components.Quad
			position *components.PositionedComponent
		}{
			id:       id,
			quad:     quad,
			position: position,
		})
}

func (qr *QuadRenderer) ComputeBuffer(q *components.Quad, position *components.PositionedComponent) []float32 {
	return []float32{
		position.X, position.Y, q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		position.X + q.Width, position.Y, q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		position.X + q.Width, position.Y + q.Height, q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		position.X, position.Y + q.Height, q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
	}
}

func (qr *QuadRenderer) Render() {
	var buffer []float32
	for _, q := range qr.quads {
		buffer = append(buffer, qr.ComputeBuffer(q.quad, q.position)...)
	}
	render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(buffer), render.NewVertexBufferLayout().
				AddLayoutFloats(2).AddLayoutFloats(4))
	qr.program.Bind()
	qr.program.SetUniformMat4f("u_MVP", qr.projectionMatrix)

	gl.DrawElements(gl.TRIANGLES, render.NewIndexBuffer(qr.generator.Generate(len(qr.quads))).GetCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))

}

func (qr *QuadRenderer) Remove(id uuid.UUID) {
	var del = -1
	for index, e := range qr.quads {
		if id == e.id {
			del = index
			break
		}
	}
	if del >= 0 {
		qr.quads = append(qr.quads[:del], qr.quads[del+1:]...)
	}
}
