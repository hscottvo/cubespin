package geometry

import (
	"math"
)

type BoundingBox struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func NewBoundingBox(vecs []Vec3) BoundingBox {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for _, i := range vecs {
		minX = min(minX, int(math.Floor(i.X)))
		minY = min(minY, int(math.Floor(i.Y)))
		maxX = max(maxX, int(math.Ceil(i.X)))
		maxY = max(maxY, int(math.Ceil(i.Y)))
	}
	return BoundingBox{MinX: minX, MinY: minY, MaxX: maxX, MaxY: maxY}
}
