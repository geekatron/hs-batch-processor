package main

import (
	"fmt"
	"os"
)

func main() {
	go PrintUsage()

	// Parse input arguments
	//	-i <input filename>
	// 	-c <change set filename>

	// If no arguments are provided assume the following:
	//	1. Input filename is 'mixtape.json'
	// 	2. Change set filename is 'changes'
	// 	2. Output filename is 'output.json'

	outputFilename := "output.json"
	inputFilename, changesetFilename := parseArguments(os.Args[1:])
	fmt.Printf("Input arguments:\n  Input Filename: %s\n  Change set filename: %s\n", inputFilename, changesetFilename)

	// Decode the input file into a MixtapeIndex (map) -> We will manipulate this with changes and Encode the changes
	// into a Mixtape struct
	mixtapeIndex := parseInputFile(inputFilename)


	// Read in the changes and apply them
	// TODO
	// NOTE: We could make this concurrent so we can apply multiple changes simultaneously.
	// 		 However, this depends if order of operations is important (e.g. preserving order of added
	//		 songs to playlist).
	applyChangeFile(changesetFilename, mixtapeIndex)

	// Persist the changes
	mixtapeIndex.persistChanges(outputFilename)
}

// Open and return the pointer of type file; if it can't be found kill the program
func openFile(filename string) *os.File {
	fmt.Println("Opening file: ", filename)
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening the file %v! \nError: %#v", filename, err)
		os.Exit(1)
	}

	return f
}

func parseInputFile(filename string) MixtapeIndex {
	f := openFile(filename)
	defer f.Close()
	// Don't build a Mixtape struct, instead build a MixtapeIndex that we will modify
	// Since we're streaming the data in, we can just build the map (index) directly and bypass the Mixtape struct
	// NOTE: Mixtape struct will just be used for testing Decoding/Encoding to make sure our custom Dec/Enc code is working
	//		 as expected
	return decodeMixtape(f)
}
