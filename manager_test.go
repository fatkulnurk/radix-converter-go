package radixconverter_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestConverterManager_NewWithoutCustom(t *testing.T) {
	m := radixconverter.NewConverterManager()
	if m == nil {
		t.Fatal("NewConverterManager() returned nil")
	}
}

func TestConverterManager_NewWithCustom(t *testing.T) {
	custom := radixconverter.NewAlphanumericUpper()
	m := radixconverter.NewConverterManager(map[string]radixconverter.IDConverter{
		"upper36": custom,
	})
	if !m.HasCustom("upper36") {
		t.Fatal("HasCustom(\"upper36\") = false, want true")
	}
}

func TestConverterManager_GetCaches(t *testing.T) {
	m := radixconverter.NewConverterManager()
	c1, err := m.Get(radixconverter.TypeBase62)
	if err != nil {
		t.Fatalf("Get(TypeBase62) error: %v", err)
	}
	c2, err := m.Get(radixconverter.TypeBase62)
	if err != nil {
		t.Fatalf("Get(TypeBase62) second call error: %v", err)
	}
	if c1 != c2 {
		t.Error("Get(TypeBase62) should return the same cached instance")
	}
}

func TestConverterManager_EncodeDecode(t *testing.T) {
	m := radixconverter.NewConverterManager()
	encoded, err := m.Encode(radixconverter.TypeBase62, 12345)
	if err != nil {
		t.Fatalf("Encode error: %v", err)
	}
	if encoded != "3d7" {
		t.Errorf("Encode(12345) = %q, want %q", encoded, "3d7")
	}
	decoded, err := m.Decode(radixconverter.TypeBase62, encoded)
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if decoded != 12345 {
		t.Errorf("Decode(%q) = %d, want 12345", encoded, decoded)
	}
}

func TestConverterManager_RegisterCustom(t *testing.T) {
	m := radixconverter.NewConverterManager()
	custom := radixconverter.NewHex()
	err := m.RegisterCustom("hex16", custom)
	if err != nil {
		t.Fatalf("RegisterCustom error: %v", err)
	}
	if !m.HasCustom("hex16") {
		t.Error("HasCustom(\"hex16\") = false, want true")
	}
}

func TestConverterManager_RegisterCustomDuplicate(t *testing.T) {
	m := radixconverter.NewConverterManager()
	custom := radixconverter.NewHex()
	_ = m.RegisterCustom("hex16", custom)
	err := m.RegisterCustom("hex16", custom)
	if err == nil {
		t.Fatal("RegisterCustom duplicate expected error, got nil")
	}
}

func TestConverterManager_GetCustomNames(t *testing.T) {
	m := radixconverter.NewConverterManager()
	_ = m.RegisterCustom("foo", radixconverter.NewHex())
	_ = m.RegisterCustom("bar", radixconverter.NewBase62())
	names := m.GetCustomNames()
	if len(names) != 2 {
		t.Fatalf("GetCustomNames() returned %d names, want 2", len(names))
	}
}

func TestConverterManager_ClearCache(t *testing.T) {
	m := radixconverter.NewConverterManager()
	_, _ = m.Get(radixconverter.TypeBase62)
	m.ClearCache()
	_, err := m.Get(radixconverter.TypeBase62)
	if err != nil {
		t.Fatalf("Get(TypeBase62) after ClearCache error: %v", err)
	}
}

func TestConverterManager_ClearAll(t *testing.T) {
	m := radixconverter.NewConverterManager()
	_ = m.RegisterCustom("myconv", radixconverter.NewHex())
	_, _ = m.Get(radixconverter.TypeBase62)
	m.ClearAll()
	if m.HasCustom("myconv") {
		t.Error("HasCustom(\"myconv\") after ClearAll = true, want false")
	}
}

func TestConverterManager_UnknownType(t *testing.T) {
	m := radixconverter.NewConverterManager()
	_, err := m.Get("unknown_type")
	if err == nil {
		t.Fatal("Get(\"unknown_type\") expected error, got nil")
	}
}
