package agent

import (
	"context"
	"fmt"
)

// Agent represents the eBPF-based monitoring agent that runs on each node
type Agent interface {
	// Start begins monitoring pod activities
	Start(ctx context.Context) error
	
	// Stop gracefully shuts down the agent
	Stop() error
	
	// GetMetrics returns current monitoring metrics
	GetMetrics() (*Metrics, error)
}

// Metrics contains runtime monitoring data
type Metrics struct {
	SyscallsMonitored   int64
	NetworkEvents       int64
	AnomaliesDetected   int64
	PoliciesEnforced    int64
}

// Config holds agent configuration
type Config struct {
	NodeName          string
	EnableSyscallTrace bool
	EnableNetworkTrace bool
	PolicyEndpoint    string
}

// New creates a new monitoring agent instance
func New(cfg *Config) (Agent, error) {
	// This is a showcase/demo implementation
	// Enterprise version uses eBPF/CO-RE for kernel tracing
	return &demoAgent{config: cfg}, nil
}

type demoAgent struct {
	config *Config
}

func (a *demoAgent) Start(ctx context.Context) error {
	fmt.Println("Demo agent: monitoring started")
	fmt.Println("Enterprise version: eBPF/CO-RE syscall and network tracing")
	return nil
}

func (a *demoAgent) Stop() error {
	fmt.Println("Demo agent: monitoring stopped")
	return nil
}

func (a *demoAgent) GetMetrics() (*Metrics, error) {
	return &Metrics{
		SyscallsMonitored: 0,
		NetworkEvents:     0,
		AnomaliesDetected: 0,
		PoliciesEnforced:  0,
	}, nil
}
