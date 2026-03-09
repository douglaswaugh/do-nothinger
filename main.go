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

var stepFuncPattern = regexp.MustCompile(`^step_(\d+)_([a-z_]+)\(\)`)

func parseStepName(scriptPath string) string {
	file, err := os.Open(scriptPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := stepFuncPattern.FindStringSubmatch(line)
		if matches != nil {
			number := matches[1]
			words := strings.ReplaceAll(matches[2], "_", " ")
			words = strings.ToUpper(words[:1]) + words[1:]
			return "Step " + number + ": " + words
		}
	}
	return ""
}

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	stepName := parseStepName(scriptPath)
	if stepName != "" {
		fmt.Fprintln(output, stepName)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && step_1_do_something")
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if input != nil {
			bufio.NewReader(input).ReadString('\n')
		}
	}

	fmt.Fprintln(output, "Done")
}
