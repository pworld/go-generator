package generatorMVC

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// Define the data structure for passing data to the template
type ControllerTemplateData struct {
	ModuleName      string
	StructName      string
	LowerStructName string
}

// StructField represents a field in a struct
type StructField struct {
	Name    string
	Type    string
	JsonTag string
}

type Method struct {
	Name               string
	SetupMock          string
	TestImplementation string
	Assertions         string
}

// Main function to generate the MVC structure
func GenerateMVC(filePath string) {
	structName, fields, err := parseFileForStruct(filePath)
	if err != nil {
		fmt.Println("Failed to get  the structName: %s\n", err)
		return
	}
	moduleName, err := getModuleName("go.mod")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(fields)
	// Generate the controller file
	writeControllerFileContent(filePath, structName, moduleName, fields)
	// Generate the Service file
	writeServiceFileContent(filePath, structName, moduleName, fields)
	// Generate the Repository file
	writeRepositoryFileContent(filePath, structName, moduleName, fields)
	// Generate the View file
	writeViewFileContent(filePath, structName, moduleName, fields)
	//// Generate the Test file
	writeTestFileContent(filePath, structName, moduleName, fields)
	// Generate the Mock file
	writeMockFileContent(filePath, structName, moduleName, fields)
}

// Parses the provided Go file and extracts the struct name
// parseFileForStruct extracts the struct name and its fields from a Go file
func parseFileForStruct(filePath string) (string, []StructField, error) {
	// Read the file
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse the file: %s", err)
	}

	var structName string
	var fields []StructField
	ast.Inspect(node, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := ts.Type.(*ast.StructType); ok {
				structName = ts.Name.Name
				for _, field := range structType.Fields.List {
					fieldType := fmt.Sprintf("%v", field.Type)
					for _, fieldName := range field.Names {
						fields = append(fields, StructField{Name: fieldName.Name, Type: fieldType, JsonTag: strings.ToLower(fieldName.Name)})
					}
				}
				return false
			}
		}
		return true
	})

	if structName == "" {
		return "", nil, fmt.Errorf("no struct found in the file")
	}

	return structName, fields, nil
}

func getModuleName(modFilePath string) (string, error) {
	file, err := os.Open(modFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			// Extract module name
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module directive not found in %s", modFilePath)
}
