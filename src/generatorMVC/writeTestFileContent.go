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

func writeTestFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	// Directory and file for the tests
	testsDir := filepath.Join(baseDir, "tests", lowerStructName+"_tests")
	testsFileName := fmt.Sprintf("%s_test.go", lowerStructName)
	testsFilePath := filepath.Join(testsDir, testsFileName)

	// Ensure the test directory exists
	if err := os.MkdirAll(testsDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create tests directory: %s\n", err))
		return err
	}

	file, err := os.Create(testsFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create test file: %s\n", err))
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			loggers.Error(fmt.Sprintf("Failed to close test file: %s\n", cerr))
		}
	}()

	// Data for template execution
	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		ModulePath      string
		Fields          []StructField
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		ModulePath:      modulePath,
		Fields:          fields,
	}

	// Parse and execute the template
	tmpl, err := template.New("test").Parse(templateMVC.TestTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating test template: %s\n", err))
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing test template: %s\n", err))
		return err
	}

	fmt.Println("Test file generated:", testsFilePath)
	return nil
}
