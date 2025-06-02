package gom8bit

import "github.com/fengdotdev/golibs-frame/sandbox/draft1/encoder"

// HeaderGom8bit is the fixed header string
const HeaderGom8bit = "gom8bit"
const VersionGom8bit = 1

func NewGom8bitBTO() *Gom8bitBTO {
	return &Gom8bitBTO{
		Header:  HeaderGom8bit,
		Version: VersionGom8bit,
		Data:    nil, // Initialize with nil or empty slice
	}
}

// Gom8bitBTO represents the binary structure
type Gom8bitBTO struct {
	Header  string // Fixed "gom8bit"
	Version uint8  // 1-byte version
	Data    []byte // Variable-length data
}

// Encode serializes the structure into a []byte
func (g *Gom8bitBTO) Encode() ([]byte, error) {
	e, err := encoder.NewEncodeObj(g.Header, g.Version, g.Data)

	if err != nil {
		return nil, err
	}
	data, err := e.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Decode deserializes the []byte into the structure
func (g *Gom8bitBTO) Decode(data []byte) error {
	e := &encoder.EncodeObj{
		Header:  g.Header,
		Version: g.Version,
		Data:    data,
	}

	if err := e.Decode(data); err != nil {
		return err
	}

	g.Data = e.Data
	return nil
}
