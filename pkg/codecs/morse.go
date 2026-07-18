package codecs

import (
	"fmt"
	"strings"
)

// MorseCodec handles Morse code encoding/decoding.
type MorseCodec struct{}

var morseMap = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
	'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---", '3': "...--",
	'4': "....-", '5': ".....", '6': "-....", '7': "--...",
	'8': "---..", '9': "----.",
	'.': ".-.-.-", ',': "--..--", '?': "..--..", '\'': ".----.",
	'!': "-.-.--", '/': "-..-.", '(': "-.--.", ')': "-.--.-",
	'&': ".-...", ':': "---...", ';': "-.-.-.", '=': "-...-",
	'+': ".-.-.", '-': "-....-", '_': "..--.-", '"': ".-..-.",
	'$': "...-..-", '@': ".--.-.",
}

var reverseMorseMap map[string]rune

func init() {
	reverseMorseMap = make(map[string]rune)
	for k, v := range morseMap {
		reverseMorseMap[v] = k
	}
}

func (c *MorseCodec) Name() string      { return "morse" }
func (c *MorseCodec) Aliases() []string { return []string{"morsecode"} }

func (c *MorseCodec) Encode(data []byte) ([]byte, error) {
	var sb strings.Builder
	upper := strings.ToUpper(string(data))
	for i, r := range upper {
		if r == ' ' {
			// Word separator - use " / " format
			if i > 0 {
				sb.WriteString(" / ")
			}
		} else {
			// Letter separator
			if i > 0 && upper[i-1] != ' ' {
				sb.WriteByte(' ')
			}
			if code, ok := morseMap[r]; ok {
				sb.WriteString(code)
			}
		}
	}
	return []byte(sb.String()), nil
}

func (c *MorseCodec) Decode(data []byte) ([]byte, error) {
	s := strings.TrimSpace(string(data))
	if len(s) == 0 {
		return []byte{}, nil
	}

	// Split into words on " / "
	words := strings.Split(s, " / ")
	var result strings.Builder
	for wi, word := range words {
		if wi > 0 {
			result.WriteByte(' ')
		}
		// Split word into letters on " "
		letters := strings.Split(strings.TrimSpace(word), " ")
		for _, letter := range letters {
			letter = strings.TrimSpace(letter)
			if letter == "" {
				continue
			}
			if r, ok := reverseMorseMap[letter]; ok {
				result.WriteRune(r)
			} else {
				return nil, fmt.Errorf("unknown morse code: %s", letter)
			}
		}
	}
	return []byte(result.String()), nil
}

func (c *MorseCodec) CanDecode(data []byte) bool {
	s := strings.TrimSpace(string(data))
	if len(s) == 0 {
		return false
	}
	// Morse code uses only dots, dashes, spaces, and slashes
	for _, ch := range s {
		if ch != '.' && ch != '-' && ch != ' ' && ch != '/' {
			return false
		}
	}
	return strings.Contains(s, ".") || strings.Contains(s, "-")
}
