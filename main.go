package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/KaviiSuri/wang-tiles/color"
	"github.com/KaviiSuri/wang-tiles/linalg"
	"github.com/KaviiSuri/wang-tiles/wang"
)

const (
	tileHeight    = 64
	tileWidth     = 64
	tileSetHeight = tileHeight * 4
	tileSetWidth  = tileWidth * 4
)

func stripes(uv linalg.Vec) color.Normalized {
	n := 20.0
	return color.Normalized{
		R: (math.Sin(uv.U()*n) + 1.0) / 2,
		G: (math.Sin((uv.U()+uv.V())*n) + 1.0) / 2,
		B: (math.Cos(uv.V()*n) + 1.0) / 2,
		A: 0.0,
	}
}

func circle(uv linalg.Vec) color.Normalized {
	center := linalg.NewVec(0.5, 0.5)
	radius := .25
	if linalg.Sub(center, uv).Len() <= radius {
		return color.Normalized{R: 1.0}
	}
	return color.Normalized{R: 1.0, G: 1.0, B: 1.0}
}

// bltr : bottom, left, top, right (order of bits)
// imagine a uint8 as a stack of bits
// 1000011 , bitmask &1 = top, bitmask >> 1 = pop, bitmask << 1 = push(0)
//func wang(bltr uint8, uv linalg.Vec) color.Normalized {
//radius := 0.5
//colors := []linalg.Vec{
//linalg.NewVec(1.0, 1.0, 0.0), // 0
//linalg.NewVec(1.0, 0.0, 1.0), // 1
//}
//sides := []linalg.Vec{
//linalg.NewVec(1.0, 0.5), // RIGHT
//linalg.NewVec(0.5, 0.0), // TOP
//linalg.NewVec(0.0, 0.5), // LEFT
//linalg.NewVec(0.5, 1.0), // BOTTOM
//}
//result := linalg.NewSizedVec(3, 0.0)
//for _, point := range sides {
//blendFactor := 1.0 - math.Min((linalg.Sub(point, uv).Len()/radius), 1.0)
//clr := colors[bltr&1]
//newClr := linalg.Add(result, linalg.Mul(clr, linalg.NewSizedVec(3, blendFactor)))
//result = linalg.Min(linalg.NewSizedVec(3, 1.0), newClr)
//bltr = bltr >> 1
//}
//return color.NewNormalizedFromVec(result)
//}

//func GenerateWangTile(bltr uint8) image.Image {
//img := image.NewRGBA(image.Rect(0, 0, tileWidth, tileHeight))
//for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
//for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
//u := float64(x) / float64(tileWidth)
//v := float64(y) / float64(tileHeight)
//clr := wang(bltr, linalg.NewVec(u, v))
//img.Set(x, y, clr)
//}
//}

//return img
//}

func main() {
	for bltr := uint8(0); bltr < 16; bltr++ {
		f, err := os.Create(fmt.Sprintf("./results/tile-%02d.jpeg", bltr))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		img := wang.NewTile(bltr, tileWidth, tileHeight)
		fmt.Printf("Generated Tile %02d\n", bltr)
		jpeg.Encode(f, img, nil)
	}
}
