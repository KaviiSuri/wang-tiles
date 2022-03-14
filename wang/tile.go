package wang

import (
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
			clr := trianglesWithHolesFragmentShader(bltr, linalg.NewVec(u, v))

			img.Set(x, y, clr)
		}
	}

	return Tile{
		bltr,
		img,
	}
}

var colors = []linalg.Vec{
	linalg.NewVec(0.1, 0.1, 0.1), // 0
	linalg.NewVec(0.9, 0.9, 0.9), // 1
}

// bltr : bottom, left, top, right (order of bits)
// imagine a uint8 as a stack of bits
// 1000011 , bitmask &1 = top, bitmask >> 1 = pop, bitmask << 1 = push(0)
func gradeintFragmentShader(bltr uint8, uv linalg.Vec) color.Normalized {
	radius := 0.5
	sides := []linalg.Vec{
		linalg.NewVec(1.0, 0.5), // RIGHT
		linalg.NewVec(0.5, 0.0), // TOP
		linalg.NewVec(0.0, 0.5), // LEFT
		linalg.NewVec(0.5, 1.0), // BOTTOM
	}
	result := linalg.NewSizedVec(3, 0.0)
	for _, point := range sides {
		blendFactor := 1.0 - math.Min((linalg.Sub(point, uv).Len()/radius), 1.0)
		//clr := colors[bltr&1]
		//newClr := linalg.Add(result, linalg.Mul(clr, linalg.NewSizedVec(3, blendFactor)))
		//result = linalg.Min(linalg.NewSizedVec(3, 1.0), newClr)
		result = linalg.Lerp(result, colors[bltr&1], linalg.NewSizedVec(3, blendFactor))
		bltr = bltr >> 1
	}
	result = linalg.Map(result, func(a float64, i int) float64 {
		return math.Pow(a, 1.0/2.2)
	})
	return color.NewNormalizedFromVec(result)
}

func trianglesFragmentShader(bltr uint8, uv linalg.Vec) color.Normalized {
	distances := []float64{
		1.0 - uv.X(), // r
		uv.Y(),       // t
		uv.X(),       // l
		1.0 - uv.Y(), // b
	}
	idx := -1
	for i := 0; i < 4; i++ {
		if idx < 0 || distances[idx] > distances[i] {
			idx = i
		}
	}

	for ; idx > 0; idx-- {
		bltr >>= 1
	}

	return color.NewNormalizedFromVec(colors[bltr&1])
}
func trianglesWithHolesFragmentShader(bltr uint8, uv linalg.Vec) color.Normalized {
	distances := []float64{
		1.0 - uv.X(), // r
		uv.Y(),       // t
		uv.X(),       // l
		1.0 - uv.Y(), // b
	}
	idx := -1
	for i := 0; i < 4; i++ {
		if idx < 0 || distances[idx] > distances[i] {
			idx = i
		}
	}

	if distances[idx] > 0.25 {
		t := 1.0 - float64(bltr)/16.0

		return color.NewNormalizedFromVec(linalg.Map(linalg.NewVec(t, t, t, 1), func(a float64, i int) float64 {
			return math.Pow(a, 1.0/2.2)
		}))
	}

	for ; idx > 0; idx-- {
		bltr >>= 1
	}

	return color.NewNormalizedFromVec(colors[bltr&1])
}
