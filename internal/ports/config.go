package ports

type Config interface {
	RunType() string
	FileInPath() string
	FileOutPath() string
	GetNumOfWorkers() int
}
