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

	for i := 1; i <= pdfInfo.PagesNumber; i++ {
		pages = append(pages, NewPage(i, textExtractor, htmlExtractor))
	}

	return &PDFReader{
		pages: pages,
	}, nil
}

func (r *PDFReader) PagesNumber() int {
	return len(r.pages)
}

func (r *PDFReader) Pages() ([]*Page, error) {
	return r.pages, nil
}
