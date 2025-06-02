package types

// frame is a matrix of data for 2d display with a third dimension for color
type Frame[T any] interface {
	Matrix[T] // Embed the Matrix interface for 2D data handling
}

type Matrix[T any] interface {
	Set(x, y uint, value T) error // Set the value at position (x, y)
	Get(x, y uint) (T, error)     // Get the value at position (x, y)
	Width() uint                  // Get the width of the matrix
	Height() uint                 // Get the height of the matrix
	X() uint                      // Get the X dimension of the matrix
	Y() uint                      // Get the Y dimension of the matrix
}
