package core

import (
	"path/filepath"
	"regexp"
	"strings"
)

// ValidatePDFPath validates if the path is a valid PDF path.
// It returns true if the path is valid.
// Example of valid paths:
// - /path/to/file.pdf
// - ./path/to/file.pdf
// - path/to/file.pdf
// - file.pdf
//
// Note: Callers that accept user-supplied paths should perform
// additional validation (e.g., restrict to a base directory) to
// prevent path traversal attacks.
func ValidatePDFPath(path string) bool {
	// Reject paths containing null bytes to prevent null byte injection.
	// Null bytes can cause path truncation in C-based tools like poppler.
	if strings.ContainsRune(path, '\x00') {
		return false
	}

	// Normalize the path to resolve redundant separators and "." components,
	// ensuring the validated path matches what the OS will actually resolve.
	cleaned := filepath.Clean(path)

	regex := `^(?:\/|\.\/)?(?:[^/\0]+\/)*[^/\0]+\.(?:pdf|PDF)$`
	match, _ := regexp.MatchString(regex, cleaned)
	return match
}
