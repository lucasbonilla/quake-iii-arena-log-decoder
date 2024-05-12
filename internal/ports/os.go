package ports

import (
	"bufio"
	"os"
)

type Os interface {
	Open(path string) error
	GetFile() *os.File
	Close() error
	Scanner() *bufio.Scanner
	SetScanner(scanner *bufio.Scanner)
	Scan() bool
	Text() string
	Err() error
}
