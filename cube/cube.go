package cube

import (
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/triangle"
)

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

//	   e ----- f
//	  /|      /|
//	 / |     / |
//	/  |    /  |
//
// a --+-- b   |
// |   |   |   |
// |   g --+--h
// |  /    | /
// | /     |/
// c ----- d
type Cube struct {
	center geometry.Vec3
	faces  []*square
}

func NewCube(center geometry.Vec3, sideLength float64) *Cube {
	off := sideLength / 2
	x := center.X
	y := center.Y
	z := center.Z
	a := geometry.NewVec3(x-off, y+off, z-off)
	b := geometry.NewVec3(x+off, y+off, z-off)
	c := geometry.NewVec3(x-off, y-off, z-off)
	d := geometry.NewVec3(x+off, y-off, z-off)
	e := geometry.NewVec3(x-off, y+off, z+off)
	f := geometry.NewVec3(x+off, y+off, z+off)
	g := geometry.NewVec3(x-off, y-off, z+off)
	h := geometry.NewVec3(x+off, y-off, z+off)

	abcd := NewSquareFromPoints(a, b, c, d, '▓')
	aecg := NewSquareFromPoints(a, e, c, g, '▒')
	bfdh := NewSquareFromPoints(b, f, d, h, '░')
	efab := NewSquareFromPoints(e, f, a, b, 'O')
	efgh := NewSquareFromPoints(e, f, g, h, '█')
	ghcd := NewSquareFromPoints(g, h, c, d, '.')

	var faces []*square

	faces = append(faces, abcd)
	faces = append(faces, aecg)
	faces = append(faces, bfdh)
	faces = append(faces, efab)
	faces = append(faces, efgh)
	faces = append(faces, ghcd)

	cube := Cube{center: center, faces: faces}
	return &cube
}

func (c Cube) Center() geometry.Vec3 {
	return c.center
}

func (c Cube) Render(p *pane.Pane) {
	for _, i := range c.faces {
		i.Render(p)
	}
}

func (c *Cube) Rotate(about geometry.Vec3, xRotation geometry.Angle, yRotation geometry.Angle, zRotation geometry.Angle) {
	for _, i := range c.faces {
		i.Rotate(about, xRotation, yRotation, zRotation)
	}

}

func (c *Cube) Move(vec geometry.Vec3) {
	for _, i := range c.faces {
		i.Move(vec)
	}
	c.center = c.center.Add3D(vec)
}

type square struct {
	a *triangle.Triangle3D
	b *triangle.Triangle3D
}

// A ----- B
// |       |
// |       |
// |       |
// C ----- D
func NewSquareFromPoints(A geometry.Vec3, B geometry.Vec3, C geometry.Vec3, D geometry.Vec3, pixel rune) *square {
	AB := B.Sub3D(A)
	AC := C.Sub3D(A)
	AD := D.Sub3D(A)

	norm := geometry.Cross3D(AB, AC)

	if geometry.Dot3D(AD, norm) != 0 {
		panic("Points must be on the same plane")
	}

	a := triangle.NewTriangle(A, C, B, pixel)
	b := triangle.NewTriangle(B, C, D, pixel)
	square := square{a: a, b: b}
	return &square
}

func (s square) Render(p *pane.Pane) {
	s.a.Render(p)
	s.b.Render(p)

}

func (s *square) Rotate(about geometry.Vec3, xRotation geometry.Angle, yRotation geometry.Angle, zRotation geometry.Angle) {
	s.a.Rotate(about, xRotation, yRotation, zRotation)
	s.b.Rotate(about, xRotation, yRotation, zRotation)
}

func (s *square) Move(vec geometry.Vec3) {
	s.a.Move(vec)
	s.b.Move(vec)
}
