package main

import (
	"fmt"
	// "hscottvo/cubespin/cube"
	"hscottvo/cubespin/geometry"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/tick"
	"hscottvo/cubespin/triangle"
	"time"
)

func main() {
	p := pane.NewPane(35, 100)
	t := triangle.NewTriangle(geometry.NewVec3(0, 0, 5), geometry.NewVec3(0, 5, 5), geometry.NewVec3(5, 5, 5), '▓')
	// t2 := triangle.NewTriangle(geometry.NewVec3(20, 20, 5), geometry.NewVec3(10, 10, 5), geometry.NewVec3(11, 13, 5), 'o')
	ticker := tick.NewTicker(144, 1)
	count := 0
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.Tick():
				fmt.Print("\033[H\033[2J")
				t.Move(geometry.NewVec3(10, 0, 0))
				t.Render(p)
				// t2.Render(p)
				fmt.Println("Frame ", count)
				str := ""
				fmt.Print(str, "\n")
				fmt.Println(p.Display())
				fmt.Print("\n")
				count += 1
			}
		}

	}()
	if ticker.Runtime == -1 {
		<-done

	} else {
		time.Sleep(time.Duration(ticker.Runtime) * time.Second)
	}

}
