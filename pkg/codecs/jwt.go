package codecs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// JWTCodec handles JWT token decoding (header and payload only, no signature verification).
type JWTCodec struct{}

func (c *JWTCodec) Name() string      { return "jwt" }
func (c *JWTCodec) Aliases() []string { return []string{"json-web-token"} }

func (c *JWTCodec) Encode(data []byte) ([]byte, error) {
	return nil, fmt.Errorf("jwt encoding requires header and payload — use jwt-decode for analysis only")
}

func (c *JWTCodec) Decode(data []byte) ([]byte, error) {
	s := strings.TrimSpace(string(data))
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid JWT format: expected 3 parts, got %d", len(parts))
	}

	// Decode header
	headerBytes, err := base64RawURLEncodeDecode(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	// Decode payload
	payloadBytes, err := base64RawURLEncodeDecode(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// Parse JSON
	var header, payload interface{}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return nil, fmt.Errorf("failed to parse header JSON: %w", err)
	}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse payload JSON: %w", err)
	}

	result := map[string]interface{}{
		"header":    header,
		"payload":   payload,
		"signature": parts[2],
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}
	return output, nil
}

func (c *JWTCodec) CanDecode(data []byte) bool {
	s := strings.TrimSpace(string(data))
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return false
	}
	// Check that all parts look like base64url
	for _, p := range parts {
		if len(p) == 0 {
			return false
		}
		for _, ch := range p {
			if !((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '-' || ch == '_') {
				return false
			}
		}
	}
	return true
}

// base64RawURLEncodeDecode decodes a base64url-encoded string without padding.
func base64RawURLEncodeDecode(s string) ([]byte, error) {
	// Add padding if necessary
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
