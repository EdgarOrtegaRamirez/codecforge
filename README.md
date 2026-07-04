# CodecForge 🔐

A universal text encoding/decoding toolkit for developers. One CLI to encode and decode across 13+ formats.

## Features

- **13 encoding formats** — Base64 (standard/URL/raw), URL, HTML entities, hex, ROT13, ROT47, binary, JWT, Unicode escapes, Morse code, Caesar cipher
- **Auto-detection** — Automatically identify encoding format from input
- **Consistent CLI** — Same interface for all codecs
- **Zero dependencies** — Uses only Go standard library
- **Fast** — Single Go binary, no runtime dependencies
- **Round-trip safe** — All codecs support encode/decode round-trips

## Installation

```bash
go install github.com/EdgarOrtegaRamirez/codecforge/cmd/codecforge@latest
```

Or build from source:

```bash
git clone https://github.com/EdgarOrtegaRamirez/codecforge
cd codecforge
go build -o codecforge ./cmd/codecforge/
```

## Usage

### Basic Encoding/Decoding

```bash
# Base64 encode
echo "Hello, World!" | codecforge encode base64
# Output: SGVsbG8sIFdvcmxkIQ==

# Base64 decode
echo "SGVsbG8sIFdvcmxkIQ==" | codecforge decode base64
# Output: Hello, World!

# URL encode
echo "hello world?key=value" | codecforge encode url
# Output: hello+world%3Fkey%3Dvalue

# URL decode
echo "hello+world%3Fkey%3Dvalue" | codecforge decode url
# Output: hello world?key=value
```

### Codec Shortcuts

```bash
# Shorthand: codecforge <codec> [input]
echo "Hello World" | codecforge base64
echo "SGVsbG8gV29ybGQ=" | codecforge base64 -d
codecforge morse "SOS"
# Output: ... --- ...
```

### Auto-Detection

```bash
echo "SGVsbG8=" | codecforge detect
# Output: Detected encodings:
#   - base64

echo "48656c6c6f" | codecforge detect
# Output: Detected encodings:
#   - hex
```

### File Input

```bash
codecforge encode base64 -f myfile.txt
codecforge decode hex -f encoded.txt
```

## Available Codecs

| Codec | Name | Aliases | Description |
|-------|------|---------|-------------|
| Base64 (standard) | `base64` | `b64` | Standard Base64 with padding |
| Base64 (URL-safe) | `base64url` | `b64url`, `b64u` | URL-safe Base64 with padding |
| Base64 (raw) | `base64raw` | `b64raw` | Raw Base64 without padding |
| URL encoding | `url` | `urlencode`, `percent` | URL percent encoding |
| HTML entities | `html` | `html-entity`, `htmlescape` | HTML entity encoding |
| Hex | `hex` | `hexadecimal` | Hexadecimal encoding |
| ROT13 | `rot13` | — | ROT13 substitution cipher |
| ROT47 | `rot47` | — | ROT47 substitution cipher |
| Binary | `binary` | `bin` | Binary bit representation |
| JWT | `jwt` | `json-web-token` | JWT token decode (header + payload) |
| Unicode escapes | `unicode` | `unicode-escape`, `u-escape` | Unicode escape sequences |
| Morse code | `morse` | `morsecode` | International Morse code |
| Caesar cipher | `caesar` | `caesar-cipher` | Caesar cipher (shift 13) |

## Examples

### JWT Analysis

```bash
echo "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.signature" | codecforge decode jwt
```

### Morse Code

```bash
codecforge morse "HELLO WORLD"
# Output: .... . .-.. .-.. --- / .-- --- .-. .-.. -..
```

### Unicode Escapes

```bash
echo "café résumé" | codecforge encode unicode
# Output: caf\u00e9 r\u00e9sum\u00e9

echo "日本語テスト" | codecforge encode unicode
# Output: \u65e5\u672c\u8a9e\u30c6\u30b9\u30c8
```

### HTML Entity Encoding

```bash
echo '<script>alert("XSS")</script>' | codecforge encode html
# Output: &lt;script&gt;alert(&#34;XSS&#34;)&lt;/script&gt;
```

## Architecture

```
codecforge/
├── cmd/codecforge/      # CLI entry point
│   └── main.go
├── pkg/codecs/          # Codec implementations
│   ├── codec.go         # Interface and registry
│   ├── base64.go        # Base64 variants
│   ├── url.go           # URL encoding
│   ├── html.go          # HTML entities
│   ├── hex.go           # Hexadecimal
│   ├── rot.go           # ROT13/ROT47
│   ├── binary.go        # Binary representation
│   ├── jwt.go           # JWT decode
│   ├── unicode.go       # Unicode escapes
│   ├── morse.go         # Morse code
│   └── caesar.go        # Caesar cipher
├── tests/               # Test suite
└── go.mod
```

## Adding Custom Codecs

Implement the `Codec` interface:

```go
type Codec interface {
    Name() string
    Aliases() []string
    Encode(data []byte) ([]byte, error)
    Decode(data []byte) ([]byte, error)
    CanDecode(data []byte) bool
}
```

Then register it in the registry.

## Testing

```bash
go test ./tests/ -v
```

## License

MIT
