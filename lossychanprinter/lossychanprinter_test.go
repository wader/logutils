package lossychanprinter

import (
	"testing"
)

func test(t *testing.T, expected string, fn func(LossyChanPrinter)) {
	c := make(chan string, 1)
	cp := New(c)
	fn(cp)
	fn(cp)
	select {
	case actual := <-c:
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
	default:
		t.Error("expected a string")
	}
	select {
	case actual := <-c:
		t.Errorf("did not expect a string, got %q", actual)
	default:
	}
}

func TestPrint(t *testing.T) {
	test(t, "abc", func(lcp LossyChanPrinter) { lcp.Print("abc") })
}

func TestPrintf(t *testing.T) {
	test(t, "abc", func(lcp LossyChanPrinter) { lcp.Printf("abc") })
}
