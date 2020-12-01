package game

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"github.com/liamg/aminal/glfont"
)

type FontRenderer struct {
	font      *glfont.Font
	textBoxes []struct {
		id      uint64
		textBox *TextBox
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

func (fr *FontRenderer) Add(id uint64, textBox *TextBox) {
	fr.textBoxes = append(fr.textBoxes,
		struct {
			id      uint64
			textBox *TextBox
		}{
			id:      id,
			textBox: textBox,
		})
}

func (fr *FontRenderer) Remove(id uint64) {
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
