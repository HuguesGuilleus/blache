package generator

import (
	"./biome"
)

type chunck struct {
	region *region
	x, z   int
	biome  imgSetRGBA

	// Minecraft data
	Level struct {
		Biomes interface{} // Minecraft biomes
		biomes [256]byte   // standard biomes
	}
}

func (c *chunck) draw() {
	if c.setBiome() {
		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				c.biome(x, z, biome.Color[c.Level.biomes[z*16+x]])
			}
		}
	}
}

// The the biome standard.
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
