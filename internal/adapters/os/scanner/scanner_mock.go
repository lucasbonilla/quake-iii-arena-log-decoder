package scanner

import "bufio"

type MockedAdapter struct {
	Lines        []string
	SetScannerFn func(scanner *bufio.Scanner)
	ScanFn       func() bool
	TextFn       func() string
	ErrFn        func() error
	cursor       int
}

func (mA *MockedAdapter) SetScanner(scanner *bufio.Scanner) {}

func (mA *MockedAdapter) Scan() bool {
	if mA.cursor >= len(mA.Lines) {
		return false
	}
	mA.cursor++
	return true
}

func (mA *MockedAdapter) Text() string {
	return mA.Lines[mA.cursor-1]
}

func (mA *MockedAdapter) Err() error {
	return mA.ErrFn()
}
