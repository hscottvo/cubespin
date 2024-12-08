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

func main() {
	p := pane.NewPane(100, 300)
	cube := cube.NewCube(geometry.NewVec3(150, 50, 100), 30)
	p.AddObject(cube)
	ticker := tick.NewTicker(144, -1)
	done := make(chan bool)
	yaw := geometry.NewAngleDegrees(1)
	pitch := geometry.NewAngleDegrees(1)
	roll := geometry.NewAngleDegrees(1)
	fmt.Println("\033[H\033[2J")
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.Tick():
				str := "\033[H"
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
