// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package blache

import (
	"fmt"
	"github.com/HuguesGuilleus/blache/region"
	"io/fs"
	"path"
)

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
	go g.parseRegion(x, z, data)
}

// Ask to region.New() to draw image and found structure from the data region.
func (g *generator) parseRegion(x, z int, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				g.addError(fmt.Errorf("Panic error (region: %d,%d): %w", x, z, err))
			} else {
				g.addError(fmt.Errorf("Panic error (region: %d,%d): %v", x, z, err))
			}
		}
	}()

	defer func() {
		g.wg.Done()
		g.regionDone++
		fmt.Fprintf(g.LogOutput, "generation: %3d%% region (%d,%d) done\n", g.regionDone*100/g.regionTotal, x, z)
	}()

	r, errorList := region.New(data)
	for _, err := range errorList {
		g.addError(fmt.Errorf("Region (%d,%d): %w", x, z, err))
	}

	imageName := fmt.Sprintf("%d.%d.png", x, z)
	g.saveFile("biome", imageName, r.Biome.BytesPNG())
	g.saveFile("bloc", imageName, r.Bloc.BytesPNG())
	g.saveFile("height", imageName, r.Height.BytesPNG())
	g.saveFile("structs", fmt.Sprintf("%d.%d.json", x, z), r.StructuresJSON())
}
