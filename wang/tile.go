package wang

import (
	"fmt"
	"image"
	"math"

	"github.com/KaviiSuri/wang-tiles/color"
	"github.com/KaviiSuri/wang-tiles/linalg"
)

type Tile struct {
	bltr uint8
	image.Image
}

func NewTile(bltr uint8, tileWidth, tileHeight int) Tile {
	img := image.NewRGBA(image.Rect(0, 0, tileWidth, tileHeight))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			u := float64(x) / float64(tileWidth)
			v := float64(y) / float64(tileHeight)
			clr := wangFragmentShader(bltr, linalg.NewVec(u, v))

			img.Set(x, y, clr)
			//r, g, b, a := clr.RGBA()
			//aR, aG, aB, aA := img.At(x, y).RGBA()
			//if r != aR || b != aB || g != aG || a != aA {
			//fmt.Printf("Mismatch r(%v, %v) g(%v, %v) b(%v, %v) a(%v, %v)\n", r, aR, g, aG, b, aB, a, aA)
			//}
		}
	}

	return Tile{
		bltr,
		img,
	}
}

// bltr : bottom, left, top, right (order of bits)
// imagine a uint8 as a stack of bits
// 1000011 , bitmask &1 = top, bitmask >> 1 = pop, bitmask << 1 = push(0)
func wangFragmentShader(bltr uint8, uv linalg.Vec) color.Normalized {
	radius := 0.5
	colors := []linalg.Vec{
		linalg.NewVec(1.0, 1.0, 0.0), // 0
		linalg.NewVec(1.0, 0.0, 1.0), // 1
	}
	sides := []linalg.Vec{
		linalg.NewVec(1.0, 0.5), // RIGHT
		linalg.NewVec(0.5, 0.0), // TOP
		linalg.NewVec(0.0, 0.5), // LEFT
		linalg.NewVec(0.5, 1.0), // BOTTOM
	}
	result := linalg.NewSizedVec(3, 0.0)
	for _, point := range sides {
		blendFactor := 1.0 - math.Min((linalg.Sub(point, uv).Len()/radius), 1.0)
		fmt.Println(" ", linalg.Sub(point, uv).Len()/0.5)
		clr := colors[bltr&1]
		newClr := linalg.Add(result, linalg.Mul(clr, linalg.NewSizedVec(3, blendFactor)))
		result = linalg.Min(linalg.NewSizedVec(3, 1.0), newClr)
		bltr = bltr >> 1
	}
	fmt.Println(result)
	return color.NewNormalizedFromVec(result)
}
