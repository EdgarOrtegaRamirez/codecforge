package codecs

import (
	"net/url"
)

// URLEncodeCodec handles URL percent encoding/decoding.
type URLEncodeCodec struct{}

func (c *URLEncodeCodec) Name() string    { return "url" }
func (c *URLEncodeCodec) Aliases() []string { return []string{"urlencode", "percent"} }

func (c *URLEncodeCodec) Encode(data []byte) ([]byte, error) {
	return []byte(url.QueryEscape(string(data))), nil
}

func (c *URLEncodeCodec) Decode(data []byte) ([]byte, error) {
	s, err := url.QueryUnescape(string(data))
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (c *URLEncodeCodec) CanDecode(data []byte) bool {
	s := string(data)
	if len(s) == 0 {
		return false
	}
	// Check for percent-encoded characters
	for i := 0; i < len(s)-2; i++ {
		if s[i] == '%' {
			return true
		}
	}
	return false
}
