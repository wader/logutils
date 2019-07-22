// Package nopprinter is a Printer that does nothing
package nopprinter

// NopPrinter is a Printer that does nothing
type NopPrinter struct{}

// New return a NopPrinter
func New() NopPrinter { return NopPrinter{} }

// Print nothing
func (NopPrinter) Print(v ...interface{}) {}

// Printf nothing
func (NopPrinter) Printf(format string, v ...interface{}) {}
