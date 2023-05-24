package core

type Page struct {
	Number int
}

func NewPage(number int) *Page {
	return &Page{
		Number: number,
	}
}

func (p *Page) Text() (string, error) {
	return "", nil
}

func (p *Page) Html() (string, error) {
	return "", nil
}
