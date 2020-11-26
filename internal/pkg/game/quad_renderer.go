package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type QuadRenderer struct {
	quads []struct {
		id   uint64
		quad *Quad
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
		id   uint64
		quad *Quad
	}{},
		buffer:           []float32{},
		program:          NewQuadShaderProgramOrPanic(),
		projectionMatrix: projectionMatrix,
		generator:        NewTriangleIndexBufferGenerator(),
	}
}

func (qr *QuadRenderer) Add(id uint64, quad *Quad) {
	qr.quads = append(qr.quads,
		struct {
			id   uint64
			quad *Quad
		}{
			id:   id,
			quad: quad,
		})
}

func (qr *QuadRenderer) Render() {
	var buffer []float32
	for _, q := range qr.quads {
		buffer = append(buffer, q.quad.ToBuffer()...)
	}
	render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(buffer), render.NewVertexBufferLayout().
				AddLayoutFloats(2).AddLayoutFloats(4))
	qr.program.Bind()
	qr.program.SetUniformMat4f("u_MVP", qr.projectionMatrix)

	gl.DrawElements(gl.TRIANGLES, render.NewIndexBuffer(qr.generator.Generate(len(qr.quads))).GetCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))

}

func (qr *QuadRenderer) Remove(id uint64) {
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
