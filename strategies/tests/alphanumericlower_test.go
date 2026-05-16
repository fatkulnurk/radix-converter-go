package strategies_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go/strategies"
)

func TestAlphanumericLower_Encode(t *testing.T) {
	c := strategies.NewAlphanumericLower()
	tests := []struct {
		input uint64
		want  string
	}{
		{0, "0"},
		{35, "z"},
		{1000, "rs"},
	}
	for _, tt := range tests {
		if got := c.Encode(tt.input); got != tt.want {
			t.Errorf("Encode(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAlphanumericLower_Decode(t *testing.T) {
	c := strategies.NewAlphanumericLower()
	if got, err := c.Decode("rs"); err != nil {
		t.Fatalf("Decode(\"rs\") error: %v", err)
	} else if got != 1000 {
		t.Errorf("Decode(\"rs\") = %d, want 1000", got)
	}
}

func TestAlphanumericLower_RoundTrip(t *testing.T) {
	c := strategies.NewAlphanumericLower()
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
