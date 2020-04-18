package main

import (
	"./generator"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	log.Println("main()")

	defer func(before time.Time) {
		log.Println("[DURATION]", time.Since(before))
	}(time.Now())

	generator.Gen(generator.Option{
		PrintDuration: time.Millisecond * time.Duration(100),
		Region:        "data/r.-1.-1.mca",
		// Region:        "data/r.0.0.mca",
	})
}
