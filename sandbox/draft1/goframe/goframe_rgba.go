package goframe

// goframe implements frame interface

type Goframe8 struct {
	initialized bool         // Flag to check if the frame is initialized
	x           uint         // X dimension
	y           uint         // Y dimension
	dots        [][]RGBADot8 // 2D slice to hold RGBADot8 values

}
