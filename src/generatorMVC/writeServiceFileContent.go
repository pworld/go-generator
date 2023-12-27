package generatorMVC

import (
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func writeServiceFileContent(filePath, structName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)
	baseDir := filepath.Dir(filepath.Dir(filePath)) // Move two directories up
	servicesDir := filepath.Join(baseDir, "services")
	serviceFileName := fmt.Sprintf("%s_service.go", lowerStructName)
	serviceFilePath := filepath.Join(servicesDir, serviceFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(servicesDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(serviceFilePath)
	if err != nil {
		fmt.Printf("Failed to create service file: %s\n", err)
		return
	}
	defer file.Close()

	// Assuming getModuleName extracts the module name from go.mod
	moduleName, err := getModuleName("go.mod")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tmpl, err := template.New("service").Parse(templateMVC.ServiceTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Service file generated:", serviceFilePath)
}
