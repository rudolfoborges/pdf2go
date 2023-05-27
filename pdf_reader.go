package pdf2go

import (
	"errors"

	"github.com/rudolfoborges/pdf2go/internal/core"
	info "github.com/rudolfoborges/pdf2go/internal/core/pdf_info"
	html "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_html"
	text "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_text"
)

// PDFReader represents a PDF file.
type PDFReader struct {
	logger        Logger
	pages         []*Page
	info          *info.PDFInfo
	textExtractor core.Extractor
	htmlExtractor core.Extractor
}

var (
	ErrInvalidPath        = errors.New("Invalid pdf path")
	ErrInvalidPagesNumber = errors.New("Invalid pages number. It must be greater than 0")
)

// NewPDFReader creates a new PDFReader.
// It returns an error if the PDFReader cannot be created.
// The path argument is the path to the PDF file.
func NewPDFReader(path string, logger Logger) (*PDFReader, error) {
	if !core.ValidatePDFPath(path) {
		return nil, ErrInvalidPath
	}

	pdfInfo, err := info.Extract(path)

	if err != nil {
		return nil, err
	}

	if pdfInfo.PagesNumber == 0 {
		return nil, ErrInvalidPagesNumber
	}

	logger.Debugf("The PDF has %d pages", pdfInfo.PagesNumber)

	pages := make([]*Page, pdfInfo.PagesNumber)

	textExtractor := text.NewTextExtractor(path, pdfInfo.PagesNumber)
	htmlExtractor := html.NewHtmlExtractor(path, pdfInfo.PagesNumber)

	for i := 0; i < pdfInfo.PagesNumber; i++ {
		pageNumber := i + 1
		pages[i] = NewPage(pageNumber, textExtractor, htmlExtractor, logger)
	}

	return &PDFReader{
		info:          pdfInfo,
		pages:         pages,
		textExtractor: textExtractor,
		htmlExtractor: htmlExtractor,
	}, nil
}

// Text returns the text of all pages from the PDF file.
// It returns an error if the text cannot be extracted.
func (r *PDFReader) Text() (string, error) {
	r.logger.Debugf("Extracting text from PDF file")
	return r.textExtractor.Extract()
}

// Html returns the html of all pages from the PDF file.
// It returns an error if the html cannot be extracted.
func (r *PDFReader) Html() (string, error) {
	r.logger.Debugf("Extracting html from PDF file")
	return r.textExtractor.Extract()
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

// Title returns the title of the PDF file or an empty string
// if the title is not defined.
func (r *PDFReader) Title() string {
	return r.info.Title
}

// Author returns the author of the PDF file or an empty string
// if the author is not defined.
func (r *PDFReader) Author() string {
	return r.info.Author
}

// Encrypted returns true if the PDF file is encrypted.
func (r *PDFReader) Encrypted() bool {
	return r.info.Encrypted
}

// CreationDate returns the creation date of the PDF file or an empty string
// if the creation date is not defined.
func (r *PDFReader) CreationDate() string {
	return r.info.CreationDate
}
