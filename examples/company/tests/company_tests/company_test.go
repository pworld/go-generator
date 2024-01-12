package company_test

import (
	"context"
	"github.com/pworld/go-generator/examples/company/models/entity"
	"github.com/pworld/go-generator/examples/company/services"
	"github.com/pworld/go-generator/examples/company/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	company := entity.Company{
		ID:       0,
		Fullname: "",
		Email:    "",
		Phone:    "",
	}

	mockRepo.On("CreateCompany", mock.Anything, mock.AnythingOfType("entity.Company")).Return(1, nil)

	service := services.NewCompanyService(mockRepo)
	id, err := service.CreateCompany(context.Background(), company)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockRepo.AssertExpectations(t)
}

func TestGetCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	testID := 1 // Example test ID
	expectedCompany := entity.Company{
		ID:       0,
		Fullname: "",
		Email:    "",
		Phone:    "",
	}

	mockRepo.On("GetCompanyByID", mock.Anything, testID).Return(expectedCompany, nil)

	service := services.NewCompanyService(mockRepo)
	resultCompany, err := service.GetCompany(context.Background(), testID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCompany, resultCompany)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	companyToUpdate := entity.Company{
		ID:       0,
		Fullname: "",
		Email:    "",
		Phone:    "",
	}

	mockRepo.On("UpdateCompany", mock.Anything, mock.AnythingOfType("entity.Company")).Return(nil)
	service := services.NewCompanyService(mockRepo)
	err := service.UpdateCompany(context.Background(), companyToUpdate)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	testID := 1 // Example test ID
	mockRepo.On("DeleteCompany", mock.Anything, testID).Return(nil)

	service := services.NewCompanyService(mockRepo)
	err := service.DeleteCompany(context.Background(), testID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompanyLists(t *testing.T) {
	mockRepo := new(mocks.MockCompanyRepository)
	page, pageSize := 1, 10

	expectedCompanys := []entity.Company{
		ID:       0,
		Fullname: "",
		Email:    "",
		Phone:    "",
	}
	expectedTotal := int64(len(expectedCompanys))

	mockRepo.On("ListCompanys", mock.Anything, page, pageSize, "", mock.Anything).Return(expectedCompanys, nil)
	mockRepo.On("TotalCompanys", mock.Anything, "", mock.Anything).Return(expectedTotal, nil)

	service := services.NewCompanyService(mockRepo)
	resultCompanys, total, err := service.CompanyLists(context.Background(), page, pageSize, "", nil)

	assert.NoError(t, err)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedCompanys, resultCompanys)
	mockRepo.AssertExpectations(t)
}
