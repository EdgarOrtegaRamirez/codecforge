package codecs

import (
	"encoding/hex"
	"strings"
)

// HexCodec handles hexadecimal encoding/decoding.
type HexCodec struct{}

func (c *HexCodec) Name() string    { return "hex" }
func (c *HexCodec) Aliases() []string { return []string{"hexadecimal"} }

func (c *HexCodec) Encode(data []byte) ([]byte, error) {
	return []byte(hex.EncodeToString(data)), nil
}

func (c *HexCodec) Decode(data []byte) ([]byte, error) {
	s := string(data)
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")
	s = strings.TrimSpace(s)
	return hex.DecodeString(s)
}

func (c *HexCodec) CanDecode(data []byte) bool {
	s := string(data)
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")
	s = strings.TrimSpace(s)
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}
