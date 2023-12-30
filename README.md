# go-generator

## Overview
`go-generator` automates the creation of MVC components in Go projects. It generates models, views, controllers, services, repositories, and test cases from Go entity files, streamlining the development process.

## Features
- **MVC Component Generation**: Automatically creates models, views, controllers, services, and repositories.
- **Test Case Automation**: Generates test cases for your services and repositories.
- **Efficient Workflow**: Reduces manual coding and accelerates project setup.

## Prerequisites
Your Go project should follow this structure:
- Entity files must be located at `internal/{package-name}/models/entity/{model-file}.go`.
- *internal* Directory is necessary and can't be change

## Installation Guide:
1. Download the Binary:
Go to the [Releases page](https://github.com/pworld/go-generator/releases/tag/v.1.0.2) of the go-generator repository. 
2. Download the binary for your operating system (go-generator.exe for Windows, go-generator-macos for macOS, go-generator-linux for Linux).
Make the Binary Executable (Linux/macOS):
3. On Linux or macOS, you might need to make the binary executable. Run the following command:
    ```bash
    chmod +x go-generator-macos  # For macOS
    chmod +x go-generator-linux  # For Linux
    Running the Tool:
    ```
    You can run the tool directly from the command line:
    ```bash
    ./go-generator-macos --generate-mvc --file path/to/entity.go  # macOS
    ./go-generator-linux --generate-mvc --file path/to/entity.go  # Linux
    ```
4. On Windows
    - Opening Command Prompt: 
      - Open the Command Prompt by searching for cmd in the Windows search bar and clicking on the Command Prompt application.
    - Navigating to the File: 
      - Once the Command Prompt is open, the user needs to navigate to the directory where go-generator.exe is located. This can be done using the cd (change directory) command.
      For example, if go-generator.exe is in the Downloads folder, the user would type cd Downloads and press Enter.
    - Running the Application:
      - To run the application from the Command Prompt, the user types the name of the executable file (go-generator.exe) and presses Enter.
      If the user wants to pass additional arguments or flags to the application (as specified in your documentation or help command), they can be added after the file name. For example: 
      ```bash
      go-generator.exe --generate-mvc --file path/to/entity.go.
      ```
## Usage
To generate MVC components in this Github Repository, run:
```bash
go run path/to/main.go --generate-mvc --file path/to/models/entity/{model-file}.go
```
Replace {model-file}.go with your actual model file name.

## Example

### Sample Directory Structures of User
```bash
- internal
  -- user
    --- models
      ---- entity
        ----- user.go
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
### Usage
```bash
go run main.go --generate-mvc --file internal/user/models/entity/user.go
```

## Expected Output Structure
After running go-generator, you can expect the following directory structure and files:
```base
- internal
  -- user
    --- models
      ---- entity
        ----- user.go
      ---- repository
        ----- user_repository.go
    --- services
      ---- user_service.go
    --- views
      ---- user_view.go
    --- controllers
      ---- user_controller.go
    --- tests
      ---- service_tests
        ----- user_service_test.go
      ---- mocks
        ----- user_mock.go

```

## Contributing
Contributions are welcome. Feel free to open issues or submit pull requests.

## License
go-generator is licensed under the Apache 2.0 License. See LICENSE for more details.