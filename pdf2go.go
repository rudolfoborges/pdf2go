package pdf2go

import (
	"github.com/rudolfoborges/pdf2go/internal/core"
)

type Options struct{}

func New(options Options) (*core.PDFReader, error) {
	pdfReader := core.NewPDFReader()
	return pdfReader, nil
}
