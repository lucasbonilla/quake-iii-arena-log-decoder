package ports

import (
	"bufio"
	"os"
)

type File interface {
	Set(file *os.File)
	Open(path string) (*os.File, error)
	GetFile() *os.File
	Close() error
	Scanner() *bufio.Scanner
}
