package scanner

import "bufio"

type Adapter struct {
	scanner *bufio.Scanner
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (bA *Adapter) SetScanner(scanner *bufio.Scanner) {
	bA.scanner = scanner
}

func (bA *Adapter) Scan() bool {
	return bA.scanner.Scan()
}

func (bA *Adapter) Text() string {
	return bA.scanner.Text()
}

func (bA *Adapter) Err() error {
	return bA.scanner.Err()
}
