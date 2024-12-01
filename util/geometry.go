package geometry

import (
	"fmt"
	"math"
)

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

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x float64, y float64, z float64) Vec3 {
	ret := Vec3{X: x, Y: y, Z: z}
	return ret
}

func (v Vec3) Rotate3D(a Angle, pX float64, pY float64) Vec3 {
	ret := v
	fmt.Print("Not implemented: Rotate3D")
	return ret
}

func (v Vec3) Add2D(other Vec3) Vec3 {
	return Vec3{X: v.X + other.X, Y: v.Y + other.Y}
}

func Dist3D(a Vec3, b Vec3) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z))
}

func Cross3D(a Vec3, b Vec3) Vec3 {
	i := a.Y*b.Z - b.Y*a.Z
	j := b.X*a.Z - a.X*b.Z
	k := a.X*b.Y - b.X*a.Y
	return NewVec3(i, j, k)
}
