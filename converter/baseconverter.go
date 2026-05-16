package converter

import (
	"github.com/fatkulnurk/radix-converter-go/radixerrors"
)

// BaseConverter provides the core radix conversion algorithm.
// It is designed to be embedded via composition in concrete strategy types.
// Deprecated: Use strategies package directly. This type is kept for backward compatibility.
type BaseConverter struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

// NewBaseConverter creates a new BaseConverter with the given charset.
// It panics if charset is empty.
func NewBaseConverter(charset string) *BaseConverter {
	if charset == "" {
		panic("radix converter: charset cannot be empty")
	}

	bc := &BaseConverter{
		charset:   charset,
		base:      uint64(len(charset)),
		charIndex: make(map[rune]int, len(charset)),
	}

	for i, r := range charset {
		bc.charIndex[r] = i
	}

	return bc
}

// Encode converts a non-negative integer to its string representation.
func (bc *BaseConverter) Encode(number uint64) string {
	if number == 0 {
		return bc.charset[:1]
	}

	var result []byte
	for number > 0 {
		remainder := number % bc.base
		result = append(result, bc.charset[remainder])
		number /= bc.base
	}

	// Reverse in place.
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Decode converts a string representation back to its integer value.
func (bc *BaseConverter) Decode(encoded string) (uint64, error) {
	if encoded == "" {
		return 0, radixerrors.ErrEmptyEncoded
	}

	var result uint64
	for _, char := range encoded {
		pos, ok := bc.charIndex[char]
		if !ok {
			return 0, radixerrors.NewInvalidCharError(char)
		}
		result = result*bc.base + uint64(pos)
	}

	return result, nil
}

// Charset returns the charset used by this converter.
func (bc *BaseConverter) Charset() string {
	return bc.charset
}

// Base returns the base (radix) of this converter.
func (bc *BaseConverter) Base() uint64 {
	return bc.base
}
