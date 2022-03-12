package main

import (
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/KaviiSuri/wang-tiles/utils"
)

const (
	height   = 512
	width    = 512
	filename = "./output.jpeg"
)

func stripes(u, v float64) utils.NormalizedColor {
	n := 20.0
	return utils.NormalizedColor{
		R: (math.Sin(u*n) + 1.0) / 2,
		G: (math.Sin((u+v)*n) + 1.0) / 2,
		B: (math.Cos(v*n) + 1.0) / 2,
		A: 0.0,
	}
}

func circles(u, v float64) utils.NormalizedColor {
	cx := 0.5
	cy := 0.5
	dx := cx - u
	dy := cy - v
	r := .25
	if point := dx*dx + dy*dy; point <= r*r {
		return utils.NormalizedColor{R: 1.0}
	}
	return utils.NormalizedColor{R: 0.0}
}

func main() {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			u := float64(x) / float64(width)
			v := float64(y) / float64(height)
			clr := circles(u, v)
			img.Set(x, y, clr)
		}
	}
	jpeg.Encode(f, img, nil)
}
