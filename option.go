// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package blache

import (
	_ "embed"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// Use it for display license.
//go:embed LICENSE
var License string

// All the options for one generation
type Option struct {
	// The regions sources. Must be defined.
	Input fs.FS

	// The output. Must be defined.
	Output FSWriter

	// Output for the log, can be unset (will be replaced by io.Discard).
	LogOutput io.Writer
}

// Error occure when do not found file in the Option.Input
var NotMapFiles = errors.New("Not found the map files")

func (option *Option) getFiles() (root string, files []fs.DirEntry, err error) {
	for _, root = range [...]string{"world/region", "region", "."} {
		files, err = fs.ReadDir(option.Input, root)
		if err == nil {
			return
		}
	}

	return "", nil, NotMapFiles
}

// Used to write asset web file and generated file.
type FSWriter interface {
	// Used to read waypoints.json file if HTTP and detecet updated files.
	fs.FS
	// Create a directory.
	MkdirAll(dir string) error
	// Create a file into the directory dir (can be empty) with the data.
	// This method can be called concurrently an many many time.
	Create(dir, name string, data []byte) error
}

// A Creator that write directory and file into the operating system into the root.
type osFSWriter struct {
	fs.FS
	Root string
	// Mutex to limit the number of concurent open writed file.
	sync.Mutex
}

// Create FSWriter that interact with the os file system.
func NewOsFSWriter(root string) FSWriter {
	return &osFSWriter{
		FS:   os.DirFS(root),
		Root: root,
	}
}

// Create the directory if it doesn't exist.
func (c *osFSWriter) MkdirAll(dir string) error {
	return os.MkdirAll(filepath.Join(c.Root, dir), 0o775)
}

// Create dir/name with the data.
func (c *osFSWriter) Create(dir, name string, data []byte) error {
	c.Lock()
	defer c.Unlock()

	return os.WriteFile(filepath.Join(c.Root, dir, name), data, 0664)
}
