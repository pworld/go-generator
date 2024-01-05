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

// Create{{.StructName}} mocks the Create{{.StructName}} method
func (m *Mock{{.StructName}}Repository) Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error) {
	args := m.Called(ctx, {{.LowerStructName}})
	return args.Int(0), args.Error(1)
}

// Get{{.StructName}}ByID mocks the Get{{.StructName}}ByID method
func (m *Mock{{.StructName}}Repository) Get{{.StructName}}ByID(ctx context.Context, id int) (entity.{{.StructName}}, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.{{.StructName}}), args.Error(1)
}

// Update{{.StructName}} mocks the Update{{.StructName}} method
func (m *Mock{{.StructName}}Repository) Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
	args := m.Called(ctx, {{.LowerStructName}})
	return args.Error(0)
}

// Delete{{.StructName}} mocks the Delete{{.StructName}} method
func (m *Mock{{.StructName}}Repository) Delete{{.StructName}}(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// List{{.StructName}} mocks the List{{.StructName}} method
func (m *Mock{{.StructName}}Repository) List{{.StructName}}(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error) {
	args := m.Called(ctx, page, pageSize, search, filters)
	return args.Get(0).([]entity.{{.StructName}}), args.Error(1)
}

// Total{{.StructName}} mocks the Total{{.StructName}} method
func (m *Mock{{.StructName}}Repository) Total{{.StructName}}(ctx context.Context, search string, filters map[string]interface{}) (int, error) {
	args := m.Called(ctx, search, filters)
	return args.Int(0), args.Error(1)
}
`
