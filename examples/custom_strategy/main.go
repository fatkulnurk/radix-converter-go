package main

import (
	"fmt"

	"github.com/fatkulnurk/radix-converter-go"
)

// BinaryConverter is a custom base-2 (binary) converter.
type BinaryConverter struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewBinaryConverter() *BinaryConverter {
	charset := "01"
	return &BinaryConverter{
		charset:   charset,
		base:      uint64(len(charset)),
		charIndex: radixconverter.BuildIndex(charset),
	}
}

func (b *BinaryConverter) Encode(number uint64) string {
	return radixconverter.Encode(number, b.charset, b.base)
}

func (b *BinaryConverter) Decode(encoded string) (uint64, error) {
	return radixconverter.Decode(encoded, b.base, b.charIndex)
}

func (b *BinaryConverter) Charset() string {
	return b.charset
}

// OctalConverter is a custom base-8 converter.
type OctalConverter struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func NewOctalConverter() *OctalConverter {
	charset := "01234567"
	return &OctalConverter{
		charset:   charset,
		base:      uint64(len(charset)),
		charIndex: radixconverter.BuildIndex(charset),
	}
}

func (o *OctalConverter) Encode(number uint64) string {
	return radixconverter.Encode(number, o.charset, o.base)
}

func (o *OctalConverter) Decode(encoded string) (uint64, error) {
	return radixconverter.Decode(encoded, o.base, o.charIndex)
}

func (o *OctalConverter) Charset() string {
	return o.charset
}

func main() {
	// ==========================================
	// 1. Using built-in Hex strategy via registry
	// ==========================================
	fmt.Println("=== Built-in Hex via Registry ===")

	hex := radixconverter.NewHex()
	_ = radixconverter.GlobalRegistry.Register("hex", hex)

	f := radixconverter.NewConverterFactory()
	hexConv, err := f.MakeByName("hex")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Hex: 255 -> %s\n", hexConv.Encode(255))

	radixconverter.GlobalRegistry.Unregister("hex")

	// ==========================================
	// 2. Custom Binary (Base-2) strategy
	// ==========================================
	fmt.Println("\n=== Custom Binary (Base-2) ===")

	binary := NewBinaryConverter()
	fmt.Printf("Binary: 42 -> %s\n", binary.Encode(42))
	fmt.Printf("Binary: 255 -> %s\n", binary.Encode(255))

	val, _ := binary.Decode("101010")
	fmt.Printf("Binary: 101010 -> %d\n", val)

	_ = radixconverter.GlobalRegistry.Register("binary", binary)
	binaryFromRegistry, _ := radixconverter.GlobalRegistry.Get("binary")
	fmt.Printf("Registry Binary: 100 -> %s\n", binaryFromRegistry.Encode(100))
	radixconverter.GlobalRegistry.Unregister("binary")

	// ==========================================
	// 3. Custom Octal (Base-8) strategy
	// ==========================================
	fmt.Println("\n=== Custom Octal (Base-8) ===")

	octal := NewOctalConverter()
	fmt.Printf("Octal: 64 -> %s\n", octal.Encode(64))
	fmt.Printf("Octal: 511 -> %s\n", octal.Encode(511))

	val, _ = octal.Decode("777")
	fmt.Printf("Octal: 777 -> %d\n", val)

	// ==========================================
	// 4. Custom Base64-like strategy
	// ==========================================
	fmt.Println("\n=== Custom Base64-like ===")

	base64Charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	charIndex := radixconverter.BuildIndex(base64Charset)

	custom64 := &customBase64{
		charset:   base64Charset,
		base:      64,
		charIndex: charIndex,
	}
	fmt.Printf("Base64-like: 1000000 -> %s\n", custom64.Encode(1000000))

	// ==========================================
	// 5. Check registered converters
	// ==========================================
	fmt.Println("\n=== Registry Management ===")

	_ = radixconverter.GlobalRegistry.Register("my_hex", radixconverter.NewHex())
	_ = radixconverter.GlobalRegistry.Register("my_binary", NewBinaryConverter())

	names := radixconverter.GlobalRegistry.GetRegisteredNames()
	fmt.Printf("Registered: %v\n", names)

	fmt.Printf("Has 'my_hex': %v\n", radixconverter.GlobalRegistry.Has("my_hex"))
	fmt.Printf("Has 'unknown': %v\n", radixconverter.GlobalRegistry.Has("unknown"))

	radixconverter.GlobalRegistry.Clear()
	fmt.Printf("After clear, registered: %v\n", radixconverter.GlobalRegistry.GetRegisteredNames())

	// ==========================================
	// 6. Using ConverterManager with custom converters
	// ==========================================
	fmt.Println("\n=== ConverterManager with Custom ===")

	m := radixconverter.NewConverterManager()
	_ = m.RegisterCustom("octal", NewOctalConverter())

	encoded, _ := m.Encode("octal", 64)
	fmt.Printf("Manager Octal: 64 -> %s\n", encoded)

	decoded, _ := m.Decode("octal", encoded)
	fmt.Printf("Manager Octal: %s -> %d\n", encoded, decoded)

	fmt.Printf("Custom names: %v\n", m.GetCustomNames())
}

type customBase64 struct {
	charset   string
	base      uint64
	charIndex map[rune]int
}

func (c *customBase64) Encode(number uint64) string {
	return radixconverter.Encode(number, c.charset, c.base)
}

func (c *customBase64) Decode(encoded string) (uint64, error) {
	return radixconverter.Decode(encoded, c.base, c.charIndex)
}

func (c *customBase64) Charset() string {
	return c.charset
}
