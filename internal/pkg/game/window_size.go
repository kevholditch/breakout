package game

type WindowSize struct {
	Width  float32
	Height float32
}

func NewWindowSize(width, height int) WindowSize {
	return WindowSize{
		Width:  float32(width),
		Height: float32(height),
	}
}
