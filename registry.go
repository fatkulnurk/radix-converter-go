package radixconverter

import "sync"

// CustomConverterRegistry is a thread-safe global registry for custom converters.
type CustomConverterRegistry struct {
	mu         sync.RWMutex
	converters map[string]IDConverter
}

var GlobalRegistry = NewCustomConverterRegistry()

func NewCustomConverterRegistry() *CustomConverterRegistry {
	return &CustomConverterRegistry{
		converters: make(map[string]IDConverter),
	}
}

func (r *CustomConverterRegistry) Register(name string, conv IDConverter) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.converters[name]; ok {
		return ErrConverterAlreadyRegistered
	}

	r.converters[name] = conv
	return nil
}

func (r *CustomConverterRegistry) Get(name string) (IDConverter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	conv, ok := r.converters[name]
	if !ok {
		return nil, ErrConverterNotRegistered
	}

	return conv, nil
}

func (r *CustomConverterRegistry) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.converters[name]
	return ok
}

func (r *CustomConverterRegistry) Unregister(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.converters[name]; !ok {
		return false
	}

	delete(r.converters, name)
	return true
}

func (r *CustomConverterRegistry) GetRegisteredNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.converters))
	for name := range r.converters {
		names = append(names, name)
	}
	return names
}

func (r *CustomConverterRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.converters = make(map[string]IDConverter)
}

func (r *CustomConverterRegistry) GetAll() map[string]IDConverter {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[string]IDConverter, len(r.converters))
	for name, conv := range r.converters {
		result[name] = conv
	}

	return result
}
