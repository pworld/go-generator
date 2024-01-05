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

### Command Init
1. Add Init using Viper
```bash
package cmd

import (
	"fmt"
	"github.com/pworld/go-mvc-boilerplate/pkg/cmd/generate"
	"github.com/pworld/loggers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   viper.GetString("APP_NAME"),
	Short: "Hello From CLI",
	Run: func(cmd *cobra.Command, args []string) {
		loggers.Info(fmt.Sprintf("APP Name: %s", viper.GetString("APP_NAME")), "", "", 0)
		loggers.Info(fmt.Sprintf("APP ENV: %s", viper.GetString("APP_ENV")), "", "", 0)
	},
}

func Initialize() {
	// Register the GenerateMVC command
	rootCmd.AddCommand(generate.GenerateMVCCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
```
2. Create generateMVC. Sample Code
```bash
package generate

import (
	generator "github.com/pworld/go-generator"
	"github.com/pworld/loggers"
	"github.com/spf13/cobra"
)

func GenerateMVCCmd() *cobra.Command {
	var filePath string

	var cmd = &cobra.Command{
		Use:   "generate-mvc",
		Short: "Generate MVC structure",
		Run: func(cmd *cobra.Command, args []string) {
			if filePath == "" {
				loggers.Warn("File path is required when generating MVC")
				return
			}
			generator.GenerateMVC(filePath)
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file")
	return cmd
} 
```
### Binary Instalations
1. Download the Binary:
Go to the [Releases page](https://github.com/pworld/go-generator/releases/tag/v1.0.2) of the go-generator repository. 
2. Download the binary for your operating system (go-generator.exe for Windows, go-generator-macos for macOS, go-generator-linux for Linux).
Make the Binary Executable (Linux/macOS):
3. On Linux or macOS, you might need to make the binary executable. Run the following command:
    ```bash
    chmod +x go-generator-macos  # For macOS
    chmod +x go-generator-linux  # For Linux
    ```
   Running the Tool:
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
go run path/to/main.go --generate-mvc --file path/to/models/entity/{models-file}.go
```
Replace {model-file}.go with your actual model file name.

## Example

### Sample Directory Structures of User
```bash
- internal
  -- company
    --- models
      ---- entity
        ----- company.go
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
go run main.go generate-mvc --file path/to/models/entity/{model-file}.go
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
Contributions to go-generator are welcome! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## How to Contribute
1. Fork the repository.
2. Create a new branch for each feature or improvement.
3. Send a pull request from each feature branch to the develop branch.

## Reporting Issues
If you find any bugs or have a feature request, please create an issue on GitHub. We appreciate detailed and accurate reports that help us identify and fix issues.

## Support
If you have any questions or need help with using go-generator, please create an issue, and we will do our best to assist you.

## License
go-generator is licensed under the Apache License 2.0. See the LICENSE file for more details.