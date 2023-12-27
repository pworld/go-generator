package templateMVC

const TestTemplate = `package service_test

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"{{.ModuleName}}/internal/{{.PackagePath}}/models/entity"
	"{{.ModuleName}}/internal/{{.PackagePath}}/services"
	"{{.ModuleName}}/tests/{{.PackagePath}}/mocks"
)

{{- range .TestCases }}
func Test{{$.StructName}}Service_{{.FunctionName}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{$.StructName}}Repository)
	svc := services.New{{$.StructName}}Service(mockRepo)

	// Setup test case specific mock data and expectations
	{{.SetupMock}}

	// Call the service method
	{{.ServiceCall}}

	// Assertions and verifications
	{{.Assertions}}
	mockRepo.AssertExpectations(t)
}
{{- end }}

`
