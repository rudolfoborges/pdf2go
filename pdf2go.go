package pdf2go

import (
	"context"

	"github.com/rudolfoborges/pdf2go/internal/core"
	"github.com/rudolfoborges/pdf2go/pkg/model"
)

func New(ctx context.Context, options *model.Options) (model.PDFReader, error) {
	pdfReader, err := core.New(ctx)

	if err != nil {
		return nil, err
	}

	return pdfReader, nil
}
