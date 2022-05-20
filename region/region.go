// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package region

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/region/minecraftColor"
)

type Region struct {
	Biome  Image
	Bloc   Image
	Height Image
	Water  Image

	// Founded structures
	Structures []Structure
}

type Structure struct {
	X    int    `json:"x"`
	Z    int    `json:"z"`
	Name string `json:"name"`
}

func New(data []byte) (r *Region, errList []error) {
	r = &Region{
		Biome:  Image{palette: minecraftColor.BiomePalette},
		Bloc:   Image{palette: minecraftColor.BlockPalette},
		Height: Image{palette: minecraftColor.HeightPalette},
		Water:  Image{palette: minecraftColor.WaterPalette},
	}

	buff := bytes.Buffer{}
	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			// Get the chunk data position into data.
			offset := 4 * (x + z*32)
			if bytesToInt(data[offset:offset+4]) == 0 {
				continue
			}
			addr := 4096 * bytesToInt(data[offset:offset+3])
			l := bytesToInt(data[addr : addr+4])
			if typeOfCompress := data[addr+4]; typeOfCompress != 2 {
				errList = append(errList, fmt.Errorf("Chunck (%d,%d): Unknown compress, expected 2, found %d", x, z, typeOfCompress))
				continue
			}
			if err := parseChunck(r, x, z, data[addr+5:addr+4+l], &buff); err != nil {
				errList = append(errList, fmt.Errorf("Chunck:(%d,%d): %w", x, z, err))
			}
		}
	}

	return r, nil
}

// Get the list of the structure encoded in JSON.
func (r *Region) StructuresJSON() []byte {
	j := []byte("[]")
	if len(r.Structures) > 0 {
		j, _ = json.Marshal(r.Structures)
	}
	return j
}

// Convert a slice of bytes to int, with bigendian.
func bytesToInt(tab []byte) (n int) {
	for _, b := range tab {
		n = n<<8 | int(b)
	}
	return n
}
