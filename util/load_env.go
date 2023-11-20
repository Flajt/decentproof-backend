package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	var envFilePath string
	var verbose bool
	envVars := make(map[string]string)
	flag.StringVar(&envFilePath, "path", ".env", "Path to the .env fiel to load in the environment")
	flag.BoolVar(&verbose, "verbose", false, "Prints out the environment variables that are loaded")
	flag.Parse()
	file, err := os.Open(envFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contentArray := strings.SplitN(scanner.Text(), "=", 2)
		if verbose {
			fmt.Println("Now loading: " + contentArray[0])
		}
		envVars[contentArray[0]] = contentArray[1]
	}

	if len(envVars) == 0 {
		log.Fatal("No environment variables found in the file")
		return
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for key, value := range envVars {
		os.Setenv(key, value)
	}
	log.Print("Environment variables loaded")
	//CREDITS: https://stackoverflow.com/a/17375838
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

}
