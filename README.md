# go-generator

## Overview
`go-generator` automates the creation of MVC components in Go projects. It generates models, views, controllers, services, repositories, and test cases from Go entity files, streamlining the development process.

## Features
- **MVC Component Generation**: Automatically creates models, views, controllers, services, and repositories.
- **Test Case Automation**: Generates test cases for your services and repositories.
- **Efficient Workflow**: Reduces manual coding and accelerates project setup.

## Prerequisites
Your Go project should follow this structure:
- Entity files must be located at `models/entity/{model-file}.go`.

### Sample Directory Structures of test
```bash
-internal
--user
---models
----entity
-----user.go
```
### Sample of User File
```go
package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"-" validate:"required"` // "-" in JSON tag to prevent sending the password hash
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

## Installation
```bash
go get github.com/yourusername/go-generator
```

## Usage
To generate MVC components, run:
```bash
go run main.go --generate-mvc --file path/to/models/entity/{model-file}.go
```

Replace {model-file}.go with your actual model file name.

## Example
### Usage
```bash
go run main.go --generate-mvc --file internal/user/models/entity/user.go
```

## Contributing
Contributions are welcome. Feel free to open issues or submit pull requests.

## License
go-generator is licensed under the Apache 2.0 License. See LICENSE for more details.