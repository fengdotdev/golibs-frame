package goframe

import "github.com/fengdotdev/golibs-frame/sandbox/draft1/gom8bit"

type GoM8Bit = gom8bit.GoM8Bit // Alias for GoM8Bit type

func NewGoM8Bit(x, y uint) GoM8Bit {
	return gom8bit.New(x, y)
}
