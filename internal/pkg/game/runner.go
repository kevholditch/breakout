package game

import (
	"fmt"
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
		panic(err)
	}

	render.UseDefaultBlending()

	for !w.ShouldClose() {
		render.Clear()

		w.SwapBuffers()
		w.PollEvents()
	}

	return nil
}
