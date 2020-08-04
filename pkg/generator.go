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
	"time"
)

type generator struct {
	Option
	region cpumutex.M
	chunck cpumutex.M
	wg     sync.WaitGroup

	// For print
	// TODO: rm these
	begin       time.Time
	nbChunckOk  int // the chunck generated
	nbChunckSum int // the total number of chunck

	// All the regions coord.
	allRegion []string

	err chan<- error
}

func (option Option) Gen() <-chan error {
	err := make(chan error)
	g := generator{
		Option: option,
		err:    err,
	}
	g.region.Init(option.CPU)
	g.chunck.Init(option.CPU)

	for _, d := range [...]string{"bloc", "biome", "height"} {
		g.Out.Dir(d)
	}

	g.makeAssets()

	// gen.print()

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
	g.nbChunckSum += 1024
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

func (g *generator) saveRegionsList() {
	sort.Strings(g.allRegion)
	data, err := json.Marshal(g.allRegion)
	if err != nil {
		g.err <- err
		return
	}
	g.Out.File("", "regions.json", data)
}

// Print the progress.
func (g *generator) print() {
	// if g.PrintDuration == 0 {
	// 	return
	// }
	//
	// g.begin = time.Now()

	bar := pb.New(0)
	bar.Format("[=> ]")
	bar.Prefix("chuncks:")
	bar.ManualUpdate = true
	bar.ShowElapsedTime = false
	bar.ShowFinalTime = false
	bar.ShowSpeed = false
	bar.ShowTimeLeft = false

	// go func() {
	// 	for range time.NewTicker(g.PrintDuration).C {
	// 		bar.Set(g.nbChunckOk)
	// 		bar.SetTotal(g.nbChunckSum)
	// 		bar.Update()
	//
	// 		// wait := float64(time.Since(g.begin)) *
	// 		// 	float64(g.nbChunckSum-g.nbChunckOk) /
	// 		// 	float64(g.nbChunckOk)
	// 		// fmt.Printf("\033[2K  %2d%%  %#6d/%-6d  %s\033[50D",
	// 		// 	g.nbChunckOk*100/g.nbChunckSum,
	// 		// 	g.nbChunckOk,
	// 		// 	g.nbChunckSum,
	// 		// 	time.Duration(wait).Round(time.Millisecond))
	// 	}
	// }()
}

// Write the web assets.
func (g *generator) makeAssets() {
	a, _ := webData.AssetDir("web")
	for _, n := range a {
		g.Out.File("", n, webData.MustAsset("web/"+n))
	}
}
