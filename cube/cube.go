package Cube

type Cube struct {
	center int
}

// center x, y, z with 6 faces
// - side length s
// each face is sxs
// located (relative) at
// - x+1/2s, y, z
// - x-1/2s, y, z
// - x, y+1/2s, z
// - x, y-1/2s, z
// - x, y, z+1/2s
// - x, y, z-1/2s
func NewCube() *Cube {
	c := Cube{center: 5}
	return &c
}

func (c Cube) Display() string {

	return `▓`
}
