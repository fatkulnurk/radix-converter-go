package radixconverter

// ConverterType represents a built-in converter strategy.
type ConverterType string

const (
	TypeBase62            ConverterType = "base62"
	TypeAlphanumericUpper ConverterType = "alphanumeric_upper"
	TypeAlphanumericLower ConverterType = "alphanumeric_lower"
	TypeAlphaOnly         ConverterType = "alpha_only"
)
