// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

// A palletted image for region tile, of any type (biome, blocks, height).
type regionImage struct {
	pixels  [32 * 32][16 * 16]uint8
	palette color.Palette
}

// The bounds of the image: always 32*16 square
func (_ *regionImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 32*16, 32*16)
}

// Return his own palette or minecraftColor.BiomePalette if not set.
func (img *regionImage) ColorModel() color.Model {
	return img.palette
}

// At returns the color of the pixel at (x, y).
func (img *regionImage) At(x, z int) color.Color {
	return img.palette[img.ColorIndexAt(x, z)]
}

// ColorIndexAt returns the palette index of the pixel at (x, y = z).
func (img *regionImage) ColorIndexAt(x, z int) uint8 {
	return img.pixels[z/16*32+x/16][(z%16)*16+x%16]
}

// Get the pixels of a specific chunck.
func (img *regionImage) chunck(chunckX, chunckZ int) []uint8 {
	return img.pixels[chunckZ*32+chunckX][:]
}

// Return the PNG encoded image.
func (img *regionImage) BytesPNG() []byte {
	img.processPalette()
	buff := bytes.Buffer{}
	png.Encode(&buff, img)
	return buff.Bytes()
}

// After draw on all chunck, select only used color. Do not used draw after.
func (img *regionImage) processPalette() {
	// Search and count used color
	var usedColors [256]bool
	for _, chunck := range img.pixels {
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

	// Create the new palette and corellation table.
	newPalette := make(color.Palette, nbUsedColor)
	colorCorrelation := [256]uint8{}
	fillingIndex := uint8(0)
	for colorIndex, used := range usedColors {
		if used {
			newPalette[fillingIndex] = img.palette[colorIndex]
			colorCorrelation[colorIndex] = fillingIndex
			fillingIndex++
		}
	}
	img.palette = newPalette

	// Apply the new pallette.
	for i := range img.pixels {
		for j, oldColor := range img.pixels[i] {
			img.pixels[i][j] = colorCorrelation[oldColor]
		}
	}
}
