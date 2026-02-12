// nolint:goconst
package domain_test

import (
	"testing"

	"github.com/fzambone/zambone-one-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewService_ValidInputs_CreatesService(t *testing.T) {
	// Arrange
	name := "user-api"
	serviceType := "rails"
	path := "./services"

	// Act
	service, err := domain.NewService(name, serviceType, path)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, "user-api", service.Name)
	assert.Equal(t, "rails", service.Type)
	assert.Equal(t, "./services", service.Path)
}

func TestNewService_EmptyName_WithError(t *testing.T) {
	// Arrange
	name := ""
	serviceType := "rails"
	path := "./services"

	// Act
	service, err := domain.NewService(name, serviceType, path)

	// Assert
	var fieldErr domain.ErrFieldEmpty
	assert.ErrorAs(t, err, &fieldErr)
	assert.Nil(t, service)
	assert.Equal(t, "service name", fieldErr.Field)
}

func TestNewService_EmptyType_WithError(t *testing.T) {
	// Arrange
	name := "users-api"
	serviceType := ""
	path := "./services"

	// Act
	service, err := domain.NewService(name, serviceType, path)

	// Assert
	var fieldError domain.ErrFieldEmpty
	assert.ErrorAs(t, err, &fieldError)
	assert.Nil(t, service)
	assert.Equal(t, "service type", fieldError.Field)
}

func TestNewService_EmptyPath_WithError(t *testing.T) {
	// Arrange
	name := "users-api"
	serviceType := "rails"
	path := ""

	// Act
	service, err := domain.NewService(name, serviceType, path)

	// Assert
	var fieldErr domain.ErrFieldEmpty
	assert.ErrorAs(t, err, &fieldErr)
	assert.Nil(t, service)
	assert.Equal(t, "service path", fieldErr.Field)
}
