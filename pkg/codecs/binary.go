package codecs

import (
	"fmt"
	"strings"
)

// BinaryCodec handles binary representation encoding/decoding.
type BinaryCodec struct{}

func (c *BinaryCodec) Name() string    { return "binary" }
func (c *BinaryCodec) Aliases() []string { return []string{"bin"} }

func (c *BinaryCodec) Encode(data []byte) ([]byte, error) {
	var sb strings.Builder
	for i, b := range data {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprintf("%08b", b))
	}
	return []byte(sb.String()), nil
}

func (c *BinaryCodec) Decode(data []byte) ([]byte, error) {
	s := strings.TrimSpace(string(data))
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	if len(s) == 0 || len(s)%8 != 0 {
		return nil, fmt.Errorf("invalid binary string length: %d (must be multiple of 8)", len(s))
	}
	result := make([]byte, 0, len(s)/8)
	for i := 0; i < len(s); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			b = b<<1 | (s[i+j] - '0')
		}
		result = append(result, b)
	}
	return result, nil
}

func (c *BinaryCodec) CanDecode(data []byte) bool {
	s := strings.TrimSpace(string(data))
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	if len(s) == 0 || len(s)%8 != 0 {
		return false
	}
	for _, ch := range s {
		if ch != '0' && ch != '1' {
			return false
		}
	}
	return true
}
