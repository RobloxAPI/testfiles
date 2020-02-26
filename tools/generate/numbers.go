package main

import (
	"fmt"
	"math"
	"strings"
)

func init() {
	Define("type_float", "Generate data/types/float* files.", func() {
		w := Open("types/float.rbxmx")
		if w == nil {
			return
		}
		w.WriteString("#output: model\n<roblox version=\"4\">\n\t<Item class=\"Folder\">\n\t\t<Properties><string name=\"Name\">float</string></Properties>\n")
		itbits("0 abbbcccd ABBBBBBBBBBCCCCCCCCCCCD", func(v uint64, i int) {
			f := math.Float32frombits(uint32(v))
			fmt.Fprintf(w, "\t\t<Item class=\"BlurEffect\"><Properties><string name=\"Name\">f%s</string><float name=\"Size\">%.20g</float></Properties></Item>\n", u32bits(uint32(v)), f)
		})
		w.WriteString("\t</Item>\n</roblox>\n")
		w.Close()
	})

	Define("type_double", "Generate data/types/double* files.", func() {
		w := Open("types/double.rbxmx")
		if w == nil {
			return
		}
		w.WriteString("#output: model\n<roblox version=\"4\">\n\t<Item class=\"Folder\">\n\t\t<Properties><string name=\"Name\">double</string></Properties>\n")
		itbits("0 abbbbcccccd ABBBBBBBBBBBBBBBBBBBBBBBBBCCCCCCCCCCCCCCCCCCCCCCCCCD", func(v uint64, i int) {
			f := math.Float64frombits(v)
			fmt.Fprintf(w, "\t\t<Item class=\"NumberValue\"><Properties><string name=\"Name\">d%s</string><double name=\"Value\">%.20g</double></Properties></Item>\n", u64bits(v), f)
		})
		w.WriteString("\t</Item>\n</roblox>\n")
		w.Close()
	})
}

// f64fmt formats a float32 string with separators between each component.
func f32fmt(s string) string {
	return s[:1] + "_" + s[1:9] + "_" + s[9:]
}

// u32bits formats the bits of a uint32 as a float32.
func u32bits(v uint32) string {
	return f32fmt(fmt.Sprintf("%032b", v))
}

// f64fmt formats a float64 string with separators between each component.
func f64fmt(s string) string {
	return s[:1] + "_" + s[1:12] + "_" + s[12:]
}

// u64bits formats the bits of a uint64 as a float64.
func u64bits(v uint64) string {
	return f64fmt(fmt.Sprintf("%064b", v))
}

// itbits iterates through a sequence of bits as specified by the bits string.
// Each unique letter indicates a field that is filled in with a 0 or 1.
// Matching letters are filled in with the same value. Every combination of
// fields is traversed. The number of iterations is equal to 2^x, where x is the
// number of unique letters. Non-letter are ignored. cb is called with each
// produced integer, as well as the current iteration.
func itbits(bits string, cb func(uint64, int)) {
	bits = strings.ReplaceAll(bits, " ", "")
	var fields [256]uint8
	var n uint8
	for _, b := range bits {
		if fields[b] != 0 {
			continue
		}
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			n++
			fields[b] = n
		}
	}
	for i := 0; i < 1<<n; i++ {
		f := fmt.Sprintf("%0*b", n, i)
		b := []byte(bits)
		var n uint64
		for i, c := range b {
			if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') {
				q := f[fields[c]-1]
				b[i] = q
				if q == '1' {
					n |= 1 << uint(len(b)-i-1)
				}
			}
		}
		cb(n, i)
	}
}
