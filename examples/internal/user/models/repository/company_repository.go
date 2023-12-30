package repository

import (
	"context"
	"database/sql"
    "github.com/pworld/go-generator/examples/internal/user/models/entity"
	"github.com/pworld/go-generator/helper"
	"github.com/pworld/loggers"
)

// CompanyRepository interface for CRUD operations on Company
type CompanyRepository interface {
    CreateCompany(ctx context.Context, company entity.Company) (int, error)
    GetCompanyByID(ctx context.Context, id int) (entity.Company, error)
    UpdateCompany(ctx context.Context, company entity.Company) error
    DeleteCompany(ctx context.Context, id int) error
    ListCompanys(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, error)
    TotalCompanys(ctx context.Context, search string, filters map[string]interface{}) (int, error)
}

// companyRepository struct implements CompanyRepository with a SQL database
type companyRepository struct {
    db *sql.DB
}

// NewCompanyRepository creates a new instance of CompanyRepository
func NewCompanyRepository(db *sql.DB) CompanyRepository {
    return &companyRepository{db: db}
}

// CreateCompany inserts a new company into the database
func (r *companyRepository) CreateCompany(ctx context.Context, company entity.Company) (int, error) {
    query := `INSERT INTO Company (ID, Fullname, Email, Phone, Username, Password, CreatedAt, UpdatedAt, DeletedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id;`

    var id int
    err := r.db.QueryRowContext(ctx, query, company.ID, company.Fullname, company.Email, company.Phone, company.Username, company.Password, company.CreatedAt, company.UpdatedAt, company.DeletedAt).Scan(&id)
    if err != nil {
        loggers.Error(fmt.Sprintf("Error creating company: %s", err))
        return 0, err
    }
    return id, nil
}

// GetCompanyByID fetches a company by their ID from the database
func (r *companyRepository) GetCompanyByID(ctx context.Context, id int) (entity.Company, error) {
    query := "SELECT * FROM Company WHERE id = ?;"

    err := r.db.QueryRowContext(ctx, query, id).Scan(&company.ID, &company.Fullname, &company.Email, &company.Phone, &company.Username, &company.Password, &company.CreatedAt, &company.UpdatedAt, &company.DeletedAt)
    if err != nil {
        loggers.ErrorLog("Error fetching company by ID", "GetCompanyByID", query, 0, err)
        return entity.Company{}, err
    }

    return company, nil
}

// UpdateCompany updates an existing company's details in the database
func (r *companyRepository) UpdateCompany(ctx context.Context, company entity.Company) error {
    query := "UPDATE Company SET Fullname = ?, Email = ?, Phone = ?, Username = ?, Password = ?, CreatedAt = ?, UpdatedAt = ?, DeletedAt = ? WHERE id = ?;"

    _, err := r.db.ExecContext(ctx, query, 
    	company.Fullname, company.Email, company.Phone, company.Username, company.Password, company.CreatedAt, company.UpdatedAt, company.DeletedAt, company.ID
	)
    if err != nil {
        loggers.ErrorLog("Error updating company", "UpdateCompany", query, 0, err)
        return err
    }

    return nil
}

// DeleteCompany removes a company from the database by their ID
func (r *companyRepository) DeleteCompany(ctx context.Context, id int) error {
    query := "DELETE FROM Company WHERE id = ?;"

    _, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        loggers.ErrorLog("Error deleting company", "DeleteCompany", query, 0, err)
        return err
    }

    return nil
}

// ListCompanys fetches a list of companys based on pagination and filters
func (r *companyRepository) ListCompanys(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, error) {
    var companys []entity.Company

    // Prepare the dynamic part of the query
    searchQuery, args := helper.PrepareSearchQuery(search, filters)
    baseQuery := "SELECT * FROM Company"
    query := baseQuery + searchQuery + " LIMIT ? OFFSET ?"
    args = append(args, pageSize, (page-1)*pageSize)

    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        loggers.ErrorLog("Error fetching list of companys", "ListCompanys", query, 0, err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var company entity.Company
        if err := rows.Scan(&company.ID, &company.Fullname, &company.Email, &company.Phone, &company.Username, &company.Password, &company.CreatedAt, &company.UpdatedAt, &company.DeletedAt); err != nil {
            return nil, err
        }
        companys = append(companys, company)
    }

    return companys, nil
}

// TotalCompanys returns the total number of companys satisfying the filters
func (r *companyRepository) TotalCompanys(ctx context.Context, search string, filters map[string]interface{}) (int, error) {
    var total int

    // Prepare the dynamic part of the query
    searchQuery, args := helper.PrepareSearchQuery(search, filters)
    query := "SELECT COUNT(*) FROM Company" + searchQuery

    err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
    if err != nil {
        loggers.ErrorLog("Error counting total companys", "TotalCompanys", query, 0, err)
        return 0, err
    }

    return total, nil
}
