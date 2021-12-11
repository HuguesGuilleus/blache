// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"compress/zlib"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
	"github.com/Tnze/go-mc/nbt"
	"image/color"
	"sort"
)

// On column of 16*16*256 blocks.
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

func (r *region) drawChunck(data []byte, x, z int) error {
	// Decompress data and parse minecraft data
	c := chunck{}
	if reader, err := zlib.NewReader(bytes.NewReader(data)); err != nil {
		return err
	} else if err := nbt.NewDecoder(reader).Decode(&c); err != nil {
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
	if err := r.biome.draw(x, z, c.Level.Biomes); err != nil {
		return err
	}

	// Draw bloc and height tiles.
	palette := c.genPalette()
	bloc := subImage(r.bloc, x, z)
	height := subImage(r.height, x, z)
	for x := 0; x < 16; x++ {
	nextBloc:
		for z := 0; z < 16; z++ {
			for _, sec := range c.Level.Sections {
				size := 64 * len(sec.BlockStates) / 4096
				m := uint64((1 << size) - 1)
				nb := 64 / size
				colors := palette[sec.Y]
				for y := 15; -1 < y; y-- {
					p := y*256 + z*16 + x
					if p/nb >= len(sec.BlockStates) {
						continue
					}
					i := uint64(sec.BlockStates[p/nb]>>(p%nb*size)) & m
					if col := colors[i]; col.A == 0xFF {
						bloc(x, z, col)
						h := sec.Y*16 + uint8(y)
						height(x, z, color.RGBA{h, h, h, 0xFF})
						continue nextBloc
					}
				}
			}
		}
	}

	return nil
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
