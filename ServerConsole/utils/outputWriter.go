package utils

import (
	"fmt"

	"ServerConsole/interfaces"
)

// outputWriter implementation type
type outputWriter struct {
}

// OutputWriterNew creates an instance of outputWriter
func OutputWriterNew() interfaces.IOutputWriter {
	writer := &outputWriter{}
	return writer
}

// Println formats using the default formats for its operands and writes to standard output.
func (writer *outputWriter) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// Printf formats according to a format specifier and writes to standard output.
func (writer *outputWriter) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}
