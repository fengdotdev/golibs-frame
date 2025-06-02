package gom8bit

func New(x, y uint) GoM8Bit {

	matrix := make([][]uint8, x)
	for i := range matrix {
		matrix[i] = make([]uint8, y) // Initialize each row with a slice of uint8
	}
	return GoM8Bit{
		initialized: true,
		x:           x,
		y:           y,
		d:           matrix, // Assign the initialized matrix
	}

}
