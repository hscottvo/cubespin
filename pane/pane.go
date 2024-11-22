package pane

type Pane struct {
	pixels [][]rune
	width  int
	height int
}

// (dist from top, dist from left)
func NewPane(height int, width int) *Pane {
	var pixels = make([][]rune, height)
	for i := range height {
		pixels[i] = make([]rune, width)
		for j := range width {
			pixels[i][j] = ' '
		}
	}
	p := Pane{pixels: pixels, width: width, height: height}

	return &p
}

func (p Pane) Display() string {
	s := ""
	s += "╔"
	for range p.width {
		s += "═"
	}
	s += "╗\n"

	for i := range p.height {
		curr_string := ""
		for j := range p.width {
			curr_string += string(p.pixels[i][j])
		}
		s += "║" + curr_string + "║\n"
	}

	s += "╚"
	for range p.width {
		s += "═"
	}
	s += "╝\n"
	return s
}
