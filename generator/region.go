package generator

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (g *generator) addRegion(file string) {
	defer g.wg.Done()

	var X, Z int
	_, f := filepath.Split(file)
	if _, err := fmt.Sscanf(f, "r.%d.%d.mca", &X, &Z); err != nil {
		log.Println("[ERROR] read X,Z from file name:", err)
		return
	}

	log.Println("[FILE]", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("[ERROR] read:", file, err)
		return
	}

	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			offset := 4 * (x + z*32)
			if bytesToInt(data[offset:offset+4]) == 0 {
				continue
			}
			addr := 4096 * (bytesToInt(data[offset : offset+3]))
			l := bytesToInt(data[addr : addr+4])
			if typeOfCompress := data[addr+4]; typeOfCompress != 2 {
				log.Print("Unknown compress (2):", typeOfCompress)
				continue
			}
			c, err := reginParseChunck(data[addr+5 : addr+4+l])
			if err != nil {
				log.Println("[ERROR]", err)
				continue
			}

			c.X = x + X*32
			c.Z = z + Z*32
			c.g = g

			g.wg.Add(1)
			g.chunck.New(c.draw, g.wg.Done)
		}
	}

}

// Decompress and parse a chunck
func reginParseChunck(data []byte) (c *chunck, _ error) {
	// Decompress data
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// Parse minecraft data
	c = &chunck{}
	return c, nbt.Unmarshal(data, c)
}

// Convert a slice of bytes to int
func bytesToInt(tab []byte) (n int) {
	for _, b := range tab {
		n = n<<8 + int(b)
	}
	return n
}
