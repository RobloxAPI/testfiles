package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/anaminus/but"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/rbxl"
	"github.com/robloxapi/rbxfile/rbxlx"
	"github.com/robloxapi/testfiles/src/golden"
)

func recurseRefs(refs map[*rbxfile.Instance]int, instances []*rbxfile.Instance) {
	for _, inst := range instances {
		if _, ok := refs[inst]; !ok {
			refs[inst] = len(refs)
			recurseRefs(refs, inst.Children)
		}
	}
}

func handler() golden.Handler {
	var refs map[*rbxfile.Instance]int
	return func(e *golden.Encoder, v interface{}) bool {
		switch v := v.(type) {
		case map[string]string:
			keys := make([]string, 0, len(v))
			for k := range v {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			a := make(golden.Array, len(keys))
			for i, k := range keys {
				a[i] = golden.Object{
					{Name: "Key", Value: k},
					{Name: "Value", Value: v[k]},
				}
			}
			e.Array(a)

		case *rbxfile.Root:
			// Prepopulate ref table.
			refs = map[*rbxfile.Instance]int{}
			recurseRefs(refs, v.Instances)

			e.Object(golden.Object{
				{Name: "Metadata", Value: v.Metadata},
				{Name: "Instances", Value: v.Instances},
			})

		case map[string]rbxfile.Value:
			props := make([]string, 0, len(v))
			for name := range v {
				props = append(props, name)
			}
			sort.Strings(props)

			a := make(golden.Array, len(props))
			for i, name := range props {
				value := v[name]
				a[i] = golden.Object{
					{Name: "Name", Value: name},
					{Name: "Type", Value: value.Type().String()},
					{Name: "Value", Value: value},
				}
			}
			e.Array(a)

		case []*rbxfile.Instance:
			a := make(golden.Array, len(v))
			for i, inst := range v {
				a[i] = inst
			}
			e.Array(a)

		case *rbxfile.Instance:
			var ref interface{}
			if r, ok := refs[v]; ok {
				ref = r
			}
			e.Object(golden.Object{
				{Name: "ClassName", Value: v.ClassName},
				{Name: "IsService", Value: v.IsService},
				{Name: "Reference", Value: ref},
				{Name: "Properties", Value: v.Properties},
				{Name: "Children", Value: v.Children},
			})

		case rbxfile.ValueString:
			e.Value(string(v))

		case rbxfile.ValueBinaryString:
			e.Value([]byte(v))

		case rbxfile.ValueProtectedString:
			e.Value(string(v))

		case rbxfile.ValueContent:
			e.Value(string(v))

		case rbxfile.ValueBool:
			e.Value(bool(v))

		case rbxfile.ValueInt:
			e.Value(int64(v))

		case rbxfile.ValueFloat:
			e.Value(float32(v))

		case rbxfile.ValueDouble:
			e.Value(float64(v))

		case rbxfile.ValueUDim:
			e.Object(golden.Object{
				{Name: "Scale", Value: v.Scale},
				{Name: "Offset", Value: v.Offset},
			})

		case rbxfile.ValueUDim2:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
			})

		case rbxfile.ValueRay:
			e.Object(golden.Object{
				{Name: "Origin", Value: v.Origin},
				{Name: "Direction", Value: v.Direction},
			})

		case rbxfile.ValueFaces:
			e.Object(golden.Object{
				{Name: "Right", Value: v.Right},
				{Name: "Top", Value: v.Top},
				{Name: "Back", Value: v.Back},
				{Name: "Left", Value: v.Left},
				{Name: "Bottom", Value: v.Bottom},
				{Name: "Front", Value: v.Front},
			})

		case rbxfile.ValueAxes:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
				{Name: "Z", Value: v.Z},
			})

		case rbxfile.ValueBrickColor:
			e.Value(uint32(v))

		case rbxfile.ValueColor3:
			e.Object(golden.Object{
				{Name: "R", Value: v.R},
				{Name: "G", Value: v.G},
				{Name: "B", Value: v.B},
			})

		case rbxfile.ValueVector2:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
			})

		case rbxfile.ValueVector3:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
				{Name: "Z", Value: v.Z},
			})

		case rbxfile.ValueCFrame:
			e.Object(golden.Object{
				{Name: "Position", Value: v.Position},
				{Name: "Rotation", Value: golden.Object{
					{Name: "R00", Value: v.Rotation[0]},
					{Name: "R01", Value: v.Rotation[1]},
					{Name: "R02", Value: v.Rotation[2]},
					{Name: "R10", Value: v.Rotation[3]},
					{Name: "R11", Value: v.Rotation[4]},
					{Name: "R12", Value: v.Rotation[5]},
					{Name: "R20", Value: v.Rotation[6]},
					{Name: "R21", Value: v.Rotation[7]},
					{Name: "R22", Value: v.Rotation[8]},
				}},
			})

		case rbxfile.ValueToken:
			e.Value(uint32(v))

		case rbxfile.ValueReference:
			if i, ok := refs[v.Instance]; ok {
				e.Value(i)
			} else {
				e.Value(nil)
			}

		case rbxfile.ValueVector3int16:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
				{Name: "Z", Value: v.Z},
			})

		case rbxfile.ValueVector2int16:
			e.Object(golden.Object{
				{Name: "X", Value: v.X},
				{Name: "Y", Value: v.Y},
			})

		case rbxfile.ValueNumberSequenceKeypoint:
			e.Object(golden.Object{
				{Name: "Time", Value: v.Time},
				{Name: "Value", Value: v.Value},
				{Name: "Envelope", Value: v.Envelope},
			})

		case rbxfile.ValueNumberSequence:
			a := make(golden.Array, len(v))
			for i, k := range v {
				a[i] = k
			}
			e.Array(a)

		case rbxfile.ValueColorSequenceKeypoint:
			e.Object(golden.Object{
				{Name: "Time", Value: v.Time},
				{Name: "Value", Value: v.Value},
				{Name: "Envelope", Value: v.Envelope},
			})

		case rbxfile.ValueColorSequence:
			a := make(golden.Array, len(v))
			for i, k := range v {
				a[i] = k
			}
			e.Array(a)

		case rbxfile.ValueNumberRange:
			e.Object(golden.Object{
				{Name: "Min", Value: v.Min},
				{Name: "Max", Value: v.Max},
			})

		case rbxfile.ValueRect:
			e.Object(golden.Object{
				{Name: "Min", Value: v.Min},
				{Name: "Max", Value: v.Max},
			})

		case rbxfile.ValuePhysicalProperties:
			if v.CustomPhysics {
				e.Object(golden.Object{
					{Name: "CustomPhysics", Value: v.CustomPhysics},
					{Name: "Density", Value: v.Density},
					{Name: "Friction", Value: v.Friction},
					{Name: "Elasticity", Value: v.Elasticity},
					{Name: "FrictionWeight", Value: v.FrictionWeight},
					{Name: "ElasticityWeight", Value: v.ElasticityWeight},
				})
			} else {
				e.Object(golden.Object{
					{Name: "CustomPhysics", Value: v.CustomPhysics},
					{Name: "Density", Value: nil},
					{Name: "Friction", Value: nil},
					{Name: "Elasticity", Value: nil},
					{Name: "FrictionWeight", Value: nil},
					{Name: "ElasticityWeight", Value: nil},
				})
			}

		case rbxfile.ValueColor3uint8:
			e.Object(golden.Object{
				{Name: "R", Value: v.R},
				{Name: "G", Value: v.G},
				{Name: "B", Value: v.B},
			})

		case rbxfile.ValueInt64:
			e.Value(int64(v))

		case rbxfile.ValueSharedString:
			e.Value([]byte(v))

		default:
			return false
		}
		return true
	}
}

