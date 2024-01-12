package repository

import (
	"context"
	"github.com/pworld/go-generator/examples/company/models/entity"
	"github.com/pworld/go-mvc-boilerplate/helper"
	"gorm.io/gorm"
)

// CompanyRepository interface for CRUD operations on Company
type CompanyRepository interface {
	CreateCompany(ctx context.Context, company entity.Company) (int, error)
	GetCompanyByID(ctx context.Context, id int) (entity.Company, error)
	UpdateCompany(ctx context.Context, company entity.Company) error
	DeleteCompany(ctx context.Context, id int) error
	ListCompanys(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, error)
	TotalCompanys(ctx context.Context, search string, filters map[string]interface{}) (int64, error)
}

// companyRepository struct implements CompanyRepository with GORM
type companyRepository struct {
	db *gorm.DB
}

// NewCompanyRepository creates a new instance of CompanyRepository
func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

// CreateCompany, GetCompanyByID, UpdateCompany, DeleteCompany, ListCompanys, TotalCompanys implementations
func (r *companyRepository) CreateCompany(ctx context.Context, company entity.Company) (int, error) {
	result := r.db.WithContext(ctx).Create(&company)
	if result.Error != nil {
		return 0, result.Error
	}
	return company.ID, nil
}

// GetCompanyByID fetches a company by their ID from the database
func (r *companyRepository) GetCompanyByID(ctx context.Context, id int) (entity.Company, error) {
	var company entity.Company
	result := r.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

// UpdateCompany updates an existing company's details in the database
func (r *companyRepository) UpdateCompany(ctx context.Context, company entity.Company) error {
	return r.db.WithContext(ctx).Save(&company).Error
}

// DeleteCompany removes a company from the database by their ID
func (r *companyRepository) DeleteCompany(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Company{}, id).Error
}

// ListCompanys fetches a list of companys based on pagination and filters
func (r *companyRepository) ListCompanys(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, error) {
	var companys []entity.Company
	query := r.db.WithContext(ctx).Model(&entity.Company{})

	// Apply dynamic search and filters
	query = helper.ApplyDynamicSearchAndFilters(query, search, filters)

	result := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&companys)
	return companys, result.Error
}

// TotalCompanys returns the total number of companys satisfying the filters
func (r *companyRepository) TotalCompanys(ctx context.Context, search string, filters map[string]interface{}) (int64, error) {
	var total int64
	query := r.db.WithContext(ctx).Model(&entity.Company{})

	// Apply dynamic search and filters
	query = helper.ApplyDynamicSearchAndFilters(query, search, filters)

	result := query.Count(&total)
	return total, result.Error
}
