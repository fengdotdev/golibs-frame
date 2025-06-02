package gom8bit

var _ Binaer = (*GoM8Bit)(nil) // Ensure GoM8Bit implements Binaer interface


// FromBinary implements trait.Binarer.
func (g *GoM8Bit) FromBinary([]byte) error {
	panic("unimplemented")
}

// ToBinary implements trait.Binarer.
func (g *GoM8Bit) ToBinary() ([]byte, error) {
	panic("unimplemented")
}

func (g *GoM8Bit) Strand() string {
	// Return the string representation of the GoM8Bit matrix
	return "GoM8Bit"
}

func (g *GoM8Bit) Data() []byte {
	// Return the raw data of the GoM8Bit matrix
	if !g.initialized {
		return nil
	}
	return Slice2DToData(g.d)
}

func Slice2DToData(s [][]uint8) []byte {
	// Convert a 2D slice of uint8 to a 1D byte slice
	if len(s) == 0 || len(s[0]) == 0 {
		return nil
	}
	data := make([]byte, len(s)*len(s[0]))
	for i := range s {
		for j := range s[i] {
			data[i*len(s[0])+j] = s[i][j]
		}
	}
	return data
}

func DataToSlice2D(data []byte, x, y uint) [][]uint8 {
	// Convert a 1D byte slice to a 2D slice of uint8
	if len(data) == 0 || x == 0 || y == 0 {
		return nil
	}
	matrix := make([][]uint8, x)
	for i := range matrix {
		matrix[i] = make([]uint8, y) // Initialize each row with a slice of uint8
	}

	for i := uint(0); i < x; i++ {
		for j := uint(0); j < y; j++ {
			index := i*y + j
			if index < uint(len(data)) {
				matrix[i][j] = data[index]
			} else {
				matrix[i][j] = 0 // Default value if data is insufficient
			}
		}
	}
	return matrix
}



