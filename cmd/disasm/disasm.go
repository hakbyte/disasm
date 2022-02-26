package main

import (
	"fmt"
	"os"

	"github.com/hakbyte/disasm/internal/disasm"
)

func main() {
	_ := disasm.New(os.Args[1])
}

// Usage prints the program usage to stdout and exit
func Usage() {
	fmt.Print("Usage:\n\tdisasm.exe <file.exe>\n")
	os.Exit(1)
}
