package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)
var SUPPORTED_MOCK_DATA = map[string]bool{
	"name": true,
	"address": true,
	"phone": true,
	"date": true,
}
func main () {
	// Define the command line flags
	var help bool
	var inputPath, outputPath string

	// Parse the command line flags
	flag.StringVar(&inputPath, "i", "", "JSON input file path (shorthand)")
	flag.StringVar(&inputPath, "input", "", "JSON input file path")

	flag.StringVar(&outputPath, "o", "", "JSON output file path (shorthand)")
	flag.StringVar(&outputPath, "output", "", "JSON output file path")

	flag.BoolVar(&help, "h", false, "Show help (shorthand)")
	flag.BoolVar(&help, "help", false, "Show help")


	flag.Parse()


	// Show help if requested
	if help || inputPath == "" || outputPath == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Check if the input file exists
	if err := validateInputFile(inputPath); err != nil {
		fmt.Printf("Input file does not exist, \n error: %s\n", err)
		os.Exit(0)
	}

	// Check if the input file exists
	if err := validateOutputPath(outputPath); err != nil {
		fmt.Printf("Output file does not exist, \n error: %s", err)
	}

	// Open the input file
	var mapping map[string]string
	if err := readInput(inputPath, &mapping); err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		os.Exit(0)
	}

	for k,v := range mapping {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}




	

}

func validateInputFile (inputPath string) error {
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

func validateOutputPath (outputPath string) error {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return err
	}
	fmt.Println("Output file already exists, do you want to overwrite it? (y/n)")
	confirmOverwrite()
	return nil
}

func confirmOverwrite () {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
	}
	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" && response != "yes"{
		fmt.Println("Exiting program")
		os.Exit(0)
	}
}

func readInput(inputPath string, mapping *map[string]string) error {
	if inputPath == "" {
		return errors.New("input file path is not valid")
	}

	if mapping == nil {
		return errors.New("mapping is not valid")
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileByte) == 0 {
		return errors.New("input file is empty")
	}

	// Unmarshal the JSON
	if err := json.Unmarshal(fileByte, mapping); err != nil {
		return err
	}

	if err := validateMapping(mapping); err != nil {
		return err
	}

	return nil
}

func validateMapping(mapping *map[string]string) error {
	for _,v := range *mapping {
		if !SUPPORTED_MOCK_DATA[v] {
			return fmt.Errorf("unsupported mock data type : %s", v)
		}
	}
	return nil
}
