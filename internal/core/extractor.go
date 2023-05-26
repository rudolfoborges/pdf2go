package core

type Extractor interface {
	Extract(pageNumber int) (string, error)
}
