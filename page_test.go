package pdf2go_test

import (
	"strings"
	"testing"

	"github.com/rudolfoborges/pdf2go"
	html "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_html"
	text "github.com/rudolfoborges/pdf2go/internal/core/pdf_to_text"
)

func TestPage(t *testing.T) {
	t.Run("should return the text from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		textExtractor := text.NewTextExtractor(path, 5)
		htmlExtractor := html.NewHtmlExtractor(path, 5)

		page := pdf2go.NewPage(2, textExtractor, htmlExtractor, pdf2go.NewDefaultLogger("debug"))

		text, err := page.Text()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.Contains(text, "Page 2") {
			t.Errorf("Expected 'Page 2', got %v", text)
		}
	})

	t.Run("should return the html from the PDF file", func(t *testing.T) {
		path := "test_files/test.pdf"
		textExtractor := text.NewTextExtractor(path, 5)
		htmlExtractor := html.NewHtmlExtractor(path, 5)

		page := pdf2go.NewPage(2, textExtractor, htmlExtractor, pdf2go.NewDefaultLogger("debug"))

		html, err := page.Html()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !strings.Contains(html, "Page 2") {
			t.Errorf("Expected 'Page 2', got %v", html)
		}
	})
}
