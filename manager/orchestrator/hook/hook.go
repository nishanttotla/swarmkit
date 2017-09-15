package hook

import (
	"github.com/docker/swarmkit/manager/state/store"
	"golang.org/x/net/context"
)

// Orchestrator does nothing.
type Orchestrator struct {
}

// NewHookOrchestrator creates a new hook Orchestrator
func NewHookOrchestrator(store *store.MemoryStore) *Orchestrator {
	return &Orchestrator{}
}

// Run contains the hook orchestrator event loop
func (g *Orchestrator) Run(ctx context.Context) error {
	return nil
}

// Stop stops the orchestrator.
func (g *Orchestrator) Stop() {
	return
}
