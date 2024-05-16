package file

import (
	"bufio"
	"os"
)

type MockedAdapter struct {
	SetFn     func(file *os.File)
	OpenFn    func(path string) (*os.File, error)
	GetFileFn func() *os.File
	CloseFn   func() error
	ScannerFn func() *bufio.Scanner
}

func (mA *MockedAdapter) Set(file *os.File) {
}

func (mA *MockedAdapter) Open(path string) (*os.File, error) {
	return mA.OpenFn(path)
}

func (mA *MockedAdapter) GetFile() *os.File {
	return mA.GetFileFn()
}

func (mA *MockedAdapter) Close() error {
	return mA.CloseFn()
}

func (mA *MockedAdapter) Scanner() *bufio.Scanner {
	return mA.ScannerFn()
}
