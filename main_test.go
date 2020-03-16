package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)



func TestOpenFileDoesNotExist(t *testing.T) {
	if os.Getenv("BAD_INPUT") == "1" {
		openFile("dne.json")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestOpenFileDoesNotExist")
	cmd.Env = append(os.Environ(), "BAD_INPUT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestOpenFileExists( t *testing.T) {
	filename := "mixtape-data.json"
	f := openFile(filename)

	if f != nil {
		fmt.Println("Successfully opened file", filename)
	} else {
		t.Errorf("Expected a pointer to a file!\n")
	}
}

// Note: We don't need to test the persistence methods since we're testing the Encoders
//			We trust the OS filesystem capability