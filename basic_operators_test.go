package lisp

import (
	"testing"
)

func testBasicOperators(t *testing.T) {
	result := eval("(quote (a b c))")
	if "(a b c)" != result {
		t.Fatalf(result)
	}
}
