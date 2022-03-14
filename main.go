package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"
	"sync"

	"github.com/KaviiSuri/wang-tiles/color"
	"github.com/KaviiSuri/wang-tiles/linalg"
	"github.com/KaviiSuri/wang-tiles/wang"
)

const (
	tileHeight       = 64
	tileWidth        = 64
	tileSetNumTilesH = 4
	tileSetNumTilesV = 4
	tileSetHeight    = tileHeight * tileSetNumTilesH
	tileSetWidth     = tileWidth * tileSetNumTilesV
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

// TODO: Maybe extract this logic to an `Atlas` or `Grid` struct
func generateTileSet() image.Image {
	atlas := image.NewRGBA(image.Rect(0, 0, tileSetWidth, tileSetHeight))
	r := image.Rect(0, 0, tileWidth, tileHeight)
	var wg sync.WaitGroup
	for bltr := uint8(0); bltr < 16; bltr++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, bltr uint8) {
			defer wg.Done()
			img := wang.NewTile(bltr, tileWidth, tileHeight)
			translateBy := image.Point{
				int(bltr%tileSetNumTilesH) * tileWidth,
				int(bltr/tileSetNumTilesH) * tileHeight,
			}
			draw.Draw(
				atlas,
				r.Add(translateBy),
				img,
				image.Point{},
				draw.Src,
			)

			fmt.Printf("Generated Tile %02d\n", bltr)
		}(&wg, bltr)
	}

	wg.Wait()
	return atlas
}

func main() {
	filename := "./temp/atlas.jpeg"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img := generateTileSet()
	jpeg.Encode(f, img, nil)
}
