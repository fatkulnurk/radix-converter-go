package radixconverter_test

import (
	"fmt"
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func BenchmarkBase62_Encode(b *testing.B) {
	c := radixconverter.NewBase62()
	for i := 0; i < b.N; i++ {
		c.Encode(123456789)
	}
}

func BenchmarkBase62_Decode(b *testing.B) {
	c := radixconverter.NewBase62()
	encoded := c.Encode(123456789)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = c.Decode(encoded)
	}
}

func BenchmarkAll_Encode(b *testing.B) {
	converters := []interface {
		Encode(uint64) string
	}{
		radixconverter.NewBase62(),
		radixconverter.NewAlphanumericUpper(),
		radixconverter.NewAlphanumericLower(),
		radixconverter.NewAlphaOnly(),
		radixconverter.NewHex(),
	}

	for _, c := range converters {
		b.Run(fmt.Sprintf("%T", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Encode(999999999)
			}
		})
	}
}

func BenchmarkAll_Decode(b *testing.B) {
	type encoderDecoder interface {
		Encode(uint64) string
		Decode(string) (uint64, error)
	}

	converters := []encoderDecoder{
		radixconverter.NewBase62(),
		radixconverter.NewAlphanumericUpper(),
		radixconverter.NewAlphanumericLower(),
		radixconverter.NewAlphaOnly(),
		radixconverter.NewHex(),
	}

	for _, c := range converters {
		b.Run(fmt.Sprintf("%T", c), func(b *testing.B) {
			encoded := c.Encode(999999999)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = c.Decode(encoded)
			}
		})
	}
}
