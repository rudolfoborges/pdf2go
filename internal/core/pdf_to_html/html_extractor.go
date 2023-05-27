package html

type HtmlExtractor struct {
	path string
}

func NewHtmlExtractor(path string) *HtmlExtractor {
	return &HtmlExtractor{
		path: path,
	}
}

func (e *HtmlExtractor) Extract(pageNumber int) (string, error) {
	return "", nil
}
