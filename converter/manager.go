package converter

import (
	"sync"

	"github.com/fatkulnurk/radix-converter-go/radixerrors"
)

// ConverterManager manages converter instances with caching.
// It is safe for concurrent use and suitable for dependency injection
// and long-running processes (e.g., servers, workers).
type ConverterManager struct {
	mu      sync.RWMutex
	cache   map[string]IDConverter
	custom  map[string]IDConverter
	factory *ConverterFactory
}

// NewConverterManager creates a new ConverterManager.
// Optional custom converters can be provided at construction time.
func NewConverterManager(customConverters ...map[string]IDConverter) *ConverterManager {
	m := &ConverterManager{
		cache:   make(map[string]IDConverter),
		custom:  make(map[string]IDConverter),
		factory: NewConverterFactory(),
	}

	if len(customConverters) > 0 {
		for name, conv := range customConverters[0] {
			m.custom[name] = conv
		}
	}

	return m
}

// Get returns a converter for the given type.
// It checks custom converters first, then the cache, then creates a new one.
func (m *ConverterManager) Get(t ConverterType) (IDConverter, error) {
	name := string(t)

	m.mu.RLock()
	// Check custom converters.
	if conv, ok := m.custom[name]; ok {
		m.mu.RUnlock()
		return conv, nil
	}

	// Check cache.
	if conv, ok := m.cache[name]; ok {
		m.mu.RUnlock()
		return conv, nil
	}
	m.mu.RUnlock()

	// Create via factory and cache.
	conv, err := m.factory.Make(t)
	if err != nil {
		return nil, err
	}

	m.mu.Lock()
	m.cache[name] = conv
	m.mu.Unlock()

	return conv, nil
}

// Encode encodes a number using the specified converter type.
func (m *ConverterManager) Encode(t ConverterType, number uint64) (string, error) {
	conv, err := m.Get(t)
	if err != nil {
		return "", err
	}
	return conv.Encode(number), nil
}

// Decode decodes a string using the specified converter type.
func (m *ConverterManager) Decode(t ConverterType, encoded string) (uint64, error) {
	conv, err := m.Get(t)
	if err != nil {
		return 0, err
	}
	return conv.Decode(encoded)
}

// RegisterCustom registers a custom converter by name.
// Returns an error if a converter with the same name is already registered.
func (m *ConverterManager) RegisterCustom(name string, conv IDConverter) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.custom[name]; ok {
		return radixerrors.ErrConverterAlreadyRegistered
	}

	m.custom[name] = conv
	return nil
}

// HasCustom checks if a custom converter with the given name is registered.
func (m *ConverterManager) HasCustom(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.custom[name]
	return ok
}

// GetCustomNames returns a slice of all registered custom converter names.
func (m *ConverterManager) GetCustomNames() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.custom))
	for name := range m.custom {
		names = append(names, name)
	}
	return names
}

// ClearCache clears the built-in converter cache.
func (m *ConverterManager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = make(map[string]IDConverter)
}

// ClearAll clears both custom converters and the cache.
func (m *ConverterManager) ClearAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = make(map[string]IDConverter)
	m.custom = make(map[string]IDConverter)
}
