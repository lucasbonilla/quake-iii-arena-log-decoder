package os

import (
	"os"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
)

type Adapter struct {
	file ports.File
}

func NewAdapter(fileP ports.File) *Adapter {
	return &Adapter{
		file: fileP,
	}
}

func (oA *Adapter) Open(path string) error {
	return oA.file.Open(path)
}
func (oA *Adapter) GetFile() *os.File {
	return oA.file.GetFile()
}
func (oA *Adapter) Close() error {
	return oA.file.Close()
}
