package generator

import (
	"./biome"
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
				c.biome(x, z, biome.Color[c.Level.biomes[z*16+x]])
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
				mask := int64(0xFFFF >> (16 - size))
				for y := 15; -1 < y; y-- {

					pos := ((y*16+z)*16 + (15 - x)) * size

					i := sec.BlockStates[pos/64]
					if shift := 64 - size - pos%64; shift < 0 {
						i = i<<(-shift) +
							sec.BlockStates[pos/64+1]>>(64+shift)
					} else {
						i >>= shift
					}
					i &= mask

					if int(i) >= len(palette[sec.Y]) {
						continue
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

// Generate the palette and standardise BlockStates
func (c *chunck) genPalette() (p [16][]color.RGBA) {
	cb := &c.region.g.colorBloc
	cb.RLock()
	defer cb.RUnlock()

	for _, sec := range c.Level.Sections {
		y := sec.Y
		l := len(sec.Palette)
		p[y] = make([]color.RGBA, l, l)

		for i, b := range sec.Palette {
			p[y][i] = cb.m[b.Name]
		}
	}

	return p
}
