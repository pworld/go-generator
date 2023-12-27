package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func writeControllerFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)

	// Calculate the directory path for the controller
	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath))) // Move two directories up
	controllerDir := filepath.Join(baseDir, "controllers")
	controllerFileName := fmt.Sprintf("%s_controller.go", lowerStructName)
	controllerFilePath := filepath.Join(controllerDir, controllerFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(controllerDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(controllerFilePath)
	if err != nil {
		fmt.Printf("Failed to create controller file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("controller").Parse(templateMVC.ControllerTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		Fields          []StructField
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		Fields:          fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Controller file generated:", controllerFilePath)
}
