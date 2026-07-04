// Package codecs provides encoding and decoding operations for various formats.
package codecs

import (
	"fmt"
	"sort"
	"strings"
)

// Codec defines the interface for an encoding/decoding format.
type Codec interface {
	// Name returns the canonical name of the codec.
	Name() string
	// Aliases returns alternative names that can be used to refer to this codec.
	Aliases() []string
	// Encode encodes the input bytes and returns the encoded result.
	Encode(data []byte) ([]byte, error)
	// Decode decodes the input bytes and returns the decoded result.
	Decode(data []byte) ([]byte, error)
	// CanDecode returns true if this codec can handle the given input.
	CanDecode(data []byte) bool
}

// Registry holds all registered codecs.
type Registry struct {
	codecs map[string]Codec
}

// NewRegistry creates a new codec registry with all built-in codecs.
func NewRegistry() *Registry {
	r := &Registry{
		codecs: make(map[string]Codec),
	}
	// Register all built-in codecs
	builtins := []Codec{
		&Base64Codec{},
		&Base64URLCodec{},
		&Base64RawCodec{},
		&URLEncodeCodec{},
		&HTMLEntityCodec{},
		&HexCodec{},
		&ROT13Codec{},
		&ROT47Codec{},
		&BinaryCodec{},
		&JWTCodec{},
		&UnicodeEscapeCodec{},
		&MorseCodec{},
		&CaesarCodec{Shift: 13},
	}
	for _, c := range builtins {
		r.Register(c)
	}
	return r
}

// Register adds a codec to the registry under its name and aliases.
func (r *Registry) Register(c Codec) {
	r.codecs[strings.ToLower(c.Name())] = c
	for _, alias := range c.Aliases() {
		r.codecs[strings.ToLower(alias)] = c
	}
}

// Get retrieves a codec by name (case-insensitive).
func (r *Registry) Get(name string) (Codec, error) {
	c, ok := r.codecs[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("unknown codec: %s", name)
	}
	return c, nil
}

// Names returns a sorted list of all registered codec names (unique).
func (r *Registry) Names() []string {
	seen := make(map[string]bool)
	var names []string
	// Walk the map; for aliases pointing to same codec, use the codec's canonical Name()
	for _, c := range r.codecs {
		n := c.Name()
		if !seen[n] {
			seen[n] = true
			names = append(names, n)
		}
	}
	sort.Strings(names)
	return names
}

// Detect tries to identify the encoding format of the given input.
func (r *Registry) Detect(data []byte) []string {
	var detected []string
	seen := make(map[string]bool)
	for _, c := range r.codecs {
		name := c.Name()
		if seen[name] {
			continue
		}
		if c.CanDecode(data) {
			detected = append(detected, name)
			seen[name] = true
		}
	}
	return detected
}
