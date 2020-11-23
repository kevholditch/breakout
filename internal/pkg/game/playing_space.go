package game

type PlayingSpace struct {
	Width  float32
	Height float32
}

func NewPlayingSpace(width, height int) PlayingSpace {
	return PlayingSpace{
		Width:  float32(width),
		Height: float32(height),
	}
}
