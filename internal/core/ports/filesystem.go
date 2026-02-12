// Package ports defines interfaces for external dependencies.
// These are the "ports" in Ports & Adapters (Hexagonal) architecture.
package ports

// FileSystem defines operations for file and directory management.
type FileSystem interface {
	// CreateDirectory creates a directory at the specified path.
	CreateDirectory(path string) error

	// WriteFile writes content to a file at the specified path.
	WriteFile(path string, content []byte) error

	// Exists checks if a file or a directory exists at the specified path.
	Exists(path string) (bool, error)
}
