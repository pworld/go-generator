package templateMVC

const TestTemplate = `package {{.LowerStructName}}_test

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"{{.ModulePath}}/models/entity"
	"{{.ModulePath}}/services"
	"{{.ModulePath}}/tests/mocks"
)

func TestCreate{{.StructName}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{.StructName}}Repository)
	{{.LowerStructName}} := entity.{{.StructName}}{
        {{- range .Fields }}
        {{ .Name }}: {{ .Value }},
        {{- end }}
	}

	mockRepo.On("Create{{.StructName}}", mock.Anything, mock.AnythingOfType("entity.{{.StructName}}")).Return(1, nil)

	service := services.New{{.StructName}}Service(mockRepo)
	id, err := service.Create{{.StructName}}(context.Background(), {{.LowerStructName}})

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockRepo.AssertExpectations(t)
}

func TestGet{{.StructName}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{.StructName}}Repository)
	testID := 1 // Example test ID
	expected{{.StructName}} := entity.{{.StructName}}{
        {{- range .Fields }}
        {{ .Name }}: {{ .Value }},
        {{- end }}
	}

	mockRepo.On("Get{{.StructName}}ByID", mock.Anything, testID).Return(expected{{.StructName}}, nil)

	service := services.New{{.StructName}}Service(mockRepo)
	result{{.StructName}}, err := service.Get{{.StructName}}(context.Background(), testID)

	assert.NoError(t, err)
	assert.Equal(t, expected{{.StructName}}, result{{.StructName}})
	mockRepo.AssertExpectations(t)
}

func TestUpdate{{.StructName}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{.StructName}}Repository)
	{{.LowerStructName}}ToUpdate := entity.{{.StructName}}{
        {{- range .Fields }}
        {{ .Name }}: {{ .Value }},
        {{- end }}
	}

	mockRepo.On("Update{{.StructName}}", mock.Anything, mock.AnythingOfType("entity.{{.StructName}}")).Return(nil)
	service := services.New{{.StructName}}Service(mockRepo)
	err:= service.Update{{.StructName}}(context.Background(), {{.LowerStructName}}ToUpdate)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete{{.StructName}}(t *testing.T) {
	mockRepo := new(mocks.Mock{{.StructName}}Repository)
	testID := 1 // Example test ID
	mockRepo.On("Delete{{.StructName}}", mock.Anything, testID).Return(nil)
	
	service := services.New{{.StructName}}Service(mockRepo)
	err := service.Delete{{.StructName}}(context.Background(), testID)
	
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test{{.StructName}}Lists(t *testing.T) {
	mockRepo := new(mocks.Mock{{.StructName}}Repository)
	page, pageSize := 1, 10

	expected{{.StructName}}s := []entity.{{.StructName}}{
		{
			{{- range .Fields }}
			{{ .Name }}: {{ .Value }},
			{{- end }}
		}
	}
	expectedTotal := int64(len(expected{{.StructName}}s))

	mockRepo.On("List{{.StructName}}s", mock.Anything, page, pageSize, "", mock.Anything).Return(expected{{.StructName}}s, nil)
	mockRepo.On("Total{{.StructName}}s", mock.Anything, "", mock.Anything).Return(expectedTotal, nil)
	
	service := services.New{{.StructName}}Service(mockRepo)
	result{{.StructName}}s, total, err := service.{{.StructName}}Lists(context.Background(), page, pageSize, "", nil)
	
	assert.NoError(t, err)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expected{{.StructName}}s, result{{.StructName}}s)
	mockRepo.AssertExpectations(t)
}
`
