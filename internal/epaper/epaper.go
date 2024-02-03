package epaper

import (
	"family-dashboard/internal/epd7in5"
	"github.com/fogleman/gg"
	"log"
)

type Print struct {
	epd *epd7in5.Epd
	ctx *gg.Context
}

func New() *Print {
	epd, _ := epd7in5.New("P1_22", "P1_24", "P1_11", "P1_18")
	cx := gg.NewContext(epd7in5.EPD_WIDTH, epd7in5.EPD_HEIGHT)
	return &Print{
		epd: epd,
		ctx: cx,
	}
}

func (p *Print) Init() {
	p.epd.Init()
	p.epd.Clear()
}

func (p *Print) Clear() {
	p.ctx.SetRGB(1, 1, 1)
	p.ctx.Clear()
	p.ctx.SetRGB(0, 0, 0)
}

// Render Renders drawing
func (p *Print) Render() {
	log.Println("Rendering...")
	p.epd.Display(p.epd.Convert(p.ctx.Image()))
}
