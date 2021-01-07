package shell

type PixelMaps [][]Pixel

func (pmaps *PixelMaps) Range(fun func(posi Posi2D, pixel *Pixel)) {
	for i := 0; i < len(*pmaps); i++ {
		for j := 0; j < len((*pmaps)[i]); j++ {
			pixel := &(*pmaps)[i][j]
			fun(Posi2D{j, i}, pixel)
		}
	}
}

func (pmaps *PixelMaps) GetPixel(posi Posi2D) (*Pixel, bool) {
	if posi.GetY() < 0 || posi.GetY() >= len(*pmaps) {
		return nil, false
	}

	if posi.GetX() < 0 || posi.GetX() >= len((*pmaps)[posi.GetY()]) {
		return nil, false
	}

	return &(*pmaps)[posi.GetY()][posi.GetX()], true
}

func (pmaps *PixelMaps) BlendMaps(posi Posi2D, maps Maps) {

}
