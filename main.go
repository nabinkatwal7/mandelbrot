package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 640
	maxIt = 128
)

var (
	palette [maxIt]byte
)

func init() {
	for i:= range palette {
		palette[i] = byte(math.Sqrt(float64(i)/float64(len(palette)))*0x80)
	}
}
