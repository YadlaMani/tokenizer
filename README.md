# Interpreter Tokenizer

This project is a simple tokenizer written in Go. It reads a source file and prints out tokens for keywords, identifiers, numbers, strings, and symbols.

## Features

- Recognizes keywords, identifiers, numbers (integers and floats), strings, and common symbols.
- Handles comments and whitespace.
- Reports errors for invalid tokens and unterminated strings.

## Usage

To run the tokenizer on a file (e.g., `test.txt`):

```bash
go run main.go tokenize test.txt
```

## Example

Given the following input in `test.txt`:

```
for i in range 10
```

The output will be:

```
FOR for null
IDENTIFIER i null
IN in null
IDENTIFIER range null
NUMBER 10 10.0
EOF  null
```

## File Structure

- `main.go` - The main tokenizer implementation.
- `test.txt` - Example input file.
- `go.mod` - Go module file.

## Error Handling

- Unterminated strings and invalid numbers are reported with line numbers.
- Unexpected characters are reported as errors.

## License

MIT License
