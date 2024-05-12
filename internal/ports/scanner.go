package ports

import "bufio"

type Scanner interface {
	SetScanner(scanner *bufio.Scanner)
	Scan() bool
	Text() string
	Err() error
}
