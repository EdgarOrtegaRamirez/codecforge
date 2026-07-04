# AGENTS.md

## Project Overview

CodecForge is a Go CLI tool for universal text encoding and decoding. It provides a consistent interface for working with 13+ encoding formats, from Base64 to Morse code.

## Architecture

```
codecforge/
├── cmd/codecforge/      # CLI entry point (cobra-style, no deps)
├── pkg/codecs/          # Codec implementations
│   ├── codec.go         # Codec interface + Registry
│   ├── base64.go        # Base64 (standard/URL/raw)
│   ├── url.go           # URL percent encoding
│   ├── html.go          # HTML entity encoding
│   ├── hex.go           # Hexadecimal
│   ├── rot.go           # ROT13/ROT47
│   ├── binary.go        # Binary bit representation
│   ├── jwt.go           # JWT decode (header+payload only)
│   ├── unicode.go       # Unicode escape sequences (\uXXXX, \UXXXXXXXX)
│   ├── morse.go         # International Morse code
│   └── caesar.go        # Caesar cipher (configurable shift)
└── tests/               # Test suite (external test package)
```

## Key Components

### Codec Interface (`pkg/codecs/codec.go`)
Every codec implements:
- `Name()` — canonical name
- `Aliases()` — alternative names
- `Encode(data) → (encoded, error)`
- `Decode(data) → (decoded, error)`
- `CanDecode(data) → bool` — heuristic detection

### Registry
- Registers all built-in codecs by name and aliases
- `Get(name)` — retrieve codec by name (case-insensitive)
- `Detect(data)` — try all codecs and return which ones can decode the input

## Commands

```bash
go build -o codecforge ./cmd/codecforge/
echo "text" | ./codecforge encode <codec>
echo "encoded" | ./codecforge decode <codec>
echo "data" | ./codecforge detect
./codecforge list
./codecforge <codec> [input]    # shorthand for encode
./codecforge <codec> -d [input] # shorthand for decode
```

## Testing

```bash
go test ./tests/ -v
```

All tests use external test package (`package codecs_test`) for black-box testing.

## Adding a New Codec

1. Create `pkg/codecs/<name>.go` implementing the `Codec` interface
2. Register it in `pkg/codecs/codec.go` `NewRegistry()`
3. Add tests in `tests/codecs_test.go`
4. Run `go test ./tests/ -v` to verify

## Dependencies

- Go standard library only (no external dependencies)
