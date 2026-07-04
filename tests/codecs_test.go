package codecs_test

import (
	"testing"

	"github.com/EdgarOrtegaRamirez/codecforge/pkg/codecs"
)

func TestBase64Encode(t *testing.T) {
	c := &codecs.Base64Codec{}
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "SGVsbG8sIFdvcmxkIQ=="},
		{"", ""},
		{"a", "YQ=="},
		{"ab", "YWI="},
		{"abc", "YWJj"},
		{"\x00\x01\x02", "AAEC"},
	}
	for _, tt := range tests {
		result, err := c.Encode([]byte(tt.input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", tt.input, err)
			continue
		}
		if string(result) != tt.expected {
			t.Errorf("Encode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	c := &codecs.Base64Codec{}
	tests := []struct {
		input    string
		expected string
	}{
		{"SGVsbG8sIFdvcmxkIQ==", "Hello, World!"},
		{"YQ==", "a"},
		{"YWI=", "ab"},
		{"YWJj", "abc"},
		{"AAEC", "\x00\x01\x02"},
	}
	for _, tt := range tests {
		result, err := c.Decode([]byte(tt.input))
		if err != nil {
			t.Errorf("Decode(%q) error: %v", tt.input, err)
			continue
		}
		if string(result) != tt.expected {
			t.Errorf("Decode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBase64RoundTrip(t *testing.T) {
	c := &codecs.Base64Codec{}
	inputs := []string{"Hello, World!", "test\x00\x01\x02", "", "abc", "special chars: !@#$%^&*()"}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestBase64CanDecode(t *testing.T) {
	c := &codecs.Base64Codec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"SGVsbG8=", true},
		{"not base64!!", false},
		{"", false},
		{"AAAA", true},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestBase64URLRoundTrip(t *testing.T) {
	c := &codecs.Base64URLCodec{}
	input := "Hello, World! 🌍"
	encoded, err := c.Encode([]byte(input))
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	decoded, err := c.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if string(decoded) != input {
		t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
	}
}

func TestURLEncodeRoundTrip(t *testing.T) {
	c := &codecs.URLEncodeCodec{}
	inputs := []string{
		"hello world",
		"key=value&other=test",
		"special chars: !@#$%^&*()",
		"unicode: café résumé",
		"",
	}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestURLEncodeCanDecode(t *testing.T) {
	c := &codecs.URLEncodeCodec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello%20world", true},
		{"key=value", false},
		{"hello+world", false},
		{"", false},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestHTMLEntityRoundTrip(t *testing.T) {
	c := &codecs.HTMLEntityCodec{}
	inputs := []string{
		"<div class=\"test\">Hello & World</div>",
		"A & B < C > D",
		"",
		"no entities here",
	}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestHTMLEncodeEntities(t *testing.T) {
	c := &codecs.HTMLEntityCodec{}
	tests := []struct {
		input    string
		expected string
	}{
		{"<", "&lt;"},
		{">", "&gt;"},
		{"&", "&amp;"},
		{`"`, "&#34;"},
		{"hello", "hello"},
	}
	for _, tt := range tests {
		result, err := c.Encode([]byte(tt.input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", tt.input, err)
			continue
		}
		if string(result) != tt.expected {
			t.Errorf("Encode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestHexRoundTrip(t *testing.T) {
	c := &codecs.HexCodec{}
	inputs := []string{
		"Hello",
		"\x00\x01\x02\x03",
		"",
		"abc",
	}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestHexDecodeWithPrefix(t *testing.T) {
	c := &codecs.HexCodec{}
	result, err := c.Decode([]byte("0x48656c6c6f"))
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if string(result) != "Hello" {
		t.Errorf("Decode(0x48656c6c6f) = %q, want %q", result, "Hello")
	}
}

func TestHexCanDecode(t *testing.T) {
	c := &codecs.HexCodec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"48656c6c6f", true},
		{"0x48656c6c6f", true},
		{"gggg", false},
		{"", false},
		{"abc", false},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestROT13RoundTrip(t *testing.T) {
	c := &codecs.ROT13Codec{}
	inputs := []string{"Hello, World!", "ABCxyz", "123!@#", ""}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestROT13KnownValues(t *testing.T) {
	c := &codecs.ROT13Codec{}
	result, _ := c.Encode([]byte("HELLO"))
	if string(result) != "URYYB" {
		t.Errorf("ROT13(HELLO) = %q, want %q", result, "URYYB")
	}
}

func TestROT47RoundTrip(t *testing.T) {
	c := &codecs.ROT47Codec{}
	inputs := []string{"Hello, World!", "ABCxyz 123", "!@#$%^&*()", ""}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestBinaryEncode(t *testing.T) {
	c := &codecs.BinaryCodec{}
	result, err := c.Encode([]byte("Hi"))
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	expected := "01001000 01101001"
	if string(result) != expected {
		t.Errorf("BinaryEncode(Hi) = %q, want %q", result, expected)
	}
}

func TestBinaryDecode(t *testing.T) {
	c := &codecs.BinaryCodec{}
	input := "01001000 01101001"
	result, err := c.Decode([]byte(input))
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if string(result) != "Hi" {
		t.Errorf("BinaryDecode(%q) = %q, want %q", input, result, "Hi")
	}
}

func TestBinaryRoundTrip(t *testing.T) {
	c := &codecs.BinaryCodec{}
	inputs := []string{"Hello", "A", "\x00\xFF", "Test 123"}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestBinaryCanDecode(t *testing.T) {
	c := &codecs.BinaryCodec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"01001000 01101001", true},
		{"0100100001101001", true},
		{"20001000", false},
		{"01", false},
		{"", false},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestJWTDecode(t *testing.T) {
	c := &codecs.JWTCodec{}
	header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	payload := "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
	sig := "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	jwt := header + "." + payload + "." + sig

	result, err := c.Decode([]byte(jwt))
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	s := string(result)
	if len(s) == 0 {
		t.Fatal("Decode returned empty result")
	}
	if !containsStr(s, "\"header\"") || !containsStr(s, "\"payload\"") || !containsStr(s, "\"signature\"") {
		t.Errorf("Decode result missing expected fields: %s", s)
	}
}

func TestJWTCannotDecode(t *testing.T) {
	c := &codecs.JWTCodec{}
	tests := []struct {
		input string
	}{
		{"not-a-jwt"},
		{"only.two"},
		{"one two three four"},
		{""},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result {
			t.Errorf("CanDecode(%q) = true, want false", tt.input)
		}
	}
}

func TestJWTCannotEncode(t *testing.T) {
	c := &codecs.JWTCodec{}
	_, err := c.Encode([]byte("test"))
	if err == nil {
		t.Error("JWT Encode should return error")
	}
}

func TestUnicodeEscapeRoundTrip(t *testing.T) {
	c := &codecs.UnicodeEscapeCodec{}
	inputs := []string{
		"Hello",
		"café",
		"日本語",
		"🌍",
		"",
		"mixed: abc 日本語 émojis 🎉",
	}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestUnicodeEscapeEncode(t *testing.T) {
	c := &codecs.UnicodeEscapeCodec{}
	result, _ := c.Encode([]byte("A"))
	if string(result) != "A" {
		t.Errorf("UnicodeEscapeEncode(A) = %q, want %q", result, "A")
	}
	result, _ = c.Encode([]byte("é"))
	if string(result) != "\\u00e9" {
		t.Errorf("UnicodeEscapeEncode(é) = %q, want %q", result, "\\u00e9")
	}
	// Characters beyond BMP use \U with 8 hex digits
	result, _ = c.Encode([]byte("🌍"))
	if string(result) != "\\U0001f30d" {
		t.Errorf("UnicodeEscapeEncode(🌍) = %q, want %q", result, "\\U0001f30d")
	}
}

func TestUnicodeEscapeCanDecode(t *testing.T) {
	c := &codecs.UnicodeEscapeCodec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"\\u00e9", true},
		{"\\U0001F30D", true},
		{"\\u1f30d", true},
		{"hello", false},
		{"", false},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestMorseEncode(t *testing.T) {
	c := &codecs.MorseCodec{}
	result, err := c.Encode([]byte("SOS"))
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	expected := "... --- ..."
	if string(result) != expected {
		t.Errorf("MorseEncode(SOS) = %q, want %q", result, expected)
	}
}

func TestMorseRoundTrip(t *testing.T) {
	c := &codecs.MorseCodec{}
	inputs := []string{"SOS", "HELLO WORLD", "A B C"}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestMorseCanDecode(t *testing.T) {
	c := &codecs.MorseCodec{}
	tests := []struct {
		input    string
		expected bool
	}{
		{"... --- ...", true},
		{".- -...", true},
		{"hello", false},
		{"", false},
		{"123", false},
	}
	for _, tt := range tests {
		result := c.CanDecode([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("CanDecode(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestCaesarRoundTrip(t *testing.T) {
	c := &codecs.CaesarCodec{Shift: 13}
	inputs := []string{"Hello, World!", "ABCXYZ", "abc123!@#"}
	for _, input := range inputs {
		encoded, err := c.Encode([]byte(input))
		if err != nil {
			t.Errorf("Encode(%q) error: %v", input, err)
			continue
		}
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) error: %v", encoded, err)
			continue
		}
		if string(decoded) != input {
			t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
		}
	}
}

func TestCaesarShift(t *testing.T) {
	c := &codecs.CaesarCodec{Shift: 3}
	result, _ := c.Encode([]byte("abc"))
	if string(result) != "def" {
		t.Errorf("Caesar(3, abc) = %q, want %q", result, "def")
	}
	result, _ = c.Decode([]byte("def"))
	if string(result) != "abc" {
		t.Errorf("CaesarDecode(3, def) = %q, want %q", result, "abc")
	}
}

func TestCaesarShiftWrap(t *testing.T) {
	c := &codecs.CaesarCodec{Shift: 3}
	result, _ := c.Encode([]byte("XYZ"))
	if string(result) != "ABC" {
		t.Errorf("Caesar(3, XYZ) = %q, want %q", result, "ABC")
	}
}

func TestRegistry(t *testing.T) {
	r := codecs.NewRegistry()
	names := r.Names()
	if len(names) == 0 {
		t.Fatal("Registry has no codecs")
	}

	c, err := r.Get("base64")
	if err != nil {
		t.Fatalf("Get(base64) error: %v", err)
	}
	if c.Name() != "base64" {
		t.Errorf("Get(base64).Name() = %q, want %q", c.Name(), "base64")
	}

	c, err = r.Get("b64")
	if err != nil {
		t.Fatalf("Get(b64) error: %v", err)
	}
	if c.Name() != "base64" {
		t.Errorf("Get(b64).Name() = %q, want %q", c.Name(), "base64")
	}

	_, err = r.Get("unknown")
	if err == nil {
		t.Error("Get(unknown) should return error")
	}
}

func TestRegistryDetect(t *testing.T) {
	r := codecs.NewRegistry()

	detected := r.Detect([]byte("SGVsbG8="))
	found := false
	for _, name := range detected {
		if name == "base64" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Detect failed to identify base64: %v", detected)
	}

	detected = r.Detect([]byte("48656c6c6f"))
	found = false
	for _, name := range detected {
		if name == "hex" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Detect failed to identify hex: %v", detected)
	}

	detected = r.Detect([]byte("01001000 01101001"))
	found = false
	for _, name := range detected {
		if name == "binary" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Detect failed to identify binary: %v", detected)
	}
}

func TestBase64RawRoundTrip(t *testing.T) {
	c := &codecs.Base64RawCodec{}
	input := "Hello"
	encoded, err := c.Encode([]byte(input))
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	if len(encoded) > 0 && encoded[len(encoded)-1] == '=' {
		t.Errorf("Raw base64 should not have padding: %q", encoded)
	}
	decoded, err := c.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if string(decoded) != input {
		t.Errorf("Roundtrip failed: %q -> %q -> %q", input, encoded, decoded)
	}
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
