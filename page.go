package pdf2go

import "github.com/rudolfoborges/pdf2go/internal/core"

type Page struct {
	Number        int
	TextExtractor core.Extractor
	HtmlExtractor core.Extractor
}

// NewPage creates a new Page.
func NewPage(number int, textExtractor, htmlExtractor core.Extractor) *Page {
	return &Page{
		Number:        number,
		TextExtractor: textExtractor,
		HtmlExtractor: htmlExtractor,
	}
}

// Text returns the text from the PDF file.
// It returns an error if the text cannot be extracted.
func (p *Page) Text() (string, error) {
	return p.TextExtractor.Extract(p.Number)
}

// Html returns the html from the PDF file.
// It returns an error if the html cannot be extracted.
func (p *Page) Html() (string, error) {
	return p.HtmlExtractor.Extract(p.Number)
}
