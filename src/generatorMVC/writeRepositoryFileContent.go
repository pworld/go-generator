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

func writeRepositoryFileContent(filePath, structName, modulePath, baseDir, moduleName string, fields []StructField) error {
	lowerStructName := strings.ToLower(structName)

	// Construct the repository directory path
	repositoryDir := filepath.Join(baseDir, "models", "repository")
	repositoryFileName := fmt.Sprintf("%s_repository.go", lowerStructName)
	repositoryFilePath := filepath.Join(repositoryDir, repositoryFileName)

	if err := os.MkdirAll(repositoryDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create repository directory: %s\n", err))
		return err
	}

	file, err := os.Create(repositoryFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create repository file: %s\n", err))
		return err
	}
	defer file.Close()

	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		ModulePath      string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		ModulePath:      modulePath,
	}

	tmpl, err := template.New("repository").Parse(templateMVC.RepositoryTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating repository template: %s\n", err))
		return err
	}
	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing repository template: %s\n", err))
		return err
	}

	fmt.Println("Repository file generated:", repositoryFilePath)
	return nil
}
