package api

type PDFReader interface {
	Open(string) error
	GetNumPages() (int, error)
	GetPage(int) (Page, error)
}
