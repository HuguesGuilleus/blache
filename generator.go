// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"encoding/json"
	"fmt"
	"github.com/HuguesGuilleus/blache/web"
	"io"
	"io/fs"
	"path"
	"sort"
	"sync"
	"time"
)

type generator struct {
	Option
	wg sync.WaitGroup

	// The region already processed.
	regionDone int64
	// Number of all regions not cached.
	regionTotal int64

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
	if g.LogOutput == nil {
		g.LogOutput = io.Discard
	}

	defer func(before time.Time) {
		fmt.Println("duration:", time.Since(before).Round(time.Millisecond))
	}(time.Now())

	if err := g.writeAssets(); err != nil {
		return []error{err}
	}

	root, files, err := g.getFiles()
	if err != nil {
		return []error{fmt.Errorf("Read directory fail: %w", err)}
	}
	for _, f := range files {
		g.readRegion(root, f)
	}

	g.saveRegionsList()

	g.wg.Wait()
	return g.errorSlice
}

// Write directory and assets.
func (g *generator) writeAssets() error {
	for _, d := range [...]string{"bloc", "biome", "height", "structs"} {
		if err := g.Output.MkdirAll(d); err != nil {
			return fmt.Errorf("Write directory %q fail: %w", d, err)
		}
	}

	for _, f := range web.List {
		g.saveFile("", f.Name, f.Data)
	}

	return nil
}

func (g *generator) readRegion(root string, entry fs.DirEntry) {
	x, z := 0, 0
	name := path.Join(root, entry.Name())
	if _, err := fmt.Sscanf(entry.Name(), "r.%d.%d.mca", &x, &z); err != nil {
		g.addError(fmt.Errorf("Error when read X end Z from file name %q: %w", name, err))
		return
	}

	outputInfo, _ := fs.Stat(g.Output, fmt.Sprintf("structs/%d.%d.json", x, z))
	inputInfo, _ := entry.Info()
	if outputInfo != nil && inputInfo != nil && outputInfo.ModTime().After(inputInfo.ModTime()) {
		fmt.Fprintf(g.LogOutput, "cache region (%d,%d)\n", x, z)
		g.allRegion = append(g.allRegion, fmt.Sprintf("%d,%d", x, z))
		return
	}

	data, err := fs.ReadFile(g.Input, name)
	if err != nil {
		g.addError(fmt.Errorf("Fail to read %q: %w", name, err))
		return
	} else if len(data) < 32*32*4 {
		g.addError(fmt.Errorf("The file for region %d,%d is too short", x, z))
		return
	}

	g.wg.Add(1)
	g.regionTotal++
	g.allRegion = append(g.allRegion, fmt.Sprintf("%d,%d", x, z))
	go parseRegion(g, x, z, data)
}

// Save all the processed region coordonates into regions.json
func (g *generator) saveRegionsList() {
	sort.Strings(g.allRegion)
	data, err := json.Marshal(g.allRegion)
	if err != nil {
		g.addError(fmt.Errorf("Generated JSON regions fail: %w", err))
		return
	}
	g.saveFile("", "regions.json", data)
}

// Save the file and manage error.
func (g *generator) saveFile(dir, name string, data []byte) {
	if err := g.Output.Create(dir, name, data); err != nil {
		g.addError(fmt.Errorf("Fail to save: %s/%s: %w", dir, name, err))
	}
}

// Add an error in errorSlice. Can be call concurently.
func (g *generator) addError(err error) {
	g.errorMutex.Lock()
	defer g.errorMutex.Unlock()
	g.errorSlice = append(g.errorSlice, err)
}