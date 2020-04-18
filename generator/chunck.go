package generator

import (
	"log"
)

type chunck struct {
	g           *generator
	X, Z        int
	DataVersion int32
	Level       struct {}
}

func (c *chunck) draw() {
	log.Println("[CHUNCK]", c.X, c.Z)
}
