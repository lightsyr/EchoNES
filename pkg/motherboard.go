package motherboard

type Motherboard struct {
	Cpu    *CPU
	Memory *Memory
}

func NewMotherboard() *Motherboard {
	return &Motherboard{Cpu: NewCPU(), Memory: NewMemory()}
}
