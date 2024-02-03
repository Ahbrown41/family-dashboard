package epaper

func (p *Print) SavePNG(file string) {
	p.ctx.SavePNG(file)
}
