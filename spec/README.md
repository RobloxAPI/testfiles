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

## Directives
Regardless of the format, any input file may begin with a number of
**directives**. A directive is a line that starts with a `#` character. The
first line in the file that isn't a directive begins the actual **content** of
the file.

Directives have two formats:

- `#flag`: Causes the "flag" directive to be enabled.
- `#key:value`: Sets the "key" directive to "value".

`flag` and `key` may contain any character that isn't `:`, `\r`, or `\n`.
`value` may contain any character that isn't `\r` or `\n`. Leading and trailing
whitespace around `flag`, `key`, and `value` is ignored.

The following directives are defined:

Directive                         | Description
----------------------------------|------------
<code>**#format:** *value*</code> | Overrides the file extension, causing the content to be interpreted as the `value` format.
<code>**#output:** *value*</code> | Sets the format of the corresponding golden file to `value`.
<code>**#begin-content**</code>   | Causes directives to stop being parsed. The next line begins the actual content of the file.

# Golden formats
A golden file is formatted as JSON, according to [RFC 8259][RFC8259]. The
specification of a golden file format describes the structure of such JSON data.
It may also provide recommendations for formatting the content of the file (i.e.
indentation).

For two golden files to be considered equal, they must be semantically
equivalent according to their JSON structure. The raw content of the files do
not need to be binary equal.

[RFC8259]: https://tools.ietf.org/html/rfc8259

## Common formatting
For readability purposes, a golden file should be formatted such that one piece
of information is on a single line.

- Newlines should consist only of the line feed (`\n`) character (not `\r\n`).
- A value within an array or a member within an object should appear on its own
  line, and should be indented with tabs (`\t`).
- A single space character should appear between the `:` and the value of an
  object member.
- Opening structural delimiters should appear on the same line as the name.
- Closing structural delimiters should appear on a new line, unless the
  structure is empty.
- Object members should be sorted by key.

Example:
```json
{
	"Array": [
		1,
		2,
		3
	],
	"Bool": true,
	"Number": 0,
	"Object": {
		"EmptyArray": [],
		"EmptyObject": {}
	},
	"String": "foobar"
}
```

Depending on the structure, a specification may recommend more specific
formatting.

## Common structure
All golden files share several common structures.

A member of an object will be described as having a "type", which may be
represented by one or more JSON types. The member *must* have a JSON value that
correctly represents this type. For example, the value of a `float` type may be
any JSON number, one of the JSON strings "Infinity", "-Infinity", or "NaN", but
not any other value.

Each golden file, unless specified otherwise, is a JSON object (the "root") with
three members:

- `Format`: A string indicating the format of the corresponding input file.
- `Output`: A string indicating the structure of the `Data` field, as determined
  by directives in the input file.
- `Data`: The value representing the content of the input file through the
  determined structure.

### Primitive value types
Formats may refer to several common primitive value types:

#### bool
A boolean value. Represented by a JSON boolean.

```json
true
false
```

#### int
A signed integer in base 10. Represented by a JSON number. The number should not
have a fractional part.

```json
10
-10
```

#### uint
An unsigned integer in base 10. Represented by a JSON number. The number should
not have a fractional part.

```json
10
```

#### float
A floating point number. Represented by a JSON number. Infinity and NaN are
represented by literal JSON strings. The number should be formatted with enough
precision to accurately reproduce the intended value.

```json
3.1415926535897931
99003.78125
-0.55555555555555558
1
"Infinity"
"-Infinity"
"NaN"
```

#### bytes
A sequence of bytes. Represented by a JSON array of strings. Each JSON string
displays a line of the sequence in a hexdump format.

For example, the following string:

	Strange game.
	The only winning move
	is not to play.

is formatted as:

	[
		"| 53 74 72 61 6e 67 65 20  67 61 6d 65 2e 0a 54 68 |Strange game..Th|",
		"| 65 20 6f 6e 6c 79 20 77  69 6e 6e 69 6e 67 20 6d |e only winning m|",
		"| 6f 76 65 0a 69 73 20 6e  6f 74 20 74 6f 20 70 6c |ove.is not to pl|",
		"| 61 79 2e                                         |ay.|"
	]

- The display is wrapped to 16 bytes.
- An extra space is inserted after the 8th byte.
- Bytes outside the inclusive range of 32-126 are displayed as `.`.

If the length of the sequence is less than 16, then the display is shortened to
the length:

	[
		"| 53 74 72 61 6e 67 65 20  67 61 6d 65 |Strange game|"
	]

#### string
A unicode string. Represented by a JSON string. May contain any unicode
character. For characters that must be escaped, the shortest possible form is
preferred.

```json
"Strange game.\nThe only winning move\nis not to play."
```

A string must instead be displayed as [bytes](#user-content-bytes) if it
contains at least one non-Graphic character (as defined by Unicode) that isn't
the following:

	U+0008  backspace        \b
	U+0009  tab              \t
	U+000A  line feed        \n
	U+000C  form feed        \f
	U+000D  carriage return  \r

## Formats
Each format has several pieces of information associated with it:

- **Output:** The value displayed in the `Output` field of the root object when
  the format is used.
- **Directive:** The directive that causes the format to be used.
- **Formats:** A list of supported input formats.

### Model
The model format displays the structure of an instance tree.

[**Specification**](golden/model.md)

- **Output:** `model`
- **Directive:** `model`
- **Formats:** `rbxl`, `rbxm`, `rbxlx`, `rbxmx`

### Binary
The binary format displays the binary structure of an instance tree file.

[**Specification**](golden/binary.md)

- **Output:** `binary`
- **Directive:** `format`
- **Formats:** `rbxl`, `rbxm`

### XML
The XML format displays the XML structure of an instance tree file.

[**Specification**](golden/xml.md)

- **Output:** `xml`
- **Directive:** `format`
- **Formats:** `rbxl` (legacy only), `rbxm` (legacy only), `rbxlx`, `rbxmx`

### Error
The error format displays errors that are expected to occur.

[**Specification**](golden/error.md)

- **Output:** `error`
- **Directive:** `error`
- **Formats:** Any
