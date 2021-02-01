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
	bar pb.ProgressBar
	// All the region coords.
	allRegion []string
}

func (option Option) Gen() {
	if option.Error == nil {
		option.Error = func(error) {}
	}
	g := generator{Option: option}
	g.cpu.Init(option.CPU)

	if err := g.initOutput(); err != nil {
		g.Error(err)
		return
	}

	if !option.NoBar {
		g.bar = *pb.New(0)
		g.bar.Format("[=> ]")
		g.bar.SetRefreshRate(time.Millisecond * 100)
		g.bar.Start()
		defer g.bar.Finish()
	}

	// TODO: rework input with fs from go1.16
	if err := g.In.Open(); err != nil {
		g.Error(fmt.Errorf("Open() error: %w", err))
		return
	}
	g.cpu.Lock()
	for {
		n, data, err := g.In.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			g.Error(fmt.Errorf("Read next file %q fail %v", n, err))
			continue
		}

		x, z := 0, 0
		if _, err := fmt.Sscanf(n, "r.%d.%d.mca", &x, &z); err != nil {
			g.Error(fmt.Errorf("Error when read X end Z from file name %q: %w", n, err))
			continue
		}

		g.wg.Add(1)
		go parseRegion(&g, x, z, data)
		g.allRegion = append(g.allRegion, fmt.Sprintf("(%d,%d)", x, z))
	}
	g.cpu.Unlock()

	g.wg.Wait()
	g.saveRegionsList()
}

// Write directory and assets.
func (g *generator) initOutput() error {
	for _, d := range [...]string{"bloc", "biome", "height", "structs"} {
		if err := g.Out.MkdirAll(d); err != nil {
			return fmt.Errorf("Write directory %q fail: %w", d, err)
		}
	}

	a, _ := webData.AssetDir("web")
	for _, n := range a {
		if err := g.Out.Create("", n, webData.MustAsset("web/"+n)); err != nil {
			return fmt.Errorf("Write assets %q fail: %w", n, err)
		}
	}

	return nil
}

// Save all the processed region coordonates into regions.json
func (g *generator) saveRegionsList() {
	sort.Strings(g.allRegion)
	data, err := json.Marshal(g.allRegion)
	if err != nil {
		g.Error(fmt.Errorf("Generated JSON regions fail: %w", err))
		return
	}
	if err := g.Out.Create("", "regions.json", data); err != nil {
		g.Error(fmt.Errorf("Write regions.json fail: %w", err))
	}
}

// Save an image of one region.
func (g *generator) saveImage(dir, f string, img *image.RGBA) {
	g.cpu.Lock()
	defer g.cpu.Unlock()
	defer g.wg.Done()

	buff := bytes.Buffer{}
	png.Encode(&buff, img)

	if err := g.Out.Create(dir, f, buff.Bytes()); err != nil {
		g.Error(fmt.Errorf("save image error: %v", err))
	}
}
