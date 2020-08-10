// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"./minecraftColor"
	"fmt"
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
		biomes     [256]byte
		Biomes     interface{}
		Sections   []section
		Structures struct {
			Starts map[string]struct{}
		}
	}
}

type section struct {
	Y       uint8
	Palette []struct {
		Name string
	}
	BlockStates []int64
}

// Draw images for one chunck.
func (c *chunck) draw() {
	c.setBiome()
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			c.biome(x, z, minecraftColor.Biome[c.Level.biomes[z*16+x]])
		}
	}

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

// Change Chunck.Level.Biomes to c.Level.biomes, a array of byte.
func (c *chunck) setBiome() {
	switch tab := c.Level.Biomes.(type) {
	case []byte:
		if l := len(tab); l == 0 {
			return
		} else if l != 256 {
			c.region.g.err <- fmt.Errorf("In chunck (%d,%d) Chunck.Level.Biome is not a 256 len byte array: len is %d", c.x+c.region.X*32, c.z+c.region.Z*32, l)
			return
		}
		copy(c.Level.biomes[:], tab)
	case []int32:
		if l := len(tab); l == 0 {
			return
		} else if l != 256 {
			c.region.g.err <- fmt.Errorf("In chunck (%d,%d) Chunck.Level.Biome is not a 256 len byte array: len is %d", c.x+c.region.X*32, c.z+c.region.Z*32, l)
			c.region.g.err <- fmt.Errorf("In chunck (%d,%d) Chunck.Level.Biome is not a 256 len int32 array: len is %d", c.x+c.region.X*32, c.z+c.region.Z*32, l)
			return
		}
		for i, b := range tab {
			c.Level.biomes[i] = byte(b)
		}
	default:
		c.region.g.err <- fmt.Errorf("Chunck.Level.Biome (%d,%d) is not a 256 len bytes array: %T", c.x+c.region.X*32, c.z+c.region.Z*32, tab)
	}
	c.Level.Biomes = nil
	return
}

// Generate the palette
func (c *chunck) genPalette() (p [16][]color.RGBA) {
	secs := make([]section, 0, len(c.Level.Sections))
	for _, s := range c.Level.Sections {
		if len(s.BlockStates) > 0 {
			secs = append(secs, s)
		}
	}
	c.Level.Sections = secs

	sort.Slice(c.Level.Sections, func(i, j int) bool {
		return c.Level.Sections[i].Y > c.Level.Sections[j].Y
	})

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
