package mocks

import (
	"ServerConsole/interfaces"
)

// outputWriter mock implementation type
type outputMockWriter struct {
}

// OutputWriterNew creates an instance of outputWriter
func OutputMockWriterNew() interfaces.IOutputWriter {
	writer := &outputMockWriter{}
	return writer
}

// Println formats using the default formats for its operands and writes to standard output.
func (writer *outputMockWriter) Println(a ...interface{}) (n int, err error) {
	return 0, nil
}

// Printf formats according to a format specifier and writes to standard output.
func (writer *outputMockWriter) Printf(format string, a ...interface{}) (n int, err error) {
	return 0, nil
}
