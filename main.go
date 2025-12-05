package main

import (
	"bufio"
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
	fmt.Fprintln(output, "Step 1: Do something")

	cmd := exec.Command("bash", "-c", "source "+scriptPath+" && step_1_do_something")
	cmd.Stdout = output
	cmd.Run()

	fmt.Fprintln(output, "Press Enter to continue...")

	if input != nil {
		bufio.NewReader(input).ReadString('\n')
	}

	fmt.Fprintln(output, "Done")
}
