// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package webData generated by go-bindata.// sources:
// web/index.html
// web/style.css
// web/app.js
package webData

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _webIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x31\x6f\xe4\x20\x10\x85\x6b\xf3\x2b\xb8\xe9\xcf\x58\x77\xcd\x15\x40\xb1\x97\x28\xe9\x92\x62\x9b\x94\x2c\x9e\x35\x24\x63\x1b\xc1\x38\xd1\xfe\xfb\x88\x25\x91\xa2\x95\x56\xda\x6a\x3c\x6f\x9e\x3f\x9e\x06\xf4\xaf\xbb\xa7\xff\xfb\x97\xe7\x7b\x19\x78\x26\x2b\x74\x2d\x92\xdc\x32\x19\xc0\x05\xac\x10\x3a\xa0\x1b\xad\xe8\xf4\x8c\xec\xa4\x0f\x2e\x17\x64\x03\x1b\x1f\x7f\xff\x83\xaa\x73\x64\x42\xbb\x23\xe7\x03\x6a\xd5\x3a\xd1\x69\x8a\xcb\x9b\xcc\x48\x06\x0a\x9f\x08\x4b\x40\x64\x90\x21\xe3\xf1\x4b\xe9\x7d\x29\x67\x40\xf1\x39\x26\x96\x25\x7b\x03\x2e\xa5\xfe\xb5\xc0\xe5\x39\x72\xc4\x23\x66\xab\x55\xf3\x5a\xa1\x55\x8b\x25\xf4\x61\x1d\x4f\x95\xb2\x91\x8c\xa3\xe1\x48\xb8\x3f\x25\x7c\xc8\xeb\x96\xac\xe8\x6a\x0e\xe9\xc9\x95\x62\xe0\x7b\x06\x3f\x8d\x3b\x5a\x7d\xf5\x75\x3a\xce\x53\xcb\x70\xa0\xd5\xab\xa1\x1f\xfa\xb4\x4c\x20\x1d\xb1\x01\x68\x96\x64\xab\x5d\xab\xdb\xc8\x71\x9d\xf1\x12\x5d\xb5\xab\xec\x3a\xbc\x11\xfe\x88\x71\x0a\x7c\x41\x0f\x67\xf1\x1a\xbe\xfd\xd2\xf8\x5a\x6d\x54\x8b\x77\xcb\xbb\x2b\x15\x0c\xed\xf3\xcf\x08\xf2\x23\x8e\x1c\x0c\xfc\x1d\x06\x90\x0d\xd9\x1a\xab\x55\x33\xd5\xf5\xb7\xb5\xd7\x7b\x38\xbf\x9b\xcf\x00\x00\x00\xff\xff\x83\xbf\x90\x80\x48\x02\x00\x00")

func webIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webIndexHtml,
		"web/index.html",
	)
}

