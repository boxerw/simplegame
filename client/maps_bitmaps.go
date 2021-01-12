package client

import "encoding/json"

type BitMaps struct {
	Matrix []Pixel
	Width  int
	Origin Posi2D
}

func (bitMaps *BitMaps) Range(fun func(posi Posi2D, pixel *Pixel)) {
	if fun == nil || bitMaps.Width <= 0 {
		return
	}

	for i := 0; i < len(bitMaps.Matrix); i++ {
		pixel := &bitMaps.Matrix[i]
		fun(Posi2D{i%bitMaps.Width - bitMaps.Origin.GetX(), i/bitMaps.Width - bitMaps.Origin.GetY()}, pixel)
	}
}

func (bitMaps *BitMaps) GetPixel(posi Posi2D) (*Pixel, bool) {
	posi.SetX(posi.GetX() + bitMaps.Origin.GetX())
	posi.SetY(posi.GetY() + bitMaps.Origin.GetY())

	if posi.GetX() < 0 || posi.GetX() >= bitMaps.Width {
		return nil, false
	}

	if posi.GetY() < 0 || posi.GetY() >= len(bitMaps.Matrix)/bitMaps.Width {
		return nil, false
	}

	return &bitMaps.Matrix[posi.GetX()+posi.GetY()*bitMaps.Width], true
}

func (bitMaps *BitMaps) BlendMaps(posi Posi2D, maps Maps) {
	if maps == nil {
		return
	}

	maps.Range(func(_posi Posi2D, _pixel *Pixel) {
		blendPos := Posi2D{posi.GetX() + _posi.GetX(), posi.GetY() + _posi.GetY()}

		blendPixel, ok := bitMaps.GetPixel(blendPos)
		if !ok {
			bitMaps.extendMap(blendPos)
			blendPixel, ok = bitMaps.GetPixel(blendPos)
			if !ok {
				panic("no pixel")
			}
		}

		blendPixel.BlendPixel(_pixel)
	})
}

func (bitMaps *BitMaps) Marshal() ([]byte, error) {
	return json.Marshal(*bitMaps)
}

func (bitMaps *BitMaps) Unmarshal(data []byte) error {
	var t BitMaps
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*bitMaps = t
	return nil
}

func (bitMaps *BitMaps) extendMap(posi Posi2D) {
	posi.SetX(posi.GetX() + bitMaps.Origin.GetX())
	posi.SetY(posi.GetY() + bitMaps.Origin.GetY())

	h := len(bitMaps.Matrix) / bitMaps.Width

	if posi.GetY() < 0 {
		delta := -posi.GetY()
		bitMaps.Matrix = append(make([]Pixel, bitMaps.Width*delta), bitMaps.Matrix...)
		bitMaps.Origin.SetY(bitMaps.Origin.GetY() + delta)
		h = len(bitMaps.Matrix) / bitMaps.Width
	} else if posi.GetY() >= h {
		delta := posi.GetY() - h + 1
		bitMaps.Matrix = append(bitMaps.Matrix, make([]Pixel, bitMaps.Width*delta)...)
		h = len(bitMaps.Matrix) / bitMaps.Width
	}

	if posi.GetX() < 0 {
		delta := -posi.GetX()
		t := make([]Pixel, h*(bitMaps.Width+delta))
		for i := 0; i < len(bitMaps.Matrix); i++ {
			t[i+(i/bitMaps.Width+1)*delta] = bitMaps.Matrix[i]
		}
		bitMaps.Matrix = t
		bitMaps.Width += delta
		bitMaps.Origin.SetX(bitMaps.Origin.GetX() + delta)
	} else if posi.GetX() >= bitMaps.Width {
		delta := posi.GetX() - bitMaps.Width + 1
		t := make([]Pixel, h*(bitMaps.Width+delta))
		for i := 0; i < len(bitMaps.Matrix); i++ {
			t[i+i/bitMaps.Width*delta] = bitMaps.Matrix[i]
		}
		bitMaps.Matrix = t
		bitMaps.Width += delta
	}
}
