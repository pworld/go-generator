package templateMVC

const ServiceTemplate = `package services

import (
    "{{.ModuleName}}/internal/user/models/repository"
)

type {{.StructName}}Service struct {
    repo repository.{{.StructName}}Repository
}

func New{{.StructName}}Service(repo repository.{{.StructName}}Repository) *{{.StructName}}Service {
    return &{{.StructName}}Service{repo: repo}
}

// Add other service methods here
`
