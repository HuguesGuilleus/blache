// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
)

// Draw the image
func (img *regionImage) drawBiome(chunckX, chunckZ int, biome interface{}) error {
	chunck := img.chunck(chunckX, chunckZ)
	switch biome := biome.(type) {
	case nil:
		fillBiomChunck(chunck)
	case []byte:
		switch l := len(biome); l {
		case 256:
			copy(chunck, biome)
		case 0:
			fillBiomChunck(chunck)
		default:
			fillBiomChunck(chunck)
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
			fillBiomChunck(chunck)
		default:
			fillBiomChunck(chunck)
			return fmt.Errorf("[]int32 length is not 2565 or 1024, it't: %d", len(biome))
		}
	default:
		fillBiomChunck(chunck)
		return fmt.Errorf("The biome is %T (expected byte or int32 array, or nothing)", biome)
	}
	return nil
}

func fillBiomChunck(chunck []uint8) {
	for i := range chunck {
		chunck[i] = minecraftColor.BiomeBlackIndex
	}
}
