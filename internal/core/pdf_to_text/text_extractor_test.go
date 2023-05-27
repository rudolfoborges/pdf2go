package text_test

import (
	"testing"

	text "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_text"
)

func TestTextExtractor(t *testing.T) {
	t.Run("should return text from pdf", func(t *testing.T) {
		path := "../../../test_files/test.pdf"
		extractor := text.NewTextExtractor(path, 5)

		txt, err := extractor.Extract()

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if txt == "" {
			t.Errorf("Expected no empty text, got %s", txt)
		}
	})

	t.Run("should return error when file is not pdf", func(t *testing.T) {
		path := "../../../test_files/test.txt"
		extractor := text.NewTextExtractor(path, 5)

		_, err := extractor.Extract()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
