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

// Generates the test file based on the struct name
func writeTestFileContent(filePath, structName, moduleName string, fields []StructField) {
	lowerStructName := strings.ToLower(structName)

	// Correctly calculate the base directory from filePath
	entityDir := filepath.Dir(filePath)
	modelsDir := filepath.Dir(entityDir)
	baseDir := filepath.Dir(modelsDir)
	packageName := filepath.Base(baseDir)

	testsDir := filepath.Join(baseDir, "tests/service_tests")
	testsFileName := fmt.Sprintf("%s_service_test.go", lowerStructName)
	testsFilePath := filepath.Join(testsDir, testsFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(testsDir, os.ModePerm); err != nil {
		loggers.Error(fmt.Sprintf("Failed to create tests directory: %s\n", err))
		return
	}

	file, err := os.Create(testsFilePath)
	if err != nil {
		loggers.Error(fmt.Sprintf("Failed to create test file: %s\n", err))
		return
	}
	defer file.Close()

	// Execute the template with the struct data
	tmpl, err := template.New("test").Parse(templateMVC.TestTemplate)
	if err != nil {
		loggers.Error(fmt.Sprintf("Error creating test template: %s\n", err))
		return
	}

	// Prepare template data
	data := struct {
		ModuleName      string
		StructName      string
		LowerStructName string
		PackagePath     string
		Fields          []StructField
		Methods         []Method
		PackageName     string
	}{
		ModuleName:      moduleName,
		StructName:      structName,
		LowerStructName: lowerStructName,
		PackagePath:     lowerStructName,
		Fields:          fields,
		Methods:         GenerateCRUDTestMethods(structName),
		PackageName:     packageName,
	}

	if err := tmpl.Execute(file, data); err != nil {
		loggers.Error(fmt.Sprintf("Error executing test template: %s\n", err))
		return
	}

	fmt.Println("Test file generated:", testsFilePath)
}

// GenerateCRUDMethods generates a slice of Method structs for standard CRUD operations
func GenerateCRUDTestMethods(structName string) []Method {
	lowerStructName := strings.ToLower(structName)

	return []Method{
		{
			Name:      "Create" + structName,
			SetupMock: fmt.Sprintf(`mockRepo.On("Create%s", mock.Anything, mock.AnythingOfType("entity.%s")).Return(1, nil)`, structName, structName),
			TestImplementation: fmt.Sprintf(`
    mock%s := entity.%s{
        // Fill The test case
    }
    %sID, err := svc.Create%s(context.Background(), mock%s)`, structName, structName, lowerStructName, structName, structName),
			Assertions: fmt.Sprintf(`
    assert.NoError(t, err)
    assert.Equal(t, 1, %sID)`, lowerStructName),
		},
		{
			Name:      "Delete" + structName,
			SetupMock: fmt.Sprintf(`mockRepo.On("Delete%s", mock.Anything, test%sID).Return(nil)`, structName, structName),
			TestImplementation: fmt.Sprintf(`
	test%sID := 1
	err := svc.Delete%s(context.Background(), test%sID)`, structName, structName, structName),
			Assertions: `
	assert.NoError(t, err)`,
		},
		{
			Name:      "Get" + structName,
			SetupMock: fmt.Sprintf(`mockRepo.On("Get%sByID", mock.Anything, test%sID).Return(mock%s, nil)`, structName, structName, structName),
			TestImplementation: fmt.Sprintf(`
	test%sID := 1
	mock%s := entity.%s{
	// Fill The test case
	}
	%s, err := svc.Get%s(context.Background(), test%sID)`, structName, structName, structName, lowerStructName, structName, structName),
			Assertions: fmt.Sprintf(`
	assert.NoError(t, err)
	assert.Equal(t, mock%s, %s)`, structName, lowerStructName),
		},
		{
			Name: "List" + structName + "s",
			SetupMock: fmt.Sprintf(`
	page, pageSize, totalData := 1, 10, 2
	mockRepo.On("List%s", mock.Anything, page, pageSize, "", mock.Anything).Return(mock%s, nil)
    mockRepo.On("Total%s", mock.Anything, "", mock.Anything).Return(totalData, nil)`, structName, structName, structName),
			TestImplementation: fmt.Sprintf(`
    total%s := 2 // Total number of %s available
    mock%s := []entity.%s{
        // Fill The test case
    }
    %s, total, err := svc.List%s(context.Background(), page, pageSize, "", nil)`, structName, structName, structName, structName, lowerStructName, structName),
			Assertions: fmt.Sprintf(`
    assert.NoError(t, err)
    assert.Equal(t, total%s, total)
    assert.Len(t, %ss, len(mock%s))
    assert.Equal(t, mock%s, %s)`, structName, lowerStructName, structName, structName, lowerStructName),
		},
	}
}
