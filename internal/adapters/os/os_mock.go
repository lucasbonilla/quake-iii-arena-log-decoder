package os

import (
	"bufio"
	"io/fs"
	"os"
)

type MockedAdapter struct {
	SetFileFn    func(file *os.File)
	OpenFileFn   func(path string) (*os.File, error)
	GetFileFn    func() *os.File
	CloseFileFn  func() error
	ScannerFn    func() *bufio.Scanner
	SetScannerFn func(scanner *bufio.Scanner)
	ScanFn       func() bool
	TextFn       func() string
	ErrFn        func() error
	CreateFn     func(fileName string) (*os.File, error)
	StatFn       func(filePath string) (fs.FileInfo, error)
	IsNotExistFn func(err error) bool
	MkdirFn      func(name string, perm fs.FileMode) error
}

func (mA *MockedAdapter) SetFile(file *os.File) {}

func (mA *MockedAdapter) OpenFile(path string) (*os.File, error) {
	return mA.OpenFileFn(path)
}

func (mA *MockedAdapter) GetFile() *os.File {
	return mA.GetFileFn()
}

func (mA *MockedAdapter) CloseFile() error {
	return mA.CloseFileFn()
}

func (mA *MockedAdapter) Scanner() *bufio.Scanner {
	return mA.ScannerFn()
}

func (mA *MockedAdapter) SetScanner(scanner *bufio.Scanner) {}

func (mA *MockedAdapter) Scan() bool {
	return mA.ScanFn()
}

func (mA *MockedAdapter) Text() string {
	return mA.TextFn()
}

func (mA *MockedAdapter) Err() error {
	return mA.ErrFn()
}

func (mA *MockedAdapter) Create(fileName string) (*os.File, error) {
	return mA.CreateFn(fileName)
}

func (mA *MockedAdapter) Stat(filePath string) (fs.FileInfo, error) {
	return mA.StatFn(filePath)
}

func (mA *MockedAdapter) IsNotExist(err error) bool {
	return mA.IsNotExistFn(err)
}

func (mA *MockedAdapter) Mkdir(name string, perm fs.FileMode) error {
	return mA.MkdirFn(name, perm)
}
