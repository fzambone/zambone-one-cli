package fakes_test

import (
	"testing"

	"github.com/fzambone/zambone-one-cli/internal/core/ports/fakes"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryFileSystem_WriteAndRead(t *testing.T) {
	fs := fakes.NewInMemoryFileSystem()

	// Write a File
	err := fs.WriteFile("test.txt", []byte("Testing Writing File"))
	assert.NoError(t, err)

	// Verify it exists
	exists, err := fs.Exists("test.txt")
	assert.NoError(t, err)
	assert.True(t, exists)

	// Read the content block
	content, ok := fs.FileContent("test.txt")
	assert.True(t, ok)
	assert.Equal(t, "Testing Writing File", string(content))
}

func TestInMemoryFileSystem_CreateDirectory(t *testing.T) {
	fs := fakes.NewInMemoryFileSystem()

	// Create a directory
	err := fs.CreateDirectory("testdir")
	assert.NoError(t, err)

	// Verify it exists
	exists, err := fs.Exists("testdir")
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestInMemoryFileSystem_FileCount(t *testing.T) {
	fs := fakes.NewInMemoryFileSystem()

	assert.Equal(t, 0, fs.FileCount())

	_ = fs.WriteFile("file1.txt", []byte("a"))
	assert.Equal(t, 1, fs.FileCount())

	_ = fs.WriteFile("file2.txt", []byte("b"))
	assert.Equal(t, 2, fs.FileCount())
}
