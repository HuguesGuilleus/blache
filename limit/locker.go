// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package limit

import (
	"sync"
)

// A sync.Locker that can be locked some times. Zero value is valid.
type locker struct {
	c chan int
}

// Create a new limit locker. If max if strict positive, it is the number
// of possible call to Locker before blocks.
func New(max int) sync.Locker {
	l := locker{}
	if max > 0 {
		l.c = make(chan int, max)
		for i := 0; i < max; i++ {
			l.c <- 0
		}
	}
	return l
}

func (l locker) Lock() {
	if l.c != nil {
		<-l.c
	}
}

func (l locker) Unlock() {
	if l.c != nil {
		l.c <- 0
	}
}
