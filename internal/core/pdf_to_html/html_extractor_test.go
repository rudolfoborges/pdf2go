package html_test

import (
	"testing"

	html "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_html"
)

func TestTextExtractor(t *testing.T) {
	t.Run("should return text from pdf", func(t *testing.T) {
		path := "../../../test_files/test.pdf"
		extractor := html.NewHtmlExtractor(path, 5)

		html, err := extractor.Extract()

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if html == "" {
			t.Errorf("Expected no empty html, got %s", html)
		}
	})

	t.Run("should return error when file is not pdf", func(t *testing.T) {
		path := "../../../test_files/test.txt"
		extractor := html.NewHtmlExtractor(path, 5)

		_, err := extractor.Extract()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
