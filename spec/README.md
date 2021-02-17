# File formats
This document describes the formats for each type of a file in a group. A group
consists of several types of file:
- [Input](#user-content-input-format): The file to be parsed.
- [Golden](#user-content-golden-format): The output file. Has the `.golden`
  extension.
- [Config](#user-content-config-format): Configures how the input file is
  parsed. Has the `.golden-config` extension.

## Input format
An input file has one of several formats, determined by the file extension.
Listed are the specifications for the various Roblox data formats:

- [`rbxl`](format/rbxl.md): Binary place.
- [`rbxlx`](format/rbxlx.md): XML place.
- [`rbxm`](format/rbxm.md): Binary model.
- [`rbxmx`](format/rbxmx.md): XML model.

## Config format
A config file is formatted as JSON, according to [RFC 8259][RFC8259]. It
consists of an object with the following fields:

Field  | Type   | Description
-------|--------|------------
Format | string | The format of the input. Overrides the input file extension.
Struct | string | The structure of output Data.

Each field is optional, with defaults that depend on the input format.

## Golden format
A golden file is formatted as JSON, according to [RFC 8259][RFC8259]. The
specification of a golden file format describes the structure of such JSON data.
It may also provide recommendations for formatting the content of the file (i.e.
indentation).

For two golden files to be considered equal, they must be semantically
equivalent according to their JSON structure. The raw content of the files do
not need to be binary equal.

[RFC8259]: https://tools.ietf.org/html/rfc8259

### Common formatting
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

### Common structure
All golden files share several common structures.

A member of an object will be described as having a "type", which may be
represented by one or more JSON types. The member *must* have a JSON value that
correctly represents this type. For example, the value of a `float` type may be
any JSON number, one of the JSON strings "Infinity", "-Infinity", or "NaN", but
not any other value.

Each golden file, unless specified otherwise, is a JSON object.

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

### Output structures
A number of structures are defined for produced golden files. Each structure has
several pieces of information associated with it:

- **Struct:** The value of the "Struct" config field that causes the structure
  to be used.
- **Formats:** A list of supported input formats.

#### Model
[**Specification**](golden/model.md)

The model structure displays the structure of an instance tree.

- **Struct:** `model`
- **Formats:** `rbxl`, `rbxm`, `rbxlx`, `rbxmx`

#### Binary
[**Specification**](golden/binary.md)

The binary structure displays the low-level binary structure of a file.

- **Struct:** `binary`
- **Formats:** `rbxl`, `rbxm`

#### XML
[**Specification**](golden/xml.md)

The XML structure displays the low-level XML structure of a file.

- **Struct:** `xml`
- **Formats:** `rbxl` (legacy only), `rbxm` (legacy only), `rbxlx`, `rbxmx`

#### Error
[**Specification**](golden/error.md)

The error structure displays errors that are expected to occur.

- **Struct:** `error`
- **Formats:** any