func defaultConfig(group golden.Group) golden.Config {
	return golden.Config{
		Format: strings.TrimPrefix(filepath.Ext(group.Input), "."),
		Struct: "model",
	}
}

func openGroup(group golden.Group) {
	config := defaultConfig(group)
	if b, err := os.ReadFile(group.Config); err == nil {
		if but.IfErrorf(json.Unmarshal(b, &config), "decode golden config") {
			return
		}
	}

	input, err := os.Open(group.Input)
	if err != nil {
		return
	}
	defer input.Close()

	var data interface{}
	var warn error
	switch r := bufio.NewReader(input); config.Format {
	case "rbxl":
		switch config.Struct {
		case "model":
			data, warn, err = rbxl.Decoder{
				Mode:  rbxl.Place,
				NoXML: true,
			}.Decode(r)
		default:
			err = fmt.Errorf("unknown struct %q for format %q", config.Struct, config.Format)
		}
	case "rbxm":
		switch config.Struct {
		case "model":
			data, warn, err = rbxl.Decoder{
				Mode:  rbxl.Model,
				NoXML: true,
			}.Decode(r)
		default:
			err = fmt.Errorf("unknown struct %q for format %q", config.Struct, config.Format)
		}
	case "rbxlx", "rbxmx":
		switch config.Struct {
		case "model":
			data, err = rbxlx.Deserialize(r)
		default:
			err = fmt.Errorf("unknown struct %q for format %q", config.Struct, config.Format)
		}
	default:
		return
	}

	if group.Golden == "" {
		group.Golden = group.Input + ".golden"
	}
	var gold *os.File
	{
		var err error
		gold, err = os.Create(group.Golden)
		but.IfFatal(err, "create golden file")
		defer gold.Close()
	}
	g := golden.NewGolden(gold, config)
	g.SetHandler(handler())
	defer g.Close()

	if warn != nil {
		g.Value("Warning", warn)
	}
	if err != nil {
		g.Value("Error", err)
	} else {
		g.Value("Data", data)
	}
	but.IfFatalf(gold.Sync(), "sync golden file")
}

func main() {
	flag.Parse()
	root := os.DirFS(".")
	var groups []golden.Group
	for _, file := range flag.Args() {
		info, err := fs.Stat(root, file)
		if but.IfError(err) {
			continue
		}
		if info.IsDir() {
			g, err := golden.Groups(root, file)
			if but.IfError(err) {
				continue
			}
			groups = append(groups, g...)
		} else {
			g := golden.GroupOf(root, file)
			if g.Input == "" {
				continue
			}
			groups = append(groups, g)
		}
	}
	for _, group := range groups {
		fmt.Println("GROUP", group)
		openGroup(group)
	}
}
