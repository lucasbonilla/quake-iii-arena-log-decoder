package config

type API struct {
	RunType string
}

func NewAPIConfig(runType string) *API {
	return &API{
		RunType: runType,
	}
}
