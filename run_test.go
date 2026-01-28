package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
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

func TestRunScriptWithOneStep_DisplaysStepCompleteAfterEnter(t *testing.T) {
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

	if !strings.Contains(output.String(), "Step 1: Do something complete") {
		t.Errorf("Expected output to contain 'Step 1: Do something complete', got: %s", output.String())
	}
}

func TestRunScriptWithMultipleStepsInOrder_ExecutesAllSteps(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_first_step() {
    echo "first step output"
}
step_2_second_step() {
    echo "second step output"
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

	// Send Enter presses in a goroutine to avoid deadlock if run() finishes early
	go func() {
		// Send Enter for first step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		// Send Enter for second step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		inputWriter.Close()
	}()

	// Wait for run() to complete
	<-done

	outputStr := output.String()
	if !strings.Contains(outputStr, "Step 1: First step") {
		t.Errorf("Expected output to contain 'Step 1: First step', got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "first step output") {
		t.Errorf("Expected output to contain 'first step output', got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Step 2: Second step") {
		t.Errorf("Expected output to contain 'Step 2: Second step', got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "second step output") {
		t.Errorf("Expected output to contain 'second step output', got: %s", outputStr)
	}
}

func TestRunScriptWithOneStep_ExecutesStepCommand(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_1_create_file() {
    echo "test content" > /tmp/do-nothinger-test-output.txt
}
`)
	scriptFile.Close()

	// Clean up test file before and after
	os.Remove("/tmp/do-nothinger-test-output.txt")
	defer os.Remove("/tmp/do-nothinger-test-output.txt")

	var output bytes.Buffer
	run(scriptFile.Name(), nil, &output)

	// Verify the file was created by the step
	content, err := os.ReadFile("/tmp/do-nothinger-test-output.txt")
	if err != nil {
		t.Errorf("Expected step to create file /tmp/do-nothinger-test-output.txt, but got error: %v", err)
	}
	if !strings.Contains(string(content), "test content") {
		t.Errorf("Expected file to contain 'test content', got: %s", string(content))
	}
}

func TestRunScriptWithMultipleStepsInDifferentOrder_ExecutesInCorrectOrder(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
step_2_second_step() {
    echo "second"
}
step_1_first_step() {
    echo "first"
}
step_3_third_step() {
    echo "third"
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

	// Send Enter presses in a goroutine to avoid deadlock if run() finishes early
	go func() {
		// Send Enter for first step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		// Send Enter for second step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		// Send Enter for third step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		inputWriter.Close()
	}()

	// Wait for run() to complete
	<-done

	outputStr := output.String()

	// Verify steps are executed in numerical order (1, 2, 3), not file order (2, 1, 3)
	firstPos := strings.Index(outputStr, "first")
	secondPos := strings.Index(outputStr, "second")
	thirdPos := strings.Index(outputStr, "third")

	if firstPos == -1 || secondPos == -1 || thirdPos == -1 {
		t.Errorf("Expected output to contain 'first', 'second', and 'third', got: %s", outputStr)
	}

	if !(firstPos < secondPos && secondPos < thirdPos) {
		t.Errorf("Expected steps to execute in order 1, 2, 3, but got different order. Output: %s", outputStr)
	}
}

func TestRunScriptWithStepsAndNonStepFunctions_OnlyExecutesSteps(t *testing.T) {
	scriptFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString(`#!/bin/bash
helper_function() {
    echo "helper called"
}
step_1_first_step() {
    echo "first"
}
another_helper() {
    echo "another helper"
}
step_2_second_step() {
    echo "second"
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

	// Send Enter presses in a goroutine to avoid deadlock if run() finishes early
	go func() {
		// Send Enter for first step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		// Send Enter for second step
		time.Sleep(100 * time.Millisecond)
		inputWriter.Write([]byte("\n"))

		inputWriter.Close()
	}()

	// Wait for run() to complete
	<-done

	outputStr := output.String()

	// Verify only step functions are executed
	if !strings.Contains(outputStr, "first") {
		t.Errorf("Expected output to contain 'first', got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "second") {
		t.Errorf("Expected output to contain 'second', got: %s", outputStr)
	}
	if strings.Contains(outputStr, "helper called") {
		t.Errorf("Expected output to NOT contain 'helper called', got: %s", outputStr)
	}
	if strings.Contains(outputStr, "another helper") {
		t.Errorf("Expected output to NOT contain 'another helper', got: %s", outputStr)
	}
}
