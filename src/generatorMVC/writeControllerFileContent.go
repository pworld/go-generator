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

// createControllerDirectory ensures the controller directory exists.
func createControllerDirectory(baseDir string) (string, error) {
	controllerDir := filepath.Join(baseDir, "controllers")
	if err := os.MkdirAll(controllerDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create controller directory: %w", err)
	}
	return controllerDir, nil
}

// executeTemplate executes the given template with provided data.
func executeTemplate(data interface{}, file *os.File) error {
	tmpl, err := template.New("controller").Parse(templateMVC.ControllerTemplate)
	if err != nil {
		return fmt.Errorf("error creating controller template: %w", err)
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing controller template: %w", err)
	}

	return nil
}

// writeControllerFileContent creates a controller file based on the provided information.
func writeControllerFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	// Create controller directory
	controllerDir, err := createControllerDirectory(baseDir)
	if err != nil {
		loggers.Error(err.Error())
		return err
	}

	// Create controller file
	controllerFileName := fmt.Sprintf("%s_controller.go", lowerStructName)
	controllerFilePath := filepath.Join(controllerDir, controllerFileName)
	file, err := os.Create(controllerFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create controller file: %s\n", err))
		return err
	}
	defer file.Close()

	// Data for template execution
	data := struct {
		ModuleName      string
		ModulePath      string
		StructName      string
		LowerStructName string
		Fields          []StructField
	}{
		ModuleName:      moduleName,
		ModulePath:      modulePath,
		StructName:      structName,
		LowerStructName: strings.ToLower(structName),
		Fields:          fields,
	}

	// Execute the template with the struct data
	if err := executeTemplate(data, file); err != nil {
		loggers.Error(err.Error())
		return err
	}

	fmt.Println("Controller file generated:", controllerFilePath)
	return nil
}
