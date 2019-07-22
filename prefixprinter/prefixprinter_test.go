package prefixprinter

import (
	"fmt"
	"testing"

	"github.com/wader/logutils/printerfunc"
)

func test(t *testing.T, expected string, fn func(pp PrefixPrinter)) {
	called := false
	pp := New(printerfunc.New(func(v ...interface{}) {
		actual := fmt.Sprint(v...)
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
		called = true
	}), "abc")
	fn(pp)

	if !called {
		t.Error("Print not called")
	}
}

func TestPrint(t *testing.T) {
	test(t, "abc123", func(pp PrefixPrinter) { pp.Print(("123")) })
}

func TestPrintf(t *testing.T) {
	test(t, "abc123", func(pp PrefixPrinter) { pp.Printf(("123")) })
}
