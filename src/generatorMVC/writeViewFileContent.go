package generatorMVC

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pworld/go-generator/src/templateMVC"
)

func writeViewFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)
	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath))) // Move two directories up
	viewsDir := filepath.Join(baseDir, "views")
	viewFileName := fmt.Sprintf("%s_view.go", lowerStructName)
	viewFilePath := filepath.Join(viewsDir, viewFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(viewsDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(viewFilePath)
	if err != nil {
		fmt.Printf("Failed to create service file: %s\n", err)
		return
	}
	defer file.Close()

	// Parse and execute the template
	tmpl, err := template.New("view").Parse(templateMVC.ViewsTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	// Prepare template data
	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackagePath     string
		Fields          []StructField
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: strings.ToLower(structName),
		PackagePath:     lowerStructName,
		Fields:          fields,
	}

	// Execute the template with the data
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("View file generated:", viewFilePath)
}
