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
	"github.com/hilmiatha/cli-mockgen/data"
)

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

	mappingRes, err := generateMockData(mapping)
	if  err != nil {
		fmt.Printf("Error generating mock data: %s\n", err)
		os.Exit(0)
	}

	
	if err := writeOutput(outputPath, mappingRes); err != nil {
		fmt.Printf("Error writing output file: %s\n", err)
		os.Exit(0)
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
		if !data.SUPPORTED_MOCK_DATA[v] {
			return fmt.Errorf("unsupported mock data type : %s", v)
		}
	}
	return nil
}


func generateMockData(mapping map[string]string) (map[string]any, error) {
	res := make(map[string]any)
	for k,v := range mapping {
		switch v {
		case "name":
			res[k] = data.Generate(v)
		case "address":
			res[k] = data.Generate(v)
		case "phone":
			res[k] = data.Generate(v)
		case "date":
			res[k] = data.Generate(v)
		}
	}

	return res, nil
}

func writeOutput(outputPath string, mapping map[string]any) error {
	if outputPath == "" {
		return errors.New("output file path is not valid")
	}

	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(outputPath, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	resultByte, err := json.MarshalIndent(mapping, "", "  ")
	if err != nil {
		return err
	}

	_, err2 := file.Write(resultByte)
	if err2 != nil {
		return err2
	}

	fmt.Println("Mock data generated successfully")
	return nil
}