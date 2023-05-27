package info_test

import (
	"testing"

	info "github.com/rudolfoborges/pdf2go/internal/core/pdf_info"
)

func TestInfoExtract(t *testing.T) {
	t.Run("should return the pdf info", func(t *testing.T) {
		path := "../../../test_files/test.pdf"
		pdfInfo, err := info.Extract(path)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if pdfInfo.PagesNumber != 5 {
			t.Errorf("Expected 5 page, got %d", pdfInfo.PagesNumber)
		}

		if pdfInfo.Title != "Test" {
			t.Errorf("Expected title 'Test PDF PDF2Go', got %s", pdfInfo.Title)
		}
	})

	t.Run("should return pdf info with no encrypted file", func(t *testing.T) {
		path := "../../../test_files/test.pdf"
		pdfInfo, err := info.Extract(path)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if pdfInfo.Encrypted != false {
			t.Errorf("Expected no encrypted file, got %v", pdfInfo.Encrypted)
		}
	})

	t.Run("should return pdf info with encrypted file", func(t *testing.T) {
		path := "../../../test_files/encrypted.pdf"
		pdfInfo, err := info.Extract(path)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if pdfInfo.Encrypted != true {
			t.Errorf("Expected encrypted file, got %v", pdfInfo.Encrypted)
		}
	})
}
