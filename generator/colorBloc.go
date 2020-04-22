package generator

import (
	"archive/zip"
	"github.com/nfnt/resize"
	"image/color"
	"image/png"
	"log"
	"regexp"
	"sync"
)

var colorBlocksInZip = regexp.MustCompile(`assets/minecraft/textures/blocks/(\w+)\.png`)

// The color from a Data Pack.
type colorBloc struct {
	sync.RWMutex
	m map[string]color.RGBA
}

func (c *colorBloc) Load(pathZip string) error {
	if c.m == nil {
		c.m = make(map[string]color.RGBA)
	}

	dataPack, err := zip.OpenReader(pathZip)
	if err != nil {
		return err
	}
	defer dataPack.Close()

	wg := sync.WaitGroup{}
	defer wg.Wait()
	for _, file := range dataPack.File {
		if !colorBlocksInZip.MatchString(file.Name) {
			continue
		}

		wg.Add(1)
		go c.addColor(file, &wg)
	}

	return nil
}

func (c *colorBloc) addColor(file *zip.File, wg *sync.WaitGroup) {
	defer wg.Done()

	id := colorBlocksInZip.ReplaceAllString(file.Name, "minecraft:$1")

	reader, err := file.Open()
	if err != nil {
		log.Printf("[ERROR IN ZIP] file='%s'; %v\n", file.Name, err)
		return
	}
	defer reader.Close()

	img, err := png.Decode(reader)
	if err != nil {
		return
	}
	r, g, b, _ := resize.Resize(1, 1, img, resize.Bilinear).At(0, 0).RGBA()

	c.Lock()
	defer c.Unlock()

	c.m[id] = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
