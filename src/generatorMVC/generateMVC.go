package generatorMVC

import (
	"bufio"
	"fmt"
	"github.com/pworld/go-generator/src/templateMVC"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

// Define the data structure for passing data to the template
type ControllerTemplateData struct {
	ModuleName      string
	StructName      string
	LowerStructName string
}

// StructField represents a field in a struct
type StructField struct {
	Name string
	Type string
}

// Main function to generate the MVC structure
func GenerateMVC(filePath string) {
	structName, fields, err := parseFileForStruct(filePath)
	if err != nil {
		fmt.Println("Failed to get  the structName: %s\n", err)
		return
	}
	fmt.Println(fields)
	// Generate the controller file
	writeControllerFileContent(filePath, structName, fields)
	// Generate the Service file
	writeServiceFileContent(filePath, structName, fields)
	//// Generate the Persistence file
	//writePersistenceFileContent(filePath,structName)
	//// Generate the View file
	//writeViewFileContent(filePath,structName)
	//// Generate the Test file
	//writeTestFileContent(filePath,structName)
	//// Generate the Mock file
	//writeMockFileContent(filePath,structName)
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
						fields = append(fields, StructField{Name: fieldName.Name, Type: fieldType})
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

// Generates the view file based on the struct name
func writeViewFileContent(structName string) {
	dirPath := "views"
	lowerStructName := strings.ToLower(structName)
	viewFileName := fmt.Sprintf("%s_view.go", lowerStructName)
	viewFilePath := fmt.Sprintf("views/%s", viewFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(viewFilePath)
	if err != nil {
		fmt.Printf("Failed to create view file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("view").Parse(templateMVC.ViewTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		StructName      string
		LowerStructName string
	}{
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("View file generated:", viewFilePath)
}

// Generates the mock file based on the struct name
func writeMockFileContent(structName string) {
	lowerStructName := strings.ToLower(structName)
	dirPath := "tests/mocks"
	mockFileName := fmt.Sprintf("%s_mock.go", lowerStructName)
	mockFilePath := fmt.Sprintf("mocks/%s", mockFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(mockFilePath)
	if err != nil {
		fmt.Printf("Failed to create mock file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("mock").Parse(templateMVC.MockTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		StructName      string
		LowerStructName string
	}{
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Mock file generated:", mockFilePath)
}

// Generates the test file based on the struct name
func writeTestFileContent(structName string) {
	lowerStructName := strings.ToLower(structName)
	dirPath := "tests/tests"
	testFileName := fmt.Sprintf("%s_test.go", lowerStructName)
	testFilePath := fmt.Sprintf("tests/%s", testFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(testFilePath)
	if err != nil {
		fmt.Printf("Failed to create test file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("test").Parse(templateMVC.TestTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		StructName      string
		LowerStructName string
	}{
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Test file generated:", testFilePath)
}

func writePersistenceFileContent(structName string) {
	// Generates the persistence file based on the struct name
	lowerStructName := strings.ToLower(structName)
	dirPath := "models/persistence"
	persistenceFileName := fmt.Sprintf("%s_persistence.go", lowerStructName)
	persistenceFilePath := fmt.Sprintf("persistence/%s", persistenceFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %s\n", err)
		return
	}

	file, err := os.Create(persistenceFilePath)
	if err != nil {
		fmt.Printf("Failed to create persistence file: %s\n", err)
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("persistence").Parse(templateMVC.PersistenceTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	data := struct {
		StructName      string
		LowerStructName string
	}{
		StructName:      structName,
		LowerStructName: lowerStructName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Persistence file generated:", persistenceFilePath)
}
