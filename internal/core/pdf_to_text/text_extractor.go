package text

type TextExtractor struct {
	path string
}

func NewTextExtractor(path string) *TextExtractor {
	return &TextExtractor{
		path: path,
	}
}

func (e *TextExtractor) Extract(pageNumber int) (string, error) {
	return "", nil
}
