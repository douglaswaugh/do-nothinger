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

var stepFuncPattern = regexp.MustCompile(`^(step_(\d+)_([a-z_]+))\(\)`)

type step struct {
	name     string
	funcName string
}

func parseSteps(scriptPath string) []step {
	data, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil
	}

	var steps []step
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		matches := stepFuncPattern.FindStringSubmatch(line)
		if matches != nil {
			funcName := matches[1]
			number := matches[2]
			words := strings.ReplaceAll(matches[3], "_", " ")
			words = strings.ToUpper(words[:1]) + words[1:]
			name := "Step " + number + ": " + words
			steps = append(steps, step{name: name, funcName: funcName})
		}
	}
	return steps
}

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	steps := parseSteps(scriptPath)
	for _, s := range steps {
		fmt.Fprintln(output, s.name)

		cmd := exec.Command("bash", "-c", "source "+scriptPath+" && "+s.funcName)
		cmd.Stdout = output
		cmd.Run()

		fmt.Fprintln(output, "Press Enter to continue...")

		if input != nil {
			bufio.NewReader(input).ReadString('\n')
		}

		fmt.Fprintln(output, s.name+" complete")
	}

	fmt.Fprintln(output, "Done")
}
