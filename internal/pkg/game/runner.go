package game

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	}

	va := render.NewVertexArray()
	ib := render.NewIndexBuffer(indices)

	box := render.NewQuad(200, 200, 300, 300, 0.7, 0.8, 0.2, 1.0)

	proj := mgl32.Ortho(0, width, 0, height, -1.0, 1.0)

	vertexBuffer := render.NewVertexBuffer(box.ToBuffer())
	va.AddBuffer(vertexBuffer, render.NewVertexBufferLayout().AddLayoutFloats(2).AddLayoutFloats(4))

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

	previousTime := glfw.GetTime()
	frameCount := 0

	inc := float32(25.0)
	holdInc := float32(50.0)

	m := mgl32.Ident4().Mul4(mgl32.Translate3D(0, 0, 0))

	w.OnKeyPress(func(key int) {
		switch key {
		case 263:
			box.Move(-inc, 0)
		case 262:
			box.Move(inc, 0)
		case 265:
			box.Move(0, inc)
		case 264:
			box.Move(0, -inc)
		}
	}, func(key int) {
		switch key {
		case 263:
			box.Move(-holdInc, 0)
		case 262:
			box.Move(holdInc, 0)
		case 265:
			box.Move(0, holdInc)
		case 264:
			box.Move(0, -holdInc)
		}
	})

	for !w.ShouldClose() {

		currentTime := glfw.GetTime()
		frameCount++

		vertexBuffer.Update(box.ToBuffer())

		if currentTime-previousTime >= 1 {
			fmt.Printf("FPS: %d\n", frameCount)

			frameCount = 0
			previousTime = currentTime
		}

		render.Clear()

		program.Bind()

		mvp := proj.Mul4(m)
		program.SetUniformMat4f("u_MVP", mvp)

		render.Render(va, ib, program)

		w.SwapBuffers()
		w.PollEvents()
	}

	return nil
}
