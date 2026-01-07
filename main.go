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
	functionName, displayName := parseStep(scriptPath)

	if functionName != "" {
		fmt.Fprintf(output, "Step 1: %s\n", displayName)

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

func parseStep(scriptPath string) (functionName string, displayName string) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", ""
	}

	re := regexp.MustCompile(`step_(\d+)_(\w+)\s*\(\)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 3 {
		return "", ""
	}

	functionName = matches[0][:len(matches[0])-2] // Remove the "()" from the match
	description := strings.ReplaceAll(matches[2], "_", " ")
	description = strings.ToUpper(string(description[0])) + description[1:]
	displayName = description

	return functionName, displayName
}
