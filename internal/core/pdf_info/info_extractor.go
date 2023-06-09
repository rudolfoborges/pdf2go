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
	Title        string
	Author       string
	PagesNumber  int
	Encrypted    bool
	CreationDate string
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

	title, _ := findInfo(regexp.MustCompile(`Title:\s+(\w+)`), text)
	author, _ := findInfo(regexp.MustCompile(`Author:\s+(\w+)`), text)
	encrypted, _ := findInfo(regexp.MustCompile(`Encrypted:\s+(\w+)`), text)
	creationDate, _ := findInfo(regexp.MustCompile(`CreationDate:\s+(\w+)`), text)

	return &PDFInfo{
		Title:        title,
		Author:       author,
		PagesNumber:  pagesNumber,
		Encrypted:    encrypted == "yes",
		CreationDate: creationDate,
	}, nil
}

func findInfo(regex *regexp.Regexp, text string) (string, error) {
	match := regex.FindStringSubmatch(text)

	if len(match) > 1 {
		return match[1], nil
	}

	return "", ErrGetPDFInfo
}
