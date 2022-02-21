package main

import (
	"os"

	"github.com/hakbyte/disasm/internal/disasm"
)

func main() {
	_ := disasm.New(os.Args[1])
}
