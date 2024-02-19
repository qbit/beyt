package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var myFont, _ = opentype.Parse(gobold.TTF)
var fontFace, _ = opentype.NewFace(myFont, &opentype.FaceOptions{
	Size:    160,
	DPI:     72,
	Hinting: font.HintingNone,
})

type myIcon struct {
	data *image.RGBA
}

func (m *myIcon) Name() string {
	return "icon.png"
}

func (m *myIcon) Content() []byte {
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, m.data)
	return buf.Bytes()
}

func buildImage(db int) *myIcon {
	i := &myIcon{}

	size := 256
	i.data = image.NewRGBA(image.Rect(0, 0, size, size))

	co := color.RGBA{A: 255}

	d := &font.Drawer{
		Dst:  i.data,
		Src:  image.NewUniform(co),
		Face: fontFace,
		Dot:  fixed.P(-3, 190),
	}

	r := []rune(fmt.Sprintf("%3s", fmt.Sprintf("%d", db)))
	l := string(r)

	d.DrawString(l)

	return i
}
