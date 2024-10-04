package memory

import (
	"github.com/ephill/gbgo/internal/hram"
	"github.com/ephill/gbgo/internal/rom"
	"github.com/ephill/gbgo/internal/vram"
	"github.com/ephill/gbgo/internal/wram"
)

const romBankStart = 0x0000
const romBankEnd = 0x7FFF
const vramStart = 0x8000
const vramEnd = 0x9FFF
const eramStart = 0xA000
const eramEnd = 0xBFFF
const wramStart = 0xC000
const wramEnd = 0xDFFF
const echoStart = 0xE000
const echoEnd = 0xFDFF
const oamStart = 0xFE00
const oamEnd = 0xFE9F
const hramStart = 0xFF80
const hramEnd = 0xFFFE

type Memory struct {
	Rom  rom.Rom
	Vram vram.Vram
	Wram wram.Wram
	Hram hram.Hram
}

func NewMemory(rom rom.Rom) Memory {
	return Memory{
		Rom:  rom,
		Vram: vram.NewVram(),
		Wram: wram.NewWram(),
		Hram: hram.NewHram(),
	}
}

func (m Memory) Read(address uint16) uint8 {
	switch {
	case address >= romBankStart && address <= romBankEnd:
		return m.Rom.Read(address)
	case address >= vramStart && address <= vramEnd:
		return m.Vram.Read(address)
	case address >= eramStart && address <= eramEnd:
		return m.Rom.Read(address)
	case address >= wramStart && address <= wramEnd:
		return m.Wram.Read(address)
	case address >= echoStart && address <= echoEnd:
		return m.Wram.Read(address)
	case address >= oamStart && address <= oamEnd:
		return m.Vram.Read(address)
	case address >= hramStart && address <= hramEnd:
		return m.Hram.Read(address)
	}

	panic("invalid read!")
}

func (m Memory) Write(address uint16, data uint8) {
	switch {
	case address >= vramStart && address <= vramEnd:
		m.Vram.Write(address, data)
	case address >= wramStart && address <= wramEnd:
		m.Wram.Write(address, data)
	case address >= echoStart && address <= echoEnd:
		m.Wram.Write(address, data)
	case address >= oamStart && address <= oamEnd:
		m.Vram.Write(address, data)
	case address >= hramStart && address <= hramEnd:
		m.Hram.Write(address, data)
	}

	panic("invalid read!")
}
