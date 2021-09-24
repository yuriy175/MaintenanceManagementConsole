package interfaces

// IOutputWriter describes output interface
type IOutputWriter interface {

	// Println formats using the default formats for its operands and writes to output.
	Println(a ...interface{}) (n int, err error)

	// Printf formats according to a format specifier and writes to output.
	Printf(format string, a ...interface{}) (n int, err error)
}
