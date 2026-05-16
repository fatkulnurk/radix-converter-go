package radixconverter_test

import (
	"errors"
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestConverterFactory_MakeBuiltins(t *testing.T) {
	f := radixconverter.NewConverterFactory()
	types := []radixconverter.ConverterType{
		radixconverter.TypeBase62,
		radixconverter.TypeAlphanumericUpper,
		radixconverter.TypeAlphanumericLower,
		radixconverter.TypeAlphaOnly,
	}
	for _, typ := range types {
		c, err := f.Make(typ)
		if err != nil {
			t.Fatalf("Make(%q) unexpected error: %v", typ, err)
		}
		if c == nil {
			t.Fatalf("Make(%q) returned nil converter", typ)
		}
		encoded := c.Encode(42)
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("Make(%q): Decode(%q) error: %v", typ, encoded, err)
		}
		if decoded != 42 {
			t.Errorf("Make(%q): roundtrip got %d, want 42", typ, decoded)
		}
	}
}

func TestConverterFactory_MakeCustom(t *testing.T) {
	custom := radixconverter.NewAlphanumericUpper()
	radixconverter.GlobalRegistry.Register("my_custom", custom)
	defer radixconverter.GlobalRegistry.Unregister("my_custom")

	f := radixconverter.NewConverterFactory()
	c, err := f.Make("my_custom")
	if err != nil {
		t.Fatalf("Make(\"my_custom\") unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("Make(\"my_custom\") returned nil converter")
	}
}

func TestConverterFactory_MakeUnknown(t *testing.T) {
	f := radixconverter.NewConverterFactory()
	_, err := f.Make("nonexistent")
	if err == nil {
		t.Fatal("Make(\"nonexistent\") expected error, got nil")
	}
	var radixErr *radixconverter.Error
	if !errors.As(err, &radixErr) {
		t.Errorf("Make(\"nonexistent\") error is not a radixconverter.Error: %v", err)
	}
}

func TestConverterFactory_EncodeDecode(t *testing.T) {
	f := radixconverter.NewConverterFactory()
	c, err := f.Make(radixconverter.TypeBase62)
	if err != nil {
		t.Fatalf("Make(Base62) error: %v", err)
	}
	encoded := c.Encode(12345)
	if encoded != "3d7" {
		t.Errorf("Encode(12345) = %q, want %q", encoded, "3d7")
	}
	decoded, err := c.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode(%q) error: %v", encoded, err)
	}
	if decoded != 12345 {
		t.Errorf("Decode(%q) = %d, want 12345", encoded, decoded)
	}
}
