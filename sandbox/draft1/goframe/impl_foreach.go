package goframe

import "fmt"

type ForEach interface {
	ForEach(func(x uint, y uint, dot RGBADot8) error) error
}

var _ ForEach = (*Goframe8)(nil)

func (g *Goframe8) ForEach(fn func(x uint, y uint, dot RGBADot8) error) error {
	if !g.initialized {
		return fmt.Errorf("frame not initialized")
	}

	for y := uint(0); y < g.y; y++ {
		for x := uint(0); x < g.x; x++ {
			dot, err := g.Get(x, y)
			if err != nil {
				return fmt.Errorf("error getting dot at (%d, %d): %w", x, y, err)
			}
			if err := fn(x, y, dot); err != nil {
				return fmt.Errorf("error processing dot at (%d, %d): %w", x, y, err)
			}
		}
	}
	return nil
}
