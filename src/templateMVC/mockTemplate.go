package templateMVC

const MockTemplate = `package mocks

import (
	"context"
	"{{.ModuleName}}/internal/{{.PackageName}}/models/entity"
	"github.com/stretchr/testify/mock"
)

// Mock{{.StructName}}Repository is a mock type for the {{.StructName}}Repository type
type Mock{{.StructName}}Repository struct {
	mock.Mock
}

// CreateUser mocks the CreateUser method
func (m *Mock{{.StructName}}Repository) CreateUser(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error) {
	args := m.Called(ctx, {{.LowerStructName}})
	return args.Int(0), args.Error(1)
}

// GetUserByID mocks the GetUserByID method
func (m *Mock{{.StructName}}Repository) GetUserByID(ctx context.Context, id int) (entity.{{.StructName}}, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.{{.StructName}}), args.Error(1)
}

// UpdateUser mocks the UpdateUser method
func (m *Mock{{.StructName}}Repository) UpdateUser(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
	args := m.Called(ctx, {{.LowerStructName}})
	return args.Error(0)
}

// DeleteUser mocks the DeleteUser method
func (m *Mock{{.StructName}}Repository) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// ListUsers mocks the ListUsers method
func (m *Mock{{.StructName}}Repository) ListUsers(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error) {
	args := m.Called(ctx, page, pageSize, search, filters)
	return args.Get(0).([]entity.{{.StructName}}), args.Error(1)
}

// TotalUsers mocks the TotalUsers method
func (m *Mock{{.StructName}}Repository) TotalUsers(ctx context.Context, search string, filters map[string]interface{}) (int, error) {
	args := m.Called(ctx, search, filters)
	return args.Int(0), args.Error(1)
}
`
