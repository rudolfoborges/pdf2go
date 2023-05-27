package text

import (
	"fmt"
	"os/exec"
)

type TextExtractor struct {
	path string
}

// NewTextExtractor creates a new TextExtractor.
// The path argument is the path to the PDF file.
func NewTextExtractor(path string) *TextExtractor {
	return &TextExtractor{
		path: path,
	}
}

// Extract extracts the text from the PDF file.
// It returns an error if the text cannot be extracted.
// The pageNumber argument is the page number to extract the text from.
func (e *TextExtractor) Extract(pageNumber int) (string, error) {
	cmd := exec.Command(
		"pdftotext",
		"-f", fmt.Sprintf("%d", pageNumber),
		"-l", fmt.Sprintf("%d", pageNumber),
		e.path,
		"-",
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
