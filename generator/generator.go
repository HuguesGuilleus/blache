package generator

import (
	// "log"
	"./cpumutex"
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
}

type generator struct {
	Option
	chunck cpumutex.M
	wg     sync.WaitGroup
}

func Gen(option Option) {
	gen := generator{
		Option: option,
	}
	defer gen.wg.Wait()

	gen.wg.Add(1)
	gen.addRegion(gen.Region)
}
