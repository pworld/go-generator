package templateMVC

const TestTemplate = `package {{.LowerStructName}}_test

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"{{.ModuleName}}/internal/{{.PackageName}}/models/entity"
	"{{.ModuleName}}/internal/{{.PackageName}}/services"
	"{{.ModuleName}}/tests/{{.PackageName}}/mocks"
)

{{range .Methods}}
func Test{{$.StructName}}Service_{{.Name}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{$.StructName}}Repository)
	svc := services.New{{$.StructName}}Service(mockRepo)
	{{.SetupMock -}}
	{{.TestImplementation -}}
	{{- .Assertions}}
	mockRepo.AssertExpectations(t)
}
{{end}}
`
