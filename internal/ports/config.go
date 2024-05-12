package ports

type Config interface {
	RunType() string
	FilePath() string
}
