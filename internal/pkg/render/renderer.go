package render

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func Clear(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Triangles(va *VertexArray, ib *IndexBuffer, shader *Program, projectionMatrix mgl32.Mat4) {
	va.Bind()
	ib.Bind()
	shader.Bind()
	shader.SetUniformMat4f("u_MVP", projectionMatrix)

	gl.DrawElements(gl.TRIANGLES, ib.count, gl.UNSIGNED_INT, gl.PtrOffset(0))
}
