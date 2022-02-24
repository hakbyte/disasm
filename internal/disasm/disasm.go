package disasm

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
	return nil
}