func webIndexHtml() (*asset, error) {
	bytes, err := webIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/index.html", size: 584, mode: os.FileMode(420), modTime: time.Unix(1588108983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x6e\xeb\x30\x0c\x44\xd7\xd6\x29\x08\xfc\xb5\x83\xe0\xff\xbf\x28\xe8\x03\xf4\x02\xbd\x80\x6c\xd1\x0e\x11\x59\x14\x24\x26\xb1\x5b\xe4\xee\x85\xec\xb8\x88\x8b\x76\x29\x81\x33\xf3\x66\x5a\x71\x33\x7c\x98\x6a\xb4\x69\xe0\x80\x70\x8c\x53\x63\xaa\x68\x9d\xe3\x30\x6c\x4f\xb9\x52\xea\xbd\xdc\x10\x4e\xec\x1c\x85\xc6\x54\xbd\x04\xad\x7b\x3b\xb2\x9f\x11\xb2\x0d\xb9\xce\x94\xb8\x6f\xcc\xdd\x98\x3f\xca\x9e\xde\xe6\x48\xaf\x49\x2e\xb1\xb8\x47\xc9\xac\x2c\x01\xa1\xe7\x89\x5c\x63\x2a\x4f\xbd\x22\xfc\xa5\x62\xdf\x8a\xaa\x8c\x08\xc7\xe6\x89\x63\x47\x51\x04\x9c\xb5\xce\x3a\x7b\x42\x08\x12\x68\x49\x3a\x6c\x49\x25\xc4\x71\x8e\xde\xce\x08\x1c\x3c\x07\xaa\x5b\x2f\xdd\xf9\x9b\x65\x2b\xc9\x51\xaa\x55\x22\x42\x16\xcf\x0e\x8e\x87\x7f\x34\x81\xb3\xe9\x3c\x24\x9a\x1b\x53\x5d\x29\x29\x77\xd6\xd7\xd6\xf3\x10\x10\x56\xba\x1d\xce\xe1\xff\xca\x6d\xbb\xf3\x90\xe4\x12\x1c\x82\xe3\xf1\xa1\x5f\x96\xc9\xfc\x4e\x08\xde\xa6\x61\x01\xfd\xe2\xc4\x53\xd9\xb2\xd0\x3e\x8b\x57\xe5\xae\x0f\x8f\xc3\xae\xd3\x56\xe6\xc6\x4e\x4f\x08\x2f\x05\x60\x27\x88\xbb\xf3\x75\xa1\x5f\xd6\x5c\x1a\x97\x2f\xa5\x49\xb7\x9a\x1d\x05\xa5\xb4\x37\x7d\xd0\xc6\x9f\x48\xee\xe6\x33\x00\x00\xff\xff\x25\xe1\x92\xc3\x3b\x02\x00\x00")

func webStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_webStyleCss,
		"web/style.css",
	)
}

