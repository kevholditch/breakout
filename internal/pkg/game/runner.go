package game

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"runtime"
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

	render.UseDefaultBlending()

	indices := []int32{
		0, 1, 2,
		0, 3, 2,
		4, 5, 6,
		4, 7, 6,
	}

	buffer := []float32{
		200, 200, 0.4, 0.3, 0.2, 1.0,
		500, 200, 0.4, 0.3, 0.2, 1.0,
		500, 500, 0.4, 0.3, 0.2, 1.0,
		200, 500, 0.4, 0.3, 0.2, 1.0,
		600, 200, 0.8, 0.2, 0.2, 1.0,
		900, 200, 0.2, 0.8, 0.2, 1.0,
		900, 500, 0.8, 0.2, 0.2, 1.0,
		600, 500, 0.2, 0.8, 0.2, 1.0,
	}

	va := render.NewVertexArray()
	ib := render.NewIndexBuffer(indices)

	proj := mgl32.Ortho(0, width, 0, height, -1.0, 1.0)

	va.AddBuffer(render.NewVertexBuffer(buffer), render.NewVertexBufferLayout().AddLayoutFloats(2).AddLayoutFloats(4))

	vertex := `#version 410 core

layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;

uniform mat4 u_MVP;

out vec4 v_Color;

void main()
{
	gl_Position = u_MVP * position;
	v_Color = color;
}`
	vs, err := render.NewShaderFromString(vertex, gl.VERTEX_SHADER)
	if err != nil {
		return err
	}

	fragment := `#version 410 core

layout(location = 0) out vec4 o_Color;

in vec4 v_Color;

void main()
{
	o_Color = v_Color;
}`
	fs, err := render.NewShaderFromString(fragment, gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		return err
	}

	program.SetUniformMat4f("u_MVP", proj)

	va.UnBind()
	ib.UnBind()
	program.UnBind()

	for !w.ShouldClose() {
		render.Clear()

		program.Bind()
		program.SetUniformMat4f("u_MVP", proj)

		render.Render(va, ib, program)

		w.SwapBuffers()
		w.PollEvents()
	}

	return nil
}
