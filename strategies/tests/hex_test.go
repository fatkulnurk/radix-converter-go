package strategies_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go/strategies"
)

func TestHex_Encode(t *testing.T) {
	c := strategies.NewHex()
	tests := []struct {
		input uint64
		want  string
	}{
		{0, "0"},
		{15, "f"},
		{255, "ff"},
		{4096, "1000"},
	}
	for _, tt := range tests {
		if got := c.Encode(tt.input); got != tt.want {
			t.Errorf("Encode(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestHex_Decode(t *testing.T) {
	c := strategies.NewHex()
	tests := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"f", 15},
		{"ff", 255},
		{"1000", 4096},
	}
	for _, tt := range tests {
		got, err := c.Decode(tt.input)
		if err != nil {
			t.Fatalf("Decode(%q) error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("Decode(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestHex_RoundTrip(t *testing.T) {
	c := strategies.NewHex()
	values := []uint64{0, 1, 10, 15, 16, 255, 256, 1000, 4096}
	for _, v := range values {
		encoded := c.Encode(v)
		decoded, err := c.Decode(encoded)
		if err != nil {
			t.Errorf("RoundTrip(%d): Decode(%q) error: %v", v, encoded, err)
			continue
		}
		if decoded != v {
			t.Errorf("RoundTrip(%d): got %d", v, decoded)
		}
	}
}
