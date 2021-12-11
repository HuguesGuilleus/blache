// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
	"image"
	"image/color"
)

type biomeImage struct {
	paletteIndex [32 * 32][16 * 16]uint8
	palette      color.Palette
}

func newBiomeImage() (img biomeImage) {
	for i := range img.paletteIndex {
		for j := range img.paletteIndex[i] {
			img.paletteIndex[i][j] = minecraftColor.BiomeBlackIndex
		}
	}
	return
}

// The boud of the image: always 32*16 square
func (_ *biomeImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 32*16, 32*16)
}

// Return his own palette or minecraftColor.BiomePalette if not set.
func (img *biomeImage) ColorModel() color.Model {
	if img.palette != nil {
		return img.palette
	}
	return minecraftColor.BiomePalette
}

// At returns the color of the pixel at (x, y).
func (img *biomeImage) At(x, z int) color.Color {
	if img.palette != nil {
		return img.palette[img.ColorIndexAt(x, z)]
	}
	return minecraftColor.BiomePalette[img.ColorIndexAt(x, z)]
}

// ColorIndexAt returns the palette index of the pixel at (x, y = z).
func (img *biomeImage) ColorIndexAt(x, z int) uint8 {
	return img.paletteIndex[z/16*32+x/16][(z%16)*16+x%16]
}

// Draw the image
func (img *biomeImage) draw(chunckX, chunckZ int, biome interface{}) error {
	chunck := img.paletteIndex[chunckZ*32+chunckX][:]
	switch biome := biome.(type) {
	case nil:
	case []byte:
		switch l := len(biome); l {
		case 256:
			copy(chunck, biome)
		case 0:
		default:
			return fmt.Errorf("Biome is bytes array with %d length (expected 256)", l)
		}
	case []int32:
		switch len(biome) {
		case 256:
			for x := 0; x < 16; x++ {
				for z := 0; z < 16; z++ {
					chunck[z*16+x] = uint8(biome[z*16+x])
				}
			}
		case 1024:
			for x := 0; x < 16; x += 4 {
				for z := 0; z < 16; z += 4 {
					c := uint8(biome[0x3F0|z|x>>2])
					for i := 0; i < 4; i++ {
						for j := 0; j < 4; j++ {
							chunck[(z+j)*16+(x+i)] = c
						}
					}
				}
			}
		case 0:
		default:
			return fmt.Errorf("[]int32 length is not 2565 or 1024, it't: %d", len(biome))
		}
	default:
		return fmt.Errorf("The biome is %T (expected byte or int32 array, or nothing)", biome)
	}
	return nil
}

// After draw on all chunck, select only used color. Do not used draw after.
func (img *biomeImage) processPalette() {
	// Seach use color
	var usedColors [256]bool
	for _, chunck := range img.paletteIndex {
		for _, c := range chunck {
			usedColors[c] = true
		}
	}
	var nbUsedColor = 0
	for _, b := range usedColors {
		if b {
			nbUsedColor++
		}
	}

	img.palette = make(color.Palette, nbUsedColor)
	var colorCorrelation [256]uint8
	var fillingIndex = uint8(0)
	for colorIndex, used := range usedColors {
		if !used {
			continue
		}
		img.palette[fillingIndex] = minecraftColor.BiomePalette[colorIndex]
		colorCorrelation[colorIndex] = fillingIndex
		fillingIndex++
	}

	for i := range img.paletteIndex {
		for j, oldColor := range img.paletteIndex[i] {
			img.paletteIndex[i][j] = colorCorrelation[oldColor]
		}
	}
}
