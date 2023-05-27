package html

import (
	"fmt"
	"os/exec"
)

type HtmlExtractor struct {
	path string
}

// NewHtmlExtractor creates a new HtmlExtractor.
// The path argument is the path to the PDF file.
func NewHtmlExtractor(path string) *HtmlExtractor {
	return &HtmlExtractor{
		path: path,
	}
}

// Extract extracts the text from the PDF file.
// It returns an error if the text cannot be extracted.
// The pageNumber argument is the page number to extract the text from.
func (e *HtmlExtractor) Extract(pageNumber int) (string, error) {
	cmd := exec.Command(
		"pdftohtml",
		"-f", fmt.Sprintf("%d", pageNumber),
		"-l", fmt.Sprintf("%d", pageNumber),
		"-s", "-i", "-stdout",
		e.path,
	)

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
