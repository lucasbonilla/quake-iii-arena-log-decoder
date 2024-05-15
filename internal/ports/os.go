package ports

import (
	"bufio"
	"os"
)

type Os interface {
	SetFile(file *os.File)
	OpenFile(path string) (*os.File, error)
	GetFile() *os.File
	CloseFile() error
	Scanner() *bufio.Scanner
	SetScanner(scanner *bufio.Scanner)
	Scan() bool
	Text() string
	Err() error
	Create(fileName string) (*os.File, error)
}
