package templateMVC

const RepositoryTemplate = `package repository

import (
    "context"
    "database/sql"
    "errors"
    "{{.ModuleName}}/helper"
    "{{.ModuleName}}/helper/loggers"
    "{{.ModuleName}}/internal/{{.LowerStructName}}/models/entity"
    "strings"
)

// {{.StructName}}Repository interface for CRUD operations on {{.StructName}}
type {{.StructName}}Repository interface {
    Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error)
    Get{{.StructName}}ByID(ctx context.Context, id int) (entity.{{.StructName}}, error)
    Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error
    Delete{{.StructName}}(ctx context.Context, id int) error
    List{{.StructName}}s(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error)
    Total{{.StructName}}s(ctx context.Context, search string, filters map[string]interface{}) (int, error)
}

// {{.LowerStructName}}Repository struct implements {{.StructName}}Repository with a SQL database
type {{.LowerStructName}}Repository struct {
    db *sql.DB
}

// New{{.StructName}}Repository creates a new instance of {{.StructName}}Repository
func New{{.StructName}}Repository(db *sql.DB) {{.StructName}}Repository {
    return &{{.LowerStructName}}Repository{db: db}
}

// Create{{.StructName}} inserts a new {{.LowerStructName}} into the database
func (r *{{.LowerStructName}}Repository) Create{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) (int, error) {
    query := ` + "`{{.Query}}`" + `

    var id int
    err := r.db.QueryRowContext(ctx, query, {{.ArgumentList}}).Scan(&id)
    if err != nil {
        loggers.ErrorLog("Error creating {{.LowerStructName}}", "Create{{.StructName}}", query, 0, err)
        return 0, err
    }
    return id, nil
}

// Get{{.StructName}}ByID fetches a {{.LowerStructName}} by their ID from the database
func (r *{{.LowerStructName}}Repository) Get{{.StructName}}ByID(ctx context.Context, id int) (entity.{{.StructName}}, error) {
    query := "SELECT * FROM {{.TableName}} WHERE id = ?;"

    err := r.db.QueryRowContext(ctx, query, id).Scan({{.ScanFields}})
    if err != nil {
        loggers.ErrorLog("Error fetching {{.LowerStructName}} by ID", "Get{{.StructName}}ByID", query, 0, err)
        return entity.{{.StructName}}{}, err
    }

    return {{.LowerStructName}}, nil
}

// Update{{.StructName}} updates an existing {{.LowerStructName}}'s details in the database
func (r *{{.LowerStructName}}Repository) Update{{.StructName}}(ctx context.Context, {{.LowerStructName}} entity.{{.StructName}}) error {
    query := "UPDATE {{.TableName}} SET {{.UpdateFields}} WHERE id = ?;"

    _, err := r.db.ExecContext(ctx, query, 
    	{{.ScanFieldsUpdate}}
	)
    if err != nil {
        loggers.ErrorLog("Error updating {{.LowerStructName}}", "Update{{.StructName}}", query, 0, err)
        return err
    }

    return nil
}

// Delete{{.StructName}} removes a {{.LowerStructName}} from the database by their ID
func (r *{{.LowerStructName}}Repository) Delete{{.StructName}}(ctx context.Context, id int) error {
    query := "DELETE FROM {{.TableName}} WHERE id = ?;"

    _, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        loggers.ErrorLog("Error deleting {{.LowerStructName}}", "Delete{{.StructName}}", query, 0, err)
        return err
    }

    return nil
}

// List{{.StructName}}s fetches a list of {{.LowerStructName}}s based on pagination and filters
func (r *{{.LowerStructName}}Repository) List{{.StructName}}s(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.{{.StructName}}, error) {
    var {{.LowerStructName}}s []entity.{{.StructName}}

    // Prepare the dynamic part of the query
    searchQuery, args := helper.PrepareSearchQuery(search, filters)
    baseQuery := "SELECT * FROM {{.TableName}}"
    query := baseQuery + searchQuery + " LIMIT ? OFFSET ?"
    args = append(args, pageSize, (page-1)*pageSize)

    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        loggers.ErrorLog("Error fetching list of {{.LowerStructName}}s", "List{{.StructName}}s", query, 0, err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var {{.LowerStructName}} entity.{{.StructName}}
        if err := rows.Scan({{.ScanFields}}); err != nil {
            return nil, err
        }
        {{.LowerStructName}}s = append({{.LowerStructName}}s, {{.LowerStructName}})
    }

    return {{.LowerStructName}}s, nil
}

// Total{{.StructName}}s returns the total number of {{.LowerStructName}}s satisfying the filters
func (r *{{.LowerStructName}}Repository) Total{{.StructName}}s(ctx context.Context, search string, filters map[string]interface{}) (int, error) {
    var total int

    // Prepare the dynamic part of the query
    searchQuery, args := helper.PrepareSearchQuery(search, filters)
    query := "SELECT COUNT(*) FROM {{.TableName}}" + searchQuery

    err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
    if err != nil {
        loggers.ErrorLog("Error counting total {{.LowerStructName}}s", "Total{{.StructName}}s", query, 0, err)
        return 0, err
    }

    return total, nil
}
`