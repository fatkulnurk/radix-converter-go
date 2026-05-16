package converter

import (
	"github.com/fatkulnurk/radix-converter-go/radixerrors"
	"github.com/fatkulnurk/radix-converter-go/strategies"
)

// ConverterFactory creates converter instances.
// It checks the CustomConverterRegistry for custom types first,
// then falls back to built-in strategies.
type ConverterFactory struct{}

// NewConverterFactory creates a new ConverterFactory.
func NewConverterFactory() *ConverterFactory {
	return &ConverterFactory{}
}

// Make returns an IDConverter for the given type.
// If type is a custom converter name, it retrieves it from the registry.
// If type is a built-in ConverterType, it creates a new instance.
// Returns an error for unknown types.
func (f *ConverterFactory) Make(t ConverterType) (IDConverter, error) {
	// Check custom converters first.
	name := string(t)
	if custom, err := CustomRegistry.Get(name); err == nil {
		return custom, nil
	}

	// Fall back to built-in converters.
	switch t {
	case Base62:
		return strategies.NewBase62(), nil
	case AlphanumericUpper:
		return strategies.NewAlphanumericUpper(), nil
	case AlphanumericLower:
		return strategies.NewAlphanumericLower(), nil
	case AlphaOnly:
		return strategies.NewAlphaOnly(), nil
	default:
		return nil, radixerrors.NewUnknownConverterError(name)
	}
}

// MakeByName is a convenience method that accepts a string directly.
func (f *ConverterFactory) MakeByName(name string) (IDConverter, error) {
	// Check custom converters first.
	if custom, err := CustomRegistry.Get(name); err == nil {
		return custom, nil
	}

	// Fall back to built-in converters.
	switch ConverterType(name) {
	case Base62:
		return strategies.NewBase62(), nil
	case AlphanumericUpper:
		return strategies.NewAlphanumericUpper(), nil
	case AlphanumericLower:
		return strategies.NewAlphanumericLower(), nil
	case AlphaOnly:
		return strategies.NewAlphaOnly(), nil
	default:
		return nil, radixerrors.NewUnknownConverterError(name)
	}
}
