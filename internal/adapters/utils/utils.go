package utils

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (uA *Adapter) PlayerExists(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}
