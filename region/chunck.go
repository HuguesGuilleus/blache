// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package region

import (
	"bytes"
	"compress/zlib"
	"github.com/HuguesGuilleus/blache/region/chunck"
	"github.com/HuguesGuilleus/blache/region/minecraftColor"
	"io"
	"sort"
)

func parseChunck(r *Region, x, z int, raw []byte, buff *bytes.Buffer) error {
	// Decompress data and parse minecraft data
	c := chunck.Chunck{}
	buff.Reset()
	if zlibReader, err := zlib.NewReader(bytes.NewReader(raw)); err != nil {
		return err
	} else if _, err := io.Copy(buff, zlibReader); err != nil {
		return err
	} else if err := c.DecodeNBT(buff.Bytes()); err != nil {
		return err
	}

	for n := range c.Level.Structures.Starts {
		r.Structures = append(r.Structures, Structure{
			X:    x,
			Z:    z,
			Name: n,
		})
	}

	if err := drawBiome(r.Biome.chunck(x, z), c.Level.Biomes); err != nil {
		return err
	}

	drawBlockAndHeight(&c, r.Bloc.chunck(x, z), r.Height.chunck(x, z), r.Water.chunck(x, z))

	return nil
}

// Draw bloc and height tiles.
func drawBlockAndHeight(c *chunck.Chunck, bloc, height, water []uint8) {
	palette := genPalette(c)
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
					switch col := colors[i]; col {
					case 0:
					case minecraftColor.WaterIndex:
						water[z*16+x] = minecraftColor.HasWater
					default:
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
func genPalette(c *chunck.Chunck) (p [16][]uint8) {
	secs := make([]chunck.Section, 0, len(c.Level.Sections))
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
