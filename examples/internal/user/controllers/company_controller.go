package controllers

import (
	"database/sql"
	"github.com/pworld/go-generator/examples/internal/user/models/entity"
	"github.com/pworld/go-generator/examples/internal/user/models/repository"
	"github.com/pworld/go-generator/examples/internal/user/services"
	"github.com/pworld/go-generator/examples/internal/user/views"
	"github.com/pworld/go-generator/helper"
	"github.com/pworld/loggers"
	"math"
	"strconv"
	"time"
)

// CompanyController structure
type CompanyController struct {
	companyService services.CompanyService
}

func NewCompanyController(db *sql.DB) *CompanyController {
	companyRepo := repository.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	return &CompanyController{
		companyService: companyService,
	}
}

// CreateUser creates a new company
func (uc *CompanyController) CreateUser(c *fiber.Ctx) error {
	var company entity.Company
	if err := c.BodyParser(&company); err != nil {
		loggers.Error(fmt.Sprintf("Error parsing request body: %s", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	company.CreatedAt = time.Now()

	if _, err := uc.companyService.CreateCompany(c.Context(), company); err != nil {
		loggers.Error(fmt.Sprintf("Error creating Company: %s", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(company)
}

// GetUser retrieves a company by ID
func (uc *CompanyController) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}

	company, err := uc.companyService.GetUser(c.Context(), id)
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
	if err := c.BodyParser(&company); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
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

// CompanyLists lists all companys
func (uc *CompanyController) CompanyLists(c *fiber.Ctx) error {
	page, pageSize, err := helper.ExtractPageAndSize(c)
	search, filters := helper.ExtractSearchAndParams(c)

	companys, total, err := uc.companyService.CompanyLists(c.Context(), page, pageSize, search, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve company list"})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	return views.CompanyListResponse(c, companys, total, totalPages)
}
