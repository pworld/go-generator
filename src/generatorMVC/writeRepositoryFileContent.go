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

func writeRepositoryFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)
	baseDir := filepath.Dir(filepath.Dir(filePath)) // This should get you to "internal/company/models"

	// Construct the repository directory path
	repositoryDir := filepath.Join(baseDir, "repository") // This should result in "internal/company/models/repository"
	repositoryFileName := fmt.Sprintf("%s_repository.go", lowerStructName)
	repositoryFilePath := filepath.Join(repositoryDir, repositoryFileName)

	if err := os.MkdirAll(repositoryDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create repository directory: %s\n", err))
		return
	}

	file, err := os.Create(repositoryFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create repository file: %s\n", err))
		return
	}
	defer file.Close()

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackageName     string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackageName:     filepath.Base(filepath.Dir(filepath.Dir(repositoryDir))),
	}

	tmpl, err := template.New("repository").Parse(templateMVC.RepositoryTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating repository template: %s\n", err))
		return
	}
	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing repository template: %s\n", err))
		return
	}

	fmt.Println("Repository file generated:", repositoryFilePath)
}
