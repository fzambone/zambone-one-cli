// Package domain contains the core business entities with zero external dependencies.
package domain

import "fmt"

// ErrFieldEmpty represents a validation error for an empty required field.
type ErrFieldEmpty struct {
	Field string
}

// Error implements the error interface.
func (e ErrFieldEmpty) Error() string {
	return fmt.Sprintf("%s cannot be empty", e.Field)
}
