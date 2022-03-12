package utils

type NormalizedColor struct {
	R, G, B, A float64
}

func (n NormalizedColor) RGBA() (r, g, b, a uint32) {
	r = uint32(n.R*255.0) << 8
	g = uint32(n.G*255.0) << 8
	b = uint32(n.B*255.0) << 8
	a = uint32(n.A*255.0) << 8
	return
}
