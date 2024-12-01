package main

import (
	"fmt"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/tick"
	"hscottvo/cubespin/triangle"
	"time"
)

// A: {5.075961234938931 15 9.131759111665328}, B: {10 10 10}, C: {10 15 10}

func main() {
	p := pane.NewPane(30, 100)
	vec := geometry.NewVec3(0, 1, 0)
	t := triangle.NewTriangle(geometry.NewVec3(0, 0, 100), geometry.NewVec3(0, 20, 100), geometry.NewVec3(20, 20, 100), '▓')
	// t2 := triangle.NewTriangle(geometry.NewVec3(15, 15, 10), geometry.NewVec3(10, 10, 10), geometry.NewVec3(10, 15, 10), '▒')
	// t3 := triangle.NewTriangle(geometry.NewVec3(5.075961234938931, 15, 9.131759111665328), geometry.NewVec3(10, 10, 10), geometry.NewVec3(10, 15, 10), '=')
	ticker := tick.NewTicker(30, 3)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.Tick():
				fmt.Print("\033[H\033[2J")
				p.Clear()
				// t.Move(geometry.NewVec3(.25, 0.289, 0))
				t.Rotate(geometry.NewVec3(20, 20, 10), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(10))
				// t2.Rotate(geometry.NewVec3(10, 10, 10), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(10), *geometry.NewAngleDegrees(0))
				t.Render(p)
				// t2.Render(p)
				// t3.Render(p)
				vec = vec.Rotate3D(geometry.NewVec3(0, 0, 0), *geometry.NewAngleDegrees(90), *geometry.NewAngleDegrees(0), *geometry.NewAngleDegrees(0))
				fmt.Printf("x: %.2f y: %.2f z: %.2f", vec.X, vec.Y, vec.Z)
				str := ""
				fmt.Print(str, "\n")
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
