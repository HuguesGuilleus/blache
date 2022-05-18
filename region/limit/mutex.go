// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package cpumutex

// A mutex for limit the number of parralelism process. Zero value is valid.
type M struct {
	c chan int
}

// Init m, if max is less 0, max will be the number of CPU.
func (m *M) Init(max int) {
	if max > 0 {
		m.c = make(chan int, max)
		for i := 0; i < max; i++ {
			m.c <- 0
		}
	}
}

func (m *M) Lock() {
	if m.c != nil {
		<-m.c
	}
}

func (m *M) Unlock() {
	if m.c != nil {
		m.c <- 0
	}
}
