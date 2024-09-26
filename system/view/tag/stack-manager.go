package tag

import (
	"sync"

	"github.com/nikolalohinski/gonja/v2/nodes"
)

type StackManager struct {
	stacks map[string]map[string][]*nodes.Wrapper
	mutex  sync.RWMutex
}

func NewStackManager() *StackManager {
	return &StackManager{
		stacks: make(map[string]map[string][]*nodes.Wrapper),
	}
}

func (sm *StackManager) Push(requestID, name string, node *nodes.Wrapper) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if _, ok := sm.stacks[requestID]; !ok {
		sm.stacks[requestID] = make(map[string][]*nodes.Wrapper)
	}

	if _, ok := sm.stacks[requestID][name]; !ok {
		sm.stacks[requestID][name] = []*nodes.Wrapper{}
	}

	sm.stacks[requestID][name] = append(sm.stacks[requestID][name], node)
}

func (sm *StackManager) Get(requestID, name string) []*nodes.Wrapper {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	if stacks, ok := sm.stacks[requestID]; ok {
		if stack, ok := stacks[name]; ok {
			return stack
		}
	}

	return nil
}

func (sm *StackManager) Clear(requestID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.stacks, requestID)
}

var StackStore = NewStackManager()
