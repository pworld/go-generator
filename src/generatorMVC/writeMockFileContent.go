package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"github.com/pworld/loggers"
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
	packageName := filepath.Base(baseDir)

	mocksDir := filepath.Join(baseDir, "tests/mocks")
	mockFileName := fmt.Sprintf("%s_mock.go", lowerStructName)
	mockFilePath := filepath.Join(mocksDir, mockFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(mocksDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create mocks directory: %s\n", err))
		return
	}

	file, err := os.Create(mockFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create mock file: %s\n", err))
		return
	}
	defer file.Close()

	// Data for template execution
	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackageName     string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: strings.ToLower(structName),
		PackageName:     packageName,
	}

	// Execute the template with the struct data
	tmpl, err := template.New("mock").Parse(templateMVC.MockTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating mock template: %s\n", err))
		return
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing mock template: %s\n", err))
		return
	}

	fmt.Println("Mock file generated:", mockFilePath)
}
