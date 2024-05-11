package ports

import "os"

type File interface {
	Open(path string) error
	GetFile() *os.File
	Close() error
}
