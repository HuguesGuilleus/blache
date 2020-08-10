// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"sync"
)

type region struct {
	X, Z   int
	g      *generator
	biome  *image.RGBA
	bloc   *image.RGBA
	height *image.RGBA
	// For waiting image generation from chunck
	wg      sync.WaitGroup
	structs []structure
}

type structure struct {
	X    int    `json:"x"`
	Z    int    `json:"z"`
	Name string `json:"name"`
}

/* PARSING */

func parseRegion(g *generator, x, z int, data []byte) {
	g.cpu.Lock()

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
			offset := 4 * (x + z*32)
			if bytesToInt(data[offset:offset+4]) == 0 {
				continue
			}
			addr := 4096 * (bytesToInt(data[offset : offset+3]))
			l := bytesToInt(data[addr : addr+4])
			if typeOfCompress := data[addr+4]; typeOfCompress != 2 {
				log.Print("Unknown compress (2):", typeOfCompress)
				continue
			}

			r.wg.Add(1)
			r.g.bar.Total += 1
			go r.addChunck(data[addr+5:addr+4+l], x, z)
		}
	}

	n := fmt.Sprintf("%d.%d.png", r.X, r.Z)
	r.g.wg.Add(4)
	g.cpu.Unlock()
	r.wg.Wait()

	go r.saveStructs()
	go r.g.saveImage("biome", n, r.biome)
	go r.g.saveImage("bloc", n, r.bloc)
	go r.g.saveImage("height", n, r.height)
	r.g.wg.Done()
}

func (r *region) addChunck(data []byte, x, z int) {
	r.g.cpu.Lock()
	defer r.g.cpu.Unlock()

	defer func() {
		r.wg.Done()
		r.g.bar.Increment()
	}()

	c, err := reginParseChunck(data)
	if err != nil {
		r.g.err <- err
		return
	}
	c.x = x
	c.z = z
	c.biome = subImage(r.biome, x, z)
	c.bloc = subImage(r.bloc, x, z)
	c.height = subImage(r.height, x, z)
	c.region = r

	for n := range c.Level.Structures.Starts {
		r.structs = append(r.structs, structure{
			X:    x,
			Z:    z,
			Name: n,
		})
	}

	c.draw()
}

func (r *region) saveStructs() {
	defer r.g.wg.Done()
	r.g.cpu.Lock()
	defer r.g.cpu.Unlock()

	var j []byte
	if len(r.structs) == 0 {
		j = []byte("[]")
	} else {
		j, _ = json.Marshal(r.structs)
	}
	r.g.Out.File("structs", fmt.Sprintf("%d.%d.json", r.X, r.Z), j)
}

// Decompress and parse a chunck
func reginParseChunck(data []byte) (*chunck, error) {
	// Decompress data
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// Parse minecraft data
	c := &chunck{}
	if err := nbt.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}

// Convert a slice of bytes to int
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
