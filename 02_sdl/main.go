package main

import (
	"github.com/karpatx/udemyTraining/02_sdl/sdl"
	"os"
	sdlx "github.com/veandco/go-sdl2/sdl"
)

func main() {
	var running bool
	var event sdlx.Event
	var stars []sdl.Star
	xs := 640
	ys := 400
	numstars := 100

	for i := 0; i < numstars; i++ {
		s := sdl.Init(xs, ys)
		stars = append(stars, s)
	}

	sdl.Sdlinit(640, 400)

	sdl.Initrender()

	running = true
	for running {
		for event = sdlx.PollEvent(); event != nil; event = sdlx.PollEvent() {
			switch event.(type) {
			case *sdlx.QuitEvent:
				running = false
			}
		}
		for i, _ := range stars {
			sdl.Move(&stars[i])
			sdl.Render(&stars[i])
		}
		sdl.Show()
	}

	sdl.Sdldeinit()
	os.Exit(0)
}
