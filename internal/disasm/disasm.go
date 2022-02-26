package disasm

import (
	"debug/pe"
	"log"
	"os"
)

// Processor modes: either 32-bit or 64-bit
const (
	mode32 = 32
	mode64 = 64
)

// Disasm represents an disassembled program in memory
type Disasm struct {
	text       []byte            // .text data
	ib         uint64            // ImageBase
	bc         uint64            // BaseOfCode
	mode       uint64            // processor mode
	disasmList map[uint64]string // disassembled .text data
}

// New opens a program file and reads its content for later disassembly
func New(filename string) *Disasm {
	if err := d.init(filename); err != nil {
		log.Fatalln(err)
	}
	return nil
}

// init initializes Disasm struct with the filename contents
func (d *Disasm) init(filename string) error {
	f, err := pe.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer f.Close()

	// Get code (for PE file it's in .text section)
	d.text, err = f.Section(".text").Data()
	if err != nil {
		return err
	}

	// Fill out remaining Disasm fields taking into consideration if it is
	// either a 32 or 64 bit binary
	switch oh := f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		d.ib = uint64(oh.ImageBase)
		d.bc = uint64(oh.BaseOfCode)
		d.mode = mode32
	case *pe.OptionalHeader64:
		d.ib = oh.ImageBase
		d.bc = uint64(oh.BaseOfCode)
		d.mode = mode64
	}

	return nil
}
