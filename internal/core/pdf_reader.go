package core

type PDFReader struct {
}

func NewPDFReader() *PDFReader {
	return &PDFReader{}
}

func (r *PDFReader) Open(path string) error {
	return nil
}

func (r *PDFReader) PagesNumber() (int, error) {
	return 0, nil
}

func (r *PDFReader) Pages(pageNum int) ([]Page, error) {
	return nil, nil
}

func (r *PDFReader) Close() error {
	return nil
}
