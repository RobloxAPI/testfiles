// The golden package provides methods for producing and comparing golden files.
package golden

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Handler receives an arbitrary value for customized encoding. v is the value
// to be encoded. e provides low-level methods for encoding the value. Must
// return whether the value was handled.
type Handler func(e *Encoder, v interface{}) bool

// Encoder is a low-level encoder for values. An initialized Encoder is received
// by a Handler.
type Encoder struct {
	w       *bufio.Writer
	lead    []byte
	handler Handler
}

// Push increases the indentation by one.
func (e *Encoder) Push() {
	e.lead = append(e.lead, '\t')
}

// Pop decreases the indentation by one.
func (e *Encoder) Pop() {
	e.lead = e.lead[:len(e.lead)-1]
}

// Byte writes the given Byte.
func (e *Encoder) Byte(c byte) {
	e.w.WriteByte(c)
}

// Literal writes the given string.
func (e *Encoder) Literal(s string) {
	e.w.WriteString(s)
}

// Newline writes a newline character followed by the current indentation.
func (e *Encoder) Newline() {
	e.Byte('\n')
	e.w.Write(e.lead)
}

// String writes a JSON string.
func (e *Encoder) String(s string) {
	// From encoding/json
	const hex = "0123456789abcdef"
	e.Byte('"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if b >= ' ' && b != '"' && b != '\\' {
				i++
				continue
			}
			if start < i {
				e.Literal(s[start:i])
			}
			e.Byte('\\')
			switch b {
			case '\\', '"':
				e.Byte(b)
			case '\n':
				e.Byte('n')
			case '\r':
				e.Byte('r')
			case '\t':
				e.Byte('t')
			default:
				e.Literal(`u00`)
				e.Byte(hex[b>>4])
				e.Byte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				e.Literal(s[start:i])
			}
			e.Literal(`\ufffd`)
			i += size
			start = i
			continue
		}
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				e.Literal(s[start:i])
			}
			e.Literal(`\u202`)
			e.Byte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		e.Literal(s[start:])
	}
	e.Byte('"')
}

// Array writes a JSON array.
func (e *Encoder) Array(v Array) {
	e.Byte('[')
	if len(v) == 0 {
		e.Byte(']')
		return
	}
	e.Push()
	e.Newline()
	e.Value(v[0])
	for i := 1; i < len(v); i++ {
		e.Byte(',')
		e.Newline()
		e.Value(v[i])
	}
	e.Pop()
	e.Newline()
	e.Byte(']')
}

// Object writes a JSON object.
func (e *Encoder) Object(v Object) {
	e.Byte('{')
	if len(v) == 0 {
		e.Byte('}')
		return
	}
	e.Push()
	e.Newline()
	e.String(v[0].Name)
	e.Literal(": ")
	e.Value(v[0].Value)
	for i := 1; i < len(v); i++ {
		e.Byte(',')
		e.Newline()
		e.String(v[i].Name)
		e.Literal(": ")
		e.Value(v[i].Value)
	}
	e.Pop()
	e.Newline()
	e.Byte('}')
}

// Literal represents a raw fragment of JSON data.
type Literal string

// Value writes an arbitrary value. The types detected by default are described
// by the Golden.Value method. Other types are handled by the Handler function.
func (e *Encoder) Value(v interface{}) {
	switch v := v.(type) {
	case nil:
		e.Literal("null")

	case bool:
		if v {
			e.Literal("true")
		} else {
			e.Literal("false")
		}

	case string:
		for _, r := range []rune(v) {
			switch r {
			case '\t', '\r', '\n', '\f', '\b':
				continue
			}
			if !unicode.IsGraphic(r) {
				e.Value([]byte(v))
				return
			}
		}
		e.String(v)

	case []byte:
		if len(v) == 0 {
			e.Literal("[]")
			break
		}
		e.Byte('[')
		e.Push()
		const width = 16
		for j := 0; j < len(v); j += width {
			e.Newline()
			e.Literal("\"| ")
			for i := j; i < j+width; {
				if i < len(v) {
					s := strconv.FormatUint(uint64(v[i]), 16)
					if len(s) == 1 {
						e.Literal("0")
					}
					e.Literal(s)
				} else if len(v) < width {
					break
				} else {
					e.Literal("  ")
				}
				i++
				if i%8 == 0 && i < j+width {
					e.Literal("  ")
				} else {
					e.Literal(" ")
				}
			}
			e.Literal("|")
			n := len(v)
			if j+width < n {
				n = j + width
			}
			for i := j; i < n; i++ {
				if 32 <= v[i] && v[i] <= 126 {
					e.w.WriteRune(rune(v[i]))
				} else {
					e.w.WriteByte('.')

				}

			}
			e.Literal("|\"")
			if j+width < len(v)-1 {
				e.Byte(',')
			}
		}
		e.Pop()
		e.Newline()
		e.Byte(']')

	case uint:
		e.Literal(strconv.FormatUint(uint64(v), 10))

	case uint8:
		e.Literal(strconv.FormatUint(uint64(v), 10))

	case uint16:
		e.Literal(strconv.FormatUint(uint64(v), 10))

	case uint32:
		e.Literal(strconv.FormatUint(uint64(v), 10))

	case uint64:
		e.Literal(strconv.FormatUint(v, 10))

	case int:
		e.Literal(strconv.FormatInt(int64(v), 10))

	case int8:
		e.Literal(strconv.FormatInt(int64(v), 10))

	case int16:
		e.Literal(strconv.FormatInt(int64(v), 10))

	case int32:
		e.Literal(strconv.FormatInt(int64(v), 10))

	case int64:
		e.Literal(strconv.FormatInt(int64(v), 10))

	case float32:
		switch {
		case math.IsInf(float64(v), 1):
			e.Literal(`"Infinity"`)
		case math.IsInf(float64(v), -1):
			e.Literal(`"-Infinity"`)
		case math.IsNaN(float64(v)):
			e.Literal(`"NaN"`)
		default:
			e.Literal(strconv.FormatFloat(float64(v), 'g', 9, 32))
		}

	case float64:
		switch {
		case math.IsInf(v, 1):
			e.Literal(`"Infinity"`)
		case math.IsInf(v, -1):
			e.Literal(`"-Infinity"`)
		case math.IsNaN(v):
			e.Literal(`"NaN"`)
		default:
			e.Literal(strconv.FormatFloat(v, 'g', 17, 64))
		}

	case json.Number:
		e.Literal(string(v))

	case error:
		e.Value(v.Error())

	case Array:
		e.Array(v)

	case []interface{}:
		e.Array(v)

	case Object:
		e.Object(v)

	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		o := make(Object, len(keys))
		for i, k := range keys {
			o[i] = Field{Name: k, Value: v[k]}
		}
		e.Object(o)
	default:
		if e.handler != nil {
			if e.handler(e, v) {
				return
			}
		}
		e.Literal("<UNKNOWN:" + reflect.TypeOf(v).String() + ">")
	}
}

