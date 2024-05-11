package config

type File struct {
	PathDEV string
	PathPRD string
}

func NewFileConfig(pathDev string, pathPrd string) *File {
	return &File{
		PathDEV: pathDev,
		PathPRD: pathPrd,
	}
}
