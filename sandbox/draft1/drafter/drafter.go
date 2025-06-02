package drafter

// drafter is simple implementation of a drawer  of lines and point for frame useful for debugging

type DrafLine struct {
	X1    int   // X1 coordinate
	Y1    int   // Y1 coordinate
	X2    int   // X2 coordinate
	Y2    int   // Y2 coordinate
	Color uint8 // Color of the line
}

type DrafPoint struct {
	X     int   // X coordinate
	Y     int   // Y coordinate
	Color uint8 // Color of the point
}
