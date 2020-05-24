package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code
type Tracer interface {
	Trace(...interface{}) //...interface means Trace accepts 0 or more args of any type
}

type tracer struct {
	out io.Writer
}

func (t tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New initializes and returns _something_
// that implements the Tracer interface
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
