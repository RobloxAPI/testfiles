# Input formats
An input file has one of several formats, determined by the file extension.
Listed are the specifications for the various Roblox data formats:

- [`rbxl`](format/rbxl.md): Binary place.
- [`rbxlx`](format/rbxlx.md): XML place.
- [`rbxm`](format/rbxm.md): Binary model.
- [`rbxmx`](format/rbxmx.md): XML model.
- [~`mesh`~](format/mesh.md): Mesh data.
- [~`terrain`~](format/terrain.md): Terrain data.
- [~`csgphs`~](format/csgphs.md): CSG physical data.

# Golden formats

## Common
The formats of golden files share a common structure:

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

Each golden file, unless specified otherwise, is a standard struct with three
fields:

- `Format`: The format of the corresponding input file.
- `Struct`: A value indicating the structure of the `Data` field, as determined
  by directives in the input file.
- `Data`: The value representing the content of the input file through the
  determined structure.

Each format has several pieces of information associated with it:

- **Struct:** The value displayed in the `Struct` field of the standard struct
  when the format is used.
- **Directive:** The directive that causes the format to be used.
- **Formats:** A list of supported input formats.

Golden files are encoded in UTF-8.

### Primitive value types
Formats may use several common primitives value types:

#### bool
A boolean value.

	Value: true
	Value: false

#### int
A signed integer in base 10.

	Value: 10
	Value: -10

#### uint
An unsigned integer in base 10.

	Value: 10

#### float
A floating point number in a decimal format, following rules similar to printf.
Single precision uses `%.9g` as the formatter, and double precision uses
`%.17g`.

	Single: 3.14159274
	Single: 99003.7812
	Single: -0.555555582
	Double: 3.1415926535897931
	Double: 99003.78125
	Double: -0.55555555555555558

Rounding currently uses half-to-even. If you have an implementation that uses
half-away-to-zero or some other odd behavior, see issue #1.

#### bytes
A sequence of bytes displayed on multiple lines. The first line is a "Length"
field whose value is the length of the sequence. Subsequent lines display the
bytes in a hex-dump format.

For example, the following string:

	Strange game.
	The only winning move
	is not to play.

is formatted as:

	Value
		Length: 51
		| 53 74 72 61 6e 67 65 20  67 61 6d 65 2e 0a 54 68 |Strange game..Th|
		| 65 20 6f 6e 6c 79 20 77  69 6e 6e 69 6e 67 20 6d |e only winning m|
		| 6f 76 65 0a 69 73 20 6e  6f 74 20 74 6f 20 70 6c |ove.is not to pl|
		| 61 79 2e                                         |ay.|

- The display is wrapped to 16 bytes.
- An extra space is inserted after the 8th byte.
- Bytes outside the inclusive range of 32-126 are displayed as `.`.

If the length of the sequence is less than 16, then the display is shortened to
the length:

	Value
		Length: 12
		| 53 74 72 61 6e 67 65 20  67 61 6d 65 |Strange game|

#### string
A string contained between `"` characters.

	Value: "Strange game.\nThe only winning move\nis not to play."

The following characters are escaped:

Character | Escape
----------|-------
`"`       | `\"`
`\`       | `\\`
newline   | `\n`
tab       | `\t`

If a string contains non-printable characters, then it is instead formatted as
bytes.

## Model format
The model format displays the structure of an instance tree.

[**Specification**](golden/model.md)

- **Struct:** `model`
- **Directive:** `model`
- **Formats:** `rbxl`, `rbxm`, `rbxlx`, `rbxmx`

## Binary format
The binary format displays the binary structure of an instance tree file.

[**Specification**](golden/binary.md)

- **Struct:** `binary`
- **Directive:** `format`
- **Formats:** `rbxl`, `rbxm`

## XML format
The XML format displays the XML structure of an instance tree file.

[**Specification**](golden/xml.md)

- **Struct:** `xml`
- **Directive:** `format`
- **Formats:** `rbxl` (legacy only), `rbxm` (legacy only), `rbxlx`, `rbxmx`

## Error format
The error format displays errors that are expected to occur.

[**Specification**](golden/error.md)

- **Struct:** `error`
- **Directive:** `error`
- **Formats:** Any
