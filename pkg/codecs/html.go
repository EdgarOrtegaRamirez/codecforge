package codecs

import (
	"html"
	"strings"
)

// HTMLEntityCodec handles HTML entity encoding/decoding.
type HTMLEntityCodec struct{}

func (c *HTMLEntityCodec) Name() string      { return "html" }
func (c *HTMLEntityCodec) Aliases() []string { return []string{"html-entity", "htmlescape"} }

func (c *HTMLEntityCodec) Encode(data []byte) ([]byte, error) {
	return []byte(html.EscapeString(string(data))), nil
}

func (c *HTMLEntityCodec) Decode(data []byte) ([]byte, error) {
	return []byte(html.UnescapeString(string(data))), nil
}

func (c *HTMLEntityCodec) CanDecode(data []byte) bool {
	s := string(data)
	if len(s) == 0 {
		return false
	}
	// Check for common HTML entities
	entities := []string{"&amp;", "&lt;", "&gt;", "&quot;", "&#", "&nbsp;"}
	lower := strings.ToLower(s)
	for _, e := range entities {
		if strings.Contains(lower, e) {
			return true
		}
	}
	return false
}
