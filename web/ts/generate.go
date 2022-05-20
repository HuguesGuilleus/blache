package main

import (
	"bytes"
	"github.com/clarkmcc/go-typescript"
	"github.com/tdewolff/minify/v2/js"
	"io/fs"
	"os"
	"strings"
)

func main() {
	// concat all .ts file into buff
	buff := bytes.Buffer{}
	fsys := os.DirFS("ts")

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if d.IsDir() || !strings.HasSuffix(d.Name(), ".ts") {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}
		buff.Write(data)

		return nil
	})
	if err != nil {
		panic(err)
	}

	// Build typescript
	output, err := typescript.Transpile(&buff, typescript.WithCompileOptions(map[string]interface{}{
		"downlevelIteration": true,
		"lib":                []string{"es2018", "dom"},
		"removeComments":     true,
		"strict":             true,
		"target":             "ES5",
	}))
	if err != nil {
		panic(err)
	}

	// Minify
	buff.Reset()
	err = js.DefaultMinifier.Minify(nil, &buff, bytes.NewReader([]byte(output)), nil)
	if err != nil {
		panic(err)
	}

	// Save it
	os.WriteFile("app.js", buff.Bytes(), 0o664)
}
