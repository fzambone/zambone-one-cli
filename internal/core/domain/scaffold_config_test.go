//nolint:goconst,revive,var-declaration
package domain_test

import (
	"testing"

	"github.com/fzambone/zambone-one-cli/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewScaffoldConfig_ValidService_CreatesConfig(t *testing.T) {
	// Arrange
	serviceName := "user-api"
	serviceType := "rails"
	servicePath := "./services"
	service, _ := domain.NewService(serviceName, serviceType, servicePath)

	// Act
	config, err := domain.NewScaffoldConfig(service)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, serviceName, config.Service.Name)
	assert.Equal(t, serviceType, config.Service.Type)
	assert.Equal(t, servicePath, config.Service.Path)
}

func TestNewScaffoldConfig_NilService_ReturnsError(t *testing.T) {
	// Arrange
	var service *domain.Service = nil

	// Act
	config, err := domain.NewScaffoldConfig(service)

	// Assert
	var fieldErr domain.ErrFieldEmpty
	assert.ErrorAs(t, err, &fieldErr)
	assert.Nil(t, config)
	assert.Equal(t, "service", fieldErr.Field)
}
