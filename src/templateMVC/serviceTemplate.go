package templateMVC

const ServiceTemplate = `package services

import (
	"context"
	"fmt"
	"{{.ModulePath}}/models/entity"
	"{{.ModulePath}}/models/repository"
	"github.com/pworld/loggers"
)

// {{.StructName}}Service interface for {{.LowerStructName}}-related business logic
type {{.StructName}}Service interface {
	Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error)
	Get{{.StructName}}(ctx context.Context, id int) (entity.{{.StructName}}, error)
	Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error
	Delete{{.StructName}}(ctx context.Context, id int) error
	{{.StructName}}Lists(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, int64, error)
}

// {{.LowerStructName}}Service struct implements {{.StructName}}Service
type {{.LowerStructName}}Service struct {
	{{.LowerStructName}}Repo repository.{{.StructName}}Repository
}

// New{{.StructName}}Service creates a new instance of {{.StructName}}Service
func New{{.StructName}}Service({{.LowerStructName}}Repo repository.{{.StructName}}Repository) {{.StructName}}Service {
	return &{{.LowerStructName}}Service{ {{.LowerStructName}}Repo: {{.LowerStructName}}Repo }
}

// Create{{.StructName}} handles the creation of a new {{.LowerStructName}}
func (s *{{.LowerStructName}}Service) Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error) {
	id, err := s.{{.LowerStructName}}Repo.Create{{.StructName}}(ctx, {{.LowerStructName}})
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create {{.LowerStructName}}:%s", err), "Create{{.StructName}}", "{{.LowerStructName}}/service", 500)
		return 0, err
	}
	return id, nil
}

// Get{{.StructName}} retrieves a {{.LowerStructName}} by their ID
func (s *{{.LowerStructName}}Service) Get{{.StructName}}(ctx context.Context, id int) (entity.{{.StructName}}, error) {
	{{.LowerStructName}}, err := s.{{.LowerStructName}}Repo.Get{{.StructName}}ByID(ctx, id)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to retrieve {{.LowerStructName}}:%s", err), "Get{{.StructName}}", "{{.LowerStructName}}/service", 500)
		return entity.{{.StructName}}{}, err
	}
	return {{.LowerStructName}}, nil
}

// Update{{.StructName}} handles updating {{.LowerStructName}} details
func (s *{{.LowerStructName}}Service) Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
	if err := s.{{.LowerStructName}}Repo.Update{{.StructName}}(ctx, {{.LowerStructName}}); err != nil {
		loggers.Error(fmt.Sprintf("Failed to update {{.LowerStructName}}:%s", err), "Update{{.StructName}}", "{{.LowerStructName}}/service", 500)
		return err
	}
	return nil
}

// Delete{{.StructName}} handles the deletion of a {{.LowerStructName}}
func (s *{{.LowerStructName}}Service) Delete{{.StructName}}(ctx context.Context, id int) error {
	if err := s.{{.LowerStructName}}Repo.Delete{{.StructName}}(ctx, id); err != nil {
		loggers.Error(fmt.Sprintf("Failed to delete {{.LowerStructName}}:%s", err), "Delete{{.StructName}}", "{{.LowerStructName}}/service", 500)
		return err
	}
	return nil
}

// {{.StructName}}Lists retrieves a list of {{.LowerStructName}}s with pagination and filtering
func (s *{{.LowerStructName}}Service) {{.StructName}}Lists(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, int64, error) {
	{{.LowerStructName}}s, err := s.{{.LowerStructName}}Repo.List{{.StructName}}s(ctx, page, pageSize, search, filters)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to list {{.LowerStructName}}s:%s", err), "{{.StructName}}Lists", "{{.LowerStructName}}/service", 500)
		return nil, 0, err
	}
	total, err := s.{{.LowerStructName}}Repo.Total{{.StructName}}s(ctx, search, filters)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to count total {{.LowerStructName}}s:%s", err), "{{.StructName}}Lists", "{{.LowerStructName}}/service", 500)
		return nil, 0, err
	}
	return {{.LowerStructName}}s, total, nil
}
`
