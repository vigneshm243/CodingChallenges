package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {

	//single object to store all the options required **Options Pattern**
	var commandLineOptions Options

	//defining the flags and setting default values
	flag.BoolVar(&commandLineOptions.countBytes, "c", false, "Count bytes")
	flag.BoolVar(&commandLineOptions.countLines, "l", false, "Count lines")
	flag.BoolVar(&commandLineOptions.countWords, "w", false, "Count words")
	flag.BoolVar(&commandLineOptions.countChars, "m", false, "Count characters")

	//reading the flags if present
	flag.Parse()
	
	//for printing default options that come up in case no flags are given
	if !commandLineOptions.countBytes &&
		!commandLineOptions.countLines &&
		!commandLineOptions.countWords &&
		!commandLineOptions.countChars {
		commandLineOptions.countBytes = true
		commandLineOptions.countWords = true
		commandLineOptions.countLines = true
	}

	//Read the file name or names
	fileNames := flag.CommandLine.Args()

	// A reader to read the data
	var reader *bufio.Reader

	// No files given read from stdio
	if len(fileNames) == 0 {
		reader = bufio.NewReader(os.Stdin)
	}
	if len(fileNames) == 1 {
		file, err := os.Open(fileNames[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		// To add handle 
		fmt.Println("many files")
	}

	wordCount := CalculateWordCount(reader)
	PrintWordCount(wordCount, commandLineOptions, fileNames)
}
