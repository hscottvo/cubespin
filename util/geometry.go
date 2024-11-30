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

func NewVec2(x float64, y float64) *Vec2 {
	ret := Vec2{X: x, Y: y}
	return &ret
}

func (v Vec2) Rotate2D(a Angle, pX float64, pY float64) *Vec2 {
	ret := v

	aboutOrigin := pX != 0 || pY != 0
	if aboutOrigin {
		ret.X -= pX
		ret.Y -= pY
	}

	ret.X = v.X*math.Cos(a.r) + v.X*math.Sin(a.r)
	ret.Y = v.Y*math.Cos(a.r) + -1*v.Y*math.Sin(a.r)

	if aboutOrigin {
		ret.X += pX
		ret.Y += pY

	}
	return &ret
}

func (v Vec2) Add2D(other Vec2) *Vec2 {
	ret := Vec2{X: v.X + other.X, Y: v.Y + other.Y}
	return &ret
}

