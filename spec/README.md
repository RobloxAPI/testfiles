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
A sequence of bytes. Represented by a JSON array of numbers, each being an
integer between 0 and 255. The array may be formatted by wrapping onto a new
line every 16 values. Spaces may be added before a value to align it to a width
of 3 characters.

For example, the following string:

	Strange game.
	The only winning move
	is not to play.

may be formatted as:

```json
[
	 83,116,114, 97,110,103,101, 32,103, 97,109,101, 46, 10, 84,104,
	101, 32,111,110,108,121, 32,119,105,110,110,105,110,103, 32,109,
	111,118,101, 10,105,115, 32,110,111,116, 32,116,111, 32,112,108,
	 97,121, 46
]
```

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
