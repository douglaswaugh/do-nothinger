package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	stepNumber, stepDescription := parseStep(scriptPath)

	if stepNumber != "" {
		functionName := formatFunctionName(stepNumber, stepDescription)
		displayName := formatDisplayName(stepDescription)

		fmt.Fprintf(output, "Step %s: %s\n", stepNumber, displayName)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && "+functionName)
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if input != nil {
			bufio.NewReader(input).ReadString('\n')
		}
	}

	fmt.Fprintln(output, "Done")
}

func parseStep(scriptPath string) (stepNumber string, stepDescription string) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", ""
	}

	re := regexp.MustCompile(`step_(\d+)_(\w+)\s*\(\)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 3 {
		return "", ""
	}

	return matches[1], matches[2]
}

func formatFunctionName(stepNumber, stepDescription string) string {
	return "step_" + stepNumber + "_" + stepDescription
}

func formatDisplayName(stepDescription string) string {
	name := strings.ReplaceAll(stepDescription, "_", " ")
	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}
