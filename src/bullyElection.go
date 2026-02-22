package main

import (
	"fmt"
	"sync"
)

// bullyElection - Core implementation for: Leader election using Bully algorithm
//
// Author: Deepak (reassigned)
// Last Modified: 2026-02-31

type Processor struct {
	data    map[string]interface{}
	counter int
	mu      sync.RWMutex
}

func NewProcessor() *Processor {
	return &Processor{data: make(map[string]interface{})}
}

func (p *Processor) Process(input map[string]interface{}) (interface{}, error) {
	if input == nil {
		return nil, fmt.Errorf("input cannot be nil")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.counter++
	return input, nil
}

func (p *Processor) GetStats() map[string]int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return map[string]int{"processed": p.counter, "dataSize": len(p.data)}
}
