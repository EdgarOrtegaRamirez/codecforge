package codecs

import (
	"fmt"
	"strconv"
	"strings"
)

// UnicodeEscapeCodec handles Unicode escape sequence encoding/decoding.
type UnicodeEscapeCodec struct{}

func (c *UnicodeEscapeCodec) Name() string      { return "unicode" }
func (c *UnicodeEscapeCodec) Aliases() []string { return []string{"unicode-escape", "u-escape"} }

func (c *UnicodeEscapeCodec) Encode(data []byte) ([]byte, error) {
	var sb strings.Builder
	for _, r := range string(data) {
		if r < 128 {
			sb.WriteRune(r)
		} else if r > 0xFFFF {
			// Use \U with 8 hex digits for characters beyond BMP
			sb.WriteString(fmt.Sprintf("\\U%08x", r))
		} else {
			// Use \u with 4 hex digits for BMP characters
			sb.WriteString(fmt.Sprintf("\\u%04x", r))
		}
	}
	return []byte(sb.String()), nil
}

func (c *UnicodeEscapeCodec) Decode(data []byte) ([]byte, error) {
	s := string(data)
	var sb strings.Builder
	i := 0
	for i < len(s) {
		if i+1 < len(s) && s[i] == '\\' {
			next := s[i+1]
			if next == 'U' && i+9 < len(s) {
				// \U with 8 hex digits
				hex := s[i+2 : i+10]
				r, err := strconv.ParseInt(hex, 16, 32)
				if err != nil {
					sb.WriteByte(s[i])
					i++
				} else {
					sb.WriteRune(rune(r))
					i += 10
				}
			} else if next == 'u' && i+5 < len(s) {
				// \u with 4 hex digits
				hex := s[i+2 : i+6]
				r, err := strconv.ParseInt(hex, 16, 32)
				if err != nil {
					sb.WriteByte(s[i])
					i++
				} else {
					sb.WriteRune(rune(r))
					i += 6
				}
			} else {
				sb.WriteByte(s[i])
				i++
			}
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}
	return []byte(sb.String()), nil
}

func (c *UnicodeEscapeCodec) CanDecode(data []byte) bool {
	s := string(data)
	return strings.Contains(s, `\u`) || strings.Contains(s, `\U`)
}
