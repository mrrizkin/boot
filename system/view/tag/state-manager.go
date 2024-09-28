package tag

import (
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nikolalohinski/gonja/v2/nodes"
)

type StateManager struct {
	stacks   map[string]map[string][]*nodes.Wrapper
	rendered map[string]map[string]bool
	mutex    sync.RWMutex
}

func NewStateManager() *StateManager {
	return &StateManager{
		stacks:   make(map[string]map[string][]*nodes.Wrapper),
		rendered: make(map[string]map[string]bool),
	}
}

func (sm *StateManager) GenerateID() (string, error) {
	return gonanoid.New()
}

func (sm *StateManager) Push(id, name string, node *nodes.Wrapper) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if _, ok := sm.stacks[id]; !ok {
		sm.stacks[id] = make(map[string][]*nodes.Wrapper)
	}

	if _, ok := sm.stacks[id][name]; !ok {
		sm.stacks[id][name] = []*nodes.Wrapper{}
	}

	sm.stacks[id][name] = append(sm.stacks[id][name], node)
}

func (sm *StateManager) Get(id, name string) []*nodes.Wrapper {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	if stacks, ok := sm.stacks[id]; ok {
		if stack, ok := stacks[name]; ok {
			return stack
		}
	}

	return nil
}

func (sm *StateManager) ShouldRender(id, name string) bool {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if _, ok := sm.rendered[id]; !ok {
		sm.rendered[id] = make(map[string]bool)
	}

	if _, ok := sm.rendered[id][name]; !ok {
		sm.rendered[id][name] = true
		return true
	}

	return false
}

func (sm *StateManager) Clear(id string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.stacks, id)
	delete(sm.rendered, id)
}

var State = NewStateManager()
