// Package lossychanprinter is a Printer that prints lossely to a string channel
package lossychanprinter

import "fmt"

// LossyChanPrinter prints lossely to a string channel
type LossyChanPrinter chan<- string

// New return a LossyChanPrinter
func New(c chan<- string) LossyChanPrinter {
	return LossyChanPrinter(c)
}

// Print to the channel if possible
func (lcp LossyChanPrinter) Print(v ...interface{}) {
	select {
	case lcp <- fmt.Sprint(v...):
	default:
	}
}

// Printf to the channel if possible
func (lcp LossyChanPrinter) Printf(format string, v ...interface{}) {
	select {
	case lcp <- fmt.Sprintf(format, v...):
	default:
	}
}
