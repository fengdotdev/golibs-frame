package rgbadot

// RGBADot8Bit represents a pixel with RGBA components, each stored as an 8-bit unsigned integer.
type RGBADot8Bit struct {
	R uint8 // Red component
	G uint8 // Green component
	B uint8 // Blue component
	A uint8 // Alpha component
}

func BlackDot() RGBADot8Bit {
	// Returns a black dot with full transparency
	return RGBADot8Bit{R: 0, G: 0, B: 0, A: 0}
}

func WhiteDot() RGBADot8Bit {
	// Returns a white dot with full opacity
	return RGBADot8Bit{R: 255, G: 255, B: 255, A: 255}
}
