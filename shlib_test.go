package shlib_test

import (
	"testing"

	"github.com/defn/shlib/v2"
)

func TestHello(t *testing.T) {
	if shlib.Hello() != "hello" {
		t.Errorf("got the wrong greeeting")
	}
}
