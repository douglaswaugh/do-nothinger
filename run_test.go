package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRunScriptWithOneStep_OutputsStepContent(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_do_something() {
    echo "hello from step"
}
`)
	scriptFile.Close()

	var output bytes.Buffer
	run(scriptFile.Name(), nil, &output)

	if !strings.Contains(output.String(), "hello from step") {
		t.Errorf("Expected output to contain 'hello from step', got: %s", output.String())
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
	run(scriptFile.Name(), nil, &output)

	if !strings.Contains(output.String(), "Done") {
		t.Errorf("Expected output to contain 'Done', got: %s", output.String())
	}
}
