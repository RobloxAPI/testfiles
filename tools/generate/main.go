// Generate test files.
package main

import (
	"bufio"
	"flag"
	"github.com/anaminus/but"
	"os"
	"path/filepath"
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
	flag.Parse()
	for _, def := range defs {
		if *def.value {
			def.gen()
		}
	}
}
