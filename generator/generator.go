package generator

import (
	"./cpumutex"
	"encoding/json"
	"fmt"
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
	// The duration between two log of the generation progress.
	// If zero, no log.
	PrintDuration time.Duration
	// The file output directory. If empty string, it will be "dist/".
	Out string
	// The path to data pack (zip format).
	DataPack string
}

type generator struct {
	Option
	colorBloc colorBloc
	region    cpumutex.M
	chunck    cpumutex.M
	wg        sync.WaitGroup

	// For print
	begin       time.Time
	nbChunckOk  int // the chunck generated
	nbChunckSum int // the total number of chunck

	// All the regions coord.
	allRegion []string
}

func Gen(option Option) {
	if option.Out == "" {
		option.Out = "dist/"
	}
	gen := generator{
		Option: option,
	}

	if err := gen.colorBloc.Load(option.DataPack); err != nil {
		log.Println("[ERROR] in load data pack:", err)
		return
	}

	gen.makeAllDir()

	gen.print()

	for r := range gen.listRegion() {
		gen.wg.Add(1)
		gen.nbChunckSum += 1024
		go gen.addRegion(r)
	}

	gen.wg.Wait()
	gen.saveRegionsList()
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
	if g.PrintDuration == 0 {
		return
	}

	g.begin = time.Now()

	go func() {
		for range time.NewTicker(g.PrintDuration).C {
			wait := float64(time.Since(g.begin)) *
				float64(g.nbChunckSum-g.nbChunckOk) /
				float64(g.nbChunckOk)
			fmt.Printf("\033[2K  %2d%%  %#6d/%-6d  %s\033[50D",
				g.nbChunckOk*100/g.nbChunckSum,
				g.nbChunckOk,
				g.nbChunckSum,
				time.Duration(wait).Round(time.Millisecond))
		}
	}()
}

func (g *generator) makeAllDir() {
	for _, d := range [...]string{"bloc", "bloc", "height"} {
		dir := filepath.Join(g.Out, d)
		if err := os.MkdirAll(dir, 0774); err != nil {
			log.Printf("[ERROR] make dir '%s': %v\n", dir, err)
			return
		}
	}
}
