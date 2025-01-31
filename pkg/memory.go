package motherboard

import "os"

type Memory struct {
	Data [0x10000]byte // 64KB memory
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) LoadROM(filepath string) error {
	romData, err := os.ReadFile(filepath)

	if err != nil {
		return err
	}

	copy(m.Data[0x8000:], romData)

	return nil
}
