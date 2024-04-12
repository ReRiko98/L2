package main

import "fmt"

type CPU struct{}

func (c *CPU) run() {
    fmt.Println("CPU is running")
}

type Memory struct{}

func (m *Memory) load() {
    fmt.Println("Memory is loaded")
}

type HardDrive struct{}

func (h *HardDrive) read() {
    fmt.Println("Hard Drive is reading")
}

type ComputerFacade struct {
    cpu      *CPU
    memory   *Memory
    hardrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:      &CPU{},
        memory:   &Memory{},
        hardrive: &HardDrive{},
    }
}

func (c *ComputerFacade) Start() {
    fmt.Println("Starting computer")
    c.cpu.run()
    c.memory.load()
    c.hardrive.read()
    fmt.Println("Computer started successfully")
}

// func main() {
//     computer := NewComputerFacade()
//     computer.Start()
// }
