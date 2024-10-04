package wram

// 32 KB
const wramSize = 32000

type Wram struct {
	bytes [wramSize]byte
}

func NewWram() Wram {
	return Wram{
		bytes: [wramSize]byte{},
	}
}

func (w *Wram) Read(address uint16) uint8 {
	return w.bytes[address]
}

func (w *Wram) Write(address uint16, data uint8) {
	w.bytes[address] = data
}
