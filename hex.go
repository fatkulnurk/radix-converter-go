package radixconverter

const hexCharset = "0123456789abcdef"

type Hex struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewHex() *Hex {
	return &Hex{
		charset:   hexCharset,
		base:      uint64(len(hexCharset)),
		charIndex: BuildIndex(hexCharset),
	}
}

func (h *Hex) Encode(number uint64) string {
	return Encode(number, h.charset, h.base)
}

func (h *Hex) Decode(encoded string) (uint64, error) {
	return Decode(encoded, h.base, h.charIndex)
}

func (h *Hex) Charset() string {
	return h.charset
}
