// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// All the options for one generation
type Option struct {
	// The regions sources.
	In Reader
	// The output. Must be defined.
	Out Creator
	// Disable bar print
	NoBar bool
	// The max number of CPU who can be used.
	// If less 0, it will the the number of CPU.
	CPU int
	// Log the error. Is not set, never output.
	Error func(error)
}

// TODO: utiliser fs.FS #go1.16
type Reader interface {
	// Open the reader.
	Open() error
	// At the end of the file list, the error is EOF.
	// Name used to get region coordonates.
	Read() (name string, data []byte, err error)
}

// An implementation of Reader who read file into a dir.
type ReaderFile struct {
	dir     string
	files   []os.FileInfo
	Verbose bool
}

func NewReaderFile(dir string) *ReaderFile {
	return &ReaderFile{dir: dir}
}
func (r *ReaderFile) Open() error {
	if r.Verbose {
		log.Printf("Open %q", r.dir)
	}
	files, err := ioutil.ReadDir(r.dir)
	if err != nil {
		return err
	}
	r.files = files
	return nil
}
func (r *ReaderFile) Read() (string, []byte, error) {
	if len(r.files) == 0 {
		return "", nil, io.EOF
	}
	n := r.files[0].Name()
	nn := filepath.Join(r.dir, n)
	r.files = r.files[1:]

	if !strings.HasSuffix(n, ".mca") {
		r.print("Read skip %q", nn)
		return r.Read()
	}

	r.print("Read %q", nn)
	data, err := ioutil.ReadFile(nn)
	if err != nil {
		return "", nil, err
	}
	return n, data, nil
}
func (r *ReaderFile) String() string { return r.dir }
func (r *ReaderFile) Set(dir string) error {
	if dir != "" {
		r.dir = dir
	} else {
		r.dir = "."
	}
	return nil
}

// print only if verbose
// TODO: remove it
func (r *ReaderFile) print(fmt string, args ...interface{}) {
	if r.Verbose {
		log.Printf(fmt, args...)
	}
}

// Used to write asset web file and generated file.
type Creator interface {
	// Create a directory.
	MkdirAll(dir string) error
	// Create a file into the directory dir (can be empty) with the data.
	// This method can be called concurrently an many many time.
	Create(dir, name string, data []byte) error
}

// A Creator that write directory and file into the operating system into the root.
type OsCreator struct {
	Root       string
	sync.Mutex // used to limit the number of concurent open file.
}

// Create an OsCreator with the root.
func NewOsCreater(root string) OsCreator {
	return OsCreator{Root: root}
}

// Create the directory if it doesn't exist.
func (c *OsCreator) MkdirAll(dir string) error {
	return os.MkdirAll(filepath.Join(c.Root, dir), 0o775)
}

// Create dir/name with the data.
func (c *OsCreator) Create(dir, name string, data []byte) error {
	c.Lock()
	defer c.Unlock()
	// TODO: use os #go1.16
	return ioutil.WriteFile(filepath.Join(c.Root, dir, name), data, 0664)
}

// Set Root. Used for the package package.
func (c *OsCreator) Set(root string) error {
	c.Root = root
	return nil
}

// String return the root.
func (c *OsCreator) String() string {
	return c.Root
}
