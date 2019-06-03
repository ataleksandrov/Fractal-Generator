package frctl

type options struct {
	numberOfGoroutines int
	outputFile         string
	quietMode          bool
}

func NewOptions(numberOfGoroutines int, outputFile string, quietMode bool) *options {
	return &options{
		numberOfGoroutines: numberOfGoroutines,
		outputFile:         outputFile,
		quietMode:          quietMode,
	}
}
