package core

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"sort"
)

type Pixel struct {
	FG, BG      termbox.Attribute
	Transparent bool
	Char        rune
}

func (p *Pixel) Width() int {
	w := runewidth.RuneWidth(p.Char)
	if w <= 0 {
		w = 1
	}
	return w
}

type Maps [][]Pixel

type Shape [][]rune

type _DrawItem struct {
	Layer, X, Y int
	Maps        Maps
}

type _DrawCache []_DrawItem

func (drawCache *_DrawCache) AddItem(item _DrawItem) {
	*drawCache = append(*drawCache, item)
}

func (drawCache *_DrawCache) Clear() {
	*drawCache = (*drawCache)[:0]
}

func (drawCache *_DrawCache) Sort() {
	sort.SliceStable(*drawCache, func(i, j int) bool {
		return (*drawCache)[i].Layer < (*drawCache)[j].Layer
	})
}

func (drawCache *_DrawCache) Drawing(fg, bg termbox.Attribute) {
	termbox.Clear(fg, bg)

	drawCache.Sort()
	for i := 0; i < len(*drawCache); i++ {
		item := &(*drawCache)[i]

		for i := 0; i < len(item.Maps); i++ {
			offsetX := 0

			for j := 0; j < len(item.Maps[i]); j++ {
				pixel := &item.Maps[i][j]

				if !pixel.Transparent {
					termbox.SetCell(item.X+offsetX, item.Y, pixel.Char, pixel.FG, pixel.BG)
				}

				offsetX += pixel.Width()
			}
		}
	}
	drawCache.Clear()

	termbox.Flush()
}
