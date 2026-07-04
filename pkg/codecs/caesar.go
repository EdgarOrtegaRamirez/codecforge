package codecs

import (
	"strings"
)

// CaesarCodec handles Caesar cipher encoding/decoding with configurable shift.
type CaesarCodec struct {
	Shift int
}

func (c *CaesarCodec) Name() string {
	return "caesar"
}

func (c *CaesarCodec) Aliases() []string {
	return []string{"caesar-cipher"}
}

func (c *CaesarCodec) Encode(data []byte) ([]byte, error) {
	return []byte(caesarShift(string(data), c.Shift)), nil
}

func (c *CaesarCodec) Decode(data []byte) ([]byte, error) {
	return []byte(caesarShift(string(data), 26-c.Shift)), nil
}

func (c *CaesarCodec) CanDecode(data []byte) bool {
	return false // Cannot reliably detect Caesar cipher
}

func caesarShift(s string, shift int) string {
	var sb strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			sb.WriteRune('a' + (r-'a'+rune(shift))%26)
		case r >= 'A' && r <= 'Z':
			sb.WriteRune('A' + (r-'A'+rune(shift))%26)
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
