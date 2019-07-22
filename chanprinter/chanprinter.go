// Package chanprinter is a Printer that prints to a string channel
package chanprinter

import "fmt"

// ChanPrinter prints to a string channel
type ChanPrinter chan<- string

// New return a ChanPrinter
func New(c chan<- string) ChanPrinter {
	return ChanPrinter(c)
}

// Print to the channel
func (cp ChanPrinter) Print(v ...interface{}) {
	cp <- fmt.Sprint(v...)
}

// Printf to the channel
func (cp ChanPrinter) Printf(format string, v ...interface{}) {
	cp <- fmt.Sprintf(format, v...)
}
