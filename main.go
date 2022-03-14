package main

import (
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
	widthInTiles  = 80
	heightInTiles = 40
	tileSetHeight = tileHeight * widthInTiles
	tileSetWidth  = tileWidth * heightInTiles
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

func main() {
	filename := "./temp/grid.jpeg"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	tileSet := wang.NewWangTileSet(tileWidth, tileHeight)
	img := wang.NewWangGrid(tileSet, widthInTiles, heightInTiles, tileWidth, tileHeight)
	//bltrs := [][]uint8{
	//{0, 1, 2, 3},
	//{4, 5, 6, 7},
	//{8, 9, 10, 11},
	//{12, 13, 14, 15},
	//}
	//img := wang.NewGridFromBLTRS(bltrs, tileSet, tileSetNumTilesH, tileSetNumTilesV, tileWidth, tileHeight)
	jpeg.Encode(f, img, nil)
}
