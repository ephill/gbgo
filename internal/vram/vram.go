package vram

// 16 KB
const vramSize = 16000

type Vram struct {
	bytes [vramSize]byte
}

func NewVram() Vram {
	return Vram{
		bytes: [vramSize]byte{},
	}
}

func (w Vram) Read(address uint16) uint8 {
	return w.bytes[address]
}

func (w Vram) Write(address uint16, data uint8) {
	w.bytes[address] = data
}
