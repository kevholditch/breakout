package game

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"github.com/google/uuid"
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
	"github.com/liamg/aminal/glfont"
)

type FontRenderer struct {
	font      *glfont.Font
	textBoxes []struct {
		id      uuid.UUID
		textBox *primitives.TextBox
	}
}

func NewFontRenderer(size WindowSize) *FontRenderer {
	box := packr.NewBox("./assets")
	fontBytes, err := box.Find("ARCADE_R.TTF")
	if err != nil {
		panic(err)
	}

	font, err := glfont.LoadFont(bytes.NewReader(fontBytes), 28, int(size.Width), int(size.Height))
	if err != nil {
		panic(err)
	}
	return &FontRenderer{font: font}
}

func (fr *FontRenderer) Render() {
	for _, t := range fr.textBoxes {
		fr.font.SetColor(t.textBox.Colour.R, t.textBox.Colour.G, t.textBox.Colour.B, t.textBox.Colour.A)
		fr.font.Print(t.textBox.Position.X(), t.textBox.Position.Y(), t.textBox.Text)
	}
}

func (fr *FontRenderer) Add(id uuid.UUID, textBox *primitives.TextBox) {
	fr.textBoxes = append(fr.textBoxes,
		struct {
			id      uuid.UUID
			textBox *primitives.TextBox
		}{
			id:      id,
			textBox: textBox,
		})
}

func (fr *FontRenderer) Remove(id uuid.UUID) {
	var del = -1
	for index, e := range fr.textBoxes {
		if id == e.id {
			del = index
			break
		}
	}
	if del >= 0 {
		fr.textBoxes = append(fr.textBoxes[:del], fr.textBoxes[del+1:]...)
	}
}
