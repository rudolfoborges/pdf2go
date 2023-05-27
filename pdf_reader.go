package pdf2go

import (
	"errors"

	info "github.com/rudolfoborges/pdf2go/internal/core/pdf_info"
	html "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_html"
	text "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_text"
)

type PDFReader struct {
	pages []*Page
}

var (
	ErrInvalidPath        = errors.New("Invalid path")
	ErrInvalidPagesNumber = errors.New("Invalid pages number")
)

// NewPDFReader creates a new PDFReader.
// It returns an error if the PDFReader cannot be created.
// The path argument is the path to the PDF file.
func NewPDFReader(path string) (*PDFReader, error) {
	pdfInfo, err := info.Extract(path)

	if err != nil {
		return nil, err
	}

	if pdfInfo.PagesNumber == 0 {
		return nil, ErrInvalidPagesNumber
	}

	pages := make([]*Page, pdfInfo.PagesNumber)

	textExtractor := text.NewTextExtractor(path)
	htmlExtractor := html.NewHtmlExtractor(path)

	for i := 0; i < pdfInfo.PagesNumber; i++ {
		pageNumber := i + 1
		pages[i] = NewPage(pageNumber, textExtractor, htmlExtractor)
	}

	return &PDFReader{
		pages: pages,
	}, nil
}

// PagesNumber returns the number of pages in the PDF file
// associated with the PDFReader.
func (r *PDFReader) PagesNumber() int {
	return len(r.pages)
}

// Pages returns a slice of pointers to Page structs.
// Each Page struct contains the page number and the
// text and html extractors.
func (r *PDFReader) Pages() ([]*Page, error) {
	return r.pages, nil
}
