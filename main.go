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
	stepName := parseStepName(scriptPath)

	if stepName != "" {
		displayName := formatStepName(stepName)
		fmt.Fprintf(output, "Step 1: %s\n", displayName)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && "+stepName)
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if input != nil {
			bufio.NewReader(input).ReadString('\n')
		}
	}

	fmt.Fprintln(output, "Done")
}

func parseStepName(scriptPath string) string {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return ""
	}

	re := regexp.MustCompile(`(step_\d+_\w+)\s*\(\)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func formatStepName(stepName string) string {
	// Remove "step_1_" prefix
	re := regexp.MustCompile(`step_\d+_`)
	name := re.ReplaceAllString(stepName, "")

	// Replace underscores with spaces
	name = strings.ReplaceAll(name, "_", " ")

	// Capitalize first letter
	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}

	return name
}
