# Radix Converter Go

A modern, idiomatic, and thread-safe **radix (base-N) converter** for Go. Encode integers to compact string representations and decode them back — perfect for URL shorteners, obfuscated IDs, and more.

Ported from the PHP library [fatkulnurk/radix-converter](https://github.com/fatkulnurk/radix-converter-php) with full Go idioms and conventions.

## Features

- **5 built-in converters**: Base62, Alphanumeric Upper (36), Alphanumeric Lower (36), Alpha Only (52), Hex (16)
- **Custom converters**: Create your own charset with a few lines of code
- **Thread-safe**: All components use `sync.RWMutex` for safe concurrent access
- **Factory pattern**: Create converters on demand
- **Manager pattern**: Cached instances, suitable for dependency injection and long-running servers
- **Custom error types**: Structured error handling with `radixconverter.Error`
- **Zero dependencies**: Pure Go, no external packages required
- **Comprehensive tests**: Full coverage across all strategies, factory, manager, and registry
- **Benchmarks**: Performance metrics included

## Requirements

- Go 1.25+

## Installation

```bash
go get github.com/fatkulnurk/radix-converter-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/fatkulnurk/radix-converter-go"
)

func main() {
    // Direct strategy usage
    b62 := radixconverter.NewBase62()
    encoded := b62.Encode(123456789)
    fmt.Printf("Encoded: %s\n", encoded)

    decoded, err := b62.Decode(encoded)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decoded: %d\n", decoded)

    // Using ConverterFactory
    f := radixconverter.NewConverterFactory()
    conv, _ := f.Make(radixconverter.TypeBase62)
    fmt.Println(conv.Encode(42)) // "g"
}
```

## Built-in Converters

| Converter                  | Type Constant                       | Charset                                    | Base |
|----------------------------|-------------------------------------|--------------------------------------------|------|
| **Base62**                 | `radixconverter.TypeBase62`         | `0-9a-zA-Z`                                | 62   |
| **Alphanumeric Upper**     | `radixconverter.TypeAlphanumericUpper` | `0-9A-Z`                             | 36   |
| **Alphanumeric Lower**     | `radixconverter.TypeAlphanumericLower` | `0-9a-z`                             | 36   |
| **Alpha Only**             | `radixconverter.TypeAlphaOnly`      | `a-zA-Z`                                   | 52   |
| **Hex**                    | (custom)                            | `0-9a-f`                                   | 16   |

### Using Each Converter

```go
import "github.com/fatkulnurk/radix-converter-go"

b62 := radixconverter.NewBase62()
fmt.Println(b62.Encode(1000000)) // "QgUK"

upper := radixconverter.NewAlphanumericUpper()
fmt.Println(upper.Encode(1000000)) // "QGLS"

lower := radixconverter.NewAlphanumericLower()
fmt.Println(lower.Encode(1000000)) // "qglu"

alpha := radixconverter.NewAlphaOnly()
fmt.Println(alpha.Encode(1000000)) // "CysW"

hex := radixconverter.NewHex()
fmt.Println(hex.Encode(255)) // "ff"
```

## URL Shortener Example

```go
b62 := radixconverter.NewBase62()
dbID := uint64(4857293)
shortCode := b62.Encode(dbID)
// https://short.link/W8nR

recovered, _ := b62.Decode(shortCode)
// recovered == 4857293
```

## Factory Pattern

The `ConverterFactory` creates new converter instances on each call. It checks the global `CustomConverterRegistry` first, then falls back to built-in types.

```go
import "github.com/fatkulnurk/radix-converter-go"

f := radixconverter.NewConverterFactory()

// Built-in type
c, err := f.Make(radixconverter.TypeBase62)

// By string name (checks custom registry first)
c, err := f.MakeByName("my_custom_converter")
```

## ConverterManager (Recommended for Servers)

The `ConverterManager` caches built-in converters and supports custom converter registration. It is **thread-safe** and ideal for dependency injection and long-running processes.

```go
import "github.com/fatkulnurk/radix-converter-go"

m := radixconverter.NewConverterManager()

_ = m.RegisterCustom("hex", radixconverter.NewHex())

encoded, _ := m.Encode(radixconverter.TypeBase62, 12345)
decoded, _ := m.Decode(radixconverter.TypeBase62, encoded)

c, _ := m.Get(radixconverter.TypeBase62)
fmt.Println(c.Encode(42))

// Management methods
m.HasCustom("hex")          // true
m.GetCustomNames()          // ["hex"]
m.ClearCache()              // clears built-in cache
m.ClearAll()                // clears everything
```

## BaseConverter for Composition

The `BaseConverter` struct provides the core algorithm and can be embedded in your own types:

```go
import "github.com/fatkulnurk/radix-converter-go"

type MyConverter struct {
    *radixconverter.BaseConverter
}

func NewMyConverter() *MyConverter {
    return &MyConverter{
        BaseConverter: radixconverter.NewBaseConverter("0123456789ABCDEF"),
    }
}
```

## Memory Safety & Dependency Injection

The `ConverterManager` is designed for safe use in **long-running server processes**. Each manager instance maintains isolated state.

```go
import "github.com/fatkulnurk/radix-converter-go"

// Each goroutine/request gets its own manager — no shared state
func handleRequest(id int) {
    m := radixconverter.NewConverterManager()
    encoded, _ := m.Encode(radixconverter.TypeBase62, uint64(id*100))
    fmt.Printf("Request %d: %d -> %s\n", id, id*100, encoded)
}

// Dependency injection
type URLShortener struct {
    manager *radixconverter.ConverterManager
}

func (s *URLShortener) Shorten(id uint64) string {
    encoded, _ := s.manager.Encode(radixconverter.TypeBase62, id)
    return encoded
}
```

## Custom Converters

Create your own charset using the shared helpers:

```go
import "github.com/fatkulnurk/radix-converter-go"

type MyConverter struct {
    charset   string
    base      uint64
    charIndex map[rune]int
}

func NewMyConverter() *MyConverter {
    charset := "ABCDEF0123456789"
    return &MyConverter{
        charset:   charset,
        base:      uint64(len(charset)),
        charIndex: radixconverter.BuildIndex(charset),
    }
}

func (m *MyConverter) Encode(number uint64) string {
    return radixconverter.Encode(number, m.charset, m.base)
}

func (m *MyConverter) Decode(encoded string) (uint64, error) {
    return radixconverter.Decode(encoded, m.base, m.charIndex)
}

func (m *MyConverter) Charset() string {
    return m.charset
}
```

Then register it:

```go
// Via global registry (NOT recommended for long-running servers)
radixconverter.GlobalRegistry.Register("my_conv", NewMyConverter())

// Via ConverterManager (recommended)
m := radixconverter.NewConverterManager()
_ = m.RegisterCustom("my_conv", NewMyConverter())
```

> **Warning**: The global `CustomConverterRegistry` is mutable global state. For long-running server processes, prefer using `ConverterManager` which encapsulates state per instance.

## Error Handling

All errors are structured as `*radixconverter.Error` for easy inspection:

```go
import (
    "errors"

    "github.com/fatkulnurk/radix-converter-go"
)

_, err := b62.Decode("")
if errors.Is(err, radixconverter.ErrEmptyEncoded) {
    // handle empty input
}

_, err = b62.Decode("!@#")
if err != nil {
    fmt.Println(err) // "radix converter: decode: invalid character: '!'"
}
```

Available sentinel errors:
- `radixconverter.ErrNegativeInput` — Negative number provided for encoding (defensive; `uint64` prevents this at compile time)
- `radixconverter.ErrEmptyEncoded` — Attempted to decode an empty string
- `radixconverter.ErrInvalidChar` — Invalid character found during decoding
- `radixconverter.ErrUnknownConverter` — Unknown converter type requested
- `radixconverter.ErrConverterAlreadyRegistered` — Duplicate custom converter registration
- `radixconverter.ErrConverterNotRegistered` — Custom converter not found

## Project Structure

```
radix-converter-go/
├── go.mod
├── README.md
├── .gitignore
├── idconverter.go           # IDConverter interface (Encode/Decode)
├── type.go                  # ConverterType constants (TypeBase62, etc.)
├── baseconverter.go         # BaseConverter (composition-friendly)
├── converter.go             # Shared Encode/Decode/BuildIndex helpers
├── factory.go               # ConverterFactory (Make/MakeByName)
├── manager.go               # ConverterManager (thread-safe, cached)
├── registry.go              # CustomConverterRegistry (thread-safe)
├── custom_converter.go      # CustomRegistry alias for GlobalRegistry
├── errors.go                # Error struct and sentinel errors
├── base62.go                # Base62 strategy
├── alphanumericlower.go     # AlphanumericLower strategy
├── alphanumericupper.go     # AlphanumericUpper strategy
├── alphaonly.go             # AlphaOnly strategy
├── hex.go                   # Hex strategy
├── *_test.go                # Tests + benchmarks
└── examples/
    ├── usage/main.go        # Basic usage examples
    └── custom_strategy/     # Custom converter & registry examples
```

## Benchmarks

```bash
go test -bench=. -benchmem .
```

Sample output:
```
BenchmarkBase62_Encode-16       19934116    53.47 ns/op    5 B/op    1 allocs/op
BenchmarkBase62_Decode-16       14563406    76.41 ns/op    0 B/op    0 allocs/op
```

## Testing

```bash
go test ./...
```

With coverage:

```bash
go test -cover ./...
```

## Comparison with PHP Version

| Feature              | PHP Version                     | Go Version                              |
|----------------------|----------------------------------|------------------------------------------|
| Package              | `Fatkulnurk\RadixConverter`     | `radixconverter`                        |
| Encoding             | `int` (signed 64-bit)           | `uint64` (unsigned 64-bit)              |
| Concurrency          | Not applicable (PHP-FPM model)  | Thread-safe with `sync.RWMutex`         |
| Custom converters    | Static global registry           | Both global registry and per-instance manager |
| Framework support    | Laravel service provider        | Pure Go library, framework-agnostic     |
| Error handling       | Exceptions (`ConverterException`) | Typed errors (`radixconverter.Error`)  |
| Octane/Swoole safety | Manager only (not static registry) | Manager recommended, registry has mutex |

## License

MIT License. See [LICENSE](LICENSE) for details.

## Author

**fatkulnurk** — [fatkulnurk@gmail.com](mailto:fatkulnurk@gmail.com)
