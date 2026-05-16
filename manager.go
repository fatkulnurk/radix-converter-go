package radixconverter

import "sync"

// ConverterManager manages converter instances with caching.
type ConverterManager struct {
	mu      sync.RWMutex
	cache   map[string]IDConverter
	custom  map[string]IDConverter
	factory *ConverterFactory
}

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

func (m *ConverterManager) Get(t ConverterType) (IDConverter, error) {
	name := string(t)

	m.mu.RLock()
	if conv, ok := m.custom[name]; ok {
		m.mu.RUnlock()
		return conv, nil
	}
	if conv, ok := m.cache[name]; ok {
		m.mu.RUnlock()
		return conv, nil
	}
	m.mu.RUnlock()

	conv, err := m.factory.Make(t)
	if err != nil {
		return nil, err
	}

	m.mu.Lock()
	m.cache[name] = conv
	m.mu.Unlock()

	return conv, nil
}

func (m *ConverterManager) Encode(t ConverterType, number uint64) (string, error) {
	conv, err := m.Get(t)
	if err != nil {
		return "", err
	}
	return conv.Encode(number), nil
}

func (m *ConverterManager) Decode(t ConverterType, encoded string) (uint64, error) {
	conv, err := m.Get(t)
	if err != nil {
		return 0, err
	}
	return conv.Decode(encoded)
}

func (m *ConverterManager) RegisterCustom(name string, conv IDConverter) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.custom[name]; ok {
		return ErrConverterAlreadyRegistered
	}

	m.custom[name] = conv
	return nil
}

func (m *ConverterManager) HasCustom(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.custom[name]
	return ok
}

func (m *ConverterManager) GetCustomNames() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.custom))
	for name := range m.custom {
		names = append(names, name)
	}
	return names
}

func (m *ConverterManager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = make(map[string]IDConverter)
}

func (m *ConverterManager) ClearAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = make(map[string]IDConverter)
	m.custom = make(map[string]IDConverter)
}
