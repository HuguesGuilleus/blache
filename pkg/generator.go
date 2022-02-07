// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg/cpumutex"
	"github.com/HuguesGuilleus/blache/web"
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

func Generate(option Option) {
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

	root, files, err := g.getFiles()
	if err != nil {
		option.Error(fmt.Errorf("Read directory fail: %w", err))
		return
	}

	g.bar.Total += (32*32 + 1 + regionNumberOfImage) * int64(len(files))
	g.wg.Add(len(files))
	defer g.wg.Wait()

	g.cpu.Lock()
	defer g.cpu.Unlock()

	for _, f := range files {
		x, z := 0, 0
		if _, err := fmt.Sscanf(f.Name(), "r.%d.%d.mca", &x, &z); err != nil {
			g.fileFail("Error when read X end Z from file name %q: %w", f.Name(), err)
			continue
		}

		n := path.Join(root, f.Name())
		data, err := fs.ReadFile(option.In, n)
		if err != nil {
			g.fileFail("Fail to read %q: %w", n, err)
			continue
		} else if len(data) < 32*32*4 {
			g.fileFail("")
			continue
		}

		g.allRegion = append(g.allRegion, fmt.Sprintf("%d,%d", x, z))
		go parseRegion(&g, x, z, data)
	}
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

// When fail to read a file, Send an error to g.Error wiith format and ars if
// format is not empty.
func (g *generator) fileFail(format string, args ...interface{}) {
	g.bar.Add(32*32 + 1)
	g.wg.Done()

	if format != "" {
		g.Error(fmt.Errorf(format, args...))
	}
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

// Encode the Image in PNG and store it.
func (g *generator) saveImage(kind, name string, img *regionImage) {
	defer g.bar.Increment()

	img.processPalette()

	buff := bytes.Buffer{}
	png.Encode(&buff, img)

	if err := g.Out.Create(kind, name, buff.Bytes()); err != nil {
		g.Error(fmt.Errorf("Fail to save image: %v", err))
	}
}
