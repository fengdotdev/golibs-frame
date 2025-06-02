package gom8bit

import "fmt"

func (g *GoM8Bit) IsInitialized() bool {
	// Return the initialization status of the matrix
	return g.initialized
}

func (g *GoM8Bit) Set(x, y uint, v uint8) error {

	// Check if the matrix is initialized
	if !g.initialized {
		return fmt.Errorf("GoM8Bit not initialized")
	}

	// Validate coordinates

	// asumes that the coordinates are valid and within bounds with a matrix of size x,y initialized
	g.d[x][y] = v
	return nil
}

func (g *GoM8Bit) Get(x, y uint) (uint8, error) {
	// Check if the matrix is initialized
	if !g.initialized {
		return 0, fmt.Errorf("GoM8Bit not initialized")
	}

	// Validate coordinates
	if !g.ValidCord(x, y) {
		return 0, fmt.Errorf("invalid coordinates")
	}

	// asumes that the coordinates are valid and within bounds with a matrix of size x,y initialized
	return g.d[x][y], nil
}

// ValidCord checks if the given coordinates are valid within the matrix bounds
// dont asume the matrix is initialized
func (g *GoM8Bit) ValidCord(x, y uint) bool {
	// Validate coordinates
	if x >= g.x || y >= g.y {
		return false
	}
	// Coordinates are valid
	return true
}

func (g *GoM8Bit) Size() (uint, uint) {
	// Check if the matrix is initialized
	if !g.initialized {
		return 0, 0
	}
	// Return the size of the matrix
	return g.x, g.y
}

func (g *GoM8Bit) SizeInt() (int, int) {
	// Check if the matrix is initialized
	if !g.initialized {
		return 0, 0
	}
	// Return the size of the matrix as integers
	return int(g.x), int(g.y)
}
