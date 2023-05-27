package pdf2go

type Options struct {
}

// New creates a new PDFReader.
// It returns an error if the PDFReader cannot be created.
// The path argument is the path to the PDF file.
func New(path string, options *Options) (*PDFReader, error) {
	return NewPDFReader(path)
}
