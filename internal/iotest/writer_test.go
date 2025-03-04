package iotest

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeT struct {
	*testing.T

	Buffer bytes.Buffer
}

func (t *fakeT) Logf(msg string, args ...interface{}) {
	fmt.Fprintln(&t.Buffer, fmt.Sprintf(msg, args...))
	// println to make sure it ends with a newline
}

func TestWriter(t *testing.T) {
	t.Parallel()

	fakeT := fakeT{T: t}
	w := Writer(&fakeT)
	_, err := io.WriteString(w, "foo")
	require.NoError(t, err)
	assert.Equal(t, "foo\n", fakeT.Buffer.String())
	// If we wanted this to be more accurate, we would have it buffer
	// the input on newlines simillar to the log-based io.Writer.
	// It doesn't matter here.
}
