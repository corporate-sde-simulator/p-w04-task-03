package main

import (
	"fmt"
	"sync"
)

// nodeManager - Supporting implementation for: Leader election using Bully algorithm
//
// Author: Deepak (reassigned)
// Last Modified: 2026-02-31

type Manager struct {
	data    map[string]interface{}
	counter int
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{data: make(map[string]interface{})}
}

func (p *Manager) Process(input map[string]interface{}) (interface{}, error) {
	if input == nil {
		return nil, fmt.Errorf("input cannot be nil")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.counter++
	return input, nil
}

func (p *Manager) GetStats() map[string]int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return map[string]int{"processed": p.counter, "dataSize": len(p.data)}
}
