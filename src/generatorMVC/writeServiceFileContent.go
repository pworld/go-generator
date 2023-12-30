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

func writeServiceFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)
	baseDir := filepath.Dir(filepath.Dir(filepath.Dir(filePath)))
	packageName := filepath.Base(baseDir)
	servicesDir := filepath.Join(baseDir, "services")
	serviceFileName := fmt.Sprintf("%s_service.go", lowerStructName)
	serviceFilePath := filepath.Join(servicesDir, serviceFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(servicesDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create services directory: %s\n", err))
		return
	}

	file, err := os.Create(serviceFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create service file: %s\n", err))
		return
	}
	defer file.Close()

	tmpl, err := template.New("service").Parse(templateMVC.ServiceTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating service template: %s\n", err))
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackageName     string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackageName:     packageName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing service template: %s\n", err))
		return
	}

	fmt.Println("Service file generated:", serviceFilePath)
}
