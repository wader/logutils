package printerfunc

import (
	"fmt"
	"testing"
)

func test(t *testing.T, expected string, fn func(p PrinterFunc)) {
	called := false
	p := New(func(v ...interface{}) {
		actual := fmt.Sprint(v...)
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
		called = true
	})
	fn(p)

	if !called {
		t.Error("Print not called")
	}
}

func TestPrint(t *testing.T) {
	test(t, "abc", func(p PrinterFunc) { p.Print("abc") })
}

func TestPrintf(t *testing.T) {
	test(t, "abc", func(p PrinterFunc) { p.Printf("abc") })
}
