package company_test

import (
	"context"
	"github.com/pworld/go-generator/examples/internal/user/models/entity"
	"github.com/pworld/go-generator/examples/internal/user/services"
	"github.com/pworld/go-generator/examples/internal/user/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCompanyService_CreateCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	svc := services.NewCompanyService(mockRepo)
	mockRepo.On("CreateCompany", mock.Anything, mock.AnythingOfType("entity.Company")).Return(1, nil)
	mockCompany := entity.Company{
		// Fill The test case
	}
	companyID, err := svc.RegisterCompany(context.Background(), mockCompany)
	assert.NoError(t, err)
	assert.Equal(t, 1, companyID)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_UpdateCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	svc := services.NewCompanyService(mockRepo)
	mockRepo.On("UpdateCompany", mock.Anything, mock.AnythingOfType("entity.Company")).Return(nil)
	mockCompany := entity.Company{
		// Fill The test case
	}
	err := svc.UpdateCompany(context.Background(), mockCompany)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_DeleteCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	svc := services.NewCompanyService(mockRepo)
	mockRepo.On("DeleteCompany", mock.Anything, testCompanyID).Return(nil)
	testCompanyID := 1
	err := svc.DeleteCompany(context.Background(), testCompanyID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_GetCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	svc := services.NewCompanyService(mockRepo)
	mockRepo.On("GetCompanyByID", mock.Anything, testCompanyID).Return(mockCompany, nil)
	testCompanyID := 1
	mockCompany := entity.Company{
		// Fill The test case
	}
	company, err := svc.GetCompany(context.Background(), testCompanyID)
	assert.NoError(t, err)
	assert.Equal(t, mockCompany, company)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_ListCompanys(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	svc := services.NewCompanyService(mockRepo)
	mockRepo.On("ListCompanys", mock.Anything, page, pageSize, "", mock.Anything).Return(mockCompanys, nil)
	mockRepo.On("TotalCompanys", mock.Anything, "", mock.Anything).Return(totalCompanys, nil)
	page, pageSize := 1, 10
	totalCompanys := 2 // Total number of Companys available
	mockCompanys := []entity.Company{
		// Fill The test case
	}
	companys, total, err := svc.ListCompanys(context.Background(), page, pageSize, "", nil)
	assert.NoError(t, err)
	assert.Equal(t, totalCompanys, total)
	assert.Len(t, companys, len(mockCompanys))
	assert.Equal(t, mockCompanys, companys)
	mockRepo.AssertExpectations(t)
}
