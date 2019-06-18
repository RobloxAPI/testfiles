# Formats

## Input formats
An input file has one of several formats, determined by the file extension.
Listed are the specifications for the various Roblox data formats:

- [`rbxl`](format/rbxl.md): Binary place.
- [`rbxlx`](format/rbxlx.md): XML place.
- [`rbxm`](format/rbxm.md): Binary model.
- [`rbxmx`](format/rbxmx.md): XML model.
- [~`mesh`~](format/mesh.md): Mesh data.
- [~`terrain`~](format/terrain.md): Terrain data.
- [~`csgphs`~](format/csgphs.md): CSG physical data.

## Golden formats
The formats of golden files have several common primitives.

Each line usually contains a **field**, which consists of a name paired with a
**value**. A field name is usually constant.

```
Field: Value
```

Depending on the content, the value may appear on multiple lines with an
additional level of indentation.

```
Field
	Value
```

A value may be a **struct**, which consists of a number of sub-fields.

```
Struct
	FieldA: Value1
	FieldB: Value2
	FieldC
		Foo: 1
		Bar: 2
```

A **list** is like a struct, except that each field is an index.

```
ListOfValues
	0: ValueA
	1: ValueB
	2: ValueC
ListOfStructs
	0
		FieldA: Value1
		FieldB: Value2
	1
		FieldA: Value1
		FieldB: Value2
```

Each golden file, unless specified otherwise, is a struct with three fields:

- `Format`: The format of the corresponding input file.
- `Struct`: A value indicating the structure of the `Data` field, as determined
  by directives in the input file.
- `Data`: The value representing the content of the input file through the
  determined structure.

### Model format
The model format displays the structure of an instance tree.

[**Specification**](golden/model.md)

Struct  | Directive | Formats
--------|-----------|--------
`model` | `model`   | `rbxl`, `rbxm`, `rbxlx`, `rbxmx`

### Binary format
The binary format displays the binary structure of an instance tree file.

[**Specification**](golden/binary.md)

Struct   | Directive | Formats
---------|-----------|--------
`binary` | `format`  | `rbxl`, `rbxm`

### XML format
The XML format displays the XML structure of an instance tree file.

[**Specification**](golden/xml.md)

Struct | Directive | Formats
-------|-----------|--------
`xml`  | `format`  | `rbxl` (legacy only), `rbxm` (legacy only), `rbxlx`, `rbxmx`

### Error format
The error format displays errors that are expected to occur.

[**Specification**](golden/error.md)

Struct  | Directive | Formats
--------|-----------|--------
`error` | `error`   | Any
