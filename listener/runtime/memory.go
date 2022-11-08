package runtime

import "github.com/narutopig/neon-lang/value"

// Memory manages variables
type Memory struct {
	state map[string]value.Value
}

func (m Memory) Get(key string) (value.Value, bool) {
	val, exists := m.state[key]
	return val, exists
}

func (m *Memory) Put(key string, val value.Value) {
	m.state[key] = val
}

func NewMemory() *Memory {
	return &Memory{state: make(map[string]value.Value)}
}
