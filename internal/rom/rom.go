package rom

import (
	"os"
)

type Rom struct {
	bytes []byte
}

func (w *Rom) Read(address uint16) uint8 {
	return w.bytes[address]
}

func LoadRom(path string) (*Rom, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &Rom{bytes: fileBytes}, nil
}
