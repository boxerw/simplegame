package model

import (
	"github.com/nsf/termbox-go"
	"sort"
)

type Text struct {
	Content string
	Fg, Bg  termbox.Attribute
}

type DrawItem struct {
	Layer int
	Posi  Posi2D
	Maps  Maps
	Text  *Text
}

type DrawCache []DrawItem

func (drawCache *DrawCache) AddItem(item *DrawItem) {
	*drawCache = append(*drawCache, *item)
}

func (drawCache *DrawCache) Clear() {
	*drawCache = (*drawCache)[:0]
}

func (drawCache *DrawCache) Sort() {
	sort.SliceStable(*drawCache, func(i, j int) bool {
		return (*drawCache)[i].Layer < (*drawCache)[j].Layer
	})
}

func (drawCache *DrawCache) Drawing(fg, bg termbox.Attribute) {
	if !termboxEx.IsInit() {
		return
	}

	termbox.Clear(fg, bg)

	drawCache.Sort()
	for i := 0; i < len(*drawCache); i++ {
		item := &(*drawCache)[i]

		if item.Maps != nil {
			item.Maps.Range(func(posi Posi2D, pixel *Pixel) {
				if pixel.Width() > 1 {
					return
				}

				x := item.Posi.GetX() + posi.GetX()
				y := item.Posi.GetY() + posi.GetY()
				cell := termbox.GetCell(x, y)

				tPixel := *pixel
				tPixel.Overlay(cell.Ch, cell.Fg, cell.Bg)

				termbox.SetCell(x, y, tPixel.Ch, tPixel.Fg, tPixel.Bg)
			})
		}

		if item.Text != nil {
			var offsetX, offsetY int
			for _, ch := range item.Text.Content {
				if ch == '\r' {
					continue
				}

				if ch == '\n' {
					offsetX = 0
					offsetY += 1
					continue
				}

				tPixel := &Pixel{
					Ch: ch,
					Fg: item.Text.Fg,
					Bg: item.Text.Bg,
				}

				x := item.Posi.GetX() + offsetX
				y := item.Posi.GetY() + offsetY

				termbox.SetCell(x, y, tPixel.Ch, tPixel.Fg, tPixel.Bg)

				offsetX += tPixel.Width()
			}
		}
	}
	drawCache.Clear()

	termbox.Flush()
}
