package policy

import (
	"context"
	"fmt"
)

// Engine evaluates security policies and makes enforcement decisions
type Engine interface {
	// Evaluate checks if an action is allowed by current policies
	Evaluate(ctx context.Context, req *EvaluationRequest) (*EvaluationResponse, error)
	
	// LoadPolicies updates the active policy set
	LoadPolicies(policies []Policy) error
	
	// GetViolations returns recent policy violations
	GetViolations() ([]Violation, error)
}

// EvaluationRequest contains context for policy evaluation
type EvaluationRequest struct {
	PodName       string
	Namespace     string
	Action        string
	Resource      string
	AnomalyScore  float64
}

// EvaluationResponse contains the policy decision
type EvaluationResponse struct {
	Allowed bool
	Reason  string
	Actions []string // Actions to take: "alert", "block", "audit"
}

// Policy represents a security policy rule
type Policy struct {
	Name        string
	Description string
	Conditions  []Condition
	Action      string
}

// Condition represents a policy condition
type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

// Violation represents a policy violation event
type Violation struct {
	Timestamp   int64
	PodName     string
	Namespace   string
	PolicyName  string
	Description string
	Severity    string
}

// New creates a new policy engine instance
func New() (Engine, error) {
	fmt.Println("Demo policy engine created")
	fmt.Println("Enterprise version: ML-based anomaly detection and advanced rule engine")
	return &demoEngine{}, nil
}

type demoEngine struct {
	policies []Policy
}

func (e *demoEngine) Evaluate(ctx context.Context, req *EvaluationRequest) (*EvaluationResponse, error) {
	// Demo implementation
	return &EvaluationResponse{
		Allowed: true,
		Reason:  "Demo mode - no real evaluation",
		Actions: []string{"audit"},
	}, nil
}

func (e *demoEngine) LoadPolicies(policies []Policy) error {
	e.policies = policies
	fmt.Printf("Loaded %d policies (demo)\n", len(policies))
	return nil
}

func (e *demoEngine) GetViolations() ([]Violation, error) {
	return []Violation{}, nil
}
