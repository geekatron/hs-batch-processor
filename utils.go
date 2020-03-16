package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"time"
)

// Notes:
// 1. Since we know the names of the input, change and output files we could have gotten away with the application
//		not taking any arguments. It'd be the responsibility of the user to make sure the files are sitting alongside
//		the binary.
// 2. Could have used a much better library but just doing some simple argument parsing in-order to ba ble to test
//	different input files without modifying variables.
func parseArguments(args []string) (string, string, bool) {
	iFilename := "mixtape.json"
	cFilename := "changes"
	performance := false

	fmt.Printf("Length of arguments: %#v\n", len(args))

	for i, arg := range args {
		switch arg {
		case "-i":
			if !validFilenameArgument(args[i+1]) {
				iFilename = args[i+1]
			} else {
				fmt.Printf("Argument %s was provided an invalid filename: %s. Using default: %s", arg, args[i+1], iFilename)
				os.Exit(1)
			}
		case "-c":
			if !validFilenameArgument(args[i+1]) {
				cFilename = args[i+1]
			} else {
				fmt.Printf("Argument %s was provided an invalid filename: %s. Using default: %s", arg, args[i+1], iFilename)
				os.Exit(1)
			}
		case "-p":
			performance = true
		default:
			fmt.Println("Unrecognized argument:", arg)
		}
	}

	return iFilename, cFilename, performance
}

// Check to see that the value after the switch flag (-e, -i) is not another command flag (-* || --*)
func validFilenameArgument(arg string) bool {
	match, err := regexp.MatchString("(^-[a-zA-Z])|^(--[a-zA-Z]*)", arg)
	if err != nil {
		fmt.Printf("Error matching provided RegEx pattern against: %s\nError: %#v\n", arg, err)
		os.Exit(1)
	}
	return match
}

// Print memory usage of the current Application; useful for debugging performance issues
func PrintUsage() {
	f, err := os.OpenFile("performance.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "[batch processor] ", log.LstdFlags)
	logger.Println("\n\n~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~\n" +
		"~-=-~ Started Batch Processor Binary!!! ~-=-~" +
		"\n~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~\n")

	for ; ; {
		time.Sleep(500 * time.Microsecond)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
		logger.Printf("Alloc = %v MiB", bytesToMegabytes(m.Alloc))
		logger.Printf("\tTotalAlloc = %v MiB", bytesToMegabytes(m.TotalAlloc))
		logger.Printf("\tHeap_inuse = %v MiB", bytesToMegabytes(m.HeapInuse))
		logger.Printf("\tFrees = %v MiB", bytesToMegabytes(m.Frees))
		logger.Printf("\tSys = %v MiB", bytesToMegabytes(m.Sys))
		logger.Printf("\tNumGC = %v\n", m.NumGC)
	}
}
// Convert bytes into Megabytes
func bytesToMegabytes(b uint64) uint64 {
	return b / 1024 / 1024
}

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Convenience functions to help Unmarshal test data
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
func loadTestMixtape(filename string) (Mixtape, error) {
	m := Mixtape{}
	f, err := os.Open(filename)
	if err != nil {
		return m, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}


