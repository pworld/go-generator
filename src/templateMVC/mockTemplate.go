package templateMVC

const MockTemplate = `package mocks

// Mock for {{.StructName}}Service
type Mock{{.StructName}}Service struct {
    // Define mock methods and fields
}
`
