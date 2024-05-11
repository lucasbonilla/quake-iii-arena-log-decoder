package file

import "os"

type Adapter struct {
	file *os.File
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (oA *Adapter) Open(path string) error {
	var err error
	oA.file, err = os.Open(path)
	if err != nil {
		return err
	}
	return nil
}

func (oA *Adapter) GetFile() *os.File {
	return oA.file
}

func (oA *Adapter) Close() error {
	return oA.file.Close()
}
