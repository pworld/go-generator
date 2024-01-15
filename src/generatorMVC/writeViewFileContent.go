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

func writeViewFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	viewsDir := filepath.Join(baseDir, "views")
	viewFileName := fmt.Sprintf("%s_view.go", lowerStructName)
	viewFilePath := filepath.Join(viewsDir, viewFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(viewsDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create views directory: %s\n", err))
		return err
	}

	file, err := os.Create(viewFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create view file: %s\n", err))
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			loggers.Error(fmt.Sprintf("Failed to close view file: %s\n", cerr))
		}
	}()

	// Use the new CompanyViewsTemplate
	tmpl, err := template.New("view").Parse(templateMVC.CompanyViewsTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating view template: %s\n", err))
		return err
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackagePath     string
		Fields          []StructField
		ModulePath      string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackagePath:     lowerStructName,
		Fields:          fields,
		ModulePath:      modulePath,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing view template: %s\n", err))
		return err
	}

	fmt.Println("View file generated:", viewFilePath)
	return nil
}
