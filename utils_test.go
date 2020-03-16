package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestParseArgumentsMissing(t *testing.T) {
	if os.Getenv("BAD_INPUT") == "1" {
		parseArguments([]string{"-i", "-e", "changes"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestParseArgumentsMissing")
	cmd.Env = append(os.Environ(), "BAD_INPUT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestParseArgumentsInputFilename(t *testing.T) {
	expectedInputFilename := "input.json"
	expectedChangesFilename := "changeset.json"
	expectedPerformanceFlag := true

	i, c, p := parseArguments([]string{"-i", expectedInputFilename,  "-c", expectedChangesFilename, "-p"})

	if i != expectedInputFilename {
		t.Errorf("Expected input filename: %s, got: %s\n", expectedInputFilename, i)
	}
	if c != expectedChangesFilename {
		t.Errorf("Expected changes filename: %s, got: %s\n", expectedChangesFilename, c)
	}
	if p != expectedPerformanceFlag {
		t.Errorf("Expected performance to be: %t, got: %t\n", expectedPerformanceFlag, p)
	}
}

func TestParseArgumentsUnrecognized(t *testing.T) {
	expectedInputFilename := "mixtape.json"
	expectedChangesFilename := "changes"
	expectedPerformanceFlag := false

	i, c, p := parseArguments([]string{"-z", "asldjasd", "--a", "fsfsdfsd"})

	if i != expectedInputFilename {
		t.Errorf("Expected input filename: %s, got: %s\n", expectedInputFilename, i)
	}
	if c != expectedChangesFilename {
		t.Errorf("Expected changes filename: %s, got: %s\n", expectedChangesFilename, c)
	}
	if p != expectedPerformanceFlag {
		t.Errorf("Expected performance to be: %t, got: %t\n", expectedPerformanceFlag, p)
	}
}

func TestParseArgumentsNoneProvided(t *testing.T) {
	expectedInputFilename := "mixtape.json"
	expectedChangesFilename := "changes"
	expectedPerformanceFlag := false

	i, c, p := parseArguments([]string{})

	if i != expectedInputFilename {
		t.Errorf("Expected input filename: %s, got: %s\n", expectedInputFilename, i)
	}
	if c != expectedChangesFilename {
		t.Errorf("Expected changes filename: %s, got: %s\n", expectedChangesFilename, c)
	}
	if p != expectedPerformanceFlag {
		t.Errorf("Expected performance to be: %t, got: %t\n", expectedPerformanceFlag, p)
	}
}
