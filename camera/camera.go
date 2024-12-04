package camera

import (
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
)

type Camera struct {
	position  geometry.Vec3
	direction geometry.Vec3
	pane      pane.Pane
}
