package ports

import "os"

type Os interface {
	Open(path string) error
	GetFile() *os.File
	Close() error
}
