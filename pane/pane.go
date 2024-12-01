package pane

import (
	"math"
)

type Pane struct {
	Pixels  [][]rune
	ZValues [][]float64
	Width   int
	Height  int
}

// (dist from top, dist from left)
func NewPane(height int, width int) *Pane {
	var pixels = make([][]rune, height)
	var zValues = make([][]float64, height)
	for i := range height {
		pixels[i] = make([]rune, width)
		zValues[i] = make([]float64, width)
		for j := range width {
			pixels[i][j] = ' '
			zValues[i][j] = math.MaxFloat64
		}
	}
	p := Pane{Pixels: pixels, ZValues: zValues, Width: width, Height: height}

	return &p
}

func (p *Pane) Display() string {
	s := ""
	s += "╔"
	for range p.Width {
		s += "═"
	}
	s += "╗\n"

	for i := range p.Height {
		curr_string := ""
		for j := range p.Width {
			curr_string += string(p.Pixels[i][j])
		}
		s += "║" + curr_string + "║\n"
	}

	s += "╚"
	for range p.Width {
		s += "═"
	}
	s += "╝\n"
	return s
}
func (p *Pane) Clear() {
	for i := range p.Height {
		for j := range p.Width {
			p.Pixels[i][j] = ' '
			p.ZValues[i][j] = math.MaxFloat64
		}
	}
}
