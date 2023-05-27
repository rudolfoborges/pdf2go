package text

import (
	"fmt"
	"os/exec"
)

type TextExtractor struct {
	path         string
	totalOfPages int
}

// NewTextExtractor creates a new TextExtractor.
// The path argument is the path to the PDF file.
// The totalOfPages argument is the total of pages of the PDF file.
func NewTextExtractor(path string, totalOfPages int) *TextExtractor {
	return &TextExtractor{
		path:         path,
		totalOfPages: totalOfPages,
	}
}

// Extract extracts the text from the PDF file.
// It returns an error if the text cannot be extracted.
func (e *TextExtractor) Extract() (string, error) {
	if e.totalOfPages < 1 {
		return "", fmt.Errorf("invalid total of pages %d", e.totalOfPages)
	}

	return e.extract(1, e.totalOfPages)
}

// Extract extracts the text from the PDF file.
// It returns an error if the text cannot be extracted.
// The pageNumber argument is the page number to extract the text from.
func (e *TextExtractor) ExtractPage(pageNumber int) (string, error) {
	if pageNumber < 1 {
		return "", fmt.Errorf("invalid page number %d", pageNumber)
	}

	return e.extract(pageNumber, pageNumber)
}

func (e *TextExtractor) extract(firstPage, lastPage int) (string, error) {
	cmd := exec.Command(
		"pdftotext",
		"-f", fmt.Sprintf("%d", firstPage),
		"-l", fmt.Sprintf("%d", lastPage),
		e.path,
		"-",
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
