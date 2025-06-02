package gom8bit

import "fmt"

func (g *GoM8Bit) String() string {
	if !g.initialized {
		return "GoM8Bit not initialized"
	}
	result := "GoM8Bit:\n"
	for i := 0; i < int(g.x); i++ {
		for j := 0; j < int(g.y); j++ {
			result += fmt.Sprintf("%d ", g.d[i][j])
		}
		result += "\n"
	}
	return result
}
