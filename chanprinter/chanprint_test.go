package chanprinter

import (
	"testing"
)

func test(t *testing.T, expected string, fn func(ChanPrinter)) {
	c := make(chan string, 1)
	cp := New(c)
	fn(cp)
	select {
	case actual := <-c:
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
	default:
		t.Error("expected a string")
	}
}

func TestPrint(t *testing.T) {
	test(t, "abc", func(cp ChanPrinter) { cp.Print("abc") })
}

func TestPrintf(t *testing.T) {
	test(t, "abc", func(cp ChanPrinter) { cp.Printf("abc") })
}
