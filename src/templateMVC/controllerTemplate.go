package templateMVC

const ControllerTemplate = `package controllers

import (
	"github.com/gofiber/fiber/v2"
	"{{.ModuleName}}/internal/{{.PackageName}}/models/entity"
	"{{.ModuleName}}/internal/{{.PackageName}}/models/repository"
	"{{.ModuleName}}/internal/{{.PackageName}}/services"
	"{{.ModuleName}}/internal/{{.PackageName}}/views"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

// {{.StructName}}Controller structure
type {{.StructName}}Controller struct {
	{{.LowerStructName}}Service services.{{.StructName}}Service
}

// New{{.StructName}}Controller creates a new instance of {{.StructName}}Controller
func New{{.StructName}}Controller(db *gorm.DB) *{{.StructName}}Controller {
	{{.LowerStructName}}Repo := repository.New{{.StructName}}Repository(db)
	{{.LowerStructName}}Service := services.New{{.StructName}}Service({{.LowerStructName}}Repo)
	return &{{.StructName}}Controller{
		{{.LowerStructName}}Service: {{.LowerStructName}}Service,
	}
}

// Create{{.StructName}} creates a new {{.LowerStructName}}
func (uc *{{.StructName}}Controller) Create{{.StructName}}(c *fiber.Ctx) error {
	var {{.LowerStructName}} entity.{{.StructName}}
	if err := parseAndValidate(c, &{{.LowerStructName}}); err != nil {
		loggers.Error("Error parsing request body:", err)
		return handleValidationError(c, err)
	}

	{{.LowerStructName}}.CreatedAt = time.Now()

	if _, err := uc.{{.LowerStructName}}Service.Create{{.StructName}}(c.Context(), {{.LowerStructName}}); err != nil {
		loggers.Error("Error creating {{.StructName}}:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON({{.LowerStructName}})
}

// Get{{.StructName}} retrieves a {{.LowerStructName}} by ID
func (uc *{{.StructName}}Controller) Get{{.StructName}}(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid {{.LowerStructName}} ID"})
	}
	
	{{.LowerStructName}}, err := uc.{{.LowerStructName}}Service.Get{{.StructName}}(c.Context(), id)
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
	if err := parseAndValidate(c, &{{.LowerStructName}}); err != nil {
		return handleValidationError(c, err)
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

// {{.StructName}}Lists lists all {{.LowerStructName}}
func (uc *{{.StructName}}Controller) {{.StructName}}Lists(c *fiber.Ctx) error {
	page, pageSize, err := helper.ExtractPageAndSize(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pagination parameters"})
	}
	search, filters := helper.ExtractSearchAndParams(c)
	
	{{.LowerStructName}}, total, err := uc.{{.LowerStructName}}Service.{{.StructName}}Lists(c.Context(), page, pageSize, search, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve {{.LowerStructName}} list"})
	}
	
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	return views.{{.StructName}}ListResponse(c, {{.LowerStructName}}, total, totalPages)
}
`
