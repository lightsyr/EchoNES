package motherboard

import "fmt"

func DebugInstruction(o Opcode, c *CPU) {
	// Print the opcode and CPU state information for debugging purposes
	fmt.Printf("DEBUG: Executing opcode %s\n", o)
	// Print the registers values
	fmt.Printf("A: %d | X: %d | Y: %d", c.A, c.X, c.Y)
}
