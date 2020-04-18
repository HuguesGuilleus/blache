package generator

import (
	// "log"
	"./cpumutex"
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

	os.MkdirAll(gen.outBiome, 0774)

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
