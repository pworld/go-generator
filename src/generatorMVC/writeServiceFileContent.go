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

func writeServiceFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	servicesDir := filepath.Join(baseDir, "services")
	serviceFileName := fmt.Sprintf("%s_service.go", lowerStructName)
	serviceFilePath := filepath.Join(servicesDir, serviceFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(servicesDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create services directory: %s\n", err))
		return err
	}

	file, err := os.Create(serviceFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create service file: %s\n", err))
		return err
	}
	defer file.Close()

	tmpl, err := template.New("service").Parse(templateMVC.ServiceTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating service template: %s\n", err))
		return err
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackageName     string
		ModulePath      string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		ModulePath:      modulePath,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing service template: %s\n", err))
		return err
	}

	fmt.Println("Service file generated:", serviceFilePath)
	return nil
}
