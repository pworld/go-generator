package main

import (
	"flag"
	"fmt"
	"github.com/pworld/go-generator/src/generatorMVC"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// Define flags
	pflag.String("file", "", "Path to the file")
	pflag.Bool("generate-mvc", false, "Flag to trigger MVC generation")

	// Merge standard flags with pflag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Parse flags
	pflag.Parse()

	// Bind the current command line flags to viper
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("Failed to bind flags:", err)
		return
	}

	// Get values using viper
	filePath := viper.GetString("file")
	generateMVC := viper.GetBool("generate-mvc")

	// Handle command
	if generateMVC {
		generatorMVC.GenerateMVC(filePath)
	} else {
		fmt.Println("Invalid or missing command")
	}
}
