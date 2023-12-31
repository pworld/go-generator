package templateMVC

const ServiceTemplate = `package services

import (
    "context"
    "{{.ModuleName}}/internal/{{.PackageName}}/models/entity"
    "{{.ModuleName}}/internal/{{.PackageName}}/models/repository"
)

// {{.StructName}}Service interface for {{.LowerStructName}}-related business logic
type {{.StructName}}Service interface {
    Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error)
    Get{{.StructName}}(ctx context.Context, id int) (entity.{{.StructName}}, error)
    Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error
    Delete{{.StructName}}(ctx context.Context, id int) error
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
	return s.{{.LowerStructName}}Repo.Create{{.StructName}}(ctx, {{.LowerStructName}})
}

// Get{{.StructName}} retrieves a {{.LowerStructName}} by their ID
func (s *{{.LowerStructName}}Service) Get{{.StructName}}(ctx context.Context, id int) (entity.{{.StructName}}, error) {
	return s.{{.LowerStructName}}Repo.Get{{.StructName}}ByID(ctx, id)
}

// Update{{.StructName}} handles updating {{.LowerStructName}} details
func (s *{{.LowerStructName}}Service) Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
	return s.{{.LowerStructName}}Repo.Update{{.StructName}}(ctx, {{.LowerStructName}})
}

// Delete{{.StructName}} handles the deletion of a {{.LowerStructName}}
func (s *{{.LowerStructName}}Service) Delete{{.StructName}}(ctx context.Context, id int) error {
	return s.{{.LowerStructName}}Repo.Delete{{.StructName}}(ctx, id)
}

// {{.StructName}}Lists retrieves a list of {{.LowerStructName}} with pagination and filtering
func (s *{{.LowerStructName}}Service) {{.StructName}}Lists(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, int, error) {
	{{.LowerStructName}}, err := s.{{.LowerStructName}}Repo.List{{.StructName}}(ctx, page, pageSize, search, filters)
	if err != nil {
		loggers.Error("Failed to list {{.LowerStructName}}", "{{.StructName}}Lists", "{{.LowerStructName}}/service", 500, err)
		return nil, 0, err
	}

	total, err := s.{{.LowerStructName}}Repo.Total{{.StructName}}(ctx, search, filters)
	if err != nil {
		loggers.Error("Failed to count total {{.LowerStructName}}", "{{.StructName}}Lists", "{{.LowerStructName}}/service", 500, err)
		return nil, 0, err
	}

	return {{.LowerStructName}}, total, nil
}

`
