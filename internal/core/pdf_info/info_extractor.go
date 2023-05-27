package info

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	ErrGetPDFInfo = errors.New("failed to get pdf info")
)

type PDFInfo struct {
	Author      string
	PagesNumber int
	Encrypted   bool
}

// Extract extracts the number of pages from the PDF file.
// It returns an error if the number of pages cannot be extracted.
// The path argument is the path to the PDF file.
func Extract(path string) (*PDFInfo, error) {
	cmd := exec.Command("pdfinfo", path)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	text := string(output)
	pageNumber, err := findInfo(regexp.MustCompile(`Pages:\s+(\d+)`), text)

	if err != nil {
		return nil, err
	}
	pagesNumber, _ := strconv.Atoi(pageNumber)

	author, _ := findInfo(regexp.MustCompile(`Author:\s+(\w+)`), text)
	encrypted, _ := findInfo(regexp.MustCompile(`Encrypted:\s+(\w+)`), text)

	return &PDFInfo{
		Author:      author,
		PagesNumber: pagesNumber,
		Encrypted:   encrypted == "yes",
	}, nil
}

func findInfo(regex *regexp.Regexp, text string) (string, error) {
	match := regex.FindStringSubmatch(text)

	if len(match) > 1 {
		return match[1], nil
	}

	return "", ErrGetPDFInfo
}
