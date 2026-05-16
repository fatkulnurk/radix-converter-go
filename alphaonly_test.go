package radixconverter_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestAlphaOnly_Encode(t *testing.T) {
	c := radixconverter.NewAlphaOnly()
	tests := []struct {
		input uint64
		want  string
	}{
		{0, "a"},
		{25, "z"},
		{51, "Z"},
		{52, "ba"},
		{100, "bW"},
	}
	for _, tt := range tests {
		if got := c.Encode(tt.input); got != tt.want {
			t.Errorf("Encode(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAlphaOnly_Decode(t *testing.T) {
	c := radixconverter.NewAlphaOnly()
	if got, err := c.Decode("bW"); err != nil {
		t.Fatalf("Decode(\"bW\") error: %v", err)
	} else if got != 100 {
		t.Errorf("Decode(\"bW\") = %d, want 100", got)
	}
}

func TestAlphaOnly_RoundTrip(t *testing.T) {
	c := radixconverter.NewAlphaOnly()
	values := []uint64{0, 1, 10, 25, 51, 52, 100, 1000}
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
