package camera

import (
	"fmt"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
)

type Perspective int

const (
	UI Perspective = iota
	Raytrace
)

type Camera struct {
	position     geometry.Vec3
	direction    geometry.Vec3
	perspective  int
	pane         pane.Pane
	height       float64
	width        float64
	paneDistance float64
}

func (c *Camera) Ray(endpoint geometry.Vec3, persp Perspective) geometry.Vec3 {
	switch persp {
	case UI:
		endpoint.Z = 0
		return endpoint
	case Raytrace:
		return endpoint.Sub3D(c.position).Normalize()

	default:
		panic(fmt.Errorf("unknown state: %d", persp))
	}
}

func (c *Camera) Position() geometry.Vec3 {
	return c.position
}

func (c *Camera) Direction() geometry.Vec3 {
	return c.direction
}

func (c *Camera) Move(vec geometry.Vec3) {
	c.position = c.position.Add3D(vec)
}

func (c *Camera) Rotate(about geometry.Vec3, xRotation geometry.Angle, yRotation geometry.Angle, zRotation geometry.Angle) {
	c.position = c.position.Rotate3D(about, xRotation, yRotation, zRotation)

	c.direction = c.direction.Rotate3D(geometry.NewVec3(0, 0, 0), xRotation, yRotation, zRotation)
}

// todos
// 1. create system obj that keeps track of objects as well as camera
// 2. each camera has its own pane(?) that uses the system obj's list
// 3. ui objects will have an array of runes that can be rendered directly, as long as it folloows the z rule
