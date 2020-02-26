package main

import (
	"fmt"
	"github.com/anaminus/but"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	// If true, write each object individually in its own file located under the
	// "spacing" directory. If false, write all objects in single "spacing"
	// file.
	const SplitFiles = false

	Define("type_xml_spacing", "Test leading and trailing spacing within type elements.", func() {
		files, err := ioutil.ReadDir(filepath.Join(DataDir, "xml", "types"))
		if but.IfError(err, "read directory") {
			return
		}
		for _, info := range files {
			if !info.IsDir() {
				continue
			}
			generateSpacing(info.Name(), SplitFiles)
		}
	})
}

// Defines how to write a value of a particular type.
type spaceTypeDef struct {
	// Name identifying the definition.
	name string
	// Class to use.
	class string
	// Property to use.
	prop string
	// String representing the value to use. Should not be equal to the default
	// value for the property.
	value string
	// Optional sub-tag to use.
	comp string
}

// How to format the value of a typeDef.
type spaceFormat struct {
	// Name identifying the case.
	name string
	// String that formats a typeDef value. Expects a string.
	str string
}

var spaceTypeDefs = map[string][]spaceTypeDef{
	"bool": {
		{"True", "BoolValue", "Value", "true", ""},
		{"False", "Frame", "Visible", "false", ""},
	},
	"Axes": {
		{"FTT", "ArcHandles", "Axes", "6", "axes"},
		{"FFF", "ArcHandles", "Axes", "0", "axes"},
	},
}

var spaceFormats = []spaceFormat{
	{"Base", "%s"},
	{"Empty", "%.s"},
	{"EmptySpace", "    %.s"},
	{"EmptyTab", "\t%.s"},
	{"EmptyNewline", "\n%.s"},
	{"EmptyReturn", "\r%.s"},
	{"EmptyVTab", "\v%.s"},
	{"EmptyFormFeed", "\f%.s"},
	{"EmptyMix", "\n \t%.s"},
	{"EmptyMixCRLF", "\r\n \t%.s"},
	{"LeadingSpace", "    %s"},
	{"LeadingTab", "\t%s"},
	{"LeadingNewline", "\n%s"},
	{"LeadingReturn", "\r%s"},
	{"LeadingVTab", "\v%s"},
	{"LeadingFormFeed", "\f%s"},
	{"LeadingMix", "\n \t%s"},
	{"LeadingMixCRLF", "\r\n \t%s"},
	{"TrailingSpace", "%s    "},
	{"TrailingTab", "%s\t"},
	{"TrailingNewline", "%s\n"},
	{"TrailingReturn", "%s\r"},
	{"TrailingVTab", "%s\v"},
	{"TrailingFormFeed", "%s\v"},
	{"TrailingMix", "%s\n \t"},
	{"TrailingMixCRLF", "%s\r\n \t"},
	{"WrappingSpace", "    %s    "},
	{"WrappingTab", "\t%s\t"},
	{"WrappingNewline", "\n%s\n"},
	{"WrappingReturn", "\r%s\r"},
	{"WrappingVTab", "\v%s\v"},
	{"WrappingFormFeed", "\f%s\f"},
	{"WrappingMix", "\n \t%s\n \t"},
	{"WrappingMixCRLF", "\r\n \t%s\r\n \t"},
	{"WrappingIndent", "\n\t\t\t\t%s\n\t\t\t"},
	{"WrappingIndentCRLF", "\r\n\t\t\t\t%s\r\n\t\t\t"},
}

func generateSpacing(typ string, split bool) {
	def, ok := spaceTypeDefs[typ]
	if !ok {
		return
	}

	const headerFormat = "" +
		"#output: model\n" +
		"## Test leading and trailing spacing within type elements.\n" +
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
		dir = filepath.Join("xml", "types", typ, "spacing")
		if but.IfError(os.MkdirAll(filepath.Join(DataDir, dir), 0755), "make directory") {
			return
		}
	} else {
		w = Open(filepath.Join("xml", "types", typ, "Spacing.rbxmx"))
		if w == nil {
			return
		}
		w.WriteString(headerFormat)
	}
	for _, def := range def {
		for _, format := range spaceFormats {
			name := def.name + format.name
			if split {
				w = Open(filepath.Join(dir, name+".rbxmx"))
				if w == nil {
					continue
				}
				w.WriteString(headerFormat)
			}
			value := fmt.Sprintf(format.str, def.value)
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
