package generator

import (
	"./cpumutex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
}

type generator struct {
	Option
	region   cpumutex.M
	chunck   cpumutex.M
	wg       sync.WaitGroup
	outBiome string

	// For print
	begin       time.Time
	nbChunckOk  int // the chunck generated
	nbChunckSum int // the total number of chunck
}

func Gen(option Option) {
	if option.Out == "" {
		option.Out = "dist/"
	}
	gen := generator{
		Option:   option,
		outBiome: filepath.Join(option.Out, "biome"),
	}
	defer gen.wg.Wait()

	if err := os.MkdirAll(gen.outBiome, 0774); err != nil {
		log.Println("[ERROR]", err)
		return
	}

	gen.print()

	for r := range gen.listRegion() {
		gen.wg.Add(1)
		gen.nbChunckSum += 1024
		go gen.addRegion(r)
	}
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
}

// Print the progress.
func (g *generator) print() {
	if g.PrintDuration == 0 {
		return
	}

	g.begin = time.Now()

	format := "\033[2K   %3d%% %5d/%-5d %s\033[50D"

	go func() {
		for range time.NewTicker(g.PrintDuration).C {
			if g.nbChunckOk == 0 {
				fmt.Printf(format, 0, 0, 0, "")
				continue
			}
			wait := (time.Since(g.begin) *
				time.Duration((g.nbChunckSum-g.nbChunckOk)/g.nbChunckOk)).
				Round(time.Millisecond)
			fmt.Printf(format,
				g.nbChunckOk*100/g.nbChunckSum,
				g.nbChunckOk,
				g.nbChunckSum,
				wait)
		}
	}()
}
