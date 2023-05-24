package pdfinfo

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

type PDFInfoExtractor struct {
}

type PDFInfo struct {
	PagesNumber int
}

func NewInfoExtractor() *PDFInfoExtractor {
	return &PDFInfoExtractor{}
}

func (e *PDFInfoExtractor) Extract(path string) (*PDFInfo, error) {
	cmd := exec.Command("pdfinfo", path)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	text := string(output)
	regex := regexp.MustCompile(`Pages:\s+(\d+)`)
	match := regex.FindStringSubmatch(text)

	if len(match) > 1 {
		pagesNumber, err := strconv.Atoi(match[1])

		if err != nil {
			return nil, err
		}

		return &PDFInfo{
			PagesNumber: pagesNumber,
		}, nil
	}

	return nil, errors.New("Page number not found")
}
