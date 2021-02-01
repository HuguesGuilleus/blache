// BSD 3-Clause License in LICENSE file at the project root.
// Copyright (c) 2020, Hugues GUILLEUS
// All rights reserved.

package main

import (
	"./meta"
	"./pkg"
	"flag"
	"fmt"
	"os"
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

	flag.Var(&out, "out", "The output Directory")
	flag.IntVar(&opt.CPU, "cpu", 0, "The number of core used, zero is for all core.")
	version := flag.Bool("version", false, "Print the version and exit")
	flag.Parse()

	opt.In = blache.NewReaderFile(flag.Arg(0))

	if *version {
		printVersion()
		return
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
		fmt.Println("  input is a directory that contain minecraft regions (*.mca)")
		flag.PrintDefaults()
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
