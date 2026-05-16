package radixconverter

type IDConverter interface {
	Encode(number uint64) string

	Decode(encoded string) (uint64, error)
}
