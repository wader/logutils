package multiprinter

import (
	"fmt"
	"testing"

	"github.com/wader/logutils/printerfunc"
)

func test(t *testing.T, expected string, fn func(mp MultiPrinter)) {
	called := 0
	p := printerfunc.New(func(v ...interface{}) {
		actual := fmt.Sprint(v...)
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
		called++
	})
	mp := New(p, p)
	fn(mp)

	if called != 2 {
		t.Errorf("Print not called two times: %d", called)
	}
}

func TestPrint(t *testing.T) {
	test(t, "abc", func(mp MultiPrinter) { mp.Print("abc") })
}

func TestPrintf(t *testing.T) {
	test(t, "abc", func(mp MultiPrinter) { mp.Printf("abc") })
}
