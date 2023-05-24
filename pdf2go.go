package pdf2go

import (
	"github.com/rudolfoborges/pdf2go/internal/core"
	"github.com/rudolfoborges/pdf2go/pkg/model"
)

func New(options *model.Options) (model.PDFReader, error) {
	pdfReader, err := core.New()

	if err != nil {
		return nil, err
	}

	return pdfReader, nil
}
