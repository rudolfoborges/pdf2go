package pdf2go

import "github.com/rudolfoborges/pdf2go/internal/core"

// Page represents a page from the PDF file.
type Page struct {
	logger        Logger
	Number        int
	TextExtractor core.Extractor
	HtmlExtractor core.Extractor
}

// NewPage creates a new Page.
func NewPage(number int, textExtractor, htmlExtractor core.Extractor, logger Logger) *Page {
	return &Page{
		logger:        logger,
		Number:        number,
		TextExtractor: textExtractor,
		HtmlExtractor: htmlExtractor,
	}
}

// Text returns the text from the PDF file.
// It returns an error if the text cannot be extracted.
func (p *Page) Text() (string, error) {
	p.logger.Debugf("Extracting text from page %d", p.Number)
	return p.TextExtractor.ExtractPage(p.Number)
}

// Html returns the html from the PDF file.
// It returns an error if the html cannot be extracted.
func (p *Page) Html() (string, error) {
	p.logger.Debugf("Extracting html from page %d", p.Number)
	return p.HtmlExtractor.ExtractPage(p.Number)
}
