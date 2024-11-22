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
