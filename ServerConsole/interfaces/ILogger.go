package interfaces

// ILogger describes logger interface
type ILogger interface {
	// Warning writes warning logs
	Warning(value string)

	// Error writes error logs
	Error(value string)

	// Info writes info logs
	Info(value string)

	// Errorf writes formatted error logs
	Errorf(format string, a ...interface{})

	// Info writes formatted info logs
	Infof(format string, a ...interface{})

	// GetZipContent returns zipped logs
	GetZipContent() ([]byte, string)
}