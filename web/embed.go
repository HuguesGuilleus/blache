// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2021, Hugues GUILLEUS
// All rights reserved.

// Contain embed file to run web map.
package web

import _ "embed"

// File index.html
//go:embed index.html
var html []byte

// File app.Js, builed by typescript
//go:embed app.js
//go:generate go run ts/generate.go
var js []byte

// File style.css
//go:embed style.css
var style []byte

// All static web files, to put at root of the output directory.
var List = []struct {
	Name string
	Data []byte
}{
	{"index.html", html},
	{"app.js", js},
	{"style.css", style},
}
