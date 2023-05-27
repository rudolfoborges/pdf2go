package core

type Extractor interface {
	Extract() (string, error)
	ExtractPage(pageNumber int) (string, error)
}
