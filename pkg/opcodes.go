package motherboard

type Opcode struct {
	Name     string
	Opcode   uint8
	Operands int
	Execute  func(*Motherboard)
	Cycles   int
}

func LDA(m *Motherboard) {
	address := uint16(m.CPU.PC + 1)
	m.Memory.Data[address] = m.CPU.A
	m.CPU.PC += 2
}

var opcodeTable = []Opcode{
	{"LDA", 0xA9, 1, LDA, 2},
}
