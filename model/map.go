package model

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

func (p *Pixel) Overlay(ch rune, fg, bg termbox.Attribute) {
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

type Maps interface {
	Range(fun func(posi Posi2D, pixel *Pixel))
	GetPixel(posi Posi2D) (*Pixel, bool)
	BlendMaps(posi Posi2D, maps Maps)
}
