package geometry

import "math"

type Angle struct {
	r float64
}

func NewAngle(r float64) *Angle {
	newRotation := math.Mod(r, 360)
	a := Angle{r: newRotation}
	return &a
}

func (a Angle) Add(r float64) {
	a.r = math.Mod(a.r+r, 360)

}

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Rotate2D(a Angle, pX float64, pY float64) {
	aboutOrigin := pX != 0 || pY != 0
	if aboutOrigin {
		v.X -= pX
		v.Y -= pY
	}

	v.X = v.X*math.Cos(a.r) + v.X*math.Sin(a.r)
	v.Y = v.Y*math.Cos(a.r) + -1*v.Y*math.Sin(a.r)

	if aboutOrigin {
		v.X += pX
		v.Y += pY

	}
}
