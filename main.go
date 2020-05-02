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
		log.Println("[DURATION]", time.Since(before).Round(time.Millisecond))
	}(time.Now())

	generator.Gen(generator.Option{
		PrintDuration: time.Millisecond * time.Duration(100),
		Regions:       "data/",
	})
}
