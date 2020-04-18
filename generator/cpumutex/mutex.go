package cpumutex

import (
	"runtime"
)

type M struct {
	init bool
	c    chan int
}

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
		m.init = true
		max := runtime.NumCPU()
		m.c = make(chan int, max)
		for i := 0; i < max; i++ {
			m.c <- 0
		}
	}
	<-m.c
}

func (m *M) Unlock() {
	m.c <- 0
}
