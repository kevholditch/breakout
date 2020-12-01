package render

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Config struct {
	MajorVersion int
	MinorVersion int
	Width        int
	Height       int
	Title        string
	SwapInterval int
}

type Window struct {
	handle *glfw.Window
}

func NewWindow(cfg Config) (*Window, error) {

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, cfg.MajorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, cfg.MinorVersion)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	window, err := glfw.CreateWindow(cfg.Width, cfg.Height, cfg.Title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		fmt.Printf("%d\n", key)

	})

	glfw.SwapInterval(cfg.SwapInterval)

	return &Window{handle: window}, nil
}

func (w *Window) OnKeyPress(onPressFunc func(key int), onHoldFunc func(key int)) {
	w.handle.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			onPressFunc(int(key))
		}
		if action == glfw.Repeat {
			onHoldFunc(int(key))
		}
	})
}

func (w *Window) ShouldClose() bool {
	return w.handle.ShouldClose()
}

func (w *Window) SwapBuffers() {
	w.handle.SwapBuffers()
}

func (w *Window) PollEvents() {
	glfw.PollEvents()
}
