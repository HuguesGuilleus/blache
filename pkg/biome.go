// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
)

// Draw biome from Chunk.Level.Biomes with c.biome image closure.
func (c *chunck) drawBiome(biome imgSetRGBA, b interface{}) error {
	switch b := b.(type) {
	case nil:
	case []byte:
		if l := len(b); l == 0 {
			return nil
		} else if l != 256 {
			return fmt.Errorf("Biome is bytes array with %d length (expected 256)", l)
		}
		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				biome(x, z, minecraftColor.Biome[b[z*16+x]])
			}
		}
	case []int32:
		switch len(b) {
		case 0:
		case 256:
			for x := 0; x < 16; x++ {
				for z := 0; z < 16; z++ {
					biome(x, z, minecraftColor.Biome[b[z*16+x]&0xFF])
				}
			}
		case 1024:
			for x := 0; x < 16; x += 4 {
				for z := 0; z < 16; z += 4 {
					c := minecraftColor.Biome[b[0x3F0|z|x>>2]]
					for i := 0; i < 4; i++ {
						for j := 0; j < 4; j++ {
							biome(x+i, z+j, c)
						}
					}
				}
			}
		default:
			return fmt.Errorf("[]int32 length is not 2565 or 1024, it't: %d", len(b))
		}
	default:
		return fmt.Errorf("The biome is %T (expected byte or int32 array, or nothing)", b)
	}
	return nil
}
