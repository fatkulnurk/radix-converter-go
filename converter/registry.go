package converter

import (
	"sync"

	"github.com/fatkulnurk/radix-converter-go/radixerrors"
)

// CustomConverterRegistry is a thread-safe global registry for custom converters.
// NOTE: Because this is a global mutable state, it is NOT safe for long-running
// server processes (e.g., Laravel Octane, Swoole) that handle requests concurrently
// without proper lifecycle management. Prefer using ConverterManager for such cases.
type CustomConverterRegistry struct {
	mu        sync.RWMutex
	converters map[string]IDConverter
}

// GlobalRegistry is the default global registry instance.
var GlobalRegistry = NewCustomConverterRegistry()

// NewCustomConverterRegistry creates a new empty registry.
func NewCustomConverterRegistry() *CustomConverterRegistry {
	return &CustomConverterRegistry{
		converters: make(map[string]IDConverter),
	}
}

// Register adds a custom converter to the registry.
// Returns an error if a converter with the same name is already registered.
func (r *CustomConverterRegistry) Register(name string, conv IDConverter) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.converters[name]; ok {
		return radixerrors.ErrConverterAlreadyRegistered
	}

	r.converters[name] = conv
	return nil
}

// Get retrieves a custom converter by name.
// Returns an error if the converter is not registered.
func (r *CustomConverterRegistry) Get(name string) (IDConverter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	conv, ok := r.converters[name]
	if !ok {
		return nil, radixerrors.ErrConverterNotRegistered
	}

	return conv, nil
}

// Has checks if a custom converter with the given name is registered.
func (r *CustomConverterRegistry) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.converters[name]
	return ok
}

// Unregister removes a custom converter from the registry.
// Returns true if the converter was found and removed, false otherwise.
func (r *CustomConverterRegistry) Unregister(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.converters[name]; !ok {
		return false
	}

	delete(r.converters, name)
	return true
}

// GetRegisteredNames returns a slice of all registered converter names.
func (r *CustomConverterRegistry) GetRegisteredNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.converters))
	for name := range r.converters {
		names = append(names, name)
	}
	return names
}

// Clear removes all registered converters.
func (r *CustomConverterRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.converters = make(map[string]IDConverter)
}

// GetAll returns a copy of all registered converters.
func (r *CustomConverterRegistry) GetAll() map[string]IDConverter {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[string]IDConverter, len(r.converters))
	for name, conv := range r.converters {
		result[name] = conv
	}

	return result
}
