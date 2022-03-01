package main

import (
	"fmt"
	"os"

	"github.com/hakbyte/disasm/internal/disasm"
)

func main() {
	// Validate command line parameters
	if len(os.Args) != 2 {
		Usage()
	}

	// Let's go!
	d := disasm.New(os.Args[1])
	d.Decode()
}

// Usage prints the program usage to stdout and exit
func Usage() {
	fmt.Print("Usage:\n\tdisasm.exe <file.exe>\n")
	os.Exit(1)
}
