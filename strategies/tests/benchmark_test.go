package strategies_test

import (
	"fmt"
	"testing"

	"github.com/fatkulnurk/radix-converter-go/strategies"
)

// BenchmarkBase62_Encode measures the performance of Base62 encoding.
func BenchmarkBase62_Encode(b *testing.B) {
	c := strategies.NewBase62()
	for i := 0; i < b.N; i++ {
		c.Encode(123456789)
	}
}

// BenchmarkBase62_Decode measures the performance of Base62 decoding.
func BenchmarkBase62_Decode(b *testing.B) {
	c := strategies.NewBase62()
	encoded := c.Encode(123456789)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = c.Decode(encoded)
	}
}

// BenchmarkAll_Encode benchmarks encoding across all converter types.
func BenchmarkAll_Encode(b *testing.B) {
	converters := []interface {
		Encode(uint64) string
	}{
		strategies.NewBase62(),
		strategies.NewAlphanumericUpper(),
		strategies.NewAlphanumericLower(),
		strategies.NewAlphaOnly(),
		strategies.NewHex(),
	}

	for _, c := range converters {
		b.Run(fmt.Sprintf("%T", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Encode(999999999)
			}
		})
	}
}

// BenchmarkAll_Decode benchmarks decoding across all converter types.
func BenchmarkAll_Decode(b *testing.B) {
	type encoderDecoder interface {
		Encode(uint64) string
		Decode(string) (uint64, error)
	}

	converters := []encoderDecoder{
		strategies.NewBase62(),
		strategies.NewAlphanumericUpper(),
		strategies.NewAlphanumericLower(),
		strategies.NewAlphaOnly(),
		strategies.NewHex(),
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
