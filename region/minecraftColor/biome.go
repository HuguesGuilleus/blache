// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package minecraftColor

import (
	"image/color"
)

var (
	BiomePalette color.Palette = func() (palette color.Palette) {
		palette = make(color.Palette, len(Biome))
		for i, c := range Biome {
			palette[i] = color.Color(c)
		}
		return
	}()
	// The index in Biome or BiomePalette of the first black color.
	BiomeBlackIndex = func() uint8 {
		for i, c := range Biome {
			if c == (color.RGBA{}) {
				return uint8(i)
			}
		}
		panic("Need a black color")
	}()
	// https://minecraft-el.gamepedia.com/Biome/ID
	Biome = [256]color.RGBA{
		color.RGBA{R: 0x00, G: 0x00, B: 0x70, A: 0xFF}, // Ocean
		color.RGBA{R: 0x8D, G: 0xB3, B: 0x60, A: 0xFF}, // Plains
		color.RGBA{R: 0xFA, G: 0x94, B: 0x18, A: 0xFF}, // Desert
		color.RGBA{R: 0x60, G: 0x60, B: 0x60, A: 0xFF}, // Extreme Hills
		color.RGBA{R: 0x05, G: 0x66, B: 0x21, A: 0xFF}, // Forest
		color.RGBA{R: 0x0B, G: 0x66, B: 0x59, A: 0xFF}, // Taiga
		color.RGBA{R: 0x07, G: 0xF9, B: 0xB2, A: 0xFF}, // Swampland
		color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}, // River
		color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}, // Hell
		color.RGBA{R: 0x80, G: 0x80, B: 0xFF, A: 0xFF}, // The End (Sky)
		color.RGBA{R: 0x90, G: 0x90, B: 0xA0, A: 0xFF}, // FrozenOcean
		color.RGBA{R: 0xA0, G: 0xA0, B: 0xFF, A: 0xFF}, // FrozenRiver
		color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Ice Plains
		color.RGBA{R: 0xA0, G: 0xA0, B: 0xA0, A: 0xFF}, // Ice Mountains
		color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xFF}, // MushroomIsland
		color.RGBA{R: 0xA0, G: 0x00, B: 0xFF, A: 0xFF}, // MushroomIslandShore
		color.RGBA{R: 0xFA, G: 0xDE, B: 0x55, A: 0xFF}, // Beach
		color.RGBA{R: 0xD2, G: 0x5F, B: 0x12, A: 0xFF}, // DesertHills
		color.RGBA{R: 0x22, G: 0x55, B: 0x1C, A: 0xFF}, // ForestHills
		color.RGBA{R: 0x16, G: 0x39, B: 0x33, A: 0xFF}, // TaigaHills
		color.RGBA{R: 0x72, G: 0x78, B: 0x9A, A: 0xFF}, // Extreme Hills Edge
		color.RGBA{R: 0x53, G: 0x7B, B: 0x09, A: 0xFF}, // Jungle
		color.RGBA{R: 0x2C, G: 0x42, B: 0x05, A: 0xFF}, // JungleHills
		color.RGBA{R: 0x62, G: 0x8B, B: 0x17, A: 0xFF}, // JungleEdge
		color.RGBA{R: 0x00, G: 0x00, B: 0x30, A: 0xFF}, // Deep Ocean
		color.RGBA{R: 0xA2, G: 0xA2, B: 0x84, A: 0xFF}, // Stone Beach
		color.RGBA{R: 0xFA, G: 0xF0, B: 0xC0, A: 0xFF}, // Cold Beach
		color.RGBA{R: 0x30, G: 0x74, B: 0x44, A: 0xFF}, // Birch Forest
		color.RGBA{R: 0x1F, G: 0x5F, B: 0x32, A: 0xFF}, // Birch Forest Hills
		color.RGBA{R: 0x40, G: 0x51, B: 0x1A, A: 0xFF}, // Roofed Forest
		color.RGBA{R: 0x31, G: 0x55, B: 0x4A, A: 0xFF}, // Cold Taiga
		color.RGBA{R: 0x24, G: 0x3F, B: 0x36, A: 0xFF}, // Cold Taiga Hills
		color.RGBA{R: 0x59, G: 0x66, B: 0x51, A: 0xFF}, // Mega Taiga
		color.RGBA{R: 0x54, G: 0x5F, B: 0x3E, A: 0xFF}, // Mega Taiga Hills
		color.RGBA{R: 0x50, G: 0x70, B: 0x50, A: 0xFF}, // Extreme Hills+
		color.RGBA{R: 0xBD, G: 0xB2, B: 0x5F, A: 0xFF}, // Savanna
		color.RGBA{R: 0xA7, G: 0x9D, B: 0x64, A: 0xFF}, // Savanna Plateau
		color.RGBA{R: 0xD9, G: 0x45, B: 0x15, A: 0xFF}, // Mesa
		color.RGBA{R: 0xB0, G: 0x97, B: 0x65, A: 0xFF}, // Mesa Plateau F
		color.RGBA{R: 0xCA, G: 0x8C, B: 0x65, A: 0xFF}, // Mesa Plateau
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{}, // void
		color.RGBA{}, // Plains M
		color.RGBA{R: 0xB5, G: 0xDB, B: 0x88, A: 0xFF}, // Sunflower Plains
		color.RGBA{R: 0xFF, G: 0xBC, B: 0x40, A: 0xFF}, // Desert M
		color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xFF}, // Extreme Hills M
		color.RGBA{R: 0x6A, G: 0x74, B: 0x25, A: 0xFF}, // Flower Forest
		color.RGBA{R: 0x59, G: 0x66, B: 0x51, A: 0xFF}, // Taiga M
		color.RGBA{R: 0x2F, G: 0xFF, B: 0xDA, A: 0xFF}, // Swampland M
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{R: 0xB4, G: 0xDC, B: 0xDC, A: 0xFF}, // Ice Plains Spikes
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{R: 0x7B, G: 0xA3, B: 0x31, A: 0xFF}, // Jungle M
		color.RGBA{},
		color.RGBA{R: 0x8A, G: 0xB3, B: 0x3F, A: 0xFF}, // JungleEdge M
		color.RGBA{},
		color.RGBA{},
		color.RGBA{},
		color.RGBA{R: 0x58, G: 0x9C, B: 0x6C, A: 0xFF}, // Birch Forest M
		color.RGBA{R: 0x47, G: 0x87, B: 0x5A, A: 0xFF}, // Birch Forest Hills M
		color.RGBA{R: 0x68, G: 0x79, B: 0x42, A: 0xFF}, // Roofed Forest M
		color.RGBA{R: 0x59, G: 0x7D, B: 0x72, A: 0xFF}, // Cold Taiga M
		color.RGBA{},
		color.RGBA{R: 0x6B, G: 0x5F, B: 0x4C, A: 0xFF}, // Mega Spruce Taiga
		color.RGBA{R: 0x6D, G: 0x77, B: 0x66, A: 0xFF}, // Redwood Taiga Hills M
		color.RGBA{R: 0x78, G: 0x98, B: 0x78, A: 0xFF}, // Extreme Hills+ M
		color.RGBA{R: 0xE5, G: 0xDA, B: 0x87, A: 0xFF}, // Savanna M
		color.RGBA{R: 0xCF, G: 0xC5, B: 0x8C, A: 0xFF}, // Savanna Plateau M
		color.RGBA{R: 0xFF, G: 0x6D, B: 0x3D, A: 0xFF}, // Mesa (Bryce)
		color.RGBA{R: 0xD8, G: 0xBF, B: 0x8D, A: 0xFF}, // Mesa Plateau F M
		color.RGBA{R: 0xF2, G: 0xB4, B: 0x8D, A: 0xFF}, // Mesa Plateau M
	}
)
