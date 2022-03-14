package wang

import (
	"fmt"
	"image"
	"math/rand"
	"sync"
)

func NewWangTileSet(tileWidth, tileHeight int) []image.Image {
	//atlas := image.NewRGBA(image.Rect(0, 0, widthInTiles*tileWidth, heightInTiles*tileHeight))
	//r := image.Rect(0, 0, tileWidth, tileHeight)
	//var wg sync.WaitGroup
	//for bltr := uint8(0); bltr < 16; bltr++ {
	//wg.Add(1)
	//go func(wg *sync.WaitGroup, bltr uint8) {
	//defer wg.Done()
	//img := NewTile(bltr, tileWidth, tileHeight)
	//translateBy := image.Point{
	//int(bltr%uint8(widthInTiles)) * tileWidth,
	//int(bltr/uint8(widthInTiles)) * tileHeight,
	//}
	//draw.Draw(
	//atlas,
	//r.Add(translateBy),
	//img,
	//image.Point{},
	//draw.Src,
	//)

	//fmt.Printf("Generated Tile %02d\n", bltr)
	//}(&wg, bltr)
	//}

	//wg.Wait()

	atlas := make([]image.Image, 16)
	var wg sync.WaitGroup
	for bltr := uint8(0); bltr < 16; bltr++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, bltr uint8) {
			defer wg.Done()
			img := NewTile(bltr, tileWidth, tileHeight)
			atlas[bltr] = img
			fmt.Printf("Generated Tile %02d\n", bltr)
		}(&wg, bltr)
	}

	wg.Wait()
	return atlas
}

// Generates a grid which follows Wang Tiles Rules
func NewWangGrid(tileSet []image.Image, widthInTiles, heightInTiles, tileWidth, tileHeight int) Grid {
	bltrs := [][]uint8{}
	for r := 0; r < heightInTiles; r++ {
		bltrs = append(bltrs, make([]uint8, widthInTiles))
	}
	// +---+---+---+
	// | m | l | l
	// +---+---+---+
	// | t |tl |tl
	// +---+---+---+
	// | t |tl |tl
	// +   +   +
	bltrs[0][0] = RandomBLTRWithConstraint(0, IGNORE_ALL)
	for c := 1; c < widthInTiles; c++ {
		bltrs[0][c] = RandomBLTRWithConstraint(bltrs[0][c-1], CONSIDER_L)
	}

	for r := 1; r < heightInTiles; r++ {
		bltrs[r][0] = RandomBLTRWithConstraint(bltrs[r-1][0], CONSIDER_T)
	}

	for r := 1; r < heightInTiles; r++ {
		for c := 1; c < widthInTiles; c++ {
			values := (bltrs[r-1][c] & CONSIDER_T) | (bltrs[r][c-1] & CONSIDER_L)
			bltrs[r][c] = RandomBLTRWithConstraint(values, CONSIDER_L|CONSIDER_T)
		}
	}

	fmt.Println(bltrs)

	grid := NewGridFromBLTRS(bltrs, tileSet, widthInTiles, heightInTiles, tileWidth, tileHeight)
	return grid
}

const (
	CONSIDER_B = uint8(1 << iota)
	CONSIDER_L
	CONSIDER_T
	CONSIDER_R
)
const (
	IGNORE_B = uint8(0 << iota)
	IGNORE_L
	IGNORE_T
	IGNORE_R
)

const (
	IGNORE_ALL = 0
)

// Generates a BLTR with given constraints
// Constraints are in the form of 2 bitmasks,
// values (0 => value 0, 1 => value 1)
// position (0 => don't apply constraint. 1 => apply constraint)
func RandomBLTRWithConstraint(values, position uint8) uint8 {
	candidates := []uint8{}
	len := 0
	for candidate := uint8(0); candidate < 16; candidate++ {
		if candidate&position == values&position {
			candidates = append(candidates, candidate)
			len++
		}
	}
	return candidates[rand.Intn(len)]
}
