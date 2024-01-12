package services

import (
	"context"
	"fmt"
	"github.com/pworld/go-generator/examples/company/models/entity"
	"github.com/pworld/go-generator/examples/company/models/repository"
	"github.com/pworld/loggers"
)

// CompanyService interface for company-related business logic
type CompanyService interface {
	CreateCompany(ctx context.Context, company entity.Company) (int, error)
	GetCompany(ctx context.Context, id int) (entity.Company, error)
	UpdateCompany(ctx context.Context, company entity.Company) error
	DeleteCompany(ctx context.Context, id int) error
	CompanyLists(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, int64, error)
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
	id, err := s.companyRepo.CreateCompany(ctx, company)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create company:%s", err), "CreateCompany", "company/service", 500)
		return 0, err
	}
	return id, nil
}

// GetCompany retrieves a company by their ID
func (s *companyService) GetCompany(ctx context.Context, id int) (entity.Company, error) {
	company, err := s.companyRepo.GetCompanyByID(ctx, id)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to retrieve company:%s", err), "GetCompany", "company/service", 500)
		return entity.Company{}, err
	}
	return company, nil
}

// UpdateCompany handles updating company details
func (s *companyService) UpdateCompany(ctx context.Context, company entity.Company) error {
	if err := s.companyRepo.UpdateCompany(ctx, company); err != nil {
		loggers.Error(fmt.Sprintf("Failed to update company:%s", err), "UpdateCompany", "company/service", 500)
		return err
	}
	return nil
}

// DeleteCompany handles the deletion of a company
func (s *companyService) DeleteCompany(ctx context.Context, id int) error {
	if err := s.companyRepo.DeleteCompany(ctx, id); err != nil {
		loggers.Error(fmt.Sprintf("Failed to delete company:%s", err), "DeleteCompany", "company/service", 500)
		return err
	}
	return nil
}

// CompanyLists retrieves a list of companys with pagination and filtering
func (s *companyService) CompanyLists(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, int64, error) {
	companys, err := s.companyRepo.ListCompanys(ctx, page, pageSize, search, filters)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to list companys:%s", err), "CompanyLists", "company/service", 500)
		return nil, 0, err
	}
	total, err := s.companyRepo.TotalCompanys(ctx, search, filters)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to count total companys:%s", err), "CompanyLists", "company/service", 500)
		return nil, 0, err
	}
	return companys, total, nil
}
