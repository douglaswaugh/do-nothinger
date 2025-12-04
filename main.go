package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdin, os.Stdout)
}

func run(scriptPath string, input io.Reader, output io.Writer) {
	fmt.Fprintln(output, "Done")
}
