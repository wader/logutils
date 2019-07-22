// Package prefixprinter is a Printer that prints with a prefix added
package prefixprinter

import "fmt"

// Printer is something that can Print/Printf(), i.e. *log.Logger
type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

// PrefixPrinter is a Printer that prints with a prefix added
type PrefixPrinter struct {
	Prefix  string
	Printer Printer
}

// New return a PrefixPrinter
func New(printer Printer, prefix string) PrefixPrinter {
	return PrefixPrinter{Printer: printer, Prefix: prefix}
}

// Print with a prefix
func (pp PrefixPrinter) Print(v ...interface{}) {
	pp.Printer.Print(pp.Prefix + fmt.Sprint(v...))
}

// Printf with a prefix
// TODO: what if prefix has format strings? but doing fmt.Sprintf() here
// pass the string instead might also be confusing
func (pp PrefixPrinter) Printf(format string, v ...interface{}) {
	pp.Printer.Printf(pp.Prefix+format, v...)
}
