package model

type Options struct{}

type Page interface {
	Text() (string, error)
	Html() (string, error)
}

type PDFReader interface {
	Open() error
	Close() error
	Pages() ([]Page, error)
	PagesNumber() (int, error)
}
