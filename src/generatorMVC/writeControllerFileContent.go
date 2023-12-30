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

func writeControllerFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)

	// Calculate the directory path for the controller
	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath))) // Move two directories up
	fmt.Print(baseDir)
	packageName := filepath.Base(baseDir)
	controllerDir := filepath.Join(baseDir, "controllers")
	controllerFileName := fmt.Sprintf("%s_controller.go", lowerStructName)
	controllerFilePath := filepath.Join(controllerDir, controllerFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(controllerDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create controller directory: %s\n", err))
		return
	}

	file, err := os.Create(controllerFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create controller file: %s\n", err))
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("controller").Parse(templateMVC.ControllerTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating controller template: %s\n", err))
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackageName     string
		Fields          []StructField
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackageName:     packageName,
		Fields:          fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing controller template: %s\n", err))
		return
	}

	fmt.Println("Controller file generated:", controllerFilePath)
}
