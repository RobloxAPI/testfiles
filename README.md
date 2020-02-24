# Roblox Test Files
The **testfiles** repository contains a database of files for testing
implementations of various [Roblox](https://corp.roblox.com) file formats.

## Philosophy

- **Specification**: How a thing is supposed to behave.
- **Implementation**: How the thing actually behaves.
- Ideally, the implementation matches the specification.
- In reality, this never quite happens.
- It is hard to have a correct implementation without a specification.
- It is harder to have a correct implementation if it *is* the specification!
- Comparing an implementation to a specification is often difficult, if not
  impossible.

### Enter: golden files
Key points:

- A program uses an implementation to write a **golden file**, which serves as a
  snapshot representing the specification.
- The golden file should be extremely readable, as it will be reviewed manually
  to determine its correctness.
- When the implementation is modified, it is compared against the current golden
  file to check for regressions.
- If the specification needs to change, the golden file can be rewritten by the
  program with the current implementation.
- Diffing can be used to inspect the correctness of the golden file after it has
  been updated.

In general, the content of a golden file should be line-based to interact better
with the diffing of version control systems. One unit of information per line.
Other than that, the content is fairly free-form; it should focus on being
parsable by human eyeballs.

## Structure

### Spec directory
Descriptions for known file formats are contained within the [`spec`](spec)
directory. Also contained are descriptions of golden file formats.

### Data directory
All test files are contained within the [`data`](data) directory. Files within
this directory are structured according the following rules:

- A **directory** is used only for organization, and is meant to be visited
  recursively.
- A **hidden file** is any file that starts with a `.`. These are ignored.
- A **golden file** is any file with the `.golden` extension.
- An **input file** is any other file.

An input file is mapped to a golden file when the file's name matches the golden
file's name without the extension. e.g. `Baseplate.rbxl` maps to
`Baseplate.rbxl.golden`.

#### Input file names
The **extension** of the file name determines the format used to decode the
file. For example, the file name:

	numberSequence.rbxl

will be decoded with the `rbxl` format.

#### Directives
An input file may begin with a number of **directives**, which control how the
file is tested. A directive is a line that begins with a `#`.

See [Directives](spec/README.md#user-content-directives) for more information.

## Testing
To test an implementation against the database, a program must be written. The
program should satisfy the following properties:

- The program should receive a directory, and iterate through the files within
  it:
	- Directories are iterated recursively.
	- Files starting with `.` are ignored.
	- Files with the `.golden` extension are golden files.
	- All other files are input files.
- The program must parse the name of an input file for the format.
- The program must parse directives within the input file according to the
  [directives spec](spec/README.md#user-content-directives).
- An input file with an unknown extension can be ignored. Unknown directives may
  also be ignored.
- For each valid input file, the program must produce a JSON structure that is
  semantically equal to the content of the corresponding golden file, in order
  for the implementation to be considered correct.
- If there is no corresponding golden file, then the content must be considered
  empty.
- If the JSON structure does not match, the program should output a
  human-readable difference between the produced structure and the content.
- If some sort of "update" flag is explicitly provided to the program, then the
  produced structure should be written to the golden file. The program should
  output the difference.

## Licensing
Files within the testfiles repository, including input files, golden files, and
documents, unless noted otherwise, are licensed to the testfiles contributors
under the [CC-BY-SA-4.0](LICENSE) license.
