package os

import (
	"bufio"
	"os"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
)

type Adapter struct {
	file    ports.File
	scanner ports.Scanner
}

func NewAdapter(fileP ports.File, scannerP ports.Scanner) *Adapter {
	return &Adapter{
		file:    fileP,
		scanner: scannerP,
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

func (oA *Adapter) Scanner() *bufio.Scanner {
	return oA.file.Scanner()
}

func (oA *Adapter) SetScanner(scanner *bufio.Scanner) {
	oA.scanner.SetScanner(scanner)
}

func (oA *Adapter) Scan() bool {
	return oA.scanner.Scan()
}

func (oA *Adapter) Text() string {
	return oA.scanner.Text()
}

func (oA *Adapter) Err() error {
	return oA.scanner.Err()
}
