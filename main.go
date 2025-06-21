package main

import (
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

//go:embed .env
var envFile string

func main() {
	envMap, err := godotenv.Parse(strings.NewReader(envFile))
	if err != nil {
		fmt.Println("error parsing embedded .env file")
	} else {
		// Set environment variables
		for key, value := range envMap {
			os.Setenv(key, value)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
