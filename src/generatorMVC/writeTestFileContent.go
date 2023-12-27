package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"strings"
	"text/template"
)

// Generates the test file based on the struct name
func writeTestFileContent(structName string) {
	lowerStructName := strings.ToLower(structName)
	dirPath := "tests/tests"
	testFileName := fmt.Sprintf("%s_test.go", lowerStructName)
	testFilePath := fmt.Sprintf("tests/%s", testFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(testFilePath)
	if err != nil {
		fmt.Printf("Failed to create test file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("test").Parse(templateMVC.TestTemplate)
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

	fmt.Println("Test file generated:", testFilePath)
}
