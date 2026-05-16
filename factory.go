package radixconverter

// ConverterFactory creates converter instances.
type ConverterFactory struct{}

func NewConverterFactory() *ConverterFactory {
	return &ConverterFactory{}
}

func (f *ConverterFactory) Make(t ConverterType) (IDConverter, error) {
	name := string(t)
	if custom, err := GlobalRegistry.Get(name); err == nil {
		return custom, nil
	}

	switch t {
	case TypeBase62:
		return NewBase62(), nil
	case TypeAlphanumericUpper:
		return NewAlphanumericUpper(), nil
	case TypeAlphanumericLower:
		return NewAlphanumericLower(), nil
	case TypeAlphaOnly:
		return NewAlphaOnly(), nil
	default:
		return nil, NewUnknownConverterError(name)
	}
}

func (f *ConverterFactory) MakeByName(name string) (IDConverter, error) {
	if custom, err := GlobalRegistry.Get(name); err == nil {
		return custom, nil
	}

	switch ConverterType(name) {
	case TypeBase62:
		return NewBase62(), nil
	case TypeAlphanumericUpper:
		return NewAlphanumericUpper(), nil
	case TypeAlphanumericLower:
		return NewAlphanumericLower(), nil
	case TypeAlphaOnly:
		return NewAlphaOnly(), nil
	default:
		return nil, NewUnknownConverterError(name)
	}
}
