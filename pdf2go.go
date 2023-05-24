package pdf2go

import (
	"context"
)

type Options struct{}

type Page interface{}

type PDFReader interface {
	Open(path string) error
	GetNumPages() (int, error)
	GetPage(pageNum int) (*Page, error)
	Close() error
}

func New(ctx context.Context, options *Options) (PDFReader, error) {
	return nil, nil
}
