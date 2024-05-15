package config

type File struct {
	PathInDev   string
	PathInProd  string
	PathOutDev  string
	PathOutProd string
}

func NewFileConfig(pathInDev string, pathInProd string, pathOutDev string, pathOutProd string) *File {
	return &File{
		PathInDev:   pathInDev,
		PathInProd:  pathInProd,
		PathOutDev:  pathOutDev,
		PathOutProd: pathOutProd,
	}
}
