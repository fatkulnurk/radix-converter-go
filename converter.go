package radixconverter

// BuildIndex creates a character-to-index map from a charset.
func BuildIndex(charset string) map[rune]int {
	index := make(map[rune]int, len(charset))
	for i, r := range charset {
		index[r] = i
	}
	return index
}

// Encode performs the radix encoding algorithm.
func Encode(number uint64, charset string, base uint64) string {
	if number == 0 {
		return charset[:1]
	}

	var result []byte
	for number > 0 {
		remainder := number % base
		result = append(result, charset[remainder])
		number /= base
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Decode performs the radix decoding algorithm.
func Decode(encoded string, base uint64, charIndex map[rune]int) (uint64, error) {
	if encoded == "" {
		return 0, ErrEmptyEncoded
	}

	var result uint64
	for _, char := range encoded {
		pos, ok := charIndex[char]
		if !ok {
			return 0, NewInvalidCharError(char)
		}
		result = result*base + uint64(pos)
	}

	return result, nil
}
