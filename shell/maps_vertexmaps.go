package shell

import "encoding/json"

type Vertex struct {
	Posi  Posi2D
	Pixel Pixel
}

type VertexMaps struct {
	List []Vertex
}

func (vertexMaps *VertexMaps) Range(fun func(posi Posi2D, pixel *Pixel)) {
	if fun == nil {
		return
	}

	for i := 0; i < len(vertexMaps.List); i++ {
		vertex := &vertexMaps.List[i]
		fun(vertex.Posi, &vertex.Pixel)
	}
}

func (vertexMaps *VertexMaps) GetPixel(posi Posi2D) (*Pixel, bool) {
	for i := 0; i < len(vertexMaps.List); i++ {
		vertex := &vertexMaps.List[i]
		if vertex.Posi == posi {
			return &vertex.Pixel, true
		}
	}

	return nil, false
}

func (vertexMaps *VertexMaps) BlendMaps(posi Posi2D, maps Maps) {
	if maps == nil {
		return
	}

	maps.Range(func(_posi Posi2D, _pixel *Pixel) {
		blendPos := Posi2D{posi.GetX() + _posi.GetX(), posi.GetY() + _posi.GetY()}

		blendPixel, ok := vertexMaps.GetPixel(blendPos)
		if !ok {
			vertexMaps.List = append(vertexMaps.List, Vertex{
				Posi:  blendPos,
				Pixel: *_pixel,
			})
			return
		}

		blendPixel.BlendPixel(_pixel)
	})
}

func (vertexMaps *VertexMaps) Marshal() ([]byte, error) {
	return json.Marshal(*vertexMaps)
}

func (vertexMaps *VertexMaps) Unmarshal(data []byte) error {
	var t VertexMaps
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*vertexMaps = t
	return nil
}
