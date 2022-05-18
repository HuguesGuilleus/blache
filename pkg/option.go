// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// All the options for one generation
type Option struct {
	// The regions sources. Must be defined.
	Input fs.FS
	// The output. Must be defined.
	Output Creator
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

	return os.WriteFile(filepath.Join(c.Root, dir, name), data, 0664)
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
