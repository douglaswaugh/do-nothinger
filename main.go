package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	scriptPath := os.Args[1]
	run(scriptPath, os.Stdout)
}

func run(scriptPath string, output io.Writer) {
	fmt.Fprintln(output, "Done")
}
