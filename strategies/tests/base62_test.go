package strategies_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go/strategies"
)

func TestBase62_EncodeZero(t *testing.T) {
	c := strategies.NewBase62()
	if got := c.Encode(0); got != "0" {
		t.Errorf("Encode(0) = %q, want %q", got, "0")
	}
}

func TestBase62_EncodeSmall(t *testing.T) {
	c := strategies.NewBase62()
	tests := []struct {
		input uint64
		want  string
	}{
		{5, "5"},
		{10, "a"},
		{36, "A"},
		{61, "Z"},
	}
	for _, tt := range tests {
		if got := c.Encode(tt.input); got != tt.want {
			t.Errorf("Encode(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestBase62_Encode12345(t *testing.T) {
	c := strategies.NewBase62()
	if got := c.Encode(12345); got != "3d7" {
		t.Errorf("Encode(12345) = %q, want %q", got, "3d7")
	}
}

func TestBase62_Decode(t *testing.T) {
	c := strategies.NewBase62()
	tests := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"5", 5},
		{"a", 10},
		{"A", 36},
		{"Z", 61},
		{"3d7", 12345},
	}
	for _, tt := range tests {
		got, err := c.Decode(tt.input)
		if err != nil {
			t.Fatalf("Decode(%q) unexpected error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("Decode(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestBase62_RoundTrip(t *testing.T) {
	c := strategies.NewBase62()
	values := []uint64{0, 1, 10, 36, 62, 100, 1000, 12345, 1000000}
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

func TestBase62_Charset(t *testing.T) {
	c := strategies.NewBase62()
	want := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if got := c.Charset(); got != want {
		t.Errorf("Charset() = %q, want %q", got, want)
	}
}

func TestBase62_DecodeEmptyString(t *testing.T) {
	c := strategies.NewBase62()
	_, err := c.Decode("")
	if err == nil {
		t.Error("Decode(\"\") expected error, got nil")
	}
}

func TestBase62_DecodeInvalidChar(t *testing.T) {
	c := strategies.NewBase62()
	_, err := c.Decode("!")
	if err == nil {
		t.Error("Decode(\"!\") expected error, got nil")
	}
}
