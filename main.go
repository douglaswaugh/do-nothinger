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

type Step struct {
	Number      string
	Description string
}

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	steps := parseSteps(scriptPath)

	var reader *bufio.Reader
	if input != nil {
		reader = bufio.NewReader(input)
	}

	for _, step := range steps {
		functionName := formatFunctionName(step)
		displayName := formatDisplayName(step)

		fmt.Fprintf(output, "Step %s: %s\n", step.Number, displayName)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && "+functionName)
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if reader != nil {
			reader.ReadString('\n')
		}

		fmt.Fprintf(output, "Step %s: %s complete\n", step.Number, displayName)
	}

	fmt.Fprintln(output, "Done")
}

func parseSteps(scriptPath string) []*Step {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil
	}

	re := regexp.MustCompile(`step_(\d+)_(\w+)\s*\(\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	var steps []*Step
	for _, match := range matches {
		if len(match) >= 3 {
			steps = append(steps, &Step{Number: match[1], Description: match[2]})
		}
	}

	return steps
}

func formatFunctionName(step *Step) string {
	return "step_" + step.Number + "_" + step.Description
}

func formatDisplayName(step *Step) string {
	name := strings.ReplaceAll(step.Description, "_", " ")
	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}
