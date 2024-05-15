package file

import (
	"bufio"
	"os"
)

type Adapter struct {
	file *os.File
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (oA *Adapter) Set(file *os.File) {
	oA.file = file
}

func (oA *Adapter) Open(path string) (*os.File, error) {
	var err error
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (oA *Adapter) GetFile() *os.File {
	return oA.file
}

func (oA *Adapter) Close() error {
	return oA.file.Close()
}

func (oA *Adapter) Scanner() *bufio.Scanner {
	return bufio.NewScanner(oA.file)
}
