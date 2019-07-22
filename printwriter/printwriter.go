// Package printwriter is a io.WriteCloser that print each lines with optional prefix.
// Can be useful when logging stdout/stderr output from a program or a line oriented
// byte stream etc.
package printwriter

import (
	"bytes"
	"unicode"
)

// Printer is something that can Print(), i.e. *log.Logger
type Printer interface {
	Print(v ...interface{})
}

// PrintWriter is a io.WriteCloser that print each lines with optional prefix
type PrintWriter struct {
	Printer Printer
	Prefix  string
	buf     bytes.Buffer
}

// NewWithPrefix return a initialized PrintWriter with a prefix
func NewWithPrefix(printer Printer, prefix string) *PrintWriter {
	return &PrintWriter{Printer: printer, Prefix: prefix}
}

// New return a initialized PrintWriter
func New(printer Printer) *PrintWriter {
	return NewWithPrefix(printer, "")
}

// same as bytes.IndexByte but with set of bytes to look for
func indexByteSet(s []byte, cs []byte) int {
	ri := -1

	for _, c := range cs {
		i := bytes.IndexByte(s, c)
		if i != -1 && (ri == -1 || i < ri) {
			ri = i
		}
	}

	return ri
}

func (wl *PrintWriter) Write(p []byte) (n int, err error) {
	wl.buf.Write(p)

	b := wl.buf.Bytes()
	pos := 0

	for {
		i := indexByteSet(b[pos:], []byte{'\n', '\r'})
		if i < 0 {
			break
		}

		// replace non-printable runes with whitespace otherwise fancy progress
		// bars etc might mess up output
		lineRunes := []rune(string(b[pos : pos+i]))
		for i, r := range lineRunes {
			if !unicode.IsPrint(r) {
				lineRunes[i] = ' '
			}
		}

		wl.Printer.Print(wl.Prefix + string(lineRunes))
		pos += i + 1
	}
	wl.buf.Reset()
	wl.buf.Write(b[pos:])

	return len(p), nil
}

// Close flushes any data left in the buffer as a line
func (wl *PrintWriter) Close() error {
	if wl.buf.Len() > 0 {
		wl.Write([]byte{'\n'})
	}
	return nil
}
