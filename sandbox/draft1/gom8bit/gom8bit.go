package gom8bit

// GoM8Bit implements a 2D matrix of 8-bit unsigned integers
type GoM8Bit struct {
	initialized bool // Indicates if the matrix is initialized
	x           uint //X
	y           uint
	d           [][]uint8 // x,y
}