// Golden is an encoder that writes the golden format.
type Golden struct {
	e *Encoder
	c bool
}

// NewGolden returns a Golden that writes to w. It is initialized by writing the
// fields in config. For correctly formatted data, Golden must be closed once
// all data has been written.
func NewGolden(w io.Writer, config Config) *Golden {
	var buf *bufio.Writer
	if bw, ok := w.(*bufio.Writer); ok {
		buf = bw
	} else {
		buf = bufio.NewWriter(w)
	}
	g := &Golden{e: &Encoder{w: buf}}
	// Write root object.
	g.e.Byte('{')
	g.e.Push()
	g.e.Newline()
	g.e.String("Format")
	g.e.Literal(": ")
	g.e.String(config.Format)
	g.e.Byte(',')
	g.e.Newline()
	g.e.String("Struct")
	g.e.Literal(": ")
	g.e.String(config.Struct)
	return g
}

// Close finishes encoding the root object, flushes any remaining data to be
// written, and closes the encoder.
func (g *Golden) Close() error {
	if g.c {
		return fmt.Errorf("already closed")
	}
	g.c = true
	// Close root object.
	g.e.Pop()
	g.e.Newline()
	g.e.Byte('}')
	g.e.Newline()
	return g.e.w.Flush()
}

// SetHandler sets a function to call when Value does not match any types. If
// the handler returns false, then an error value is written instead.
func (g *Golden) SetHandler(handler Handler) {
	g.e.handler = handler
}

// Array represents a JSON array.
type Array []interface{}

// Array writes a Data field as a JSON array. Each argument is an element in the
// array.
//
// No attempt is made to ensure that the field name is unique.
//
// Does nothing if the encoder is closed.
func (g *Golden) Array(name string, v ...interface{}) {
	if g.c {
		return
	}
	g.e.Array(v)
}

// Object represents a JSON object.
type Object []Field

// Field represents the field of a JSON object.
type Field struct {
	Name  string
	Value interface{}
}

// Object writes a Data field as a JSON object. Each argument is a field in the
// object.
//
// No attempt is made to ensure that the field name is unique.
//
// Does nothing if the encoder is closed.
func (g *Golden) Object(name string, v ...Field) {
	if g.c {
		return
	}
	g.e.Object(v)
}

// Value writes a Data field as an arbitrary value. The following types are
// detected by default:
//
//     - nil is written as a JSON null.
//     - bool is written as a JSON bool.
//     - primitive numeric types are written as a JSON number.
//     - Infinity and NaN are written as a JSON string.
//     - Literal and json.Number are written as-is. Note that an improperly
//       formatted value can cause the resulting JSON to become invalid.
//     - []byte is written as a JSON array of human-readable strings.
//     - string is written as a JSON string, or according to []byte if the
//       string contains non-graphic characters.
//     - error is written according to string.
//     - Array and []interface{} are written as a JSON array.
//     - Object is written as a JSON object with ordered fields. No attempt is
//       made to make fields unique.
//     - map[string]interface{} is written as a JSON object with fields ordered
//       lexicographically.
//
// If the type is not known, then the function set by SetHandler is called. If
// the handler returns false, or no handler is set, then `<UNKNOWN:type>` is
// emitted, causing the resulting JSON to be invalid.
//
// No attempt is made to ensure that the field name is unique.
//
// Does nothing if the encoder is closed.
func (g *Golden) Value(name string, value interface{}) {
	if g.c {
		return
	}
	g.e.Byte(',')
	g.e.Newline()
	g.e.String(name)
	g.e.Literal(": ")
	g.e.Value(value)
}

// Config contains options that configure a group.
type Config struct {
	// The format of the input file.
	Format string `json:",omitempty"`
	// The structure of the output file.
	Struct string `json:",omitempty"`
}
