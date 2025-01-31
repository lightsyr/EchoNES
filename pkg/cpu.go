package motherboard

type CPU struct {
	A  uint8  // Acumulator
	X  uint8  // Index register
	Y  uint8  // Y register
	PC uint16 // Program Counter
	SP uint16 // Stack Pointer
	P  uint8  // Status Register
}

// Create and return a new CPU structure with default values
func NewCPU() *CPU {
	return &CPU{
		A:  0x00,   // Initial value of the accumulator
		X:  0x00,   // Initial value of the X index register
		Y:  0x00,   // Initial value of the Y index register
		PC: 0x8000, // Initial value of the Program Counter
		SP: 0x01FF, // Initial value of the Stack Pointer
		P:  0x34,   // Initial value of the Status Register (P)
	}
}
