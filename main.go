package main

import (
	"fmt"
	"log"
	"os"

	motherboard "github.com/lightsyr/EchoNES/pkg"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go <path-to-rom>")
	}

	motherboard := motherboard.NewMotherboard()

	romPath := os.Args[1]

	err := motherboard.Memory.LoadROM(romPath)

	if err != nil {
		log.Fatalf("Error loading ROM: %v", err)
	}

	fmt.Printf("Value at address 0x8000: %x\n", motherboard.Memory.Data[0x8000])
}
