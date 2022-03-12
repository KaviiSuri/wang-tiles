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

func circles(uv utils.Vec2) utils.NormalizedColor {
	center := utils.Vec2{X: 0.5, Y: 0.5}
	diff := center.Sub(&uv)
	radius := .25
	if math.Pow(diff.Len(), 2) <= radius*radius {
		return utils.NormalizedColor{R: 1.0}
	}
	return utils.NormalizedColor{R: 1.0, G: 1.0, B: 1.0}
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
			clr := circles(utils.Vec2{X: u, Y: v})
			img.Set(x, y, clr)
		}
	}
	jpeg.Encode(f, img, nil)
}
