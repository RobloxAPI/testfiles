package main

import (
	"fmt"
	"github.com/anaminus/but"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// If true, write each object individually in its own file located under the
	// "integers" directory. If false, write all objects in single "integers"
	// file.
	const SplitFiles = false

	Define("type_xml_integers", "Test integer primitives within type elements.", func() {
		files, err := ioutil.ReadDir(filepath.Join(DataDir, "xml", "types"))
		if but.IfError(err, "read directory") {
			return
		}
		for _, info := range files {
			if !info.IsDir() {
				continue
			}
			generateIntegers(info.Name(), SplitFiles)
		}
	})
}

// Defines how to write a value of a particular type.
type integerTypeDef struct {
	// Class to use.
	class string
	// Property to use.
	prop string
	// Optional bounds.
	min, max int
	// Optional sub-tag to use.
	comp string
}

// How to generate the integer. An integer is generated with a number of `bits`
// with the formula ((sign << bits) + offs).
type integerFormat struct {
	sign int
	offs int
	// Name identifying the case. Must contain `%d` to format (bits + 1).
	name string
}

var integerTypeDefs = map[string]integerTypeDef{
	"Axes": {"ArcHandles", "Axes", 0, 7, "axes"},
}

var integerFormats = []integerFormat{
	{-1, -2, "Int%dMinUnderflowMore"},
	{-1, -1, "Int%dMinUnderflow"},
	{-1, +0, "Int%dMin"},
	{-1, +1, "Int%dMinIn"},
	{+1, -2, "Int%dMaxIn"},
	{+1, -1, "Int%dMax"},
	{+1, +0, "Int%dMaxOverflow"},
	{+1, +1, "Int%dMaxOverflowMore"},
}

var integerBits = []int{
	1,
	8*1 - 1, 8 * 1,
	8*2 - 1, 8 * 2,
	8*3 - 1, 8 * 3,
	8*4 - 1, 8 * 4,
	8*5 - 1, 8 * 5,
	8*6 - 1, 8 * 6,
	8*7 - 1, 8 * 7,
	8*8 - 1, 8 * 8,
}

func generateIntegers(typ string, split bool) {
	def, ok := integerTypeDefs[typ]
	if !ok {
		return
	}

	const headerFormat = "" +
		"#output: model\n" +
		"## Test constraints and overflow of integer type elements.\n" +
		"<roblox version=\"4\">\n"
	const itemFormat = "" +
		"\t<Item class=\"%[1]s\">\n" +
		"\t\t<Properties>\n" +
		"\t\t\t<string name=\"Name\">%[2]s</string>\n" +
		"\t\t\t<%[4]s name=\"%[3]s\">%[5]s</%[4]s>\n" +
		"\t\t</Properties>\n" +
		"\t</Item>\n"
	const compFormat = "" +
		"<%[1]s>%[2]s</%[1]s>"
	const footerFormat = "" +
		"</roblox>\n"

	var w *file
	var dir string
	if split {
		dir = filepath.Join("xml", "types", typ, "integers")
		if but.IfError(os.MkdirAll(filepath.Join(DataDir, dir), 0755), "make directory") {
			return
		}
	} else {
		w = Open(filepath.Join("xml", "types", typ, "Integers.rbxmx"))
		if w == nil {
			return
		}
		w.WriteString(headerFormat)
	}
	for _, bits := range integerBits {
		for _, format := range integerFormats {
			name := fmt.Sprintf(format.name, bits+1)
			if split {
				w = Open(filepath.Join(dir, name+".rbxmx"))
				if w == nil {
					continue
				}
				w.WriteString(headerFormat)
			}

			num := big.NewInt(int64(format.sign))
			num.Lsh(num, uint(bits))
			num.Add(num, big.NewInt(int64(format.offs)))
			value := num.String()
			name += "_" + value
			if def.comp != "" {
				value = fmt.Sprintf(compFormat, def.comp, value)
			}
			fmt.Fprintf(w, itemFormat, def.class, name, def.prop, typ, value)
			if split {
				w.WriteString(footerFormat)
				w.Close()
			}
		}
	}
	if def.min != 0 || def.max != 0 {
		for _, format := range integerFormats {
			name := "Range" + strings.ReplaceAll(format.name, "%d", "")
			if split {
				w = Open(filepath.Join(dir, name+".rbxmx"))
				if w == nil {
					continue
				}
				w.WriteString(headerFormat)
			}
			num := new(big.Int)
			if format.sign < 0 {
				num.SetInt64(int64(def.min))
			} else {
				num.SetInt64(int64(def.max + 1))
			}
			num.Add(num, big.NewInt(int64(format.offs)))
			value := num.String()
			name += "_" + value
			if def.comp != "" {
				value = fmt.Sprintf(compFormat, def.comp, value)
			}
			fmt.Fprintf(w, itemFormat, def.class, name, def.prop, typ, value)
			if split {
				w.WriteString(footerFormat)
				w.Close()
			}
		}
	}
	if !split {
		w.WriteString(footerFormat)
		w.Close()
	}
}
