package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type ModelField struct {
	Name string
	Type string
}

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
		return fmt.Sprintf("func validate%s%s(value string) bool {\n    return len(value) > 0\n}", modelName, field.Name)
	case "int":
		return fmt.Sprintf("func validate%s%s(value int) bool {\n    return value >= 0\n}", modelName, field.Name)
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <path-to-model-file>")
		os.Exit(1)
	}

	//modelPath := os.Args[1]
	modelPath := "model/User.go"
	modelName := extractModelName(modelPath) // Implement this function to extract model name from file path

	fields, err := parseModelFile(modelPath)
	if err != nil {
		fmt.Printf("Error parsing model file: %v\n", err)
		os.Exit(1)
	}

	validations := generateValidationFunctions(modelName, fields)

	directory := "./model_validations"
	filename := strings.ToLower(modelName) + "_validations.go"

	if err := writeToFile(directory, filename, validations); err != nil {
		fmt.Printf("Error writing validation functions to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Validation functions for model '%s' have been written to %s/%s\n", modelName, directory, filename)
}
