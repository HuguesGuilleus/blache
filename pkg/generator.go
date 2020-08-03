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
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// All the options for one generation
type Option struct {
	// One region file.
	Regions string
	// Disable bar print
	NoBar bool
	// The file output directory. If empty string, it will be "dist/".
	Out string
	// The max number of CPU who can be used.
	// If less 0, it will the the number of CPU.
	CPU int
}

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
	if option.Out == "" {
		option.Out = "dist/"
	}
	err := make(chan error)
	gen := generator{
		Option: option,
		err:    err,
	}
	gen.region.Init(option.CPU)
	gen.chunck.Init(option.CPU)

	gen.makeAllDir()
	gen.makeAssets()

	// gen.print()

	for r := range gen.listRegion() {
		gen.wg.Add(1)
		go gen.addRegion(r)
	}

	go func() {
		gen.wg.Wait()
		gen.saveRegionsList()
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
	data, _ := json.Marshal(g.allRegion)
	ioutil.WriteFile(filepath.Join(g.Out, "regions.json"), data, 0664)
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

func (g *generator) makeAllDir() {
	for _, d := range [...]string{"bloc", "biome", "height"} {
		dir := filepath.Join(g.Out, d)
		if err := os.MkdirAll(dir, 0775); err != nil {
			log.Printf("[ERROR] make dir '%s': %v\n", dir, err)
			return
		}
	}
}

// Write the web assets.
func (g *generator) makeAssets() {
	a, _ := webData.AssetDir("web")
	for _, n := range a {
		name := filepath.Join(g.Out, n)
		// f, err := os.Create(name)
		f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
		if err != nil {
			log.Printf("[ERROR] on create assets file '%s' %v\n", name, err)
			return
		}
		defer f.Close()
		f.Write(webData.MustAsset("web/" + n))
	}
}
