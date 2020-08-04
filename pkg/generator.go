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
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
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
	err := make(chan error)
	g := generator{
		Option: option,
		err:    err,
		bar:    *pb.New(0),
	}
	g.region.Init(option.CPU)
	g.chunck.Init(option.CPU)
	g.bar.Format("[=> ]")
	g.bar.Prefix("chuncks:")
	g.bar.Start()

	for _, d := range [...]string{"bloc", "biome", "height"} {
		g.Out.Dir(d)
	}
	g.makeAssets()

	for r := range g.listRegion() {
		g.wg.Add(1)
		go g.addRegion(r)
	}

	go func() {
		g.wg.Wait()
		g.saveRegionsList()
		close(err)
	}()
	return err
}

// List the regions from MAC file into Option.Regions
func (g *generator) listRegion() (r chan string) {
	r = make(chan string, 2)

	go func() {
		defer close(r)

		f, err := ioutil.ReadDir(g.Regions)
		if err != nil {
			log.Println("[ERROR] read ", g.Regions, err)
			return
		}

		for _, f := range f {
			n := f.Name()
			if f.IsDir() || !strings.HasSuffix(n, ".mca") {
				continue
			}
			r <- filepath.Join(g.Regions, n)
		}
	}()

	return
}

// Add a new region
func (g *generator) addRegion(file string) {
	g.bar.Total += int64(1024)
	defer g.wg.Done()

	g.region.Lock()
	defer g.region.Unlock()

	r := region{
		file: file,
		g:    g,
	}
	r.parse()

	g.allRegion = append(g.allRegion, fmt.Sprintf("(%d,%d)", r.X, r.Z))
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
