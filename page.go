package pdf2go

import "github.com/rudolfoborges/pdf2go/internal/core"

type Page struct {
	Number        int
	TextExtractor core.Extractor
	HtmlExtractor core.Extractor
}

func NewPage(number int, textExtractor, htmlExtractor core.Extractor) *Page {
	return &Page{
		Number:        number,
		TextExtractor: textExtractor,
		HtmlExtractor: htmlExtractor,
	}
}

func (p *Page) Text() (string, error) {
	return p.TextExtractor.Extract(p.Number)
}

func (p *Page) Html() (string, error) {
	return p.HtmlExtractor.Extract(p.Number)
}
