package core

type PDFReader struct {
	pages []Page
}

func NewPDFReader() *PDFReader {
	return &PDFReader{}
}

func (r *PDFReader) Open() error {
	return nil
}

func (r *PDFReader) PagesNumber() int {
	return len(r.pages)
}

func (r *PDFReader) Pages(pageNum int) ([]Page, error) {
	return r.pages, nil
}

func (r *PDFReader) Close() error {
	return nil
}
