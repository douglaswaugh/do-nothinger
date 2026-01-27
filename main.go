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
	step := parseStep(scriptPath)

	if step != nil {
		functionName := formatFunctionName(step)
		displayName := formatDisplayName(step)

		fmt.Fprintf(output, "Step %s: %s\n", step.Number, displayName)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && "+functionName)
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if input != nil {
			bufio.NewReader(input).ReadString('\n')
		}

		fmt.Fprintf(output, "Step %s: %s complete\n", step.Number, displayName)
	}

	fmt.Fprintln(output, "Done")
}

func parseStep(scriptPath string) *Step {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil
	}

	re := regexp.MustCompile(`step_(\d+)_(\w+)\s*\(\)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 3 {
		return nil
	}

	return &Step{Number: matches[1], Description: matches[2]}
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
