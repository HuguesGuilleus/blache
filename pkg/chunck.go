// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"compress/zlib"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
	"github.com/Tnze/go-mc/nbt"
	"io"
	"sort"
)

// On column of 16*16*256 block = 16 section.
type chunck struct {
	Level struct {
		Biomes     interface{}
		Sections   []section
		Structures struct {
			Starts map[string]struct{}
		}
	}
}

// One section, 16*16*16 block.
type section struct {
	Y       uint8
	Palette []struct {
		Name string
	}
	BlockStates []int64
}

func drawChunck(r *region, raw []byte, x, z int) error {
	// Decompress data and parse minecraft data
	c := chunck{}
	r.buff.Reset()
	if zlibReader, err := zlib.NewReader(bytes.NewReader(raw)); err != nil {
		return err
	} else if _, err := io.Copy(&r.buff, zlibReader); err != nil {
		return err
	} else if err := nbt.Unmarshal(r.buff.Bytes(), &c); err != nil {
		return err
	}

	// Save structure
	for n := range c.Level.Structures.Starts {
		r.structs = append(r.structs, structure{
			X:    x,
			Z:    z,
			Name: n,
		})
	}

	// Draw biome tile.
	if err := r.biome.drawBiome(x, z, c.Level.Biomes); err != nil {
		return err
	}

	drawBlockAndHeight(&c, r.bloc.chunck(x, z), r.height.chunck(x, z))

	return nil
}

// Draw bloc and height tiles.
func drawBlockAndHeight(c *chunck, bloc, height []uint8) {
	palette := c.genPalette()
	for x := 0; x < 16; x++ {
	nextBloc:
		for z := 0; z < 16; z++ {
			for _, sec := range c.Level.Sections {
				size := 64 * len(sec.BlockStates) / 4096
				mask := uint64((1 << size) - 1)
				nb := 64 / size
				colors := palette[sec.Y]
				for y := 15; -1 < y; y-- {
					p := y*256 + z*16 + x
					if p/nb >= len(sec.BlockStates) {
						continue
					}
					i := uint64(sec.BlockStates[p/nb]>>(p%nb*size)) & mask
					if col := colors[i]; col != 0 {
						bloc[z*16+x] = col
						h := sec.Y*16 + uint8(y)
						height[z*16+x] = h
						continue nextBloc
					}
				}
			}
		}
	}
}

// Generate the palette
func (c *chunck) genPalette() (p [16][]uint8) {
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
		p[sec.Y] = make([]uint8, 1<<(64*len(sec.BlockStates)/4096))
		for i, b := range sec.Palette {
			p[sec.Y][i] = minecraftColor.BlocColorIndex(b.Name)
		}
	}

	return p
}
