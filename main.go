package main

import (
	"fmt"
	"hscottvo/cubespin/cube"
	"hscottvo/cubespin/pane"
	"hscottvo/cubespin/tick"
	"time"
)

func main() {
	p := pane.NewPane(35, 100)
	c := cube.NewCube()
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
				fmt.Println("Frame ", count)
				str := ""
				for range count {
					str += c.Display()
				}
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
