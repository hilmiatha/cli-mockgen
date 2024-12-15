package main

import (
	"flag"
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
		return
	}



	

}

