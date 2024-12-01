package square

import (
	"hscottvo/cubespin/geometry"
)

type Square struct {
	posX     float32
	posY     float32
	s        int
	rotation geometry.Angle
}

func NewSquare(posX float32, posY float32, s int, rotation geometry.Angle) *Square {
	square := Square{
		posX:     posX,
		posY:     posY,
		s:        s,
		rotation: rotation,
	}
	return &square
}
