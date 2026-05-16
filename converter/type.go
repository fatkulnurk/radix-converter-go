package converter

// ConverterType represents a built-in converter strategy.
type ConverterType string

const (
	Base62             ConverterType = "base62"
	AlphanumericUpper  ConverterType = "alphanumeric_upper"
	AlphanumericLower  ConverterType = "alphanumeric_lower"
	AlphaOnly          ConverterType = "alpha_only"
)
