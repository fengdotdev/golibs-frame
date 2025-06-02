package goframe

func New(x, y uint) Goframe8 {

	matrix := make([][]RGBADot8, y)
	for i := range matrix {
		matrix[i] = make([]RGBADot8, x) // Initialize each row with a slice of RGBADot8
	}
	return Goframe8{
		initialized: true,
		x:           x,
		y:           y,
		dots:        matrix, // Assign the initialized matrix
	}
}

func NewWithBgColor(x, y uint, bgColor RGBADot8) Goframe8 {
	g := New(x, y)

	g.ForEach(func(x uint, y uint, dot RGBADot8) error {
		return g.Set(x, y, bgColor)
	})

	return g
}

func NewWH(width, height uint) Goframe8 {
	return New(width, height)
}
