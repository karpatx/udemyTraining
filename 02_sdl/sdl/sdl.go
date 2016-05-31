package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"os"
)

var window *sdl.Window
var renderer *sdl.Renderer
var width int
var height int

func Sdlinit(xs int, ys int) bool {
	width = xs
	height = ys
	window, err := sdl.CreateWindow("sdl test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		xs, ys, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return false
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()

	return true
}

func Sdldeinit() {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

func Initrender() {
	//renderer.Clear()
	var rect = sdl.Rect{0, 0, int32(width), int32(height)}
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.FillRect(&rect)
}

func Show() {
	renderer.Present()
	sdl.Delay(50)
}

func Render(star *Star) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(star.xcoord, star.ycoord)
}