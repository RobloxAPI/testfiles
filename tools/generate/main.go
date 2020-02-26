// Generate test files.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anaminus/but"
)

const DataDir = `../../data`

var defs []def

type def struct {
	value *bool
	gen   func()
}

func Define(name, desc string, gen func()) {
	value := flag.Bool(name, false, desc)
	defs = append(defs, def{value: value, gen: gen})
}

type file struct {
	f *os.File
	*bufio.Writer
}

func (f *file) Close() {
	if but.IfError(f.Flush(), "flush buffer") {
		return
	}
	if but.IfError(f.f.Sync(), "sync file") {
		return
	}
	if but.IfError(f.f.Close(), "close file") {
		return
	}
}

func Open(name string) *file {
	f, err := os.Create(filepath.Join(DataDir, name))
	if but.IfError(err, "create new file") {
		return nil
	}
	return &file{f: f, Writer: bufio.NewWriter(f)}
}

func main() {
	const usage = `
The generate command must be run from tools/generate within the testfiles
repository.
`
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()
	}
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	flag.Parse()
	for _, def := range defs {
		if *def.value {
			def.gen()
		}
	}
}
