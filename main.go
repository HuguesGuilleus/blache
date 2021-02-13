// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/blache/meta"
	"github.com/HuguesGuilleus/blache/pkg"
	"os"
	"strings"
	"time"
)

func main() {
	out := blache.NewOsCreater("public")
	opt := blache.Option{
		Out: &out,
		Error: func(err error) {
			fmt.Fprintf(os.Stderr, "\033[1G\033[K%v\n", err)
		},
	}

	flag.BoolVar(&opt.NoBar, "bar", false, "Disable progress bar")
	flag.IntVar(&opt.CPU, "cpu", 0, "The number of core used, zero is for all core.")
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
		opt.In = r
	} else {
		opt.In = os.DirFS(a)
	}

	defer func(before time.Time) {
		fmt.Println("[DURATION]", time.Since(before).Round(time.Millisecond*10))
	}(time.Now())
	opt.Gen()
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

func printVersion() {
	fmt.Println("Blache")
	fmt.Println()
	fmt.Println("\tGit: ", meta.Git)
	fmt.Println("\tDate:", meta.Date)
	fmt.Println()
	fmt.Println(`BSD 3-Clause License
Copyright (c) 2020, 2021, Hugues GUILLEUS
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`)
}
