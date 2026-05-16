package radixconverter_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestAlphanumericUpper_Encode(t *testing.T) {
	c := radixconverter.NewAlphanumericUpper()
	tests := []struct {
		input uint64
		want  string
	}{
		{0, "0"},
		{35, "Z"},
		{1000, "RS"},
	}
	for _, tt := range tests {
		if got := c.Encode(tt.input); got != tt.want {
			t.Errorf("Encode(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAlphanumericUpper_Decode(t *testing.T) {
	c := radixconverter.NewAlphanumericUpper()
	if got, err := c.Decode("RS"); err != nil {
		t.Fatalf("Decode(\"RS\") error: %v", err)
	} else if got != 1000 {
		t.Errorf("Decode(\"RS\") = %d, want 1000", got)
	}
}

func TestAlphanumericUpper_RoundTrip(t *testing.T) {
	c := radixconverter.NewAlphanumericUpper()
	values := []uint64{0, 1, 10, 35, 36, 100, 1000, 10000}
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
