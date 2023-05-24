package core

type Extractor interface {
	Extract() (string, error)
}
