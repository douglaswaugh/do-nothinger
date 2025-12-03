package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("Hello, Do Nothinger!")
}

func run(scriptPath string, output io.Writer) {
	fmt.Fprintln(output, "Done")
}
