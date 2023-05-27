package html

import (
	"fmt"
	"os/exec"
)

type HtmlExtractor struct {
	path         string
	totalOfPages int
}

// NewHtmlExtractor creates a new HtmlExtractor.
// The path argument is the path to the PDF file.
// The totalOfPages argument is the total of pages of the PDF file.
func NewHtmlExtractor(path string, totalOfPages int) *HtmlExtractor {
	return &HtmlExtractor{
		path:         path,
		totalOfPages: totalOfPages,
	}
}

// Extract extracts the html from the PDF file.
// It returns an error if the html cannot be extracted.
func (e *HtmlExtractor) Extract() (string, error) {
	if e.totalOfPages < 1 {
		return "", fmt.Errorf("invalid total of pages %d", e.totalOfPages)
	}

	return e.extract(1, e.totalOfPages)
}

// Extract extracts the text from the PDF file.
// It returns an error if the html cannot be extracted.
// The pageNumber argument is the page number to extract the html from.
func (e *HtmlExtractor) ExtractPage(pageNumber int) (string, error) {
	if pageNumber < 1 {
		return "", fmt.Errorf("invalid page number %d", pageNumber)
	}

	return e.extract(pageNumber, pageNumber)
}

func (e *HtmlExtractor) extract(firstPage, lastPage int) (string, error) {
	cmd := exec.Command(
		"pdftohtml",
		"-f", fmt.Sprintf("%d", firstPage),
		"-l", fmt.Sprintf("%d", lastPage),
		"-s", "-i", "-stdout",
		e.path,
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
