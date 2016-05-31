package sdl

import (
	"math/rand"
)

type Star struct {
	xcoord int
	ycoord int
	speed int
}

func Init(xs int, ys int) Star {
	var star Star
	star.xcoord = rand.Int() % xs
	star.ycoord = rand.Int() % ys
	star.speed = -3 - (rand.Int() % 10)
	return star
}

func Move(s *Star) {
	s.xcoord += s.speed
}
