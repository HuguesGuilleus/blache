// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// All the options for one generation
type Option struct {
	// One region file.
	Regions string
	// Disable bar print
	NoBar bool
	// The file output directory. If empty string, it will be "dist/".
	Out Writer
	// The max number of CPU who can be used.
	// If less 0, it will the the number of CPU.
	CPU int
}

// Used to save image or web assets. dir is the directory and can be empty.
type Writer interface {
	Dir(dir string) error
	// The directory can be empty. It must be safe for multiple goroutine.
	File(dir, name string, data []byte) error
}

// Write the file into
// A Writer to save data into files. It implement flag.Value
type WriterFile struct {
	root string
	m    sync.Mutex // used to limit the number of concurent open file.
}

func NewWriterFile(root string) *WriterFile { return &WriterFile{root: root} }
func (w *WriterFile) Dir(dir string) error {
	w.m.Lock()
	defer w.m.Unlock()
	return os.MkdirAll(filepath.Join(w.root, dir), 0775)
}
func (w *WriterFile) File(dir, name string, data []byte) error {
	w.m.Lock()
	defer w.m.Unlock()
	return ioutil.WriteFile(filepath.Join(w.root, dir, name), data, 0664)
}
func (w *WriterFile) String() string { return w.root }
func (w *WriterFile) Set(root string) error {
	w.root = root
	return nil
}
