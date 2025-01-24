package main

import (
	"fmt"
	"hscottvo/cubespin/cube"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/tick"
	// "hscottvo/cubespin/triangle"
	"math"
	"time"
)

func main() {
	planePos := geometry.NewVec3(0, 0, 0)
	planeNorm := geometry.NewVec3(0, 0, 1)
	angle := geometry.NewAngleDegrees(0)
	p := pane.NewPane(100, 300, planePos, planeNorm)
	cube := cube.NewCube(geometry.NewVec3(150, 50, 100), 30)
	p.AddObject(cube)
	ticker := tick.NewTicker(144, -1)
	done := make(chan bool)
	yaw := geometry.NewAngleDegrees(1)
	pitch := geometry.NewAngleDegrees(1)
	roll := geometry.NewAngleDegrees(1.5)
	fmt.Println("\033[H\033[2J")
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.Tick():
				str := "\033[H"
				angle.Add(.01)
				// fmt.Println(angle.Degree())
				cube.Move(geometry.NewVec3(0, math.Cos(angle.Degree())*0.2, 0))
				// fmt.Print("\033[H\033[2J")
				cube.Rotate(cube.Center(), *yaw, *pitch, *roll)
				fmt.Println(str + p.Display())
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
