package templateMVC

const ControllerTemplate = `package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"{{.ModuleName}}/helper"
	"{{.ModuleName}}/internal/{{.LowerStructName}}/models/entity"
	"{{.ModuleName}}/internal/{{.PackagePath}}/models/repository"
	"{{.ModuleName}}/internal/{{.PackagePath}}/services"
	"{{.ModuleName}}/internal/{{.PackagePath}}/views"
	"math"
	"strconv"
	"time"
)

// {{.StructName}}Controller structure
type {{.StructName}}Controller struct {
	{{.LowerStructName}}Service services.{{.StructName}}Service
}

func New{{.StructName}}Controller(db *sql.DB) *{{.StructName}}Controller {
	{{.LowerStructName}}Repo := repository.New{{.StructName}}Repository(db)
	{{.LowerStructName}}Service := services.New{{.StructName}}Service({{.LowerStructName}}Repo)
	return &{{.StructName}}Controller{
		{{.LowerStructName}}Service: {{.LowerStructName}}Service,
	}
}

// CreateUser creates a new {{.LowerStructName}}
func (uc *{{.StructName}}Controller) CreateUser(c *fiber.Ctx) error {
	var {{.LowerStructName}} entity.{{.StructName}}
	if err := c.BodyParser(&{{.LowerStructName}}); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	{{.LowerStructName}}.CreatedAt = time.Now()

	if _, err := uc.{{.LowerStructName}}Service.Create{{.StructName}}(c.Context(), {{.LowerStructName}}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON({{.LowerStructName}})
}

// GetUser retrieves a {{.LowerStructName}} by ID
func (uc *{{.StructName}}Controller) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid {{.LowerStructName}} ID"})
	}
	
	{{.LowerStructName}}, err := uc.{{.LowerStructName}}Service.GetUser(c.Context(), id)
	if err != nil {
		return views.{{.StructName}}ErrorResponse(c, fiber.StatusInternalServerError, "{{.StructName}} not found")
	}
	return views.{{.StructName}}Response(c, {{.LowerStructName}})
}

// Update{{.StructName}} updates a {{.LowerStructName}}
func (uc *{{.StructName}}Controller) Update{{.StructName}}(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid {{.LowerStructName}} ID"})
	}
	
	var {{.LowerStructName}} entity.{{.StructName}}
	if err := c.BodyParser(&{{.LowerStructName}}); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	{{.LowerStructName}}.ID = id
	{{.LowerStructName}}.UpdatedAt = time.Now()
	
	if err := uc.{{.LowerStructName}}Service.Update{{.StructName}}(c.Context(), {{.LowerStructName}}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON({{.LowerStructName}})
}

// Delete{{.StructName}} deletes a {{.LowerStructName}} by ID
func (uc *{{.StructName}}Controller) Delete{{.StructName}}(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid {{.LowerStructName}} ID"})
	}
	
	if err := uc.{{.LowerStructName}}Service.Delete{{.StructName}}(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.SendStatus(fiber.StatusNoContent)
}

// {{.StructName}}Lists lists all {{.LowerStructName}}s
func (uc *{{.StructName}}Controller) {{.StructName}}Lists(c *fiber.Ctx) error {
page, pageSize, err := helper.ExtractPageAndSize(c)
	search, filters := helper.ExtractSearchAndParams(c)
	
	{{.LowerStructName}}s, total, err := uc.{{.LowerStructName}}Service.{{.StructName}}Lists(c.Context(), page, pageSize, search, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve {{.LowerStructName}} list"})
	}
	
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	return views.{{.StructName}}ListResponse(c, {{.LowerStructName}}s, total, totalPages)
}
`
