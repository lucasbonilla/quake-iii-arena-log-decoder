package config

type API struct {
	RunType string
	Workers int
}

func NewAPIConfig(runType string, workers int) *API {
	return &API{
		RunType: runType,
		Workers: workers,
	}
}
