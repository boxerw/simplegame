package model

type Vertex struct {
	Posi  Posi2D
	Pixel Pixel
}

type VertexMaps []Vertex

func (vmaps *VertexMaps) Range(fun func(posi Posi2D, pixel *Pixel)) {
	for i := 0; i < len(*vmaps); i++ {
		vertex := &(*vmaps)[i]
		fun(vertex.Posi, &vertex.Pixel)
	}
}

func (vmaps *VertexMaps) GetPixel(posi Posi2D) (*Pixel, bool) {
	for i := 0; i < len(*vmaps); i++ {
		vertex := &(*vmaps)[i]
		if vertex.Posi == posi {
			return &vertex.Pixel, true
		}
	}

	return nil, false
}

func (vmaps *VertexMaps) BlendMaps(posi Posi2D, maps Maps) {
	if maps == nil {
		return
	}

	maps.Range(func(_posi Posi2D, _pixel *Pixel) {
		posi := Posi2D{posi.GetX() + _posi.GetX(), posi.GetY() + _posi.GetY()}
		pixel, ok := vmaps.GetPixel(posi)
		if ok {
			tPixel := *_pixel
			tPixel.Overlay(pixel.Ch, pixel.Bg, pixel.Fg)
			*pixel = tPixel
		} else {
			*vmaps = append(*vmaps, Vertex{
				Posi:  posi,
				Pixel: *_pixel,
			})
		}
	})
}
