package goframe

import (
	"fmt"

	"github.com/fengdotdev/golibs-frame/sandbox/draft1/dots/rgbadot"
	"github.com/fengdotdev/golibs-frame/sandbox/draft1/types"
)

var _ types.Matrix[RGBADot8] = (*Goframe8)(nil)

func (g *Goframe8) Set(x uint, y uint, value RGBADot8) error {
	if !g.initialized {
		return fmt.Errorf("frame not initialized")
	}

	if x >= g.x || y >= g.y {
		return fmt.Errorf("coordinates (%d, %d) out of bounds for frame size (%d, %d)", x, y, g.x, g.y)
	}
	g.dots[y][x] = value
	return nil
}

func (g *Goframe8) Get(x uint, y uint) (RGBADot8, error) {

	if !g.initialized {
		return rgbadot.WhiteDot(), fmt.Errorf("frame not initialized")
	}

	if x >= g.x || y >= g.y {
		return rgbadot.WhiteDot(), fmt.Errorf("coordinates (%d, %d) out of bounds for frame size (%d, %d)", x, y, g.x, g.y)
	}
	return g.dots[y][x], nil
}

func (g *Goframe8) Width() uint {
	if !g.initialized {
		return 0
	}
	return g.x
}

func (g *Goframe8) Height() uint {
	if !g.initialized {
		return 0
	}
	return g.y
}

func (g *Goframe8) X() uint {
	return g.Width()
}
func (g *Goframe8) Y() uint {
	return g.Height()
}
