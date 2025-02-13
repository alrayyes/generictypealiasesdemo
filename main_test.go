package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestCmdOutput(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	got := buf.String()

	want := `Hi, my name is Peter Integer and my ID is 1
Hi, my name is Peter String and my ID is a
Hi, my name is Peter Struct and my ID is {a}
`

	assert.Equal(t, want, got)
}
