package main

import (
	"fmt"
	"log"
	"math"

	"github.com/fatkulnurk/radix-converter-go/converter"
	"github.com/fatkulnurk/radix-converter-go/strategies"
)

func main() {
	// ==========================================
	// 1. Using ConverterFactory
	// ==========================================
	fmt.Println("=== ConverterFactory ===")

	f := converter.NewConverterFactory()

	// Create a Base62 converter
	base62, err := f.Make(converter.Base62)
	if err != nil {
		log.Fatal(err)
	}

	encoded := base62.Encode(123456789)
	fmt.Printf("Base62: 123456789 -> %s\n", encoded)

	decoded, err := base62.Decode(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base62: %s -> %d\n", encoded, decoded)

	// ==========================================
	// 2. Using strategies directly
	// ==========================================
	fmt.Println("\n=== Direct Strategy Usage ===")

	b62 := strategies.NewBase62()
	fmt.Printf("Base62: 42 -> %s\n", b62.Encode(42))

	upper := strategies.NewAlphanumericUpper()
	fmt.Printf("Upper36: 42 -> %s\n", upper.Encode(42))

	lower := strategies.NewAlphanumericLower()
	fmt.Printf("Lower36: 42 -> %s\n", lower.Encode(42))

	alpha := strategies.NewAlphaOnly()
	fmt.Printf("Alpha52: 42 -> %s\n", alpha.Encode(42))

	hex := strategies.NewHex()
	fmt.Printf("Hex16: 255 -> %s\n", hex.Encode(255))

	// ==========================================
	// 3. Comparing all converter types
	// ==========================================
	fmt.Println("\n=== All Converters (number=1000000) ===")

	number := uint64(1000000)
	types := []converter.ConverterType{
		converter.Base62,
		converter.AlphanumericUpper,
		converter.AlphanumericLower,
		converter.AlphaOnly,
	}

	for _, typ := range types {
		c, _ := f.Make(typ)
		enc := c.Encode(number)
		dec, _ := c.Decode(enc)
		fmt.Printf("%-20s: %d -> %s -> %d\n", typ, number, enc, dec)
	}

	// ==========================================
	// 4. URL Shortener use case
	// ==========================================
	fmt.Println("\n=== URL Shortener Example ===")

	dbID := uint64(4857293)
	shortCode := b62.Encode(dbID)
	fmt.Printf("Database ID %d -> Short URL: https://short.link/%s\n", dbID, shortCode)

	// Decode back
	recovered, _ := b62.Decode(shortCode)
	fmt.Printf("Short code %s -> Database ID: %d\n", shortCode, recovered)

	// ==========================================
	// 5. Error handling
	// ==========================================
	fmt.Println("\n=== Error Handling ===")

	_, err = b62.Decode("")
	if err != nil {
		fmt.Printf("Empty decode error: %v\n", err)
	}

	_, err = b62.Decode("!@#")
	if err != nil {
		fmt.Printf("Invalid char error: %v\n", err)
	}

	// ==========================================
	// 6. Round-trip with max uint64
	// ==========================================
	fmt.Println("\n=== Round-trip with large numbers ===")

	largeNumber := uint64(math.MaxUint32)
	enc := b62.Encode(largeNumber)
	dec, _ := b62.Decode(enc)
	fmt.Printf("Round-trip %d -> %s -> %d (match: %v)\n", largeNumber, enc, dec, dec == largeNumber)
}
