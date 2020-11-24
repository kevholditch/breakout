package game

var (
	colourLimeGreen = Colour{0, 0.705, 0.705, 1}
	colourDarkBlue  = Colour{0.054, 0.094, 0.172, 1}
	colourWhite     = Colour{1, 1, 1, 1}
)

type Colour struct {
	R, G, B, A float32
}
