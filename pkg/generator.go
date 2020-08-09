// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"../web/webData"
	"./cpumutex"
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/cheggaaa/pb.v3"
	"image"
	"image/png"
	"io"
	"sort"
	"sync"
	"time"
)

type generator struct {
	Option
	cpu cpumutex.M
	wg  sync.WaitGroup
	err chan<- error
	bar pb.ProgressBar
	// All the region coords.
	allRegion []string
}

func (option Option) Gen() <-chan error {
	ch := make(chan error)
	g := generator{
		Option: option,
		err:    ch,
	}
	g.cpu.Init(option.CPU)

	for _, d := range [...]string{"bloc", "biome", "height"} {
		g.Out.Dir(d)
	}
	g.makeAssets()

	go func() {
		defer close(ch)

		if err := g.In.Open(); err != nil {
			g.err <- fmt.Errorf("Open() error: %v", err)
			return
		}

		g.bar = *pb.New(0)
		g.bar.Format("[=> ]")
		g.bar.Prefix("chuncks:")
		g.bar.RefreshRate = time.Millisecond * 50
		g.bar.Start()
		defer g.bar.Finish()

		g.cpu.Lock()
		for {
			n, data, err := g.In.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				g.err <- err
				continue
			}

			x, z := 0, 0
			if _, err := fmt.Sscanf(n, "r.%d.%d.mca", &x, &z); err != nil {
				g.err <- fmt.Errorf("Error when read X end Z from file name %q: %v", n, err)
				continue
			}

			g.wg.Add(1)
			go parseRegion(&g, x, z, data)
			g.allRegion = append(g.allRegion, fmt.Sprintf("(%d,%d)", x, z))
		}
		g.cpu.Unlock()

		g.wg.Wait()
		g.saveRegionsList()
	}()

	return ch
}

// Save all the processed region coordonates into regions.json
func (g *generator) saveRegionsList() {
	sort.Strings(g.allRegion)
	data, err := json.Marshal(g.allRegion)
	if err != nil {
		g.err <- err
		return
	}
	g.Out.File("", "regions.json", data)
}

// Write the web assets.
func (g *generator) makeAssets() {
	a, _ := webData.AssetDir("web")
	for _, n := range a {
		g.Out.File("", n, webData.MustAsset("web/"+n))
	}
}

// Save an image of one region.
func (g *generator) saveImage(dir, f string, img *image.RGBA) {
	g.cpu.Lock()
	defer g.cpu.Unlock()
	defer g.wg.Done()

	buff := bytes.Buffer{}
	png.Encode(&buff, img)

	if err := g.Out.File(dir, f, buff.Bytes()); err != nil {
		g.err <- fmt.Errorf("save image error: %v", err)
	}
}
