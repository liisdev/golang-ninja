package cache

type memory struct {
	data map[string]interface{}
}

type cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

func New() *memory {
	return &memory{
		data: make(map[string]interface{}),
	}
}

func (m *memory) Set(key string, value interface{}) {
	m.data[key] = value
}

func (m *memory) Get(key string) interface{} {
	value, exists := m.data[key]
	if !exists {
		return nil
	}
	return value
}

func (m *memory) Delete(key string) {
	delete(m.data, key)
}
