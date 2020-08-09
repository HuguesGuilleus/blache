// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package blache

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// All the options for one generation
type Option struct {
	// The regions sources.
	In Reader
	// Disable bar print
	NoBar bool
	// The file output directory. If empty string, it will be "dist/".
	Out Writer
	// The max number of CPU who can be used.
	// If less 0, it will the the number of CPU.
	CPU int
}

type Reader interface {
	// Open the reader.
	Open() error
	// At the end of the file list, the error is EOF.
	// Name used to get region coordonates.
	Read() (name string, data []byte, err error)
}

// An implementation of Reader who read file into a dir.
type ReaderFile struct {
	dir   string
	files []os.FileInfo
}

func NewReaderFile(dir string) *ReaderFile {
	return &ReaderFile{dir: dir}
}
func (r *ReaderFile) Open() error {
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
	r.files = r.files[1:]

	if !strings.HasSuffix(n, ".mca") {
		return r.Read()
	}

	data, err := ioutil.ReadFile(filepath.Join(r.dir, n))
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
