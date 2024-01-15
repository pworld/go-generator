package templateMVC

const RepositoryTemplate = `package repository

import (
	"context"
	"{{.ModuleName}}/helper"
	"{{.ModulePath}}/models/entity"
	"gorm.io/gorm"
)

// {{.StructName}}Repository interface for CRUD operations on {{.StructName}}
type {{.StructName}}Repository interface {
	Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error)
	Get{{.StructName}}ByID(ctx context.Context, id int) (entity.{{.StructName}}, error)
	Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error
	Delete{{.StructName}}(ctx context.Context, id int) error
	List{{.StructName}}s(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error)
	Total{{.StructName}}s(ctx context.Context, search string, filters map[string]interface{}) (int64, error)
}

// {{.LowerStructName}}Repository struct implements {{.StructName}}Repository with GORM
type {{.LowerStructName}}Repository struct {
	db *gorm.DB
}

// New{{.StructName}}Repository creates a new instance of {{.StructName}}Repository
func New{{.StructName}}Repository(db *gorm.DB) {{.StructName}}Repository {
	return &{{.LowerStructName}}Repository{db: db}
}

// Create{{.StructName}}, Get{{.StructName}}ByID, Update{{.StructName}}, Delete{{.StructName}}, List{{.StructName}}s, Total{{.StructName}}s implementations
func (r *{{.LowerStructName}}Repository) Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error) {
    result := r.db.WithContext(ctx).Create(&{{.LowerStructName}})
    if result.Error != nil {
        return 0, result.Error
    }
    return {{.LowerStructName}}.ID, nil
}

// Get{{.StructName}}ByID fetches a {{.LowerStructName}} by their ID from the database
func (r *{{.LowerStructName}}Repository) Get{{.StructName}}ByID(ctx context.Context, id int) (entity.{{.StructName}}, error) {
	var {{.LowerStructName}} entity.{{.StructName}}
	result := r.db.WithContext(ctx).First(&{{.LowerStructName}}, id)
	return {{.LowerStructName}}, result.Error
}

// Update{{.StructName}} updates an existing {{.LowerStructName}}'s details in the database
func (r *{{.LowerStructName}}Repository) Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
	return r.db.WithContext(ctx).Save(&{{.LowerStructName}}).Error
}

// Delete{{.StructName}} removes a {{.LowerStructName}} from the database by their ID
func (r *{{.LowerStructName}}Repository) Delete{{.StructName}}(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.{{.StructName}}{}, id).Error
}

// List{{.StructName}}s fetches a list of {{.LowerStructName}}s based on pagination and filters
func (r *{{.LowerStructName}}Repository) List{{.StructName}}s(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error) {
	var {{.LowerStructName}}s []entity.{{.StructName}}
	query := r.db.WithContext(ctx).Model(&entity.{{.StructName}}{})

	// Apply dynamic search and filters
	query = helper.ApplyDynamicSearchAndFilters(query, search, filters)

	result := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&{{.LowerStructName}}s)
	return {{.LowerStructName}}s, result.Error
}

// Total{{.StructName}}s returns the total number of {{.LowerStructName}}s satisfying the filters
func (r *{{.LowerStructName}}Repository) Total{{.StructName}}s(ctx context.Context, search string, filters map[string]interface{}) (int64, error) {
	var total int64
	query := r.db.WithContext(ctx).Model(&entity.{{.StructName}}{})

	// Apply dynamic search and filters
	query = helper.ApplyDynamicSearchAndFilters(query, search, filters)

	result := query.Count(&total)
	return total, result.Error
}
`
