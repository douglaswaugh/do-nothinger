package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRunScriptWithOneStep_WaitsForEnterBeforeShowingDone(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_do_something() {
    echo "hello"
}
`)
	scriptFile.Close()

	inputReader, inputWriter := io.Pipe()
	var output bytes.Buffer

	done := make(chan bool)
	go func() {
		run(scriptFile.Name(), inputReader, &output)
		done <- true
	}()

	// Give run() time to reach the input wait
	time.Sleep(100 * time.Millisecond)

	// Send Enter
	inputWriter.Write([]byte("\n"))

	// Wait for run() to complete
	<-done

	// Clean up
	inputWriter.Close()

	// Now "Done" should be in the output
	if !strings.Contains(output.String(), "Done") {
		t.Errorf("Expected output to contain 'Done' after Enter, got: %s", output.String())
	}
}

func TestRunScriptWithDifferentStepName_DisplaysParsedStepName(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_check_queue_count() {
    echo "checking queue"
}
`)
	scriptFile.Close()

	var output bytes.Buffer
	run(scriptFile.Name(), nil, &output)

	if !strings.Contains(output.String(), "Step 1: Check queue count") {
		t.Errorf("Expected output to contain 'Step 1: Check queue count', got: %s", output.String())
	}
}

func TestRunScriptWithOneStep_DisplaysStepName(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_do_something() {
    echo "hello"
}
`)
	scriptFile.Close()

	var output bytes.Buffer
	run(scriptFile.Name(), nil, &output)

	if !strings.Contains(output.String(), "Step 1: Do something") {
		t.Errorf("Expected output to contain 'Step 1: Do something', got: %s", output.String())
	}
}

func TestRunScriptWithOneStep_DisplaysPressEnterToContinue(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_do_something() {
    echo "hello"
}
`)
	scriptFile.Close()

	var output bytes.Buffer
	run(scriptFile.Name(), nil, &output)

	if !strings.Contains(output.String(), "Press Enter to continue...") {
		t.Errorf("Expected output to contain 'Press Enter to continue...', got: %s", output.String())
	}
}

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
