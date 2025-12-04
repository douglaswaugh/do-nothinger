package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	cmd := exec.Command("bash", "-c", "source "+scriptPath+" && step_1_do_something")
	cmd.Stdout = output
	cmd.Run()

	fmt.Fprintln(output, "Done")
}
