package client

type Maps interface {
	Range(fun func(posi Posi2D, pixel *Pixel))
	GetPixel(posi Posi2D) (*Pixel, bool)
	BlendMaps(posi Posi2D, maps Maps)
}
