package wang

import (
	"image"
	"image/draw"
	"log"
	"sync"
)

type Grid struct {
	bltrs         [][]uint8
	widthInTiles  int
	heightInTiles int
	tileWidth     int
	tileHeight    int
	*image.RGBA
}

// Generates an Empty Grid from Given BLTRS Slice
func NewGridFromBLTRS(bltrs [][]uint8, tileSet []image.Image, widthInTiles, heightInTiles, tileWidth, tileHeight int) Grid {
	g := Grid{
		bltrs:         bltrs,
		widthInTiles:  widthInTiles,
		heightInTiles: heightInTiles,
		tileWidth:     tileWidth,
		tileHeight:    tileHeight,
	}
	img := image.NewRGBA(image.Rect(0, 0, widthInTiles*tileWidth, heightInTiles*tileHeight))
	g.RGBA = img
	g.render(tileSet)
	return g
}

func (g Grid) render(tileSet []image.Image) {
	rect := image.Rect(0, 0, g.tileWidth, g.tileHeight)
	var wg sync.WaitGroup
	for r := 0; r < g.heightInTiles; r++ {
		for c := 0; c < g.widthInTiles; c++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, r, c int) {
				wg.Done()
				tile := tileSet[g.bltrs[r][c]]
				translateBy := image.Point{c * g.tileHeight, r * g.tileWidth}
				draw.Draw(
					g,
					rect.Add(translateBy),
					tile,
					image.Point{},
					draw.Src,
				)

			}(&wg, r, c)
		}
	}

	wg.Wait()
}

func (g Grid) checkBLTRSAndDimentions() {
	if g.heightInTiles != len(g.bltrs) {
		log.Fatalf("Height Mismatch %v != %v", g.heightInTiles, len(g.bltrs))
	}

	for i, bltr := range g.bltrs {
		if g.widthInTiles != len(bltr) {
			log.Fatalf("Width Mismatch in row %v %v != %v", i, g.heightInTiles, len(bltr))
		}
	}
}
