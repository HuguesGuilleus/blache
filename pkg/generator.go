// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"../web/webData"
	"./cpumutex"
	"encoding/json"
	"fmt"
	"gopkg.in/cheggaaa/pb.v3"
	"io"
	"sort"
	"sync"
)

type generator struct {
	Option
	region cpumutex.M
	chunck cpumutex.M
	wg     sync.WaitGroup
	err    chan<- error
	bar    pb.ProgressBar
	// All the regions coord.
	allRegion []string
}

func (option Option) Gen() <-chan error {
	ch := make(chan error)
	g := generator{
		Option: option,
		err:    ch,
	}
	g.region.Init(option.CPU)
	g.chunck.Init(option.CPU)

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
		g.bar.Start()
		defer g.bar.Finish()

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
