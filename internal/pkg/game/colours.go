package game

var (
	colourTeal      = Colour{0, 0.705, 0.705, 1}
	colourBlue      = Colour{0.125, 0.611, 0.933, 1}
	colourDarkNavy  = Colour{0.054, 0.094, 0.172, 1}
	colourCoral     = Colour{0.952, 0.439, 0.439, 1}
	colourLightNavy = Colour{0.274, 0.274, 0.352, 1}

	colourWhite = Colour{1, 1, 1, 1}
)

type Colour struct {
	R, G, B, A float32
}
