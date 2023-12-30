package views

import (
	"github.com/pworld/go-generator/examples/internal/user/models/entity"
)

// FormatCompanyDetails formats the company data for the response
func FormatCompanyDetails(company entity.Company) map[string]interface{} {
	return map[string]interface{}{
		// Fields to format here...
		"id":        company.ID,
		"fullname":  company.Fullname,
		"email":     company.Email,
		"phone":     company.Phone,
		"username":  company.Username,
		"password":  company.Password,
		"createdat": company.CreatedAt,
		"updatedat": company.UpdatedAt,
		"deletedat": company.DeletedAt,
	}
}

// CompanyResponse formats a successful company-related response
func CompanyResponse(c *fiber.Ctx, company entity.Company) error {
	formattedCompany := FormatCompanyDetails(company)
	return c.JSON(fiber.Map{
		"success": true,
		"company": formattedCompany,
	})
}

// CompanyListResponse formats a response for a list of companys
func CompanyListResponse(c *fiber.Ctx, companys []entity.Company, total, totalPages int) error {
	formattedCompanys := make([]map[string]interface{}, 0)
	for _, company := range companys {
		formattedCompanys = append(formattedCompanys, FormatCompanyDetails(company))
	}
	return c.JSON(fiber.Map{
		"success":    true,
		"items":      formattedCompanys,
		"total":      total,
		"totalPages": totalPages,
	})
}

// CompanyErrorResponse formats an error response specific to company operations
func CompanyErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   message,
		"context": "company operation",
	})
}
