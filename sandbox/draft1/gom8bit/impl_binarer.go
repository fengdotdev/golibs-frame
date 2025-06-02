package gom8bit

var _ Binaer = (*GoM8Bit)(nil) // Ensure GoM8Bit implements Binaer interface

const HeaderGom8bit = "gom8bit"

// <header, version, length, sha 256, data>

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
	data := make([]byte, g.x*g.y)
	for i := uint(0); i < g.x; i++ {
		for j := uint(0); j < g.y; j++ {
			data[i*g.y+j] = g.d[i][j]
		}
	}
	return data
}

