// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
)

type region struct {
	X, Z    int
	g       *generator
	biome   *image.RGBA
	bloc    *image.RGBA
	height  *image.RGBA
	structs []structure
}

type structure struct {
	X    int    `json:"x"`
	Z    int    `json:"z"`
	Name string `json:"name"`
}

/* PARSING */

func parseRegion(g *generator, x, z int, data []byte) {
	defer func() {
		switch err := recover(); err.(type) {
		case nil:
		case error:
			g.Error(fmt.Errorf("Panic error (region: %d,%d): %w", x, z, err))
		default:
			g.Error(fmt.Errorf("Panic error (region: %d,%d): %v", x, z, err))
		}
	}()

	g.cpu.Lock()
	defer g.cpu.Unlock()
	defer g.wg.Done()

	r := region{
		g:      g,
		X:      x,
		Z:      z,
		biome:  image.NewRGBA(image.Rect(0, 0, 32*16, 32*16)),
		bloc:   image.NewRGBA(image.Rect(0, 0, 32*16, 32*16)),
		height: image.NewRGBA(image.Rect(0, 0, 32*16, 32*16)),
	}

	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			r.g.bar.Increment()
			// Get the chunk data position into data.
			offset := 4 * (x + z*32)
			if bytesToInt(data[offset:offset+4]) == 0 {
				continue
			}
			addr := 4096 * bytesToInt(data[offset:offset+3])
			l := bytesToInt(data[addr : addr+4])
			if typeOfCompress := data[addr+4]; typeOfCompress != 2 {
				g.Error(fmt.Errorf("In region (%d,%d), in chunck (%d,%d) Unknown compress, expected 2, found %d", r.X, r.Z, x, z, typeOfCompress))
				continue
			}
			if err := r.drawChunck(data[addr+5:addr+4+l], x, z); err != nil {
				g.Error(fmt.Errorf("In region (%d,%d), in chunck (%d,%d) %w", r.X, r.Z, x, z, err))
			}
		}
	}
	r.g.bar.Increment()

	r.g.wg.Add(4)
	n := fmt.Sprintf("%d.%d.png", r.X, r.Z)
	go r.g.saveImage("biome", n, r.biome)
	go r.g.saveImage("bloc", n, r.bloc)
	go r.g.saveImage("height", n, r.height)
	go r.saveStructs()
}

// Save the list of the structure in JSON.
func (r *region) saveStructs() {
	defer r.g.wg.Done()
	r.g.cpu.Lock()
	defer r.g.cpu.Unlock()

	var j []byte
	if len(r.structs) == 0 {
		j = []byte("[]")
	} else {
		var err error
		j, err = json.Marshal(r.structs)
		if err != nil {
			r.g.Error(fmt.Errorf("Chunck (%d,%d), JSON structures list genration fail: %w", r.X, r.Z, err))
			return
		}
	}
	n := fmt.Sprintf("%d.%d.json", r.X, r.Z)
	if err := r.g.Out.Create("structs", n, j); err != nil {
		r.g.Error(fmt.Errorf("Write list of structure file %q fail: %w", n, err))
	}
}

// Convert a slice of bytes to int, with bigendian.
func bytesToInt(tab []byte) (n int) {
	for _, b := range tab {
		n = n<<8 + int(b)
	}
	return n
}

/* IMAGE */

// A function to set a pixel with a predefined offset.
type imgSetRGBA func(x, z int, c color.RGBA)

func subImage(img *image.RGBA, chunckX, chunckZ int) imgSetRGBA {
	return func(x, z int, c color.RGBA) {
		img.SetRGBA(x+16*chunckX, z+16*chunckZ, c)
	}
}
