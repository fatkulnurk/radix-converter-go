package converter

// IDConverter defines the interface for radix (base-N) encoding and decoding.
type IDConverter interface {
	// Encode converts a non-negative integer to its string representation in the converter's base.
	Encode(number uint64) string

	// Decode converts a string representation back to its original integer value.
	Decode(encoded string) (uint64, error)
}
