// Package printerfunc is a Printer that prints using a function (similar to http.HandlerFunc)
package printerfunc

import "fmt"

// PrinterFunc is a Printer that prints using a function (similar to http.HandlerFunc)
type PrinterFunc func(v ...interface{})

// New return a PrinterFunc
func New(fn func(v ...interface{})) PrinterFunc {
	return PrinterFunc(fn)
}

// Print using function
func (fn PrinterFunc) Print(v ...interface{}) {
	fn(v...)
}

// Printf using function
func (fn PrinterFunc) Printf(format string, v ...interface{}) {
	fn(fmt.Sprintf(format, v...))
}
