// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package main

import (
	"./pkg"
	"flag"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.Println("main()")
}

func main() {
	opt := blache.Option{
		Regions: "data1/",
	}
	flag.IntVar(&opt.CPU, "cpu", 0, "The number of core used, zero is for all core.")
	flag.Parse()

	defer func(before time.Time) {
		log.Println("[DURATION]", time.Since(before).Round(time.Millisecond*10))
	}(time.Now())

	for err := range opt.Gen() {
		log.Println(err)
	}
}
