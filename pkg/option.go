// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// All the options for one generation
type Option struct {
	// The regions sources.
	In fs.FS
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
