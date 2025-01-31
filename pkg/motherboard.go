package motherboard

type Motherboard struct {
	CPU    *CPU
	Memory *Memory
}

func NewMotherboard() *Motherboard {
	return &Motherboard{CPU: NewCPU(), Memory: NewMemory()}
}
