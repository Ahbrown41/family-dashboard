package epaper

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"strings"
)

// DrawText Draws text on screens
func (p *Print) DrawText(x float64, y float64, size float64, str string) (float64, float64) {
	p.ctx.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: size,
	})
	lines := strings.Split(str, "\n")
	p.ctx.SetFontFace(face)
	w, h := p.ctx.MeasureMultilineString(str, 0)
	for i, line := range lines {
		p.ctx.DrawStringAnchored(line, x, y+(float64(i)*h), 0.0, 0.0)
	}
	if len(lines) > 1 {
		return w, h * float64(len(lines))
	}
	return w, h
}

// DrawRectangle Draws a rectangle
func (p *Print) DrawRectangle(x float64, y float64, w float64, h float64) {
	p.ctx.DrawRectangle(x, y, w, h)
	p.ctx.Stroke()
}
