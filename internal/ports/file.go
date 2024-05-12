package ports

import (
	"bufio"
	"os"
)

type File interface {
	Open(path string) error
	GetFile() *os.File
	Close() error
	Scanner() *bufio.Scanner
}
