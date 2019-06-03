package frctl

type options struct {
	numberOfGoroutines int
	outputFile         string
	quietMode          bool
}

// NewOptions returns new options e.g. numberOfGoroutines, outputFile, quietMode
func NewOptions(numberOfGoroutines int, outputFile string, quietMode bool) *options {
	return &options{
		numberOfGoroutines: numberOfGoroutines,
		outputFile:         outputFile,
		quietMode:          quietMode,
	}
}
