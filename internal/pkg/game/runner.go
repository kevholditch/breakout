package game

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"github.com/kevholditch/breakout/internal/pkg/systems"
	"runtime"
	"time"
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

	box := render.NewQuad(200, 50, 300, 50, 0.7, 0.8, 0.2, 1.0)

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

	speed := float32(0)
	inc := float32(1)
	smallInc := float32(0.5)

	m := mgl32.Ident4().Mul4(mgl32.Translate3D(0, 0, 0))

	world := ecs.World{}
	world.AddSystem(systems.NewFrameRateSystem())

	w.OnKeyPress(func(key int) {
		switch key {
		case 263:
			if speed > 0 {
				speed = 0
			} else {
				speed -= inc
			}
		case 262:
			if speed < 0 {
				speed = 0
			} else {
				speed += inc
			}
			//case 265:
			//	box.Move(0, inc)
			//case 264:
			//	box.Move(0, -inc)
		}
	}, func(key int) {
		switch key {
		case 263:
			if speed > 0 {
				speed = 0
			} else {
				speed -= smallInc
			}
		case 262:
			if speed < 0 {
				speed = 0
			} else {
				speed += smallInc
			}
			//case 265:
			//	box.Move(0, holdInc)
			//case 264:
			//	box.Move(0, -holdInc)
		}
	})

	last := time.Now()

	for !w.ShouldClose() {

		// 100 ms elasped
		// 0.1

		elapsed := time.Now().Sub(last)
		last = time.Now()
		world.Update(float32(elapsed.Milliseconds()) / float32(1000))

		currentTime := glfw.GetTime()
		frameCount++

		vertexBuffer.Update(box.Move(float32(elapsed.Milliseconds())*speed, 0).ToBuffer())

		if currentTime-previousTime >= 1 {

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
