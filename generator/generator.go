package generator

import (
	"./cpumutex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// All the options for one generation
type Option struct {
	// One region file.
	Region string
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

	gen.wg.Add(1)
	go gen.addRegion(gen.Region)
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
	g.nbChunckSum = 1024

	go func() {
		for range time.NewTicker(g.PrintDuration).C {
			if g.nbChunckOk == 0 {
				continue
			}
			wait := (time.Since(g.begin) *
				time.Duration((g.nbChunckSum-g.nbChunckOk)/g.nbChunckOk)).
				Round(time.Millisecond)
			fmt.Printf("\033[2K   %3d%% %5d/%-5d %s\033[50D",
				g.nbChunckOk*100/g.nbChunckSum,
				g.nbChunckOk,
				g.nbChunckSum,
				wait)
		}
	}()
}
