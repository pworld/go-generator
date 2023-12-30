package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/pworld/go-generator/src/generatorMVC"
	"github.com/pworld/loggers"
)

func main() {
	// Define and parse flags
	pflag.String("file", "", "Path to the file")
	pflag.Bool("generate-mvc", false, "Flag to trigger MVC generation")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Bind command line flags to viper
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		loggers.Error(fmt.Sprintf("Error binding flags: %v\n", err))
		os.Exit(1)
	}

	// Extract flag values using viper
	filePath := viper.GetString("file")
	generateMVC := viper.GetBool("generate-mvc")

	// Execute the command based on flags
	if generateMVC {
		if filePath == "" {
			loggers.Warn("File path is required when generating MVC")
			os.Exit(1)
		}
		generatorMVC.GenerateMVC(filePath)
	} else {
		loggers.Error("Invalid or missing command")
		os.Exit(1)
	}
}
