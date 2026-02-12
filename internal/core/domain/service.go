// Package domain contains the core business entities with zero external dependencies.
package domain

// Service represents a macro service being scaffolded.
type Service struct {
	Name string
	Type string
	Path string
}

// NewService creates a new Service with validation.
func NewService(name, serviceType, path string) (*Service, error) {
	if name == "" {
		return nil, ErrFieldEmpty{Field: "service name"}
	}
	if serviceType == "" {
		return nil, ErrFieldEmpty{Field: "service type"}
	}
	if path == "" {
		return nil, ErrFieldEmpty{Field: "service path"}
	}

	return &Service{
		Name: name,
		Type: serviceType,
		Path: path,
	}, nil
}
