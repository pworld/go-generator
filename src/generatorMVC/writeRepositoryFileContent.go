package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func writeRepositoryFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)

	// Correctly calculate the base directory from filePath
	entityDir := filepath.Dir(filePath)
	modelsDir := filepath.Dir(entityDir)
	baseDir := filepath.Dir(modelsDir)

	repositoryDir := filepath.Join(baseDir, "models/repository")
	repositoryFileName := fmt.Sprintf("%s_repository.go", lowerStructName)
	repositoryFilePath := filepath.Join(repositoryDir, repositoryFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(repositoryDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(repositoryFilePath)
	if err != nil {
		fmt.Printf("Failed to create repository file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("repository").Parse(templateMVC.RepositoryTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	fieldNames, placeholders := buildQueryParts(structName, fields)
	query := "INSERT INTO " + structName + " (" + fieldNames + ") VALUES (" + placeholders + ") RETURNING id;"

	// Generate argument list as a string
	var argumentParts []string
	for _, field := range fields {
		argumentParts = append(argumentParts, lowerStructName+"."+field.Name)
	}
	argumentList := strings.Join(argumentParts, ", ")
	scanFields := generateScanFields(lowerStructName, fields)

	// generateUpdateQuery
	updateFields := generateUpdateFields(fields)
	scanFieldsUpdate := generateUpdateArguments(lowerStructName, fields)
	scanFieldsListsUpdate := generateListArguments(lowerStructName, fields)

	// Before executing the template
	fmt.Printf("Debug - scanFieldsUpdate: %s\n", scanFieldsUpdate)

	// Prepare data for the template
	data := struct {
		ModuleName            string
		StructName            string
		LowerStructName       string
		TableName             string
		Query                 string
		UpdateFields          string
		ArgumentList          string
		ScanFields            string
		ScanFieldsUpdate      string
		ScanFieldsListsUpdate string
	}{
		ModuleName:            moduleName,
		StructName:            structName,
		LowerStructName:       strings.ToLower(structName),
		TableName:             structName,
		Query:                 query,
		UpdateFields:          updateFields,
		ArgumentList:          argumentList,
		ScanFields:            scanFields,
		ScanFieldsUpdate:      scanFieldsUpdate,
		ScanFieldsListsUpdate: scanFieldsListsUpdate,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Repository file generated:", repositoryFilePath)
}

// buildQueryParts creates lists of field names and placeholders for an SQL query
func buildQueryParts(structName string, fields []StructField) (string, string) {
	var fieldNames, placeholders []string

	for _, field := range fields {
		fieldNames = append(fieldNames, field.Name)
		placeholders = append(placeholders, "?") // Using "?" as a placeholder
	}

	return strings.Join(fieldNames, ", "), strings.Join(placeholders, ", ")
}

// getStructFieldValues returns the values of all fields in a struct as a slice of interfaces.
func generateScanFields(structName string, fields []StructField) string {
	var scanFields []string
	for _, field := range fields {
		scanFields = append(scanFields, "&"+structName+"."+field.Name)
	}
	return strings.Join(scanFields, ", ")
}

func generateUpdateFields(fields []StructField) string {
	var setParts []string
	for _, field := range fields {
		if field.Name != "ID" { // Skip "ID" or any other fields you don't want to update
			setParts = append(setParts, fmt.Sprintf("%s = ?", field.Name))
		}
	}
	return strings.Join(setParts, ", ")
}

func generateUpdateArguments(structName string, fields []StructField) string {
	var argumentParts []string
	for _, field := range fields {
		if field.Name != "ID" { // Include all fields except ID for update
			argumentParts = append(argumentParts, structName+"."+field.Name)
		}
	}
	argumentParts = append(argumentParts, structName+".ID") // Add ID at the end for WHERE clause
	return strings.Join(argumentParts, ", ")
}

func generateListArguments(structName string, fields []StructField) string {
	var argumentParts []string
	for _, field := range fields {
		if field.Name != "ID" { // Include all fields except ID for update
			argumentParts = append(argumentParts, "&"+structName+"."+field.Name)
		}
	}
	argumentParts = append(argumentParts, structName+".ID") // Add ID at the end for WHERE clause
	return strings.Join(argumentParts, ", ")
}
