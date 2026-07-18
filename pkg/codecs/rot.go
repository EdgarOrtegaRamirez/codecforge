package codecs

import (
	"strings"
)

// ROT13Codec handles ROT13 encoding/decoding.
type ROT13Codec struct{}

func (c *ROT13Codec) Name() string      { return "rot13" }
func (c *ROT13Codec) Aliases() []string { return []string{} }

func (c *ROT13Codec) Encode(data []byte) ([]byte, error) {
	return []byte(rot(string(data), 13)), nil
}

func (c *ROT13Codec) Decode(data []byte) ([]byte, error) {
	return []byte(rot(string(data), 13)), nil // ROT13 is its own inverse
}

func (c *ROT13Codec) CanDecode(data []byte) bool {
	return false // ROT13 detection is unreliable
}

// ROT47Codec handles ROT47 encoding/decoding.
type ROT47Codec struct{}

func (c *ROT47Codec) Name() string      { return "rot47" }
func (c *ROT47Codec) Aliases() []string { return []string{} }

func (c *ROT47Codec) Encode(data []byte) ([]byte, error) {
	return []byte(rot47(string(data))), nil
}

func (c *ROT47Codec) Decode(data []byte) ([]byte, error) {
	return []byte(rot47(string(data))), nil // ROT47 is its own inverse
}

func (c *ROT47Codec) CanDecode(data []byte) bool {
	return false
}

// rot applies a ROT shift to alphabetic characters.
func rot(s string, shift int) string {
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

// rot47 applies ROT47 to printable ASCII characters (33-126).
func rot47(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if r >= 33 && r <= 126 {
			sb.WriteRune(33 + ((r - 33 + 47) % 94))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
