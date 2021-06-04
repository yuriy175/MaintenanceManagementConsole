package interfaces

// ILogger describes logger interface
type ILogger interface {
	// Warning writes warning logs
	Warning(value string)

	// Error writes error logs
	Error(value string)

	// Info writes info logs
	Info(value string)
}