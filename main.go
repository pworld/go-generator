package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type ModelField struct {
	Name string
	Type string
}

// ------------START Validations
// parseModelFile reads a Go source file and extracts model fields.
func parseModelFile(filePath string) ([]ModelField, error) {
	fset := token.NewFileSet()

	// Read the file contents
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the source string
	node, err := parser.ParseFile(fset, "", string(content), parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var fields []ModelField
	// Inspect the AST and find Struct Types
	ast.Inspect(node, func(n ast.Node) bool {
		// Find StructType nodes
		t, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range t.Fields.List {
			var fieldType string
			if len(field.Names) > 0 {
				fieldType = getTypeAsString(field.Type)
				for _, name := range field.Names {
					fields = append(fields, ModelField{Name: name.Name, Type: fieldType})
				}
			}
		}
		return false
	})

	return fields, nil
}

// getTypeAsString converts the type of a field into a string.
func getTypeAsString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getTypeAsString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + getTypeAsString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeAsString(t.Elt)
		// Add more cases as necessary to handle more complex types
	}

	return ""
}

// generateValidationFunctions creates basic validation functions for the given fields.
func generateValidationFunctions(modelName string, fields []ModelField) string {
	var validations []string

	// Start with the package declaration
	validations = append(validations, "package model")

	for _, field := range fields {
		validation := generateValidationForField(modelName, field)
		if validation != "" {
			validations = append(validations, validation)
		}
	}

	return strings.Join(validations, "\n\n")
}

// generateValidationForField creates a validation function for a single field.
func generateValidationForField(modelName string, field ModelField) string {
	switch field.Type {
	case "string":
		return fmt.Sprintf("func Validate%s%s(value string) bool {\n    return len(value) > 0\n}", modelName, field.Name)
	case "int":
		return fmt.Sprintf("func Validate%s%s(value int) bool {\n    return value >= 0\n}", modelName, field.Name)
	// Add more cases for different types
	default:
		return ""
	}
}

// writeToFile writes the given content to a file, creating the file if it does not exist.
func writeToFile(directory, filename, content string) error {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(directory, filename)
	return os.WriteFile(filePath, []byte(content), 0644)
}

// Extracts the model name from the file path (assuming the file name is the model name)
func extractModelName(filePath string) string {
	base := filepath.Base(filePath)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// ------------END Validations

// ------------ Controller
func toTitleCase(str string) string {
	if str == "" {
		return ""
	}
	r := []rune(str)
	return string(unicode.ToUpper(r[0])) + strings.ToLower(string(r[1:]))
}

func generateController(modelName string) string {
	modelNameCapitalized := toTitleCase(modelName) // Replacing strings.Title with our function
	// ... rest of your generateController function ...

	return fmt.Sprintf(
		`package controllers

import (
    "github.com/gofiber/fiber/v2"
    "model/user" 
)

type UserController struct {
    // DB CON
}

// NewUserController creates a new UserController instance.
func NewUserController() *UserController {
    return &UserController{
        // Initialize dependencies here
    }
}

// CreateUser handles POST requests to create a new user.
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Logic to save user
    // ...

    return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser handles GET requests to retrieve a user by ID.
func (uc *UserController) GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    
    // Logic to retrieve the user from the database using id

    user := &models.User{ID: 1, Fullname: "John Doe", Email: "john@example.com"}

    return c.JSON(user)
}

// UpdateUser handles PUT requests to update a user.
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")

    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Logic update user in the database 


    return c.JSON(user)
}

// DeleteUser handles DELETE requests to delete a user.
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")

    // Logic delete the user from the database

    return c.SendString("User successfully deleted")
}
`,
		modelNameCapitalized, modelName, modelNameCapitalized, modelNameCapitalized, modelNameCapitalized,
		modelNameCapitalized, modelNameCapitalized, modelName, modelName)
}

// ------------ END Controller

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: go run main.go <path-to-model-file>")
	// 	os.Exit(1)
	// }

	//modelPath := os.Args[1]
	path := "model/user/"
	NameFile := "user.go"

	modelPath := filepath.Join(path, NameFile)

	modelName := extractModelName(modelPath) // Implement this function to extract model name from file path

	fields, err := parseModelFile(modelPath)
	if err != nil {
		fmt.Printf("Error parsing model file: %v\n", err)
		os.Exit(1)
	}

	// ------------ Validations
	validations := generateValidationFunctions(modelName, fields)

	directory := "./model/user"
	filename := strings.ToLower(modelName) + "_validations.go"

	if err := writeToFile(directory, filename, validations); err != nil {
		fmt.Printf("Error writing validation functions to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Validation functions for model '%s' have been written to %s/%s\n", modelName, directory, filename)
	// ------------END Validations

	// ------------ Controller
	// Generate and write the controller file
	controllerContent := generateController(modelName)
	controllerDir := "./controllers"
	controllerFilename := strings.ToLower(modelName) + "_controller.go"

	if err := writeToFile(controllerDir, controllerFilename, controllerContent); err != nil {
		fmt.Printf("Error writing controller file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Controller for model '%s' has been written to %s/%s\n", modelName, controllerDir, controllerFilename)
	// ------------ END Controller
}
