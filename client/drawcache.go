package client

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

type DrawCache interface {
	AddItem(item *DrawItem)
	Clear()
	Sort()
	Size() int
	Drawing(fg, bg termbox.Attribute)
}

func NewDrawCache() DrawCache {
	return &_DrawCache{}
}

type _DrawCache struct {
	items []DrawItem
}

func (drawCache *_DrawCache) AddItem(item *DrawItem) {
	drawCache.items = append(drawCache.items, *item)
}

func (drawCache *_DrawCache) Clear() {
	drawCache.items = nil
}

func (drawCache *_DrawCache) Sort() {
	sort.SliceStable(drawCache.items, func(i, j int) bool {
		return drawCache.items[i].Layer < drawCache.items[j].Layer
	})
}

func (drawCache *_DrawCache) Size() int {
	return len(drawCache.items)
}

func (drawCache *_DrawCache) Drawing(fg, bg termbox.Attribute) {
	if !termboxEx.IsInit() {
		return
	}

	termbox.Clear(fg, bg)

	drawCache.Sort()
	for i := 0; i < len(drawCache.items); i++ {
		item := &drawCache.items[i]

		if item.Maps != nil {
			item.Maps.Range(func(posi Posi2D, pixel *Pixel) {
				if pixel.Width() > 1 {
					return
				}

				x := item.Posi.GetX() + posi.GetX()
				y := item.Posi.GetY() + posi.GetY()

				w, h := termbox.Size()
				if x < 0 || y < 0 || x >= w || y >= h {
					return
				}

				cell := termbox.GetCell(x, y)

				t := *pixel
				t.OverlayCell(cell.Ch, cell.Fg, cell.Bg)

				termbox.SetCell(x, y, t.Ch, t.Fg, t.Bg)
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

				t := &Pixel{
					Ch: ch,
					Fg: item.Text.Fg,
					Bg: item.Text.Bg,
				}

				x := item.Posi.GetX() + offsetX
				y := item.Posi.GetY() + offsetY

				w, h := termbox.Size()
				if x < 0 || y < 0 || x >= w || y >= h {
					continue
				}

				termbox.SetCell(x, y, t.Ch, t.Fg, t.Bg)

				offsetX += t.Width()
			}
		}
	}
	drawCache.Clear()

	termbox.Flush()
}
