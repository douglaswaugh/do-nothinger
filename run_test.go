package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

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
