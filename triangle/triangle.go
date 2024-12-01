package triangle

import (
	// "errors"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"math"
)

const (
	RENDER_THRESH float64 = 0.0001
)

type Triangle3D struct {
	A     geometry.Vec3
	B     geometry.Vec3
	C     geometry.Vec3
	Norm  geometry.Vec3
	k     float64
	pixel rune
}

func NewTriangle(a geometry.Vec3, b geometry.Vec3, c geometry.Vec3, pixel rune) Triangle3D {
	AB := b.Sub3D(a)
	AC := c.Sub3D(a)
	norm := geometry.Cross3D(AB, AC)
	k := a.X*norm.X + a.Y*norm.Y + a.Z*norm.Z
	return Triangle3D{A: a, B: b, C: c, Norm: norm, k: k, pixel: pixel}
}

func (t Triangle3D) Render(p *pane.Pane) {
	array := []geometry.Vec3{t.A, t.B, t.C}
	box := geometry.NewBoundingBox(array)

	box.MinX = int(math.Max(float64(box.MinX), 0))
	box.MinY = int(math.Max(float64(box.MinY), 0))

	box.MaxX = int(math.Min(float64(box.MaxX), float64(p.Width)))
	box.MaxY = int(math.Min(float64(box.MaxY), float64(p.Height)))

	for i := box.MinX; i < box.MaxX; i++ {
		for j := box.MinY; j < box.MaxY; j++ {
			ray := geometry.NewVec3(float64(i), float64(j), 0)
			hit := t.Hit(ray)
			if hit != nil && hit.Z > 0 && hit.Z < p.ZValues[p.Height-j-1][i] {
				p.Pixels[p.Height-j-1][i] = t.pixel
				p.ZValues[p.Height-j-1][i] = hit.Z

			}

		}

	}
}

func (t Triangle3D) Hit(P geometry.Vec3) *geometry.Vec3 {
	z := P.ProjectPlaneZ(t.Norm, t.k)
	P.Z = z
	alpha, beta, gamma := t.barycentric2D(P)
	if math.Abs(alpha+beta+gamma-1) < RENDER_THRESH {
		return &P
	}
	return nil
}

func (t *Triangle3D) Move(vec geometry.Vec3) {
	t.A = t.A.Add3D(vec)
	t.B = t.B.Add3D(vec)
	t.C = t.C.Add3D(vec)
}

// P has to be on the same plane as t
func (t Triangle3D) barycentric2D(P geometry.Vec3) (float64, float64, float64) {
	// areaABC := ma
	AB := t.B.Sub3D(t.A)
	AC := t.C.Sub3D(t.A)
	CA := t.A.Sub3D(t.C)
	BC := t.C.Sub3D(t.B)

	// PA := t.A.Sub3D(P)//
	AP := P.Sub3D(t.A)
	BP := P.Sub3D(t.B)
	CP := P.Sub3D(t.C)

	areaABC := geometry.Cross3D(AB, AC).Magnitude3D()

	areaPBC := geometry.Cross3D(BP, BC).Magnitude3D()
	areaPCA := geometry.Cross3D(CP, CA).Magnitude3D()
	areaPAB := geometry.Cross3D(AP, AB).Magnitude3D()

	alpha := areaPBC / areaABC
	beta := areaPCA / areaABC
	gamma := areaPAB / areaABC

	return alpha, beta, gamma
}
