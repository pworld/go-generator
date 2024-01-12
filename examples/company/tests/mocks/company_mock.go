package mocks

import (
	"context"
	"github.com/pworld/go-generator/examples/company/models/entity"
	"github.com/stretchr/testify/mock"
)

// MockCompanyRepository is a mock type for the CompanyRepository type
type MockCompanyRepository struct {
	mock.Mock
}

// CreateCompany mocks the CreateCompany method
func (m *MockCompanyRepository) CreateCompany(ctx context.Context, company entity.Company) (int, error) {
	args := m.Called(ctx, company)
	return args.Int(0), args.Error(1)
}

// GetCompanyByID mocks the GetCompanyByID method
func (m *MockCompanyRepository) GetCompanyByID(ctx context.Context, id int) (entity.Company, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Company), args.Error(1)
}

// UpdateCompany mocks the UpdateCompany method
func (m *MockCompanyRepository) UpdateCompany(ctx context.Context, company entity.Company) error {
	args := m.Called(ctx, company)
	return args.Error(0)
}

// DeleteCompany mocks the DeleteCompany method
func (m *MockCompanyRepository) DeleteCompany(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// ListCompany mocks the ListCompany method
func (m *MockCompanyRepository) ListCompany(ctx context.Context, page, pageSize int, search string, filters map[string]interface{}) ([]entity.Company, error) {
	args := m.Called(ctx, page, pageSize, search, filters)
	return args.Get(0).([]entity.Company), args.Error(1)
}

// TotalCompany mocks the TotalCompany method
func (m *MockCompanyRepository) TotalCompany(ctx context.Context, search string, filters map[string]interface{}) (int64, error) {
	args := m.Called(ctx, search, filters)
	return int64(args.Int(0)), args.Error(1)
}
