package drafter

import "github.com/fengdotdev/golibs-frame/sandbox/draft1/goframe"

// drafter is simple implementation of a drawer  of lines and point for frame useful for debugging

type Drafter interface {
	Paint(f goframe.Goframe8) error // Paint draws the frame using the drafter)
}

var _ Drafter = (*DrafLine[any])(nil)

type DrafLine[T any] struct {
	X1    int // X1 coordinate
	Y1    int // Y1 coordinate
	X2    int // X2 coordinate
	Y2    int // Y2 coordinate
	Color T   // Color of the line
}

// Paint implements Drafter.
func (d *DrafLine[T]) Paint(f goframe.Goframe8) error {
	return nil
}

type DrafPoint[T any] struct {
	X     int // X coordinate
	Y     int // Y coordinate
	Color T   // Color of the point

}



func NewDrafLine[T any](x1, y1, x2, y2 int, color T) DrafLine[T] {
	return DrafLine[T]{X1: x1, Y1: y1, X2: x2, Y2: y2, Color: color}
}
