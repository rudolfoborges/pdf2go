package pdf2go_test

import (
	"strings"
	"testing"

	"github.com/rudolfoborges/pdf2go"
)

func TestPDFReader(t *testing.T) {
	t.Run("should return the number of pages from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		pdfReader, err := pdf2go.NewPDFReader(path, pdf2go.NewDefaultLogger("debug"))

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if pdfReader.PagesNumber() != 5 {
			t.Errorf("Expected 5 pages, got %d", pdfReader.PagesNumber())
		}
	})

	t.Run("should return title from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		pdfReader, err := pdf2go.NewPDFReader(path, pdf2go.NewDefaultLogger("debug"))

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if pdfReader.Title() != "Test" {
			t.Errorf("Expected title 'Test', got %s", pdfReader.Title())
		}
	})

	t.Run("should return the text from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		pdfReader, err := pdf2go.NewPDFReader(path, pdf2go.NewDefaultLogger("debug"))

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		text, err := pdfReader.Text()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.Contains(text, "Page 1") {
			t.Errorf("Expected 'Page 1', got %v", text)
		}
	})

	t.Run("should return the html from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		pdfReader, err := pdf2go.NewPDFReader(path, pdf2go.NewDefaultLogger("debug"))

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		html, err := pdfReader.Html()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.Contains(html, "Page 1") {
			t.Errorf("Expected 'Page 1', got %v", html)
		}
	})

	t.Run("should return an error if the path is invalid", func(t *testing.T) {
		path := "test_files/invalid.txt"
		_, err := pdf2go.NewPDFReader(path, pdf2go.NewDefaultLogger("debug"))

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}

		if err != pdf2go.ErrInvalidPath {
			t.Errorf("Expected %v, got %v", pdf2go.ErrInvalidPath, err)
		}
	})
}
