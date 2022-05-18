// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/minecraftColor"
)

type region struct {
	X, Z    int
	g       *generator
	biome   regionImage
	bloc    regionImage
	height  regionImage
	structs []structure
	// To uncompress chunck data
	buff bytes.Buffer
}

type structure struct {
	X    int    `json:"x"`
	Z    int    `json:"z"`
	Name string `json:"name"`
}

func parseRegion(g *generator, x, z int, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				g.addError(fmt.Errorf("Panic error (region: %d,%d): %w", x, z, err))
			} else {
				g.addError(fmt.Errorf("Panic error (region: %d,%d): %v", x, z, err))
			}
		}
	}()

	defer func() {
		g.wg.Done()
		g.regionDone++
		fmt.Fprintf(g.LogOutput, "generation: %3d%% region (%d,%d) done\n", g.regionDone*100/g.regionTotal, x, z)
	}()

	r := region{
		g:      g,
		X:      x,
		Z:      z,
		biome:  regionImage{palette: minecraftColor.BiomePalette},
		bloc:   regionImage{palette: minecraftColor.BlockPalette},
		height: regionImage{palette: minecraftColor.HeightPalette},
	}

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
				g.addError(fmt.Errorf("Error region:(%d,%d) chunck:(%d,%d): Unknown compress, expected 2, found %d", r.X, r.Z, x, z, typeOfCompress))
				continue
			}
			if err := drawChunck(&r, data[addr+5:addr+4+l], x, z); err != nil {
				g.addError(fmt.Errorf("Error region:(%d,%d) chunck:(%d,%d): %w", r.X, r.Z, x, z, err))
			}
		}
	}

	name := fmt.Sprintf("%d.%d.png", r.X, r.Z)
	g.saveFile("biome", name, r.biome.BytesPNG())
	g.saveFile("bloc", name, r.bloc.BytesPNG())
	g.saveFile("height", name, r.height.BytesPNG())
	r.saveStructs()
}

// Save the list of the structure in JSON.
func (r *region) saveStructs() {
	j := []byte("[]")
	if len(r.structs) > 0 {
		var err error
		j, err = json.Marshal(r.structs)
		if err != nil {
			r.g.addError(fmt.Errorf("Chunck (%d,%d), JSON structures list genration fail: %w", r.X, r.Z, err))
			return
		}
	}
	r.g.saveFile("structs", fmt.Sprintf("%d.%d.json", r.X, r.Z), j)
}

// Convert a slice of bytes to int, with bigendian.
func bytesToInt(tab []byte) (n int) {
	for _, b := range tab {
		n = n<<8 | int(b)
	}
	return n
}
