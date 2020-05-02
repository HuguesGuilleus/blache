package generator

import (
	"./minecraftColor"
	"image/color"
	"sort"
)

type chunck struct {
	region *region
	x, z   int
	biome  imgSetRGBA
	bloc   imgSetRGBA
	height imgSetRGBA

	// Minecraft data
	Level struct {
		Biomes   interface{} // Minecraft biomes
		biomes   [256]byte   // standard biomes
		Sections []struct {
			Y       uint8
			Palette []struct {
				Name string
			}
			BlockStates []int64
		}
	}
}

// Draw images for one chunck.
func (c *chunck) draw() {
	if c.setBiome() {
		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				c.biome(x, z, minecraftColor.Biome[c.Level.biomes[z*16+x]])
			}
		}
	}

	sort.Slice(c.Level.Sections, func(i, j int) bool {
		return c.Level.Sections[i].Y > c.Level.Sections[j].Y
	})
	palette := c.genPalette()

	for x := 0; x < 16; x++ {
	nextBloc:
		for z := 0; z < 16; z++ {
			for _, sec := range c.Level.Sections {
				size := 64 * len(sec.BlockStates) / 4096
				mask := uint64((1 << size) - 1)
				for y := 15; -1 < y; y-- {

					// Source: https://wiki.vg/Chunk_Format#Chunk_Section_structure
					position := (y*16+z)*16 + x
					start := position * size / 64
					end := ((position+1)*size - 1) / 64
					offset := position * size % 64
					i := uint64(0)
					if start == end {
						i = uint64(sec.BlockStates[start]>>offset) & mask
					} else {
						i = uint64(sec.BlockStates[start]>>offset|
							sec.BlockStates[end]<<(64-offset)) & mask
					}

					col := palette[sec.Y][i]
					if col.A == 0xFF {
						c.bloc(x, z, col)
						h := sec.Y*16 + uint8(y)
						c.height(x, z, color.RGBA{h, h, h, 0xFF})
						continue nextBloc
					}
				}
			}
		}
	}
}

// Set the biome standard.
func (c *chunck) setBiome() (ok bool) {
	switch tab := c.Level.Biomes.(type) {
	case []byte:
		if len(tab) != 256 {
			return false
		}
		for i, b := range tab {
			c.Level.biomes[i] = b
		}
		return true
	case []int32:
		if len(tab) != 256 {
			return false
		}
		for i, b := range tab {
			c.Level.biomes[i] = byte(b)
		}
		return true
	}
	return false
}

// Generate the palette
func (c *chunck) genPalette() (p [16][]color.RGBA) {
	for _, sec := range c.Level.Sections {
		y := sec.Y
		size := 64 * len(sec.BlockStates) / 4096
		l := 1 << size
		p[y] = make([]color.RGBA, l, l)

		for i, b := range sec.Palette {
			p[y][i] = minecraftColor.GetBloc(b.Name)
		}
	}

	return p
}
