package codecs

import (
	"encoding/base64"
)

// Base64Codec handles standard Base64 encoding/decoding.
type Base64Codec struct{}

func (c *Base64Codec) Name() string    { return "base64" }
func (c *Base64Codec) Aliases() []string { return []string{"b64"} }

func (c *Base64Codec) Encode(data []byte) ([]byte, error) {
	result := base64.StdEncoding.EncodeToString(data)
	return []byte(result), nil
}

func (c *Base64Codec) Decode(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func (c *Base64Codec) CanDecode(data []byte) bool {
	s := string(data)
	if len(s) == 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil && len(s)%4 == 0
}

// Base64URLCodec handles URL-safe Base64 encoding/decoding.
type Base64URLCodec struct{}

func (c *Base64URLCodec) Name() string    { return "base64url" }
func (c *Base64URLCodec) Aliases() []string { return []string{"b64url", "b64u"} }

func (c *Base64URLCodec) Encode(data []byte) ([]byte, error) {
	result := base64.URLEncoding.EncodeToString(data)
	return []byte(result), nil
}

func (c *Base64URLCodec) Decode(data []byte) ([]byte, error) {
	return base64.URLEncoding.DecodeString(string(data))
}

func (c *Base64URLCodec) CanDecode(data []byte) bool {
	s := string(data)
	if len(s) == 0 {
		return false
	}
	_, err := base64.URLEncoding.DecodeString(s)
	return err == nil
}

// Base64RawCodec handles raw Base64 (no padding) encoding/decoding.
type Base64RawCodec struct{}

func (c *Base64RawCodec) Name() string    { return "base64raw" }
func (c *Base64RawCodec) Aliases() []string { return []string{"b64raw"} }

func (c *Base64RawCodec) Encode(data []byte) ([]byte, error) {
	result := base64.RawStdEncoding.EncodeToString(data)
	return []byte(result), nil
}

func (c *Base64RawCodec) Decode(data []byte) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(string(data))
}

func (c *Base64RawCodec) CanDecode(data []byte) bool {
	s := string(data)
	if len(s) == 0 {
		return false
	}
	_, err := base64.RawStdEncoding.DecodeString(s)
	return err == nil
}
