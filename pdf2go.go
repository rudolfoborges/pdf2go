package pdf2go

// Config is the configuration used by the PDFReader.
// It is used to configure the logger and the log level.
type Config struct {
	// Logger is the logger used by the PDFReader.
	// If nil, the DefaultLogger is used.
	Logger Logger

	// LogLevel is the log level used by the logger.
	// It can be "error" or "debug".
	LogLevel string
}

// New creates a new PDFReader.
// It returns an error if the PDFReader cannot be created.
// The path argument is the path to the PDF file.
func New(path string, config Config) (*PDFReader, error) {
	var logger Logger

	if config.Logger == nil {
		logger = NewDefaultLogger(config.LogLevel)
	} else {
		logger = config.Logger
	}

	return NewPDFReader(path, logger)
}
