package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"strings"
	"text/template"
)

// Generates the mock file based on the struct name
func writeMockFileContent(structName string) {
	lowerStructName := strings.ToLower(structName)
	dirPath := "tests/mocks"
	mockFileName := fmt.Sprintf("%s_mock.go", lowerStructName)
	mockFilePath := fmt.Sprintf("mocks/%s", mockFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(mockFilePath)
	if err != nil {
		fmt.Printf("Failed to create mock file: %s\n", err)
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
		StructName      string
		LowerStructName string
	}{
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Mock file generated:", mockFilePath)
}
