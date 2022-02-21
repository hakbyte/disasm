package disasm

// Processor modes: either 32-bit or 64-bit
const (
	mode32 = 32
	mode64 = 64
)

// Disasm represents an disassembled program in memory
type Disasm struct {
}

// New opens a program file and reads its content for later disassembly
func New(filename string) *Disasm {
	return nil
}
