package services

import (
	"context"
	"github.com/pworld/go-generator/examples/internal/user/models/entity"
	"github.com/pworld/go-generator/examples/internal/user/models/repository"
)

// CompanyService interface for company-related business logic
type CompanyService interface {
	CreateCompany(ctx context.Context, company entity.Company) (int, error)
	GetCompany(ctx context.Context, id int) (entity.Company, error)
	UpdateCompany(ctx context.Context, company entity.Company) error
	DeleteCompany(ctx context.Context, id int) error
}

// companyService struct implements CompanyService
type companyService struct {
	companyRepo repository.CompanyRepository
}

// NewCompanyService creates a new instance of CompanyService
func NewCompanyService(companyRepo repository.CompanyRepository) CompanyService {
	return &companyService{companyRepo: companyRepo}
}

// CreateCompany handles the creation of a new company
func (s *companyService) CreateCompany(ctx context.Context, company entity.Company) (int, error) {
	return s.companyRepo.CreateCompany(ctx, company)
}

// GetCompany retrieves a company by their ID
func (s *companyService) GetCompany(ctx context.Context, id int) (entity.Company, error) {
	return s.companyRepo.GetCompanyByID(ctx, id)
}

// UpdateCompany handles updating company details
func (s *companyService) UpdateCompany(ctx context.Context, company entity.Company) error {
	return s.companyRepo.UpdateCompany(ctx, company)
}

// DeleteCompany handles the deletion of a company
func (s *companyService) DeleteCompany(ctx context.Context, id int) error {
	return s.companyRepo.DeleteCompany(ctx, id)
}
