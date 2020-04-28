package cpumutex

import (
	"runtime"
)

// A mutex for limit the number of parralelism process. Zero value is valid.
type M struct {
	init bool
	c    chan int
}

// Init m, if max is less 0, max will be the number of CPU.
func (m *M) Init(max int) {
	m.init = true
	if max < 1 {
		max = runtime.NumCPU()
	}
	m.c = make(chan int, max)
	for i := 0; i < max; i++ {
		m.c <- 0
	}
}

// Lock m, run in a new goroutine each functions, and unlock m.
func (m *M) New(ff ...func()) {
	m.Lock()
	go func() {
		defer m.Unlock()
		for _, f := range ff {
			f()
		}
	}()
}

func (m *M) Lock() {
	if !m.init {
		m.Init(0)
	}
	<-m.c
}

func (m *M) Unlock() {
	m.c <- 0
}
