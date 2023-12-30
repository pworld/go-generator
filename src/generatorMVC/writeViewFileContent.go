package generatorMVC

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pworld/go-generator/src/templateMVC"
	"github.com/pworld/loggers"
)

func writeViewFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)
	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath))) // Move two directories up
	viewsDir := filepath.Join(baseDir, "views")
	viewFileName := fmt.Sprintf("%s_view.go", lowerStructName)
	viewFilePath := filepath.Join(viewsDir, viewFileName)
	packageName := filepath.Base(baseDir)

	// Ensure the directory exists
	if err := os.MkdirAll(viewsDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create views directory: %s\n", err))
		return
	}

	file, err := os.Create(viewFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create view file: %s\n", err))
		return
	}
	defer file.Close()

	tmpl, err := template.New("view").Parse(templateMVC.ViewsTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating view template: %s\n", err))
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackagePath     string
		Fields          []StructField
		PackageName     string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackagePath:     lowerStructName,
		Fields:          fields,
		PackageName:     packageName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing view template: %s\n", err))
		return
	}

	fmt.Println("View file generated:", viewFilePath)
}
