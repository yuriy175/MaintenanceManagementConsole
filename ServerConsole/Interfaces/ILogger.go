package interfaces

// ILogger describes logger interface
type ILogger interface {
	Warning(value string)
	Error(value string)
}
