package templateMVC

const CompanyViewsTemplate = `package views

import (
	"github.com/gofiber/fiber/v2"
	"{{.ModulePath}}/models/entity"
)

// Format{{.StructName}}Details formats the {{.LowerStructName}} data for the response
func Format{{.StructName}}Details({{.LowerStructName}} entity.{{.StructName}}) map[string]interface{} {
    return map[string]interface{}{
        // Fields to format here...
        {{- range .Fields }}
        "{{ .JsonTag }}": {{ $.LowerStructName }}.{{ .Name }},
        {{- end }}
    }
}

// {{.StructName}}Response formats a successful {{.LowerStructName}}-related response
func {{.StructName}}Response(c *fiber.Ctx, {{.LowerStructName}} entity.{{.StructName}}) error {
	formatted{{.StructName}} := Format{{.StructName}}Details({{.LowerStructName}})
	return c.JSON(fiber.Map{
		"success": true,
		"{{.LowerStructName}}": formatted{{.StructName}},
	})
}

// {{.StructName}}ListResponse formats a response for a list of {{.LowerStructName}}
func {{.StructName}}ListResponse(c *fiber.Ctx, {{.LowerStructName}} []entity.{{.StructName}}, total, totalPages int) error {
    formatted{{.StructName}} := make([]map[string]interface{}, 0)
    for _, {{.LowerStructName}} := range {{.LowerStructName}} {
        formatted{{.StructName}} = append(formatted{{.StructName}}, Format{{.StructName}}Details({{.LowerStructName}}))
    }
    return c.JSON(fiber.Map{
        "success":    true,
        "items":      formatted{{.StructName}},
        "total":      total,
        "totalPages": totalPages,
    })
}

// {{.StructName}}ErrorResponse formats an error response specific to {{.LowerStructName}} operations
func {{.StructName}}ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
    return c.Status(statusCode).JSON(fiber.Map{
        "success": false,
        "error":   message,
        "context": "{{.LowerStructName}} operation",
    })
}
`
