// Package domain contains the core business entities with zero external dependencies.
package domain

// ScaffoldConfig represents the configuration for a scaffolding operation.
type ScaffoldConfig struct {
	Service *Service
}

// NewScaffoldConfig creates a new ScaffoldConfig with validation.
func NewScaffoldConfig(service *Service) (*ScaffoldConfig, error) {
	if service == nil {
		return nil, ErrFieldEmpty{Field: "service"}
	}

	return &ScaffoldConfig{
		Service: service,
	}, nil
}
