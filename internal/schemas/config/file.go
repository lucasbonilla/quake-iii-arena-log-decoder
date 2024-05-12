package config

type File struct {
	Path string
}

func NewFileConfig(path string) *File {
	return &File{
		Path: path,
	}
}
