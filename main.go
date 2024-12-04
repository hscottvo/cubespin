package main

import (
	"fmt"
	"hscottvo/cubespin/cube"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/tick"
	// "hscottvo/cubespin/triangle"
	"time"
)

// A: {5.075961234938931 15 9.131759111665328}, B: {10 10 10}, C: {10 15 10}

func main() {
	p := pane.NewPane(35, 100)
	// a := geometry.NewVec3(5, 10, 120)
	// b := geometry.NewVec3(10, 10, 120)
	// c := geometry.NewVec3(5, 5, 120)
	// d := geometry.NewVec3(10, 5, 120)
	// square := cube.NewSquareFromPoints(a, b, c, d, 'x')
	// square.Move(geometry.NewVec3(80, 10, 0))
	// square.Rotate3D(*geometry.NewAngleDegrees(5))
	// t := triangle.NewTriangle(geometry.NewVec3(10, 0, 100), geometry.NewVec3(10, 30, 100), geometry.NewVec3(50, 15, 150), '▓')
	cube := cube.NewCube(geometry.NewVec3(20, 20, 20), 10)
	cube.Rotate(cube.Center(), *geometry.NewAngleDegrees(45), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(0))
	// p.AddObject(square)
	// p.AddObject(t)
	p.AddObject(cube)
	ticker := tick.NewTicker(30, 3)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.Tick():
				fmt.Print("\033[H\033[2J")
				cube.Rotate(cube.Center(), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(10), *geometry.NewAngleDegrees(0))
				// p.Clear()
				// t.Rotate(geometry.NewVec3(20, 20, 100), *geometry.NewAngleDegrees(10), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(10))
				// t.Render(p)
				// square.Move(geometry.NewVec3(-1, 0, 0))
				// square.Render(p)
				fmt.Println(p.Display())
				fmt.Print("\n")
			}
		}

	}()
	if ticker.Runtime == -1 {
		<-done

	} else {
		time.Sleep(time.Duration(ticker.Runtime) * time.Second)
	}

}
