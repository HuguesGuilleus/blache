// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/cpumutex"
	"github.com/HuguesGuilleus/blache/web"
	"image"
	"image/png"
	"io/fs"
	"path"
	"sort"
	"sync"
)

type generator struct {
	Option
	cpu cpumutex.M
	wg  sync.WaitGroup
	bar bar
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
		go g.bar.Start()
		defer g.bar.Finish()
	}

	var (
		root  string
		files []fs.DirEntry
		err   error
	)
	for _, p := range []string{"world/region", "region", "."} {
		files, err = fs.ReadDir(option.In, p)
		if err == nil {
			root = p
			break
		}
	}
	if err != nil {
		option.Error(fmt.Errorf("Read directory fail: %w", err))
		return
	}

	g.bar.Total += (32*32 + 1) * int64(len(files))
	g.cpu.Lock()
	for _, f := range files {
		n := path.Join(root, f.Name())
		data, err := fs.ReadFile(option.In, n)
		if err != nil {
			option.Error(fmt.Errorf("Fail to read %q: %w", n, err))
			break
		}

		x, z := 0, 0
		if _, err := fmt.Sscanf(f.Name(), "r.%d.%d.mca", &x, &z); err != nil {
			g.Error(fmt.Errorf("Error when read X end Z from file name %q: %w", f.Name(), err))
			continue
		}

		if len(data) < 32*32*4 {
			g.bar.Add(32*32 + 1)
			continue
		}

		g.wg.Add(1)
		go parseRegion(&g, x, z, data)
		g.allRegion = append(g.allRegion, fmt.Sprintf("%d,%d", x, z))
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

	for _, f := range web.List {
		if err := g.Out.Create("", f.Name, f.Data); err != nil {
			return fmt.Errorf("Fail to write web file %q: %v", f.Name, err)
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
