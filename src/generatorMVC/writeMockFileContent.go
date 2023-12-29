package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Generates the mock file based on the struct name
func writeMockFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)

	// Correctly calculate the base directory from filePath
	entityDir := filepath.Dir(filePath)
	modelsDir := filepath.Dir(entityDir)
	baseDir := filepath.Dir(modelsDir)

	testsDir := filepath.Join(baseDir, "tests/mocks")
	testsFileName := fmt.Sprintf("%s_service_test.go", lowerStructName)
	testsFilePath := filepath.Join(testsDir, testsFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(testsDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(testsFilePath)
	if err != nil {
		fmt.Printf("Failed to create repository file: %s\n", err)
		return
	}
	defer file.Close()
	// Execute the template with the struct data
	tmpl, err := template.New("mock").Parse(templateMVC.MockTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
	}{
		ModuleName:      structName,
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Mock file generated:", testsFilePath)
}
