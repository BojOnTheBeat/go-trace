package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {

	// It's probably better to omit Tracer type and let type inferrence work here
	// but I'll leave this for education purposes
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
