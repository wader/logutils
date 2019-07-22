// Package multiprinter is a Printer that prints to many Printers
package multiprinter

// Printer is something that can Print/Printf(), i.e. *log.Logger
type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

// MultiPrinter prints to many Printers
type MultiPrinter []Printer

// New return a MultiPrinter
func New(printers ...Printer) MultiPrinter {
	return MultiPrinter(printers)
}

// Print to printers
func (mp MultiPrinter) Print(v ...interface{}) {
	for _, p := range mp {
		p.Print(v...)
	}
}

// Printf to printers
func (mp MultiPrinter) Printf(format string, v ...interface{}) {
	for _, p := range mp {
		p.Printf(format, v...)
	}
}
