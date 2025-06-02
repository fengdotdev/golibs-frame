package goframe

// goframe implements frame interface

type GoframeRGBA struct {
	x uint    // X dimension
	y uint    // Y dimension
	r GoM8Bit // Red channel
	g GoM8Bit // Green channel
	b GoM8Bit // Blue channel
	a GoM8Bit // Alpha channel
}

func NewGoframeRGBA(x, y uint) GoframeRGBA {
	return GoframeRGBA{
		x: x,
		y: y,
		r: NewGoM8Bit(x, y),
		g: NewGoM8Bit(x, y),
		b: NewGoM8Bit(x, y),
		a: NewGoM8Bit(x, y),
	}
}


func (g *GoframeRGBA) X() uint {
	return g.x
}
func (g *GoframeRGBA) Y() uint {
	return g.y
}