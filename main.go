// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package main

import (
	"archive/zip"
	_ "embed"
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/blache/pkg"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	out := blache.NewOsCreater("public")
	opt := blache.Option{
		Output: &out,
	}

	flag.BoolVar(&opt.NoBar, "bar", false, "Disable progress bar")
	flag.Var(&out, "out", "The output Directory")
	version := flag.Bool("version", false, "Print the version and exit")
	flag.Parse()

	if *version {
		printVersion()
		return
	}

	if a := flag.Arg(0); a == "" {
		flag.Usage()
		return
	} else if strings.HasSuffix(a, ".zip") {
		r, err := zip.OpenReader(a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[1G\033[KFail to open zip file: %q: %v\n", a, err)
			os.Exit(1)
		}
		opt.Input = r
	} else {
		opt.Input = os.DirFS(a)
	}

	defer func(before time.Time) {
		fmt.Println("[DURATION]", time.Since(before).Round(time.Millisecond*10))
	}(time.Now())

	errors := blache.Generate(opt)
	for _, err := range errors {
		fmt.Fprintln(os.Stderr, err)
	}
	if len(errors) > 0 {
		os.Exit(1)
	}
}

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: $ blache [OPTION ...] input")
		fmt.Println()
		fmt.Println("Input:")
		fmt.Println("  - a directory")
		fmt.Println("  - a zip file")
		fmt.Println("  Inside all minecraft regions *.mca must in one of this repository:")
		fmt.Println("  - world/region")
		fmt.Println("  - region")
		fmt.Println("  - direcly in the repository")
		fmt.Println()
		fmt.Println("Option:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  $ blache -out=www my_world.zip")
		fmt.Println("  $ blache -out=www ~/.minecraft/saves/new_world/")
		fmt.Println()
	}
}

//go:embed LICENSE
var license string

func printVersion() {
	if info, _ := debug.ReadBuildInfo(); info != nil {
		fmt.Println("Blache", info.Main.Version)
		fmt.Println(info.Main.Sum)
	} else {
		fmt.Println("Blache")
	}
	fmt.Println()
	fmt.Println(license)
}
