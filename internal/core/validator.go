package core

import "regexp"

// ValidatePDFPath validates if the path is a valid PDF path.
// It returns true if the path is valid.
// Example of valid paths:
// - /path/to/file.pdf
// - ./path/to/file.pdf
// - path/to/file.pdf
// - file.pdf
func ValidatePDFPath(path string) bool {
	regex := `^(?:\/|\.\/)?(?:[^/\0]+\/)*[^/\0]+\.(?:pdf|PDF)$`
	match, _ := regexp.MatchString(regex, path)
	return match
}
