// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"os"
	"sync/atomic"
	"time"
)

const (
	barSize    = 70
	barRefresh = time.Millisecond * 100
	barBegin   = 13
	barEnd     = 4
)

type bar struct {
	Nb, Total int64
	closer    chan<- struct{}
}

func (b *bar) Increment() {
	atomic.AddInt64(&b.Nb, 1)
}

func (b *bar) Start() {
	if b.Total == 0 {
		b.Total = 1
	}
	c := make(chan struct{})
	b.closer = c
	t := time.NewTicker(barRefresh)
	defer t.Stop()

	var buff [barBegin + barEnd + barSize]byte
	copy(buff[:], "\033[K [    % ] ")
	copy(buff[:barBegin+barSize], "\033[1G")
	for i := barBegin; i < barSize+barBegin; i++ {
		buff[i] = '.'
	}

	for {
		select {
		case <-c:
			return
		case <-t.C:
			n := b.Nb * 100 / b.Total
			buff[6] = byte(n/100) + '0'
			buff[7] = byte(n/10%10) + '0'
			buff[8] = byte(n%10) + '0'
			max := barBegin + b.Nb*barSize/b.Total
			for i := int64(barBegin); i < max; i++ {
				buff[i] = '/'
			}
			os.Stdout.Write(buff[:])
		}
	}
}

func (b *bar) Finish() {
	b.closer <- struct{}{}
	os.Stdout.WriteString("\033[1G\033[K")
}
