package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain_CallsRunWithScriptPath(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString("#!/bin/bash\n")
	scriptFile.Close()

	cmd := exec.Command("go", "run", ".", scriptFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v, output: %s", err, output)
	}

	if !strings.Contains(string(output), "Done") {
		t.Errorf("Expected output to contain 'Done', got: %s", output)
	}
}

func TestRunScriptWithZeroSteps_DisplaysDone(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString("#!/bin/bash\n")
	scriptFile.Close()

	var output bytes.Buffer
	run(scriptFile.Name(), &output)

	if !strings.Contains(output.String(), "Done") {
		t.Errorf("Expected output to contain 'Done', got: %s", output.String())
	}
}
