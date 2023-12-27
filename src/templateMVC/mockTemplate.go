package templateMVC

const MockTemplate = `package mocks

import (
	"context"
	"{{.ModuleName}}/internal/{{.PackagePath}}/models/entity"
	"github.com/stretchr/testify/mock"
)

// Mock{{.StructName}}Repository is a mock type for the {{.StructName}}Repository type
type Mock{{.StructName}}Repository struct {
	mock.Mock
}

{{- range .Methods }}
// {{.Name}} mocks the {{.Name}} method
func (m *Mock{{$.StructName}}Repository) {{.Name}}({{.Parameters}}) {{.ReturnTypes}} {
	args := m.Called({{.CallParameters}})
	{{- if .HasMultipleReturnValues }}
	return {{.ReturnValues}}
	{{- else }}
	return args.Get(0).({{.SingleReturnValue}})
	{{- end }}
}
{{- end }}

`
