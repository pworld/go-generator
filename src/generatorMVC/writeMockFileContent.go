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
func writeMockFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	mocksDir := filepath.Join(baseDir, "tests/mocks")
	mockFileName := fmt.Sprintf("%s_mock.go", lowerStructName)
	mockFilePath := filepath.Join(mocksDir, mockFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(mocksDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create mocks directory: %s\n", err))
		return err
	}

	file, err := os.Create(mockFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create mock file: %s\n", err))
		return err
	}
	defer file.Close()

	// Data for template execution
	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		ModulePath      string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: strings.ToLower(structName),
		ModulePath:      modulePath,
	}

	// Execute the template with the struct data
	tmpl, err := template.New("mock").Parse(templateMVC.MockTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating mock template: %s\n", err))
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing mock template: %s\n", err))
		return err
	}

	fmt.Println("Mock file generated:", mockFilePath)
	return nil
}
