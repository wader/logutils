package printwriter

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestWriteLogger(t *testing.T) {
	for _, c := range []struct {
		prefix   string
		writes   [][]byte
		expected []byte
	}{
		{
			"",
			[][]byte{
				[]byte("aaa\n"),
				[]byte("bbb\n"),
			},
			[]byte("aaa\nbbb\n"),
		},
		{
			">",
			[][]byte{
				[]byte("aaa\n"),
				[]byte("bbb\n"),
			},
			[]byte(">aaa\n>bbb\n"),
		},
		{
			">",
			[][]byte{
				[]byte("download\n"),
				[]byte("\rprogress 1"),
				[]byte("\rprogress 2"),
				[]byte("\rprogress 3"),
				[]byte("\n"),
			},
			[]byte(">download\n>\n>progress 1\n>progress 2\n>progress 3\n"),
		},
		{
			">",
			[][]byte{
				[]byte("a"),
				[]byte("b\n"),
			},
			[]byte(">ab\n"),
		},
		{
			">",
			[][]byte{
				[]byte("a"),
				[]byte("\n"),
			},
			[]byte(">a\n"),
		},
		{
			">",
			[][]byte{
				[]byte("a\nb\nc\n"),
			},
			[]byte(">a\n>b\n>c\n"),
		},
		{
			">",
			[][]byte{
				[]byte("a"),
				[]byte("\b\n"),
			},
			[]byte(">a \n"),
		},
		{
			">",
			[][]byte{
				[]byte("🐹"),
				[]byte("\n"),
			},
			[]byte(">🐹\n"),
		},
		{
			">",
			[][]byte{
				[]byte("a\n"),
				[]byte("b"),
			},
			[]byte(">a\n>b\n"),
		},
	} {

		actualBuf := &bytes.Buffer{}
		log := log.New(actualBuf, "", 0)
		var pw *PrintWriter
		if c.prefix != "" {
			pw = NewWithPrefix(log, c.prefix)
		} else {
			pw = New(log)
		}

		for _, w := range c.writes {
			pw.Write(w)
		}
		pw.Close()

		if !reflect.DeepEqual(actualBuf.Bytes(), c.expected) {
			t.Errorf("writes %#v, expected %#v, actual %#v", c.writes, string(c.expected), actualBuf.String())
		}
	}
}

func ExampleNewWithPrefix() {
	log := log.New(os.Stdout, "", 0)
	stdoutPW := NewWithPrefix(log, "stdout> ")
	cmd := exec.Command("echo", "hello")
	cmd.Stdout = stdoutPW
	cmd.Run()
	stdoutPW.Close()
	// Output: stdout> hello
}
