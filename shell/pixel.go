package shell

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Transparent struct {
	Ch, Fg, Bg, Attr bool
}

type Blend struct {
	Attr bool
}

type Pixel struct {
	Ch          rune
	Fg, Bg      termbox.Attribute
	Transparent Transparent
	Blend       Blend
}

func (p *Pixel) Width() int {
	return runewidth.RuneWidth(p.Ch)
}

func (p *Pixel) OverlayCell(ch rune, fg, bg termbox.Attribute) {
	if p.Transparent.Ch {
		p.Ch = ch
	}

	if p.Transparent.Fg {
		p.Fg = fg & termbox.Attribute(uint64(1<<9)-1)
	}

	if p.Transparent.Bg {
		p.Bg = bg
	}

	if p.Transparent.Attr {
		p.Fg = (fg & termbox.Attribute(^(uint64(1<<9) - 1))) | (p.Fg & termbox.Attribute(uint64(1<<9)-1))
	} else {
		if p.Blend.Attr {
			p.Fg |= fg & termbox.Attribute(^(uint64(1<<9) - 1))
		}
	}
}

func (p *Pixel) BlendPixel(pixel *Pixel) {
	t := *pixel
	t.OverlayCell(p.Ch, p.Bg, p.Fg)

	p.Ch = t.Ch
	p.Bg = t.Bg
	p.Fg = t.Fg

	if !t.Transparent.Ch {
		p.Transparent.Ch = t.Transparent.Ch
	}
	if !t.Transparent.Fg {
		p.Transparent.Fg = t.Transparent.Fg
	}
	if !t.Transparent.Bg {
		p.Transparent.Bg = t.Transparent.Bg
	}
	if !t.Transparent.Attr {
		p.Transparent.Attr = t.Transparent.Attr
	}
	if !t.Blend.Attr {
		p.Blend.Attr = t.Blend.Attr
	}
}