func webStyleCss() (*asset, error) {
	bytes, err := webStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/style.css", size: 571, mode: os.FileMode(420), modTime: time.Unix(1587760277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x18\x6b\x93\xdb\xb6\xf1\xbb\x7f\x85\xcc\x4e\x25\x20\xc2\xf1\x24\x9d\xe3\xce\x90\x86\x3d\x79\xdc\xa4\x6e\x9d\xd8\x73\x97\x34\xe9\x69\x54\x0d\x44\x42\x12\x72\x14\xc1\x21\xa1\xb7\xf8\xdf\x3b\x0b\x80\x20\xf5\xb8\xcb\xb5\x93\x2f\x12\xb8\x58\xec\x2e\xf6\xbd\x58\xb1\xbc\x35\x1e\xb3\x35\x13\x8a\xe7\x14\xa9\xb9\x28\xda\x6d\xf8\xf5\x1d\x14\x1f\x0e\xd3\x65\x1a\x29\x21\x53\xbd\xff\x4d\x3e\x23\x63\x96\xcf\x96\x0b\x9e\xaa\x82\x7c\x21\x33\x9e\xf2\x9c\x29\x99\xe3\x7d\x85\xd8\x62\xb1\xcc\x14\x5a\xb1\x64\xc9\xf1\x3e\xe7\x6a\x99\xa7\x2d\xfd\xd5\x12\x69\xa1\x58\x1a\x71\x39\x6d\x7d\xf9\xa0\x41\x41\xca\xd7\xad\x2f\xc8\x31\xc9\x79\x21\x93\x95\x3e\xa7\x17\x96\x4c\x58\xe2\xb0\x7c\x65\x69\xa5\x7c\x8d\xbe\x1c\x0e\xe8\x0b\xfd\x92\xcb\x85\x28\x38\xc6\x67\x04\x48\xce\x7f\xe7\x91\x6a\x48\x35\x5d\x26\x53\x91\x24\x3c\xae\x24\x53\xf9\x76\x5f\x28\x9e\x21\x77\x07\x3f\xe5\x9b\x4a\x70\x1c\x96\x11\x53\xd1\x1c\x69\x59\x80\x16\x02\x31\xca\x57\x8e\xa0\x81\x3e\x47\x6f\xe8\xa9\x79\x2e\xd7\xde\xe8\xc5\x34\xf5\xf9\x9c\x17\xcb\x44\x69\x15\x2c\x13\xe5\xc7\x32\xe5\x1f\x2a\x75\x58\x98\xa1\x17\x18\x4d\x1f\xc1\x7c\x35\xe7\x29\x72\x97\x25\x95\x94\xa0\x3f\x4d\xbd\x16\x8f\xd6\x17\x67\x59\x96\x6c\x2f\x58\xf8\x70\x18\x8e\x30\x36\x7a\xc1\xc6\x0a\xa1\x71\x9b\x9a\xca\xb1\xe3\xd4\x0e\x71\xc1\x75\x26\x32\xde\xe2\xbd\x26\x40\xf7\x09\x9b\xf0\x24\xe8\x91\x82\xa7\x2a\x70\xa8\x78\x2f\xa6\x48\x0d\x7b\xa3\x76\x1f\x6b\xed\xb5\xd4\xb0\x3f\x0a\xad\xe9\xf5\xba\x24\x2a\xdf\x16\xc1\x70\x44\x64\x06\x7f\x25\x99\x92\x2d\x51\x64\x56\x61\xcd\xe8\x1e\x24\x0e\x56\x3c\x9f\xa0\x1e\x26\xd6\x0c\xe6\xbb\x8f\x89\x67\xf0\x2c\x60\x80\x4b\xa2\xb6\x19\x38\xe5\xfd\x76\x31\x91\x09\xa5\xd4\xab\xe4\xf1\xda\x6d\x34\x1b\x1a\xb8\x0f\x11\x01\x57\x1b\xd1\x86\xb8\x95\x64\x73\x51\x84\x25\x26\xb3\xd0\x19\x53\x53\x4f\x1d\x86\x3b\xb3\x72\x20\x6d\x90\x61\x4a\x56\x23\x50\xec\xa9\x1f\xc8\x4c\x2b\x63\x6a\xf5\x00\x91\xf2\xf3\x36\xe3\xb7\x79\x2e\x73\xe4\xfd\x50\x29\xba\x25\x8a\x16\x4b\x72\xce\xe2\x6d\x8b\x6f\x78\xb4\x54\x22\x9d\xf9\x1e\x0e\xd7\x73\x91\x70\x34\xc6\xe0\x98\x40\x87\xf6\xc9\xb6\xdd\x46\x8a\xca\x0c\xf4\x3b\xf8\xb0\x1d\x56\x9a\x18\x05\x1a\x06\x10\xeb\xb2\x87\x03\x42\x8a\x36\x30\x70\xbb\xad\xfc\x88\x25\x09\xda\x62\xd2\xc3\xc1\x56\x7b\x05\x6e\xb7\x5f\x23\x45\xab\x1d\x22\xb3\x61\x1f\x3c\x06\x9c\x16\x57\xaa\x09\xc5\x14\x6d\x69\x8f\x28\x2c\x33\x3a\xb4\xdc\x89\xf5\xd8\x51\x58\xac\x05\x04\x85\x86\xe3\x7d\xc4\x0a\xde\xea\x05\xfa\xaf\x1f\x80\xb0\xe1\x24\xe7\xec\x31\xd4\x90\x37\xc1\xd8\xd7\x8e\xd3\xed\x5a\x6b\xef\x4d\x1e\xd1\x8c\x09\xb0\x0d\xa6\x2c\x29\x78\x69\xf0\xbf\x6e\xe0\x6f\xa9\x46\x0a\x41\x88\xde\x28\x8c\x64\xaa\x44\xba\xe4\x06\xf1\x6f\x81\xcc\xe8\xd8\x97\x59\xe1\x67\x32\x43\x38\x1c\xfb\xe0\x65\xf6\xc3\xe1\xc6\x7c\xca\x96\x89\x0a\xc4\x14\xc1\xbd\x0d\x12\x01\x05\x24\x3c\x9d\xa9\xf9\xfb\x5e\xbb\xad\x86\xd5\xd7\x55\x1f\xd4\x66\x6e\x46\x29\x7d\x7b\x38\x54\xcb\x01\xc6\xfb\x31\xed\xd5\x84\xcb\x57\x62\xea\x10\x6f\xda\x6d\xf4\x5a\x1d\x0e\x48\x0b\xfc\x5e\x87\x43\x5b\xaf\xdf\xa9\xe1\xcd\x08\xc3\x61\x73\x2d\x7b\x25\xa3\xa1\x23\x1a\x6f\xdb\x6d\x8b\xf3\x0e\xa2\xa6\x3e\xa1\x63\xa8\xa1\x57\x7d\x4a\x35\xb1\x07\x47\xd8\x83\x51\x68\xf5\xb2\x2c\xc0\x4a\xf8\xe8\x1c\x20\xff\xb1\xda\xca\x57\x32\xa3\x10\xfd\xc6\x4f\x5c\xa2\x69\x66\x44\x30\xcb\x5b\xc2\x47\xe1\x96\xf6\xc2\x72\x2a\x52\x96\x24\xdb\xfd\x94\x2a\xf8\x74\x37\x6b\x7f\x6d\xc3\xc1\x5c\xfc\xd4\x07\x7a\xa3\x0f\x7a\x23\x58\x49\x11\xb7\x7a\xc6\x21\x54\xbe\xe4\x65\x58\x96\x75\x68\x2e\x0b\x9e\x7f\x2f\xd7\x69\x22\x59\x8c\x26\x24\x65\x0b\x6e\xf2\xd2\x92\xfe\x72\xf7\xc9\x8f\x72\xce\x14\xff\x3c\x81\xd4\xf9\xcb\xdd\x27\x34\xc1\x3a\xeb\x31\x1a\xcb\x48\x27\x46\x8b\x71\x9b\x70\xf8\x42\x1d\xd6\xc1\xa1\xdb\xd3\xf7\x64\x59\xc6\xd3\xf8\xbb\xb9\x48\x62\xc4\x70\xc8\xfc\x79\xce\xa7\x74\x19\x32\x3f\xb6\x7c\x29\x70\x0d\x99\x1f\x25\x22\x7a\x44\x80\x92\xf3\x85\x5c\x71\x84\x43\x90\x21\xe7\x2b\xf9\xd8\x90\x61\x09\x19\x1c\xa4\xb8\xbb\xfd\xe1\xe3\xe7\x9f\xc6\xf7\x1f\x1f\x6e\x69\xff\xed\x57\x37\x03\x2d\xdb\xcf\x22\xe1\x90\x1d\xc2\xba\x14\x56\x20\xbc\xaf\x56\x43\x6f\x92\xc8\xc8\x1b\x51\xf3\x1f\x36\xe0\x42\x2e\xb8\xde\xd0\x8b\xc6\xce\x9c\x8b\xd9\x5c\xc1\x96\x5d\x85\x25\x76\xa4\x0f\x07\xb7\xa4\xfb\x12\x1b\x35\x3d\x48\xb9\x68\x88\x01\x9f\x78\x0f\xbf\x43\xfd\xe3\x89\xd4\x1b\xd1\xde\x88\xc2\x22\x6c\xc0\xe5\x12\xd8\xf4\x47\x54\xaf\x80\x0d\xc0\x0f\x07\xfd\x57\x93\xff\x97\xe0\x6b\xe8\x57\x1a\x39\xd8\xd9\xd5\xec\x21\x11\x93\xb9\xad\x33\xe0\x6a\x54\x27\x67\x5d\xa0\x32\x59\xfc\x46\x7b\x6e\xfd\x50\xad\x0b\xb1\xe3\xb4\xa1\x59\x03\x35\x09\x5f\xc8\xb4\x42\x4b\x44\xa1\xee\xf8\x4c\xc8\xb4\xa0\xc3\x91\x81\x45\x2c\x5d\xb1\xa2\x76\x8d\x19\x57\xd6\x2f\xbe\xdd\x7e\x8c\x91\x88\x31\xe4\xbf\x26\x2a\xa5\xe9\x32\x49\xf0\x2b\xe3\xc8\x9e\x88\x5b\xa9\x6c\x2d\x20\x0e\x3c\x4b\x52\x6d\x68\xe3\x00\x90\xfc\x4e\xa6\x0a\xaa\x70\x67\x10\x77\x1a\x04\xd5\xe6\x94\x5a\x2a\x5b\xe6\x54\x2b\xd2\x47\x2c\x49\x91\x0a\xf5\xa9\x96\x1e\x61\x03\x9e\xb3\x62\x7e\xcf\x15\x9a\xdb\x6f\x70\xc0\x1f\xe5\xb2\xe0\x15\x42\xce\x41\x35\x08\x87\x6b\x91\xc6\x72\xed\xb3\x38\xbe\x5d\xf1\x54\xd3\x82\x0a\x84\x3c\xf0\x65\x8f\x9c\x57\xc4\xf1\xf1\xf9\xf2\x19\x12\x06\xe9\x65\x44\xca\x57\xc6\xc8\x7e\x96\x4b\x25\xa1\x70\xdb\xed\x66\x51\x6e\x2a\x6f\x2d\x62\x35\xa7\x96\xb7\x48\x53\x9e\xff\x0a\x90\xa6\xf1\x7c\xe3\xda\x47\x48\x7f\xd7\x20\x83\x15\xe7\x6c\xfd\x4d\x92\x80\x00\xe1\x19\x77\xbb\x79\xc6\xde\x39\x4f\xb7\x5b\xe7\x86\x44\x46\x1a\xa6\x15\x4f\x9d\x09\x7e\xe0\xaa\x52\x78\xa4\x36\x3e\xf4\x6e\xf7\x6a\x9b\x70\x08\x54\x16\x3d\x7a\xc7\x5b\x77\xd0\x39\xf6\x48\x8f\x9c\x5d\x93\x9c\x5f\x0a\x87\x53\x99\x23\x88\x85\x0d\xbd\x72\xbe\x1e\x6e\xde\x9d\x1d\xee\x36\x76\xbb\xd4\x7d\xe0\x7d\x45\x60\xd7\x24\xb0\x7b\x77\xce\xab\x41\x61\x77\x44\xc1\x69\xf1\xe3\x82\xcd\x38\xda\x90\x1d\xf4\xbd\x4f\x28\x53\xe3\xd4\xea\x34\x1c\x7e\x23\xe6\xff\xa1\x76\x8d\x6a\x48\xd1\xa1\x40\x6c\xc2\xb7\x7f\x0d\x63\x80\xe4\x42\x91\x84\xc4\x1b\x12\xef\xc8\x86\xec\x88\x58\xcc\x08\x87\x36\x6a\xdc\x0f\x1d\x35\xd7\xb9\x1a\x7a\x8e\xc2\x98\xe1\xbd\x6d\x51\xc6\xcc\x54\x45\xd7\xa5\x08\x45\x8f\x6d\x1d\x26\xf5\xb5\xc3\x78\x43\xad\xf0\x57\x3f\x32\x35\xf7\xd9\xa4\x40\x2e\x0b\xfd\x35\xc1\x61\xbc\xb3\x08\x0f\xe7\x08\x0f\x80\x50\x31\xa4\xfd\xd0\x36\x44\x63\x66\x0b\x2c\x14\xe3\x61\x9f\xdc\x10\xf2\x66\x84\x43\x9b\x2f\x0a\x15\x7f\x27\x65\x1e\xd7\x6c\xac\xda\x7e\xc3\xe1\xee\x09\x94\x07\xa7\xd9\x2a\xe2\x75\x86\xb8\xdd\x88\x42\x19\x4b\x19\x15\x0d\xdf\x90\xaa\x74\x01\xd8\xb8\x9a\xaa\x6a\xcc\xc8\x08\x38\x08\xc4\x62\x46\xc7\xcc\x87\xae\x1e\xd5\x89\xca\x29\xe8\x35\xa5\x42\x61\x3b\xca\x41\x6f\xe1\x3c\x5b\x80\xd9\xef\x17\x52\xaa\xb9\x48\x67\xb7\x29\x9b\x24\x3c\xa6\xba\x99\xab\x91\x6a\x1f\x02\x1b\x1a\x8b\x26\x24\x69\xc4\x4e\xa1\x72\xf9\xc8\x6d\xf4\xe4\x3c\x6e\xc4\x4e\x22\x52\xae\x63\x9f\xde\xf8\xbd\xea\x56\x37\xe4\x8d\x95\xfd\x26\xb0\x3e\xf1\xff\xc9\x7f\xc4\x58\xe6\x2c\x9d\xf1\xa7\xd8\xf7\xfc\xc1\x39\xfb\x37\xc1\x09\x29\xa4\x67\x63\xa6\xe6\x83\xef\x91\xf7\xa3\xd7\x8d\x37\x5d\xaf\xe5\x75\xe3\x5d\xd7\x6b\xad\xbc\x6e\xd2\xf5\x5a\x73\xf3\xb7\xf2\xba\x57\xf0\xbf\xf3\xb0\xb3\xd6\x60\x14\x96\xa5\x1d\xd9\xce\xc2\xac\x72\x82\x3a\xca\x32\x59\x58\x3f\x70\x01\xa6\x5d\x72\x9a\x48\x99\x23\xd8\xee\xda\xed\xeb\x3a\xae\x2f\x91\x3e\x29\x34\x17\x46\xa5\xff\x25\x70\xc7\xec\x05\x01\x3a\xa9\x03\x74\x72\x12\xa0\x63\x66\x8a\xbf\x73\xe1\x29\x07\x3c\xcf\x3a\xb9\xff\x7b\x21\x53\xaf\x72\xde\x7e\xe0\xd0\x80\x92\x71\x01\xac\x71\x50\xed\xe0\x10\x93\x8d\xfb\x39\xc4\xd3\x52\xf1\x12\x3b\x34\x63\xad\xd6\x14\x04\x1d\x4c\x6c\xaf\x4f\x7b\x0e\x5f\xa4\x51\xb2\x8c\x79\x81\x3c\xe4\x75\x37\x5d\x8f\x78\xdd\x5d\xd7\xc3\x1e\x86\x14\xab\xcb\xbf\xa1\xd8\x8a\x25\x2f\x5a\xa9\x54\x2d\x0e\xa4\xbd\xf0\x52\xae\xb5\x85\xe7\x82\x85\x26\x4a\x32\xf4\x8f\xfb\xcf\x3f\x81\x27\x8a\x74\x26\xa6\x5b\xb4\xdf\x04\x75\x52\xd9\x05\x75\xf6\x28\x02\xe7\x10\x44\x05\x47\x59\x01\x5a\xb6\x27\x18\xdf\x37\x19\xdb\x66\x2d\x0a\x61\x54\x8d\xa8\xe6\x9c\xb1\xbc\xe0\x88\x29\x39\x41\x73\x3f\xe7\x59\xc2\x22\x8e\xae\xff\xf3\x97\x6b\xe2\x79\x18\xe3\xc3\x61\x0f\x13\xb3\x99\x1e\xc6\x78\x1f\x51\xfd\x7d\xc4\x9e\xa2\xc8\x57\x2d\x91\xba\xe6\x18\x7f\x88\x7c\x15\x54\x5f\x3e\x34\xc1\x8d\xb6\xf0\xa7\xe5\x62\xc2\x73\x14\xf9\x1b\x7c\x38\x34\x7b\x44\xb7\xb1\xab\x37\x74\xc7\xe1\x36\x0a\x7c\x38\x9c\x75\x8f\xcf\x76\x0d\xd1\x1c\x12\x84\x6b\x9e\xeb\xa7\x12\xf3\x00\x42\xe9\x71\x7a\xb5\x09\x27\x3c\xbe\xdf\x4b\xda\x93\x9d\x94\x8b\x23\x65\xef\x5c\xb4\xec\x6c\x98\xd4\x6d\x79\x6d\x4a\xaa\xa3\x7f\x21\xcc\xe3\x8d\x06\x7d\x35\x20\x8d\x3b\x7e\xd5\xef\xe1\xe6\x64\x0e\x44\x7c\xb9\x54\x67\x24\xd8\xa6\x26\x71\x7d\x44\xe2\xba\xff\xd6\x0d\x93\x17\x24\x87\x5e\xf4\x9f\x7c\x5b\x4b\xfe\xc8\xb7\x67\x5d\x3d\x7c\x4e\xe9\xbe\xf3\x4d\x9e\xcb\xf5\x27\x3e\x55\x9d\xe0\xa9\x2e\x12\x8c\x7c\xd5\x6c\xf2\xaf\x07\x61\x49\xcc\xc9\x3b\x68\x5f\x9e\x3f\xda\x7d\xe2\xe8\x2f\xd9\xb3\xe7\x1e\x9e\x62\x09\xf3\xe7\xf3\x27\x2f\x70\xbc\x7a\xfa\x84\x35\x34\xaa\x0c\x81\x01\xbf\xfb\x32\x7c\x63\x7d\x7d\xa2\xd7\x09\x9a\xa3\xd5\xb8\x39\x36\x8d\x9b\x73\xd3\xf8\xf2\xe0\x54\x62\xd2\x29\x9e\xe6\x6a\xfb\x45\x25\xbf\x4d\xe4\xa4\x66\x34\x71\x68\x27\xa3\xf9\x85\xce\x59\x16\x0a\x46\xe7\xae\x37\xf6\xba\xe3\xe3\x9e\xa6\x96\x96\xf4\x70\xd7\xf3\x9f\x46\x78\x30\x08\x59\x3a\xf3\x4c\x72\x2e\x87\x8f\x7c\x3b\x82\x82\xff\x7a\xea\xc2\x6d\x7a\x9e\xd6\x9f\x70\x54\x3d\x34\xd1\xd3\xea\x75\xec\xa7\x09\x2b\xd4\xaf\x73\xce\x13\x0a\x95\xfd\x7b\xa6\xdc\x94\x65\xb5\x72\x36\x19\x75\xd6\x80\xde\xa9\x6b\x1d\x87\x6d\x9d\x22\xf4\xca\x8f\x79\xa2\xd8\xbf\x29\xa5\x3d\x27\x33\x70\x4a\xe5\xba\xc9\x43\x4c\x51\x2a\xd7\xe6\xbd\xed\xf3\x14\xe1\x2b\x27\x49\x0d\x7b\xdf\xef\xf5\x2a\x73\x57\xae\xd1\xe4\xf1\xbe\xf7\xc1\x85\x78\xd3\x65\xc6\xa7\xfa\xd1\x95\xce\x1a\xd3\xa8\x2a\x2c\x11\xc6\xf5\xe3\xcb\x71\x17\xe9\x0c\xaf\xbb\x1d\xf3\xa0\xff\xf4\x73\xbe\x6e\xe6\xf5\xd5\x4c\x17\x88\x43\xe1\x17\x79\x44\x55\xd7\xbb\xd6\x15\xd1\xd7\x15\x51\xdb\x35\x14\xbe\x34\x2f\x2d\xe7\xce\x58\xbd\xa7\x0b\x10\x16\xf0\x74\xeb\x77\x11\x51\xbf\xd3\xeb\x11\x3b\x2c\xcd\x00\x0a\x42\xac\x04\x37\x2a\xb6\x4f\x0e\x1d\x63\xc2\x41\xdc\xb9\xe4\xb2\xac\x98\x37\x1e\x88\x4e\x5e\x08\x3a\x55\x3e\xff\x36\x91\x51\x07\x5f\x98\x8f\xf5\x13\xd1\xa5\xf1\x18\xa4\x38\x29\x25\xe8\xa8\xbc\x69\xdf\xfe\x63\xc6\x42\x2e\xf8\x9f\xc9\x19\xe8\xbd\x8c\xb5\x19\xb3\xff\x44\xde\xd5\xd4\xfb\xdc\x7b\xc3\x23\xdf\x82\x13\x7a\x67\x71\xd5\x64\x60\xcb\x8f\x0d\x02\xa8\x3d\x40\xf3\xbf\x01\x00\x00\xff\xff\xa7\x08\xb5\x3a\x35\x1b\x00\x00")

func webAppJsBytes() ([]byte, error) {
	return bindataRead(
		_webAppJs,
		"web/app.js",
	)
}

func webAppJs() (*asset, error) {
	bytes, err := webAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/app.js", size: 6965, mode: os.FileMode(420), modTime: time.Unix(1588113565, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"web/index.html": webIndexHtml,
	"web/style.css":  webStyleCss,
	"web/app.js":     webAppJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"web": &bintree{nil, map[string]*bintree{
		"app.js":     &bintree{webAppJs, map[string]*bintree{}},
		"index.html": &bintree{webIndexHtml, map[string]*bintree{}},
		"style.css":  &bintree{webStyleCss, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}