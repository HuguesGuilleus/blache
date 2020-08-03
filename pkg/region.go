// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

type region struct {
	X, Z   int
	g      *generator
	file   string // The MCA file
	biome  *image.RGBA
	bloc   *image.RGBA
	height *image.RGBA
	// For waiting image generation from chunck
	wg sync.WaitGroup
}

/* PARSING */

func (r *region) parse() {
	_, f := filepath.Split(r.file)
	if _, err := fmt.Sscanf(f, "r.%d.%d.mca", &r.X, &r.Z); err != nil {
		r.g.err <- fmt.Errorf("Error when read X end Z from file name: %v", err)
		r.g.nbChunckOk += 1024
		return
	}

	data, err := ioutil.ReadFile(r.file)
	if err != nil {
		log.Println("[ERROR] read:", r.file, err)
		r.g.nbChunckOk += 1024
		return
	}

	r.biome = image.NewRGBA(image.Rect(0, 0, 32*16, 32*16))
	r.bloc = image.NewRGBA(image.Rect(0, 0, 32*16, 32*16))
	r.height = image.NewRGBA(image.Rect(0, 0, 32*16, 32*16))

	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			offset := 4 * (x + z*32)
			if bytesToInt(data[offset:offset+4]) == 0 {
				r.g.nbChunckOk++
				continue
			}
			addr := 4096 * (bytesToInt(data[offset : offset+3]))
			l := bytesToInt(data[addr : addr+4])
			if typeOfCompress := data[addr+4]; typeOfCompress != 2 {
				log.Print("Unknown compress (2):", typeOfCompress)
				r.g.nbChunckOk++
				continue
			}

			r.wg.Add(1)
			go r.addChunck(data[addr+5:addr+4+l], x, z)
		}
	}

	r.wg.Wait()
	r.imgSave()
}

func (r *region) addChunck(data []byte, x, z int) {
	r.g.chunck.Lock()
	defer r.g.chunck.Unlock()

	defer func() {
		r.wg.Done()
		r.g.nbChunckOk++
	}()

	c, err := reginParseChunck(data)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	c.x = x
	c.z = z
	c.biome = subImage(r.biome, x, z)
	c.bloc = subImage(r.bloc, x, z)
	c.height = subImage(r.height, x, z)
	c.region = r

	c.draw()
}

// Decompress and parse a chunck
func reginParseChunck(data []byte) (c *chunck, _ error) {
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
	c = &chunck{}
	return c, nbt.Unmarshal(data, c)
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

// Save the images.
func (r *region) imgSave() {
	n := fmt.Sprintf("%d.%d.png", r.X, r.Z)
	r.g.wg.Add(3)
	go saveImage(filepath.Join(r.g.Out, "biome", n), r.biome, &r.g.wg)
	go saveImage(filepath.Join(r.g.Out, "bloc", n), r.bloc, &r.g.wg)
	go saveImage(filepath.Join(r.g.Out, "height", n), r.height, &r.g.wg)
}

// A mutex for limit number of open image file.
var limitFile sync.Mutex

func saveImage(f string, img image.Image, wg *sync.WaitGroup) {
	defer wg.Done()

	buff := &bytes.Buffer{}
	png.Encode(buff, img)

	limitFile.Lock()
	defer limitFile.Unlock()

	err := ioutil.WriteFile(f, buff.Bytes(), 0664)
	if err != nil {
		log.Printf("[ERROR] on save '%s': %v\n", f, err)
	}
}
