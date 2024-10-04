package hram

// TODO (EvanP): What is the right size for this?
const hramSize = 32000

type Hram struct {
	bytes [hramSize]byte
}

func NewHram() Hram {
	return Hram{
		bytes: [hramSize]byte{},
	}
}

func (w Hram) Read(address uint16) uint8 {
	return w.bytes[address]
}

func (w Hram) Write(address uint16, data uint8) {
	w.bytes[address] = data
}
