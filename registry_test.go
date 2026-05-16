package radixconverter_test

import (
	"testing"

	"github.com/fatkulnurk/radix-converter-go"
)

func TestCustomConverterRegistry_RegisterAndGet(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	defer r.Clear()

	conv := radixconverter.NewAlphanumericUpper()
	err := r.Register("upper", conv)
	if err != nil {
		t.Fatalf("Register error: %v", err)
	}

	got, err := r.Get("upper")
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	if got != conv {
		t.Error("Get() returned different instance")
	}
}

func TestCustomConverterRegistry_RegisterDuplicate(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	defer r.Clear()

	conv := radixconverter.NewHex()
	_ = r.Register("hex", conv)
	err := r.Register("hex", conv)
	if err == nil {
		t.Fatal("Register duplicate expected error, got nil")
	}
}

func TestCustomConverterRegistry_GetUnregistered(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	_, err := r.Get("nonexistent")
	if err == nil {
		t.Fatal("Get unregistered expected error, got nil")
	}
}

func TestCustomConverterRegistry_Has(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	defer r.Clear()

	if r.Has("foo") {
		t.Error("Has(\"foo\") = true, want false (not registered)")
	}

	_ = r.Register("foo", radixconverter.NewBase62())
	if !r.Has("foo") {
		t.Error("Has(\"foo\") = false, want true")
	}
}

func TestCustomConverterRegistry_Unregister(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	_ = r.Register("temp", radixconverter.NewHex())

	if !r.Unregister("temp") {
		t.Error("Unregister(\"temp\") = false, want true")
	}
	if r.Has("temp") {
		t.Error("Has(\"temp\") after unregister = true, want false")
	}
}

func TestCustomConverterRegistry_UnregisterNonExistent(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	if r.Unregister("nope") {
		t.Error("Unregister(\"nope\") = true, want false")
	}
}

func TestCustomConverterRegistry_GetRegisteredNames(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	defer r.Clear()

	_ = r.Register("a", radixconverter.NewHex())
	_ = r.Register("b", radixconverter.NewBase62())

	names := r.GetRegisteredNames()
	if len(names) != 2 {
		t.Fatalf("GetRegisteredNames() = %d, want 2", len(names))
	}
}

func TestCustomConverterRegistry_Clear(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	_ = r.Register("x", radixconverter.NewAlphaOnly())
	r.Clear()

	if r.Has("x") {
		t.Error("Has(\"x\") after Clear = true, want false")
	}
	names := r.GetRegisteredNames()
	if len(names) != 0 {
		t.Errorf("GetRegisteredNames() after Clear = %d, want 0", len(names))
	}
}

func TestCustomConverterRegistry_GetAll(t *testing.T) {
	r := radixconverter.NewCustomConverterRegistry()
	defer r.Clear()

	conv := radixconverter.NewHex()
	_ = r.Register("hex", conv)

	all := r.GetAll()
	if len(all) != 1 {
		t.Fatalf("GetAll() = %d entries, want 1", len(all))
	}
	if all["hex"] != conv {
		t.Error("GetAll()[\"hex\"] returned wrong instance")
	}
}
