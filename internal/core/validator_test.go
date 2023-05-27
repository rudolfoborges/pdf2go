package core_test

import (
	"testing"

	"github.com/rudolfoborges/pdf2go/internal/core"
)

func TestValidator(t *testing.T) {
	t.Run("should return path valid", func(t *testing.T) {
		path := "../../test_files/test.pdf"
		isValid := core.ValidatePDFPath(path)

		if isValid != true {
			t.Errorf("Expected true, got %v", isValid)
		}
	})

	t.Run("should return path invalid", func(t *testing.T) {
		path := "../../test_files/test.txt"
		isValid := core.ValidatePDFPath(path)

		if isValid != false {
			t.Errorf("Expected false, got %v", isValid)
		}
	})
}
