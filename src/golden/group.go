package golden

import (
	"io/fs"
	"path"
	"sort"
	"strings"
)

type Group struct {
	Input  string
	Config string
	Golden string
}

// Groups returns a list of file groups produced by walking the given directory.
//
// Sibling files form a group when their names match. The part of the name used
// to match depends on the type:
//
//     - An input file uses its full name.
//     - A golden file uses its name without the `.golden` extension.
//     - A config file uses its name without the `.golden-config` extension.
//
// Each resulting Group will have an Input, file other fields are set only if
// the file exists. Paths in each Group are relative to the input directory.
func Groups(f fs.FS, dir string) (groups []Group, err error) {
	gs := map[string]Group{}
	err = fs.WalkDir(f, dir, func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		base := path.Base(name)
		// Ignore hidden files.
		if strings.HasPrefix(base, ".") && base != "." {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		ext := path.Ext(name)

		stem := base[:len(base)-len(ext)]
		input := path.Join(path.Dir(name), stem)
		switch ext {
		case ".golden":
			g := gs[input]
			g.Golden = name
			gs[input] = g
		case ".golden-config":
			g := gs[input]
			g.Config = name
			gs[input] = g
		default:
			g := gs[name]
			g.Input = name
			gs[name] = g
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	groups = make([]Group, 0, len(gs))
	for _, g := range gs {
		if g.Input == "" {
			continue
		}
		groups = append(groups, g)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Input < groups[j].Input
	})
	return groups, nil
}

// Returns whether the file is valid and is not a directory.
func validGroupFile(f fs.FS, file string) bool {
	info, err := fs.Stat(f, file)
	return err == nil && !info.IsDir()
}

// GroupOf returns the group corresponding to the given file. The resulting
// group will be empty if the file does not exist or cannot be a part of a
// group.
func GroupOf(f fs.FS, file string) (group Group) {
	base := path.Base(file)
	if strings.HasPrefix(base, ".") {
		// Hidden files are ignored.
		return group
	}
	if !validGroupFile(f, file) {
		return group
	}
	dir := path.Dir(file)
	switch ext := path.Ext(file); ext {
	case ".golden":
		group.Golden = file
		base := base[:len(base)-len(ext)]
		if p := path.Join(dir, base); validGroupFile(f, p) {
			group.Input = p
		}
		if p := path.Join(dir, base+".golden-config"); validGroupFile(f, p) {
			group.Config = p
		}
	case ".golden-config":
		group.Config = file
		base := base[:len(base)-len(ext)]
		if p := path.Join(dir, base); validGroupFile(f, p) {
			group.Input = p
		}
		if p := path.Join(dir, base+".golden"); validGroupFile(f, p) {
			group.Golden = p
		}
	default:
		group.Input = file
		if p := path.Join(dir, base+".golden"); validGroupFile(f, p) {
			group.Golden = p
		}
		if p := path.Join(dir, base+".golden-config"); validGroupFile(f, p) {
			group.Config = p
		}
	}
	return group
}
