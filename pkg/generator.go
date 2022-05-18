// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/web"
	"image/png"
	"io/fs"
	"path"
	"sort"
	"sync"
)

type generator struct {
	Option
	wg  sync.WaitGroup
	bar bar
	// All the region coords.
	allRegion []string

	// All errors occure durring generation.
	errorSlice []error
	errorMutex sync.Mutex
}

// Read minecraft data file from Option.Input and save image to Option.Output.
//
// Multiples errors can occure, so Generate can return multiple errors.
func Generate(option Option) []error {
	g := generator{Option: option}

	if err := g.initOutput(); err != nil {
		return []error{err}
	}

	if !option.NoBar {
		go g.bar.Start()
		defer g.bar.Finish()
	}

	root, files, err := g.getFiles()
	if err != nil {
		return []error{fmt.Errorf("Read directory fail: %w", err)}
	}

	g.bar.Total += (32*32 + 1 + regionNumberOfImage) * int64(len(files))
	g.wg.Add(len(files))

	for _, f := range files {
		x, z := 0, 0
		if _, err := fmt.Sscanf(f.Name(), "r.%d.%d.mca", &x, &z); err != nil {
			g.fileFail("Error when read X end Z from file name %q: %w", f.Name(), err)
			continue
		}

		n := path.Join(root, f.Name())
		data, err := fs.ReadFile(option.Input, n)
		if err != nil {
			g.fileFail("Fail to read %q: %w", n, err)
			continue
		} else if len(data) < 32*32*4 {
			g.fileFail("The file for region %d,%d is too short", x, z)
			continue
		}

		g.allRegion = append(g.allRegion, fmt.Sprintf("%d,%d", x, z))
		go parseRegion(&g, x, z, data)
	}
	g.saveRegionsList()
	g.wg.Wait()

	return g.errorSlice
}

// Write directory and assets.
func (g *generator) initOutput() error {
	for _, d := range [...]string{"bloc", "biome", "height", "structs"} {
		if err := g.Output.MkdirAll(d); err != nil {
			return fmt.Errorf("Write directory %q fail: %w", d, err)
		}
	}

	for _, f := range web.List {
		if err := g.Output.Create("", f.Name, f.Data); err != nil {
			return fmt.Errorf("Fail to write web file %q: %v", f.Name, err)
		}
	}

	return nil
}

// When fail to read a file, Send an error to g.Error with format and args if
// format is not empty.
func (g *generator) fileFail(format string, args ...interface{}) {
	g.bar.Add(32*32 + 1)
	g.wg.Done()

	if format != "" {
		g.addError(fmt.Errorf(format, args...))
	}
}

// Save all the processed region coordonates into regions.json
func (g *generator) saveRegionsList() {
	sort.Strings(g.allRegion)
	data, err := json.Marshal(g.allRegion)
	if err != nil {
		g.addError(fmt.Errorf("Generated JSON regions fail: %w", err))
		return
	}
	if err := g.Output.Create("", "regions.json", data); err != nil {
		g.addError(fmt.Errorf("Write regions.json fail: %w", err))
	}
}

// Encode the Image in PNG and store it.
func (g *generator) saveImage(kind, name string, img *regionImage) {
	defer g.bar.Increment()

	img.processPalette()

	buff := bytes.Buffer{}
	png.Encode(&buff, img)

	if err := g.Output.Create(kind, name, buff.Bytes()); err != nil {
		g.addError(fmt.Errorf("Fail to save image: %v", err))
	}
}

// Add an error in errorSlice. Can be call concurently.
func (g *generator) addError(err error) {
	g.errorMutex.Lock()
	defer g.errorMutex.Unlock()
	g.errorSlice = append(g.errorSlice, err)
}
