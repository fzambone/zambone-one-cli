// Package fakes handwritten test doubles for port interfaces.
// These fakes are simple, in-memory implementations used across all tests.
package fakes

import "sync"

// InMemoryFileSystem implements ports.FileSystem using in-memory maps.
// It's used for testing - files are stored in RAM, not in disk.
type InMemoryFileSystem struct {
	mu    sync.RWMutex
	files map[string][]byte
	dirs  map[string]bool
}

// NewInMemoryFileSystem creates a new in-memory Filesystem.
func NewInMemoryFileSystem() *InMemoryFileSystem {
	return &InMemoryFileSystem{
		files: make(map[string][]byte),
		dirs:  make(map[string]bool),
	}
}

// CreateDirectory records that a directory was created
func (f *InMemoryFileSystem) CreateDirectory(path string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.dirs[path] = true
	return nil
}

// WriteFile stores the file content in memory
func (f *InMemoryFileSystem) WriteFile(path string, content []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.files[path] = content
	return nil
}

// Exists check if a file or directory exists in memory
func (f *InMemoryFileSystem) Exists(path string) (bool, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if _, ok := f.files[path]; ok {
		return true, nil
	}

	if _, ok := f.dirs[path]; ok {
		return true, nil
	}

	return false, nil
}

// FileContent returns the value of a file if it exists.
// Used by tests to verify what was written.
func (f *InMemoryFileSystem) FileContent(path string) ([]byte, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	content, ok := f.files[path]
	return content, ok
}

// FileCount returns how many files are stored.
// Used by tests to verify the exact number of files.
func (f *InMemoryFileSystem) FileCount() int {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return len(f.files)
}

// SeedFile pre-populates a file (for testing "file already exists" scenarios).
func (f *InMemoryFileSystem) SeedFile(path string, content []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.files[path] = content
}

// SeedDir pre-populates a directory.
func (f *InMemoryFileSystem) SeedDir(path string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.dirs[path] = true
}
