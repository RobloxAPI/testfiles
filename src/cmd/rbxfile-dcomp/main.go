// The rbxfile-dcomp command rewrites a rbxl/rbxm file with decompressed
// chunks.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/anaminus/but"
	"github.com/robloxapi/rbxfile/rbxl"
)

const usage = `usage: rbxfile-dcomp [INPUT] [OUTPUT]

Reads a binary RBXL or RBXM file from INPUT, and writes to OUTPUT the same file,
but with uncompressed chunks.

INPUT and OUTPUT are paths to files. If INPUT is "-" or unspecified, then stdin
is used. If OUTPUT is "-" or unspecified, then stdout is used. Warnings and
errors are written to stderr.
`

func main() {
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	flag.Usage = func() { fmt.Fprintf(flag.CommandLine.Output(), usage) }
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 && args[0] != "-" {
		in, err := os.Open(args[0])
		but.IfFatal(err, "open input")
		input = in
		defer in.Close()
	}
	if len(args) >= 2 && args[1] != "-" {
		out, err := os.Create(args[1])
		but.IfFatal(err, "create output")
		defer out.Close()
		defer func() {
			but.IfFatal(out.Sync(), "sync output")
		}()
		output = out
	}

	warn, err := rbxl.Decoder{}.Decompress(output, input)
	if warn != nil {
		but.Log("warning:", warn)
	}
	if err != nil {
		but.Log("error:", err)
	}
}
