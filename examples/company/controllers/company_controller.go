package controllers

import (
	"github.com/pworld/go-generator/examples/company/models/entity"
	"github.com/pworld/go-generator/examples/company/models/repository"
	"github.com/pworld/go-generator/examples/company/services"
	"github.com/pworld/go-generator/examples/company/views"
	"github.com/pworld/loggers"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

// CompanyController structure
type CompanyController struct {
	companyService services.CompanyService
}

// NewCompanyController creates a new instance of CompanyController
func NewCompanyController(db *gorm.DB) *CompanyController {
	companyRepo := repository.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	return &CompanyController{
		companyService: companyService,
	}
}

// CreateCompany creates a new company
func (uc *CompanyController) CreateCompany(c *fiber.Ctx) error {
	var company entity.Company
	if err := parseAndValidate(c, &company); err != nil {
		loggers.Error("Error parsing request body:", err)
		return handleValidationError(c, err)
	}

	company.CreatedAt = time.Now()

	if _, err := uc.companyService.CreateCompany(c.Context(), company); err != nil {
		loggers.Error("Error creating Company:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(company)
}

// GetCompany retrieves a company by ID
func (uc *CompanyController) GetCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}

	company, err := uc.companyService.GetCompany(c.Context(), id)
	if err != nil {
		return views.CompanyErrorResponse(c, fiber.StatusInternalServerError, "Company not found")
	}
	return views.CompanyResponse(c, company)
}

// UpdateCompany updates a company
func (uc *CompanyController) UpdateCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}

	var company entity.Company
	if err := parseAndValidate(c, &company); err != nil {
		return handleValidationError(c, err)
	}

	company.ID = id
	company.UpdatedAt = time.Now()

	if err := uc.companyService.UpdateCompany(c.Context(), company); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(company)
}

// DeleteCompany deletes a company by ID
func (uc *CompanyController) DeleteCompany(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}

	if err := uc.companyService.DeleteCompany(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CompanyLists lists all company
func (uc *CompanyController) CompanyLists(c *fiber.Ctx) error {
	page, pageSize, err := helper.ExtractPageAndSize(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pagination parameters"})
	}
	search, filters := helper.ExtractSearchAndParams(c)

	company, total, err := uc.companyService.CompanyLists(c.Context(), page, pageSize, search, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve company list"})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	return views.CompanyListResponse(c, company, total, totalPages)
}
