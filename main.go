package main

import (
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/KaviiSuri/wang-tiles/internal"
	"github.com/KaviiSuri/wang-tiles/linalg"
)

const (
	height   = 512
	width    = 512
	filename = "./output.jpeg"
)

func stripes(uv linalg.Vec) internal.NormalizedColor {
	n := 20.0
	return internal.NormalizedColor{
		R: (math.Sin(uv.U()*n) + 1.0) / 2,
		G: (math.Sin((uv.U()+uv.V())*n) + 1.0) / 2,
		B: (math.Cos(uv.V()*n) + 1.0) / 2,
		A: 0.0,
	}
}

func circle(uv linalg.Vec) internal.NormalizedColor {
	center := linalg.NewVec(0.5, 0.5)
	radius := .25
	if center.Sub(uv).Len() <= radius {
		return internal.NormalizedColor{R: 1.0}
	}
	return internal.NormalizedColor{R: 1.0, G: 1.0, B: 1.0}
}

func wang(uv linalg.Vec) internal.NormalizedColor {
	radius := 0.5
	centers := []struct {
		point linalg.Vec
		color linalg.Vec
	}{
		{point: linalg.NewVec(0.5, 0.0), color: linalg.NewVec(1.0, 0.0, 0.0)},
		{point: linalg.NewVec(0.5, 1.0), color: linalg.NewVec(1.0, 0.0, 1.0)},
		{point: linalg.NewVec(1.0, 0.5), color: linalg.NewVec(0.0, 1.0, 1.0)},
		{point: linalg.NewVec(0.0, 0.5), color: linalg.NewVec(1.0, 1.0, 0.0)},
	}
	result := linalg.NewSizedVec(3, 0.0)
	for _, center := range centers {
		blendFactor := 1.0 - math.Min((center.point.Sub(uv).Len()/radius), 1.0)
		newClr := result.Add(center.color.Mul(linalg.NewSizedVec(3, blendFactor)))
		result = linalg.NewSizedVec(3, 1.0).Min(newClr)
	}
	return internal.NewNormalizedColorFromVec(result)
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
			clr := wang(linalg.NewVec(u, v))
			img.Set(x, y, clr)
		}
	}
	jpeg.Encode(f, img, nil)
}
