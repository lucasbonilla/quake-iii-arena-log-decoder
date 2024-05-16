package os

import (
	"bufio"
	"io/fs"
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

func (oA *Adapter) SetFile(file *os.File) {
	oA.file.Set(file)
}

func (oA *Adapter) OpenFile(path string) (*os.File, error) {
	return oA.file.Open(path)
}

func (oA *Adapter) GetFile() *os.File {
	return oA.file.GetFile()
}

func (oA *Adapter) CloseFile() error {
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

func (oA *Adapter) Create(fileName string) (*os.File, error) {
	return os.Create(fileName)
}

func (oA *Adapter) Stat(filePath string) (fs.FileInfo, error) {
	return os.Stat(filePath)
}

func (oA *Adapter) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (oA *Adapter) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}
